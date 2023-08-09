// Copyright 2022 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package operator

import (
	"bytes"
	jsonv1 "encoding/json"
	"errors"
	"fmt"
	"reflect"

	"github.com/mongodb/mongodb-atlas-cli/internal/cli/atlas/datafederation"
	"github.com/mongodb/mongodb-atlas-cli/internal/kubernetes/operator/dbusers"
	"github.com/mongodb/mongodb-atlas-cli/internal/kubernetes/operator/deployment"
	"github.com/mongodb/mongodb-atlas-cli/internal/kubernetes/operator/features"
	"github.com/mongodb/mongodb-atlas-cli/internal/kubernetes/operator/project"
	"github.com/mongodb/mongodb-atlas-cli/internal/kubernetes/operator/resources"
	"github.com/mongodb/mongodb-atlas-cli/internal/store"
	atlasV1 "github.com/mongodb/mongodb-atlas-kubernetes/pkg/api/v1"
	"github.com/mongodb/mongodb-atlas-kubernetes/pkg/api/v1/common"
	"github.com/mongodb/mongodb-atlas-kubernetes/pkg/api/v1/status"
	atlasv2 "go.mongodb.org/atlas-sdk/v20230201003/admin"
	"go.mongodb.org/atlas/mongodbatlas"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer/json"
	"k8s.io/client-go/kubernetes/scheme"
)

const (
	yamlSeparator        = "---\r\n"
	maxClusters          = 500
	DefaultClustersCount = 10
)

type ConfigExporter struct {
	featureValidator        features.FeatureValidator
	dataProvider            store.AtlasOperatorGenericStore
	credsProvider           store.CredentialsGetter
	projectID               string
	clusters                []string
	targetNamespace         string
	operatorVersion         string
	includeSecretsData      bool
	orgID                   string
	dictionaryForAtlasNames map[string]string
}

var (
	ErrClusterNotFound        = errors.New("cluster not found")
	ErrNoCloudManagerClusters = errors.New("can not get 'advanced clusters' object")
)

func NewConfigExporter(dataProvider store.AtlasOperatorGenericStore, credsProvider store.CredentialsGetter, projectID, orgID string) *ConfigExporter {
	return &ConfigExporter{
		dataProvider:            dataProvider,
		credsProvider:           credsProvider,
		projectID:               projectID,
		clusters:                []string{},
		targetNamespace:         "",
		includeSecretsData:      false,
		orgID:                   orgID,
		dictionaryForAtlasNames: resources.AtlasNameToKubernetesName(),
	}
}

func (e *ConfigExporter) WithClustersNames(clusters []string) *ConfigExporter {
	e.clusters = clusters
	return e
}

func (e *ConfigExporter) WithTargetNamespace(namespace string) *ConfigExporter {
	e.targetNamespace = namespace
	return e
}

func (e *ConfigExporter) WithSecretsData(enabled bool) *ConfigExporter {
	e.includeSecretsData = enabled
	return e
}

func (e *ConfigExporter) WithTargetOperatorVersion(version string) *ConfigExporter {
	e.operatorVersion = version
	return e
}

func (e *ConfigExporter) WithFeatureValidator(validator features.FeatureValidator) *ConfigExporter {
	e.featureValidator = validator
	return e
}

func (e *ConfigExporter) Run() (string, error) {
	// TODO: Add REST to OPERATOR entities matcher

	output := bytes.NewBufferString(yamlSeparator)
	var resources []runtime.Object

	serializer := json.NewSerializerWithOptions(json.DefaultMetaFactory, scheme.Scheme, scheme.Scheme,
		json.SerializerOptions{Yaml: true, Pretty: true})

	projectResources, projectName, err := e.exportProject()
	if err != nil {
		return "", err
	}
	resources = append(resources, projectResources...)

	deploymentsResources, err := e.exportDeployments(projectName)
	if err != nil {
		return "", err
	}
	resources = append(resources, deploymentsResources...)

	dataFederationResource, err := e.exportDataFedertaion(projectName)
	if err != nil {
		return "", err
	}
	resources = append(resources, dataFederationResource...)

	for _, res := range resources {
		err = serializer.Encode(res, output)
		if err != nil {
			return "", err
		}
		output.WriteString(yamlSeparator)
	}

	return output.String(), nil
}

func (e *ConfigExporter) exportProject() ([]runtime.Object, string, error) {
	var resources []runtime.Object

	// Project
	projectData, err := project.BuildAtlasProject(e.dataProvider, e.featureValidator, e.orgID, e.projectID, e.targetNamespace, e.includeSecretsData, e.dictionaryForAtlasNames, e.operatorVersion)
	if err != nil {
		return nil, "", err
	}
	resources = append(resources, projectData.Project)
	for _, secret := range projectData.Secrets {
		resources = append(resources, secret)
	}

	// Teams
	for _, team := range projectData.Teams {
		resources = append(resources, team)
	}

	// Project secret with credentials
	resources = append(resources, project.BuildProjectConnectionSecret(e.credsProvider,
		projectData.Project.Name,
		projectData.Project.Namespace, e.orgID, e.includeSecretsData, e.dictionaryForAtlasNames))

	// DB users
	usersData, relatedSecrets, err := dbusers.BuildDBUsers(e.dataProvider, e.projectID, projectData.Project.Name, e.targetNamespace, e.dictionaryForAtlasNames, e.operatorVersion)
	if err != nil {
		return nil, "", err
	}
	for _, user := range usersData {
		resources = append(resources, user)
	}
	for _, secret := range relatedSecrets {
		resources = append(resources, secret)
	}

	return resources, projectData.Project.Name, nil
}

func (e *ConfigExporter) exportDeployments(projectName string) ([]runtime.Object, error) {
	var result []runtime.Object

	if len(e.clusters) == 0 {
		clusters, err := fetchClusterNames(e.dataProvider, e.projectID)
		if err != nil {
			return nil, err
		}
		e.clusters = clusters
	}

	for _, deploymentName := range e.clusters {
		// Try advanced cluster first
		if advancedCluster, err := deployment.BuildAtlasAdvancedDeployment(e.dataProvider, e.featureValidator, e.projectID, projectName, deploymentName, e.targetNamespace, e.dictionaryForAtlasNames, e.operatorVersion); err == nil {
			if advancedCluster != nil {
				// Append deployment to result
				result = append(result, advancedCluster.Deployment)
				// Append backup schedule
				if advancedCluster.BackupSchedule != nil {
					result = append(result, advancedCluster.BackupSchedule)
				}
				// Append backup policies (one)
				for _, policy := range advancedCluster.BackupPolicies {
					if policy != nil {
						result = append(result, policy)
					}
				}
			}
			continue
		}

		// Try serverless cluster next
		if serverlessCluster, err := deployment.BuildServerlessDeployments(e.dataProvider, e.featureValidator, e.projectID, projectName, deploymentName, e.targetNamespace, e.dictionaryForAtlasNames, e.operatorVersion); err == nil {
			if serverlessCluster != nil {
				result = append(result, serverlessCluster)
			}
			continue
		}
		return nil, fmt.Errorf("%w: %s(%s)", ErrClusterNotFound, deploymentName, e.projectID)
	}
	return result, nil
}

func fetchClusterNames(clustersProvider store.AtlasAllClustersLister, projectID string) ([]string, error) {
	result := make([]string, 0, DefaultClustersCount)
	response, err := clustersProvider.ProjectClusters(projectID, &mongodbatlas.ListOptions{ItemsPerPage: maxClusters})
	if err != nil {
		return nil, err
	}

	if clusters, ok := response.(*atlasv2.PaginatedAdvancedClusterDescription); ok {
		if clusters == nil {
			return nil, ErrNoCloudManagerClusters
		}

		for i := range clusters.Results {
			cluster := clusters.Results[i]
			if reflect.ValueOf(cluster).IsZero() {
				continue
			}
			result = append(result, cluster.GetName())
		}
	}

	serverlessInstances, err := clustersProvider.ServerlessInstances(projectID, &mongodbatlas.ListOptions{ItemsPerPage: maxClusters})
	if err != nil {
		return nil, err
	}

	if serverlessInstances == nil {
		return result, nil
	}

	for i := range serverlessInstances.Results {
		cluster := serverlessInstances.Results[i]
		result = append(result, *cluster.Name)
	}

	return result, nil
}

func (e *ConfigExporter) exportDataFedertaion(projectName string) ([]runtime.Object, error) {
	var result []runtime.Object
	fmt.Println("In exporter")
	fmt.Println("ProjectID: ", e.projectID)
	// Fetching list of data federations
	dataFederations, err := e.dataProvider.DataLakes(e.projectID)
	if err != nil {
		panic(err)
	}
	// TODO: delete vvv
	j, _ := jsonv1.MarshalIndent(dataFederations, "", " ")
	fmt.Println(">>>", string(j))

	for _, obj := range dataFederations {
		x, err = datafederation.BuildAtlasDataFederation(projectName)
	}

	// Storing each
	convertCloudProviderConfig := func(cloudProviderConfig *mongodbatlas.CloudProviderConfig) *atlasV1.CloudProviderConfig {
		if cloudProviderConfig == nil {
			return &atlasV1.CloudProviderConfig{
				AWS: &atlasV1.AWSProviderConfig{
					RoleID:       "",
					TestS3Bucket: "",
				},
			}
		}
		return &atlasV1.CloudProviderConfig{
			AWS: &atlasV1.AWSProviderConfig{
				RoleID:       cloudProviderConfig.AWSConfig.RoleID,
				TestS3Bucket: cloudProviderConfig.AWSConfig.TestS3Bucket,
			},
		}
	}

	convertDataFederationSpec := func(dataFederationSpec *mongodbatlas.DataLake) *atlasV1.DataFederationSpec {
		return &atlasV1.DataFederationSpec{
			Project:             common.ResourceRefNamespaced{Name: dataFederationSpec.GroupID, Namespace: e.targetNamespace},
			Name:                dataFederationSpec.Name,
			CloudProviderConfig: convertCloudProviderConfig(&dataFederationSpec.CloudProviderConfig),
			DataProcessRegion:   &atlasV1.DataProcessRegion{CloudProvider: dataFederationSpec.DataProcessRegion.CloudProvider, Region: dataFederationSpec.DataProcessRegion.Region},
		}
	}

	for _, obj := range dataFederations {
		AtlasDataFederation := &atlasV1.AtlasDataFederation{
			TypeMeta: v1.TypeMeta{
				APIVersion: "atlas.mongodb.com/v1",
				Kind:       "AtlasDataFederation",
			},
			ObjectMeta: v1.ObjectMeta{
				Name:      resources.NormalizeAtlasName(fmt.Sprintf("%s-%s", projectName, obj.Name), e.dictionaryForAtlasNames),
				Namespace: e.targetNamespace,
				Labels: map[string]string{
					features.ResourceVersion: e.operatorVersion,
				},
			},
			Spec: *convertDataFederationSpec(&obj),
			Status: status.DataFederationStatus{
				Common: status.Common{
					Conditions: []status.Condition{},
				},
			},
		}
		result = append(result, AtlasDataFederation)
	}
	return result, nil
}

func getDatabases(database []mongodbatlas.DataLakeDatabase) []atlasV1.Database {
	var result []atlasV1.Database
	for _, obj := range database {
		atlasDatabase := atlasV1.Database{
			Collections:            getCollection(obj.Collections),
			MaxWildcardCollections: int(*obj.MaxWildcardCollections),
			Name:                   obj.Name,
			Views:                  getViews(obj.Views),
		}
		result = append(result, atlasDatabase)
	}
	return result
}

func getCollection(collections []mongodbatlas.DataLakeCollection) []atlasV1.Collection {
	var result []atlasV1.Collection
	for _, obj := range collections {
		collection := atlasV1.Collection{
			DataSources: getDataSources(obj.DataSources),
			Name:        obj.Name,
		}
		result = append(result, collection)
	}
	return result
}

func getDataSources(dataSources []mongodbatlas.DataLakeDataSource) []atlasV1.DataSource {
	var result []atlasV1.DataSource
	var emptyarr []string
	for _, obj := range dataSources {
		dataSource := atlasV1.DataSource{
			AllowInsecure:       false,
			Collection:          "",
			CollectionRegex:     "",
			Database:            "",
			DatabaseRegex:       "",
			DefaultFormat:       obj.DefaultFormat,
			Path:                obj.Path,
			ProvenanceFieldName: "",
			StoreName:           obj.StoreName,
			Urls:                emptyarr,
		}
		result = append(result, dataSource)
	}
	return result
}

func getViews(views []mongodbatlas.DataLakeDatabaseView) []atlasV1.View {
	var result []atlasV1.View
	for _, obj := range views {
		view := atlasV1.View{
			Name:     obj.Name,
			Pipeline: obj.Pipeline,
			Source:   obj.Source,
		}
		result = append(result, view)
	}
	return result
}
