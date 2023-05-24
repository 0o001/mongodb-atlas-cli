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

//go:generate mockgen -destination=../../mocks/atlas/api_cloud_migration_service_store_mock.go -package=atlas github.com/mongodb/mongodb-atlas-cli/internal/store/atlas CreateLinkTokenOperation, CreatePushMigrationOperation, CutoverMigrationOperation, DeleteLinkTokenOperation, GetPushMigrationOperation, GetValidationStatusOperation, ListSourceProjectsOperation, ValidateMigrationOperation

type CreateLinkTokenOperation interface {
	CreateLinkToken (*atlasv2.CreateLinkTokenApiParams) (*atlasv2.TargetOrg, error)
}

// CreateLinkToken encapsulates the logic to manage different cloud providers.
func (s *Store) CreateLinkToken(params *atlasv2.CreateLinkTokenApiParams) (*atlasv2.TargetOrg, error) {
	result, _, err := s.clientv2.CloudMigrationServiceApi.CreateLinkTokenWithParams(s.ctx, params).Execute()
	return &result, err
}

type CreatePushMigrationOperation interface {
	CreatePushMigration (*atlasv2.CreatePushMigrationApiParams) (*atlasv2.LiveMigrationResponse, error)
}

// CreatePushMigration encapsulates the logic to manage different cloud providers.
func (s *Store) CreatePushMigration(params *atlasv2.CreatePushMigrationApiParams) (*atlasv2.LiveMigrationResponse, error) {
	result, _, err := s.clientv2.CloudMigrationServiceApi.CreatePushMigrationWithParams(s.ctx, params).Execute()
	return &result, err
}

type CutoverMigrationOperation interface {
	CutoverMigration (*atlasv2.CutoverMigrationApiParams) error
}

// CutoverMigration encapsulates the logic to manage different cloud providers.
func (s *Store) CutoverMigration(params *atlasv2.CutoverMigrationApiParams) (error) {
	_, err := s.clientv2.CloudMigrationServiceApi.CutoverMigrationWithParams(s.ctx, params).Execute()
	return err
}

type DeleteLinkTokenOperation interface {
	DeleteLinkToken (*atlasv2.DeleteLinkTokenApiParams) (*atlasv2.map[string]interface{}, error)
}

// DeleteLinkToken encapsulates the logic to manage different cloud providers.
func (s *Store) DeleteLinkToken(params *atlasv2.DeleteLinkTokenApiParams) (*atlasv2.map[string]interface{}, error) {
	result, _, err := s.clientv2.CloudMigrationServiceApi.DeleteLinkTokenWithParams(s.ctx, params).Execute()
	return &result, err
}

type GetPushMigrationOperation interface {
	GetPushMigration (*atlasv2.GetPushMigrationApiParams) (*atlasv2.LiveMigrationResponse, error)
}

// GetPushMigration encapsulates the logic to manage different cloud providers.
func (s *Store) GetPushMigration(params *atlasv2.GetPushMigrationApiParams) (*atlasv2.LiveMigrationResponse, error) {
	result, _, err := s.clientv2.CloudMigrationServiceApi.GetPushMigrationWithParams(s.ctx, params).Execute()
	return &result, err
}

type GetValidationStatusOperation interface {
	GetValidationStatus (*atlasv2.GetValidationStatusApiParams) (*atlasv2.Validation, error)
}

// GetValidationStatus encapsulates the logic to manage different cloud providers.
func (s *Store) GetValidationStatus(params *atlasv2.GetValidationStatusApiParams) (*atlasv2.Validation, error) {
	result, _, err := s.clientv2.CloudMigrationServiceApi.GetValidationStatusWithParams(s.ctx, params).Execute()
	return &result, err
}

type ListSourceProjectsOperation interface {
	ListSourceProjects (*atlasv2.ListSourceProjectsApiParams) (*atlasv2.[]AvailableProject, error)
}

// ListSourceProjects encapsulates the logic to manage different cloud providers.
func (s *Store) ListSourceProjects(params *atlasv2.ListSourceProjectsApiParams) (*atlasv2.[]AvailableProject, error) {
	result, _, err := s.clientv2.CloudMigrationServiceApi.ListSourceProjectsWithParams(s.ctx, params).Execute()
	return &result, err
}

type ValidateMigrationOperation interface {
	ValidateMigration (*atlasv2.ValidateMigrationApiParams) (*atlasv2.Validation, error)
}

// ValidateMigration encapsulates the logic to manage different cloud providers.
func (s *Store) ValidateMigration(params *atlasv2.ValidateMigrationApiParams) (*atlasv2.Validation, error) {
	result, _, err := s.clientv2.CloudMigrationServiceApi.ValidateMigrationWithParams(s.ctx, params).Execute()
	return &result, err
}

