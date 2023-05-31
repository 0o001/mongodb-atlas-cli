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

// This code was autogenerated at 2023-04-12T16:00:40+01:00. Note: Manual updates are allowed, but may be overwritten.

package runs

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongodb-atlas-cli/internal/flag"
	mocks "github.com/mongodb/mongodb-atlas-cli/internal/mocks/atlas"
	"github.com/mongodb/mongodb-atlas-cli/internal/pointer"
	"github.com/mongodb/mongodb-atlas-cli/internal/test"
	atlasv2 "go.mongodb.org/atlas-sdk/admin"
)

func TestWatch_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockPipelineRunsDescriber(ctrl)

	opts := &WatchOpts{
		id:           "test",
		pipelineName: "pipeline1",
		store:        mockStore,
	}

	expected := &atlasv2.IngestionPipelineRun{
		Id:          pointer.Get("1a5cbd92c036a0eb288"),
		DatasetName: pointer.Get("pipeline 1"),
		State:       pointer.Get("DONE"),
	}

	mockStore.
		EXPECT().
		PipelineRun(opts.ProjectID, opts.pipelineName, opts.id).
		Return(expected, nil).
		Times(1)

	if err := opts.Run(); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
}

func TestWatcherBuilder(t *testing.T) {
	test.CmdValidator(
		t,
		WatchBuilder(),
		0,
		[]string{flag.ProjectID, flag.Output},
	)
}
