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

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package store

import (
	"context"
	atlasv2 "go.mongodb.org/atlas/mongodbatlasv2"
)

//go:generate mockgen -destination=../../mocks/atlas/api_rolling_index_store_mock.go -package=atlas github.com/mongodb/mongodb-atlas-cli/internal/store/atlas CreateRollingIndexOperation

type CreateRollingIndexOperation interface {
	CreateRollingIndex (*atlasv2.CreateRollingIndexApiParams) error
}

// CreateRollingIndex encapsulates the logic to manage different cloud providers.
func (s *Store) CreateRollingIndex(params *atlasv2.CreateRollingIndexApiParams) (error) {
	_, err := s.clientv2.RollingIndexApi.CreateRollingIndexWithParams(s.ctx, params).Execute()
	return err
}

