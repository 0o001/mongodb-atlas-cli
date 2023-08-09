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

	"github.com/mongodb/mongodb-atlas-cli/internal/kubernetes/operator/features"
	"github.com/mongodb/mongodb-atlas-cli/internal/kubernetes/operator/resources"
	atlasV1 "github.com/mongodb/mongodb-atlas-kubernetes/pkg/api/v1"
	"github.com/mongodb/mongodb-atlas-kubernetes/pkg/api/v1/common"
	"github.com/mongodb/mongodb-atlas-kubernetes/pkg/api/v1/status"
	"go.mongodb.org/atlas/mongodbatlas"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

func (e *ConfigExporter) BuildAtlasDataFederation(projectName string, dataFederation mongodbatlas.DataLake) (runtime.Object, error) {

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
		Spec: getDataFederationSpec(&dataFederation),
		Status: status.DataFederationStatus{
			Common: status.Common{
				Conditions: []status.Condition{},
			},
		},
	}
	return AtlasDataFederation, nil
}

func getCloudProviderConfig(cloudProviderConfig *mongodbatlas.CloudProviderConfig) *atlasV1.CloudProviderConfig {
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

func getDataFederationSpec(dataFederationSpec *mongodbatlas.DataLake) atlasV1.DataFederationSpec {
	return atlasV1.DataFederationSpec{
		Project:             common.ResourceRefNamespaced{Name: dataFederationSpec.GroupID, Namespace: e.targetNamespace},
		Name:                dataFederationSpec.Name,
		CloudProviderConfig: getCloudProviderConfig(&dataFederationSpec.CloudProviderConfig),
		DataProcessRegion:   &atlasV1.DataProcessRegion{CloudProvider: dataFederationSpec.DataProcessRegion.CloudProvider, Region: dataFederationSpec.DataProcessRegion.Region},
	}
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
