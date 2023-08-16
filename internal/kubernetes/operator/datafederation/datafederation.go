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
	atlas "go.mongodb.org/atlas/mongodbatlas"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

const (
	DeletingState = "DELETING"
	DeletedState  = "DELETED"
)

func BuildAtlasDataFederation( /*dataFederationStore store.AtlasOperatorDataFederationStore,*/ dataFederation *atlas.DataFederationInstance, projectName, operatorVersion, targetNamespace string, dictionary map[string]string) (runtime.Object, error) {
	if !isDataFederationExportable(dataFederation) {
		return nil, nil
	}
	AtlasDataFederation := &atlasV1.AtlasDataFederation{
		TypeMeta: v1.TypeMeta{
			APIVersion: "atlas.mongodb.com/v1",
			Kind:       "AtlasDataFederation",
		},
		ObjectMeta: v1.ObjectMeta{
			Name:      resources.NormalizeAtlasName(fmt.Sprintf("%s-%s", projectName, dataFederation.Name), dictionary),
			Namespace: targetNamespace,
			Labels: map[string]string{
				features.ResourceVersion: operatorVersion,
			},
		},
		Spec: GetDataFederationSpec(dataFederation, targetNamespace, projectName),
		Status: status.DataFederationStatus{
			Common: status.Common{
				Conditions: []status.Condition{},
			},
		},
	}
	return AtlasDataFederation, nil
}

func isDataFederationExportable(dataFederation *atlas.DataFederationInstance) bool {
	if dataFederation.State == DeletingState || dataFederation.State == DeletedState {
		return false
	}
	return true
}

func GetDataFederationSpec(dataFederationSpec *atlas.DataFederationInstance, targetNamespace, projectName string) atlasV1.DataFederationSpec {
	return atlasV1.DataFederationSpec{
		Project:             common.ResourceRefNamespaced{Name: projectName, Namespace: targetNamespace},
		Name:                dataFederationSpec.Name,
		CloudProviderConfig: GetCloudProviderConfig(dataFederationSpec.CloudProviderConfig),
		DataProcessRegion:   &atlasV1.DataProcessRegion{CloudProvider: dataFederationSpec.DataProcessRegion.CloudProvider, Region: dataFederationSpec.DataProcessRegion.Region},
		Storage:             &atlasV1.Storage{Databases: GetDatabases(dataFederationSpec.Storage.Databases), Stores: GetStores(dataFederationSpec.Storage.Stores)},
	}
}

func GetCloudProviderConfig(cloudProviderConfig *atlas.CloudProviderConfig) *atlasV1.CloudProviderConfig {
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

func GetDatabases(database []*atlas.DataFederationDatabase) []atlasV1.Database {
	var result []atlasV1.Database
	for _, obj := range database {
		atlasDatabase := atlasV1.Database{
			Collections:            GetCollection(obj.Collections),
			MaxWildcardCollections: int(obj.MaxWildcardCollections),
			Name:                   obj.Name,
			Views:                  GetViews(obj.Views),
		}
		result = append(result, atlasDatabase)
	}
	return result
}

func GetCollection(collections []*atlas.DataFederationCollection) []atlasV1.Collection {
	var result []atlasV1.Collection
	for _, obj := range collections {
		collection := atlasV1.Collection{
			DataSources: GetDataSources(obj.DataSources),
			Name:        obj.Name,
		}
		result = append(result, collection)
	}
	return result
}

func GetDataSources(dataSources []*atlas.DataFederationDataSource) []atlasV1.DataSource {
	var result []atlasV1.DataSource
	for _, obj := range dataSources {
		dataSource := atlasV1.DataSource{
			AllowInsecure:       aws.ToBool(obj.AllowInsecure),
			Collection:          obj.Collection,
			CollectionRegex:     obj.CollectionRegex,
			Database:            obj.Database,
			DatabaseRegex:       obj.DatabaseRegex,
			DefaultFormat:       obj.DefaultFormat,
			Path:                obj.Path,
			ProvenanceFieldName: obj.ProvenanceFieldName,
			StoreName:           obj.StoreName,
			Urls:                aws.ToStringSlice(obj.Urls),
		}
		result = append(result, dataSource)
	}
	return result
}

func GetViews(views []*atlas.DataFederationDatabaseView) []atlasV1.View {
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

func GetStores(stores []*atlas.DataFederationStore) []atlasV1.Store {
	var result []atlasV1.Store
	for _, obj := range stores {
		store := atlasV1.Store{
			Name:                     obj.Name,
			Provider:                 obj.Provider,
			AdditionalStorageClasses: aws.ToStringSlice(obj.AdditionalStorageClasses),
			Bucket:                   obj.Bucket,
			Delimiter:                obj.Delimiter,
			IncludeTags:              aws.ToBool(obj.IncludeTags),
			Prefix:                   obj.Prefix,
			Public:                   aws.ToBool(obj.Public),
			Region:                   obj.Region,
		}
		result = append(result, store)
	}
	return result
}
