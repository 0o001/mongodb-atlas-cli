// Copyright 2023 MongoDB Inc
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

package store

import (
	atlas "go.mongodb.org/atlas/mongodbatlas"
)

type DataFederationStore interface {
	DataFederationLister
	// DataFederationCreator
	// DataFederationDeleter
	// DataFederationDescriber
	// DataFederationUpdater
}

type DataFederationLister interface {
	DataFederationList(string, string) ([]*atlas.DataFederationInstance, error)
}

// type DataFederationCreator interface {
// 	CreateDataFederation(string, *atlas.DataFederationInstance) (*atlas.DataFederationInstance, error)
// }

// type DataFederationDeleter interface {
// 	DeleteDataFederation(string, string) error
// }

// type DataFederationDescriber interface {
// 	DataFederation(string, string) (*atlas.DataFederationInstance, error)
// }

// type DataFederationUpdater interface {
// 	UpdateDataFederation(string, string, *atlas.DataFederationInstance) (*atlas.DataFederationInstance, error)
// }

// DataFederationList encapsulates the logic to manage different cloud providers.
func (s *Store) DataFederationList(projectID string, typeFlag string) ([]*atlas.DataFederationInstance, error) {
	result, _, err := s.client.(*atlas.Client).DataFederation.List(s.ctx, projectID)
	return result, err
}

// DataFederation encapsulates the logic to manage different cloud providers.
// func (s *Store) DataFederation(projectID, id string) (*atlas.DataFederationInstance, error) {
// 	result, _, err := s.client.(*atlas.Client).DataFederation.Get(s.ctx, projectID, id)
// 	return result, err
// }

// // CreateDataFederation encapsulates the logic to manage different cloud providers.
// func (s *Store) CreateDataFederation(projectID string, opts *atlas.DataFederationInstance) (*atlas.DataFederationInstance, error) {
// 	result, _, err := s.client.(*atlas.Client).DataFederation.Create(s.ctx, projectID, opts)
// 	return result, err
// }

// // DeleteDataFederation encapsulates the logic to manage different cloud providers.
// func (s *Store) DeleteDataFederation(projectID, id string) error {
// 	_, err := s.client.(*atlas.Client).DataFederation.Delete(s.ctx, projectID, id)
// 	return err
// }

// // UpdateDataFederation encapsulates the logic to manage different cloud providers.
// func (s *Store) UpdateDataFederation(projectID, id string, opts *atlas.DataFederationInstance) (*atlas.DataFederationInstance, error) {
// 	validation := &atlas.DataFederationUpdateOptions{SkipRoleValidation: false}
// 	result, _, err := s.client.(*atlas.Client).DataFederation.Update(s.ctx, projectID, id, opts, validation)
// 	return result, err
// }
