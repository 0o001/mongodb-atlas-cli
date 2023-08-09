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

//go:build unit

// This code was autogenerated at 2023-06-22T17:46:28+01:00. Note: Manual updates are allowed, but may be overwritten.

package privateendpoints

import (
	"bytes"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/internal/flag"
	mocks "github.com/mongodb/mongodb-atlas-cli/internal/mocks/atlas"
	"github.com/mongodb/mongodb-atlas-cli/internal/pointer"
	"github.com/mongodb/mongodb-atlas-cli/internal/test"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/atlas-sdk/v20230201004/admin"
)

func TestCreate_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockDataFederationPrivateEndpointCreator(ctrl)

	buf := new(bytes.Buffer)
	expected := &admin.PaginatedPrivateNetworkEndpointIdEntry{
		TotalCount: pointer.Get(1),
		Results: []admin.PrivateNetworkEndpointIdEntry{
			{
				EndpointId: "id",
				Comment:    pointer.Get("comment"),
			},
		},
	}

	createOpts := &CreateOpts{
		store: mockStore,
		OutputOpts: cli.OutputOpts{
			Template:  createTemplate,
			OutWriter: buf,
		},
	}

	mockStore.
		EXPECT().
		CreateDataFederationPrivateEndpoint(createOpts.ConfigProjectID(), createOpts.newCreateRequest()).Return(expected, nil).
		Times(1)

	if err := createOpts.Run(); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}

	assert.Equal(t, `Data federation private endpoint id created.`, buf.String())
	t.Log(buf.String())
}

func TestCreateBuilder(t *testing.T) {
	test.CmdValidator(
		t,
		CreateBuilder(),
		0,
		[]string{flag.ProjectID, flag.Output},
	)
}
