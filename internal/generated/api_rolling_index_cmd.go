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

package generated

import (
	"context"
	"os"
	"time"

	"github.com/spf13/cobra"
	"go.mongodb.org/atlas-sdk/admin"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
)

type createRollingIndexOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	groupId string
	clusterName string
	
}

func (opts *createRollingIndexOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = newClientWithAuth()
		return err
	}
}

func (opts *createRollingIndexOpts) Run(ctx context.Context) error {
	params := &admin.CreateRollingIndexApiParams{
		GroupId: opts.groupId,
		ClusterName: opts.clusterName,
		
	}
	_, err := opts.client.RollingIndexApi.CreateRollingIndexWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(nil)
}

func createRollingIndexBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := createRollingIndexOpts{}
	cmd := &cobra.Command{
		Use: "createRollingIndex",
		Short: "Create One Rolling Index",
		Annotations: map[string]string{
			"output":      template,
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.initClient(),
				opts.InitOutput(cmd.OutOrStdout(), template),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "groupId", "", `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`)
	cmd.Flags().StringVar(&opts.clusterName, "clusterName", "", `Human-readable label that identifies the cluster on which MongoDB Cloud creates an index.`)
	

	cmd.Flags().CollationVar(&opts.collation, "collation", , ``)

	cmd.Flags().StringVar(&opts.collection, "collection", "", `Human-readable label of the collection for which MongoDB Cloud creates an index.`)

	cmd.Flags().StringVar(&opts.db, "db", "", `Human-readable label of the database that holds the collection on which MongoDB Cloud creates an index.`)

	cmd.Flags().ArraySliceVar(&opts.keys, "keys", nil, `List that contains one or more objects that describe the parameters that you want to index.`)

	cmd.Flags().IndexOptionsVar(&opts.options, "options", , ``)


	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("clusterName")
	return cmd
}

func rollingIndexBuilder() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "rollingIndex",
		Short:   `Creates one index to a database deployment in a rolling manner. You can&#39;t create a rolling index on an &#x60;M0&#x60; free cluster or &#x60;M2/M5&#x60; shared cluster.`,
	}
	cmd.AddCommand(
		createRollingIndexBuilder(),
	)
	return cmd
}

