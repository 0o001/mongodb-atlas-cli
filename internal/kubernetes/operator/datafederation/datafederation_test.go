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

//go:build unit

package datafederation

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/internal/kubernetes/operator/features"
	"github.com/mongodb/mongodb-atlas-cli/internal/kubernetes/operator/resources"
	"github.com/mongodb/mongodb-atlas-cli/internal/pointer"
	atlasV1 "github.com/mongodb/mongodb-atlas-kubernetes/pkg/api/v1"
	"github.com/mongodb/mongodb-atlas-kubernetes/pkg/api/v1/common"
	"github.com/mongodb/mongodb-atlas-kubernetes/pkg/api/v1/status"
	atlas "go.mongodb.org/atlas/mongodbatlas"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const projectName = "testProject-1"
const dataFederationName = "testDataFederation-1"
const targetNamespace = "test-namespace-1"
const resourceVersion = "x.y.z"

func Test_BuildAtlasDataFederation(t *testing.T) {
	dictionary := resources.AtlasNameToKubernetesName()

	t.Run("Can import Data Federations", func(t *testing.T) {
		dataFederation := &atlas.DataFederationInstance{

			CloudProviderConfig: &atlas.CloudProviderConfig{
				AWSConfig: atlas.AwsCloudProviderConfig{
					RoleID:       "TestRoleID",
					TestS3Bucket: "TestBucket",
				},
			},
			DataProcessRegion: &atlas.DataProcessRegion{
				CloudProvider: "TestProvider",
				Region:        "TestRegion",
			},
			Hostnames: []string{"TestHosname"},
			Name:      dataFederationName,
			State:     "TestState",
			Storage: &atlas.DataFederationStorage{
				Databases: []*atlas.DataFederationDatabase{
					{
						Collections: []*atlas.DataFederationCollection{
							{
								DataSources: []*atlas.DataFederationDataSource{
									{
										AllowInsecure:       pointer.Get(true),
										Collection:          "TestCollection",
										CollectionRegex:     "TestCollectionRegex",
										Database:            "TestDatabase",
										DatabaseRegex:       "TestDatabaseRegex",
										DefaultFormat:       "TestFormat",
										Path:                "TestPath",
										ProvenanceFieldName: "TestFieldName",
										StoreName:           "TestStoreName",
										Urls:                []*string{pointer.Get("TestUrl")},
									},
								},
								Name: "TestName",
							},
						},
						MaxWildcardCollections: 10,
						Name:                   "TestName",
						Views: []*atlas.DataFederationDatabaseView{
							{
								Name:     "TestName",
								Pipeline: "TestPipeline",
								Source:   "TestSource",
							},
						},
					},
				},
				Stores: []*atlas.DataFederationStore{
					{
						Name:                     "TestName",
						Provider:                 "TestProvider",
						AdditionalStorageClasses: []*string{pointer.Get("TestClasses")},
						Bucket:                   "TestBucket",
						Delimiter:                "TestDelimiter",
						IncludeTags:              pointer.Get(true),
						Prefix:                   "TestPrefix",
						Public:                   pointer.Get(true),
						Region:                   "TestRegion",
					},
				},
			},
		}

		expected := &atlasV1.AtlasDataFederation{
			TypeMeta: v1.TypeMeta{
				Kind:       "AtlasDataFederation",
				APIVersion: "atlas.mongodb.com/v1",
			},
			ObjectMeta: v1.ObjectMeta{
				Name:      resources.NormalizeAtlasName(fmt.Sprintf("%s-%s", projectName, dataFederation.Name), dictionary),
				Namespace: targetNamespace,
				Labels: map[string]string{
					features.ResourceVersion: resourceVersion,
				},
			},
			Spec: atlasV1.DataFederationSpec{
				Project: common.ResourceRefNamespaced{
					Name:      projectName,
					Namespace: targetNamespace,
				},
				Name: dataFederationName,
				CloudProviderConfig: &atlasV1.CloudProviderConfig{
					AWS: &atlasV1.AWSProviderConfig{
						RoleID:       dataFederation.CloudProviderConfig.AWSConfig.RoleID,
						TestS3Bucket: dataFederation.CloudProviderConfig.AWSConfig.TestS3Bucket,
					},
				},
				DataProcessRegion: &atlasV1.DataProcessRegion{
					CloudProvider: dataFederation.DataProcessRegion.CloudProvider,
					Region:        dataFederation.DataProcessRegion.Region,
				},
				Storage: &atlasV1.Storage{
					Databases: []atlasV1.Database{
						{
							Collections: []atlasV1.Collection{
								{
									DataSources: []atlasV1.DataSource{
										{
											AllowInsecure:       true,
											Collection:          dataFederation.Storage.Databases[0].Collections[0].DataSources[0].Collection,
											CollectionRegex:     dataFederation.Storage.Databases[0].Collections[0].DataSources[0].CollectionRegex,
											Database:            dataFederation.Storage.Databases[0].Collections[0].DataSources[0].Database,
											DatabaseRegex:       dataFederation.Storage.Databases[0].Collections[0].DataSources[0].DatabaseRegex,
											DefaultFormat:       dataFederation.Storage.Databases[0].Collections[0].DataSources[0].DefaultFormat,
											Path:                dataFederation.Storage.Databases[0].Collections[0].DataSources[0].Path,
											ProvenanceFieldName: dataFederation.Storage.Databases[0].Collections[0].DataSources[0].ProvenanceFieldName,
											StoreName:           dataFederation.Storage.Databases[0].Collections[0].DataSources[0].StoreName,
											Urls:                []string{*dataFederation.Storage.Databases[0].Collections[0].DataSources[0].Urls[0]},
										},
									},
									Name: dataFederation.Storage.Databases[0].Collections[0].Name,
								},
							},
							MaxWildcardCollections: int(dataFederation.Storage.Databases[0].MaxWildcardCollections),
							Name:                   dataFederation.Storage.Databases[0].Name,
							Views: []atlasV1.View{
								{
									Name:     dataFederation.Storage.Databases[0].Views[0].Name,
									Pipeline: dataFederation.Storage.Databases[0].Views[0].Pipeline,
									Source:   dataFederation.Storage.Databases[0].Views[0].Source,
								},
							},
						},
					},
					Stores: []atlasV1.Store{
						{
							Name:                     dataFederation.Storage.Stores[0].Name,
							Provider:                 dataFederation.Storage.Stores[0].Provider,
							AdditionalStorageClasses: []string{*dataFederation.Storage.Stores[0].AdditionalStorageClasses[0]},
							Bucket:                   dataFederation.Storage.Stores[0].Bucket,
							Delimiter:                dataFederation.Storage.Stores[0].Delimiter,
							IncludeTags:              *dataFederation.Storage.Stores[0].IncludeTags,
							Prefix:                   dataFederation.Storage.Stores[0].Prefix,
							Public:                   *dataFederation.Storage.Stores[0].Public,
							Region:                   dataFederation.Storage.Stores[0].Region,
						},
					},
				},
			},
			Status: status.DataFederationStatus{
				Common: status.Common{
					Conditions: []status.Condition{},
				},
			},
		}

		got, err := BuildAtlasDataFederation(dataFederation, projectName, resourceVersion, targetNamespace, dictionary)
		if err != nil {
			t.Fatalf("%v", err)
		}

		if !reflect.DeepEqual(expected, got) {
			expJs, _ := json.MarshalIndent(expected, "", " ")
			gotJs, _ := json.MarshalIndent(got, "", " ")
			fmt.Printf("E:%s\r\n; G:%s\r\n", expJs, gotJs)
			t.Fatalf("Advanced data federation mismatch.\r\nexpected: %v\r\ngot: %v\r\n", expected, got)
		}
	})
}
