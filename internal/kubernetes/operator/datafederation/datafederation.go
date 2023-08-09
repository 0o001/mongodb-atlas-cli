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

package datafederation

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/mongodb/mongodb-atlas-cli/internal/kubernetes/operator/features"
	"github.com/mongodb/mongodb-atlas-cli/internal/kubernetes/operator/resources"
	atlasV1 "github.com/mongodb/mongodb-atlas-kubernetes/pkg/api/v1"
	"github.com/mongodb/mongodb-atlas-kubernetes/pkg/api/v1/common"
	"github.com/mongodb/mongodb-atlas-kubernetes/pkg/api/v1/status"
	atlasv2 "go.mongodb.org/atlas-sdk/v20230201003/admin"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	// "github.com/mongodb/mongodb-atlas-cli/internal/store"
)

func BuildAtlasDataFederation(projectName string, dataFederation atlasv2.DataLakeTenant, operatorVersion, targetNamespace string, dictionary map[string]string) (runtime.Object, error) {

	AtlasDataFederation := &atlasV1.AtlasDataFederation{
		TypeMeta: v1.TypeMeta{
			APIVersion: "atlas.mongodb.com/v1",
			Kind:       "AtlasDataFederation",
		},
		ObjectMeta: v1.ObjectMeta{
			Name:      resources.NormalizeAtlasName(fmt.Sprintf("%s-%s", projectName, *dataFederation.Name), dictionary),
			Namespace: targetNamespace,
			Labels: map[string]string{
				features.ResourceVersion: operatorVersion,
			},
		},
		Spec: GetDataFederationSpec(&dataFederation, targetNamespace, projectName),
		Status: status.DataFederationStatus{
			Common: status.Common{
				Conditions: []status.Condition{},
			},
		},
	}
	return AtlasDataFederation, nil
}

func GetDataFederationSpec(dataFederationSpec *atlasv2.DataLakeTenant, targetNamespace, projectName string) atlasV1.DataFederationSpec {
	return atlasV1.DataFederationSpec{
		Project:             common.ResourceRefNamespaced{Name: projectName, Namespace: targetNamespace},
		Name:                aws.ToString(dataFederationSpec.Name),
		CloudProviderConfig: GetCloudProviderConfig(dataFederationSpec.CloudProviderConfig),
		DataProcessRegion:   &atlasV1.DataProcessRegion{CloudProvider: dataFederationSpec.DataProcessRegion.CloudProvider, Region: dataFederationSpec.DataProcessRegion.Region},
		Storage:             &atlasV1.Storage{Databases: GetDatabases(dataFederationSpec.Storage.Databases), Stores: GetStores(dataFederationSpec.Storage.Stores)},
		// PrivateEndpoints: GetPrivateEndpoints(),
	}
}

func GetCloudProviderConfig(cloudProviderConfig *atlasv2.DataLakeCloudProviderConfig) *atlasV1.CloudProviderConfig {
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
			RoleID:       cloudProviderConfig.Aws.RoleId,
			TestS3Bucket: cloudProviderConfig.Aws.TestS3Bucket,
		},
	}
}

func GetDatabases(database []atlasv2.DataLakeDatabaseInstance) []atlasV1.Database {
	var result []atlasV1.Database
	for _, obj := range database {
		atlasDatabase := atlasV1.Database{
			Collections:            GetCollection(obj.Collections),
			MaxWildcardCollections: aws.ToInt(obj.MaxWildcardCollections),
			Name:                   aws.ToString(obj.Name),
			Views:                  GetViews(obj.Views),
		}
		result = append(result, atlasDatabase)
	}
	return result
}

func GetCollection(collections []atlasv2.DataLakeDatabaseCollection) []atlasV1.Collection {
	var result []atlasV1.Collection
	for _, obj := range collections {
		collection := atlasV1.Collection{
			DataSources: GetDataSources(obj.DataSources),
			Name:        aws.ToString(obj.Name),
		}
		result = append(result, collection)
	}
	return result
}

func GetDataSources(dataSources []atlasv2.DataLakeDatabaseDataSourceSettings) []atlasV1.DataSource {
	var result []atlasV1.DataSource
	for _, obj := range dataSources {
		dataSource := atlasV1.DataSource{
			AllowInsecure:       aws.ToBool(obj.AllowInsecure),
			Collection:          aws.ToString(obj.Collection),
			CollectionRegex:     aws.ToString(obj.CollectionRegex),
			Database:            aws.ToString(obj.Database),
			DatabaseRegex:       aws.ToString(obj.DatabaseRegex),
			DefaultFormat:       aws.ToString(obj.DefaultFormat),
			Path:                aws.ToString(obj.Path),
			ProvenanceFieldName: aws.ToString(obj.ProvenanceFieldName),
			StoreName:           aws.ToString(obj.StoreName),
			Urls:                obj.Urls,
		}
		result = append(result, dataSource)
	}
	return result
}

func GetViews(views []atlasv2.DataLakeApiBase) []atlasV1.View {
	var result []atlasV1.View
	for _, obj := range views {
		view := atlasV1.View{
			Name:     aws.ToString(obj.Name),
			Pipeline: aws.ToString(obj.Pipeline),
			Source:   aws.ToString(obj.Source),
		}
		result = append(result, view)
	}
	return result
}

func GetStores(stores []atlasv2.DataLakeStoreSettings) []atlasV1.Store {
	var result []atlasV1.Store
	for _, obj := range stores {
		store := atlasV1.Store{
			Name:                     aws.ToString(obj.Name),
			Provider:                 obj.Provider,
			AdditionalStorageClasses: obj.AdditionalStorageClasses,
			Bucket:                   aws.ToString(obj.Bucket),
			Delimiter:                aws.ToString(obj.Delimiter),
			IncludeTags:              aws.ToBool(obj.IncludeTags),
			Prefix:                   aws.ToString(obj.Prefix),
			Public:                   aws.ToBool(obj.Public),
			Region:                   aws.ToString(obj.Region),
		}
		result = append(result, store)
	}
	return result
}

// func GetPrivateEndpoints(privateEndpoint []atlasv2.ListDataFederationPrivateEndpointsApiParams) []atlasV1.DataFederationPE {
// 	hold, err := store.DataLakePrivateEndpointLister.DataLakePrivateEndpoints(privateEndpoint)
// }
