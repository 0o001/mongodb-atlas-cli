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

//go:generate mockgen -destination=../../mocks/atlas/api_cloud_provider_access_store_mock.go -package=atlas github.com/mongodb/mongodb-atlas-cli/internal/store/atlas AuthorizeCloudProviderAccessRoleOperation, CreateCloudProviderAccessRoleOperation, DeauthorizeCloudProviderAccessRoleOperation, GetCloudProviderAccessRoleOperation, ListCloudProviderAccessRolesOperation

type AuthorizeCloudProviderAccessRoleOperation interface {
	AuthorizeCloudProviderAccessRole (*atlasv2.AuthorizeCloudProviderAccessRoleApiParams) (*atlasv2.CloudProviderAccessRole, error)
}

// AuthorizeCloudProviderAccessRole encapsulates the logic to manage different cloud providers.
func (s *Store) AuthorizeCloudProviderAccessRole(params *atlasv2.AuthorizeCloudProviderAccessRoleApiParams) (*atlasv2.CloudProviderAccessRole, error) {
	result, _, err := s.clientv2.CloudProviderAccessApi.AuthorizeCloudProviderAccessRoleWithParams(s.ctx, params).Execute()
	return &result, err
}

type CreateCloudProviderAccessRoleOperation interface {
	CreateCloudProviderAccessRole (*atlasv2.CreateCloudProviderAccessRoleApiParams) (*atlasv2.CloudProviderAccessRole, error)
}

// CreateCloudProviderAccessRole encapsulates the logic to manage different cloud providers.
func (s *Store) CreateCloudProviderAccessRole(params *atlasv2.CreateCloudProviderAccessRoleApiParams) (*atlasv2.CloudProviderAccessRole, error) {
	result, _, err := s.clientv2.CloudProviderAccessApi.CreateCloudProviderAccessRoleWithParams(s.ctx, params).Execute()
	return &result, err
}

type DeauthorizeCloudProviderAccessRoleOperation interface {
	DeauthorizeCloudProviderAccessRole (*atlasv2.DeauthorizeCloudProviderAccessRoleApiParams) error
}

// DeauthorizeCloudProviderAccessRole encapsulates the logic to manage different cloud providers.
func (s *Store) DeauthorizeCloudProviderAccessRole(params *atlasv2.DeauthorizeCloudProviderAccessRoleApiParams) (error) {
	_, err := s.clientv2.CloudProviderAccessApi.DeauthorizeCloudProviderAccessRoleWithParams(s.ctx, params).Execute()
	return err
}

type GetCloudProviderAccessRoleOperation interface {
	GetCloudProviderAccessRole (*atlasv2.GetCloudProviderAccessRoleApiParams) (*atlasv2.CloudProviderAccess, error)
}

// GetCloudProviderAccessRole encapsulates the logic to manage different cloud providers.
func (s *Store) GetCloudProviderAccessRole(params *atlasv2.GetCloudProviderAccessRoleApiParams) (*atlasv2.CloudProviderAccess, error) {
	result, _, err := s.clientv2.CloudProviderAccessApi.GetCloudProviderAccessRoleWithParams(s.ctx, params).Execute()
	return &result, err
}

type ListCloudProviderAccessRolesOperation interface {
	ListCloudProviderAccessRoles (*atlasv2.ListCloudProviderAccessRolesApiParams) (*atlasv2.CloudProviderAccess, error)
}

// ListCloudProviderAccessRoles encapsulates the logic to manage different cloud providers.
func (s *Store) ListCloudProviderAccessRoles(params *atlasv2.ListCloudProviderAccessRolesApiParams) (*atlasv2.CloudProviderAccess, error) {
	result, _, err := s.clientv2.CloudProviderAccessApi.ListCloudProviderAccessRolesWithParams(s.ctx, params).Execute()
	return &result, err
}

