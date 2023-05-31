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
	"github.com/spf13/cobra"
	"go.mongodb.org/atlas-sdk/admin"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
)

type DownloadSharedClusterBackupOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	clusterName string
	groupId string
}

func (opts *DownloadSharedClusterBackupOpts) initClient(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *DownloadSharedClusterBackupOpts) Run(ctx context.Context) error {
	params := &admin.DownloadSharedClusterBackupApiParams{
		ClusterName: opts.clusterName,
		GroupId: opts.groupId,
	}
	resp, _, err := opts.client.SharedTierSnapshotsApi.DownloadSharedClusterBackupWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func DownloadSharedClusterBackupBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := DownloadSharedClusterBackupOpts{}
	cmd := &cobra.Command{
		Use:     "<<use>>",
		// Aliases: []string{"?"},
		Short:   "Download One M2 or M5 Cluster Snapshot",
		Long:    fmt.Sprintf(usage.RequiredRole, "Project Read Only"), // how to tell?
		Args:    require.NoArgs,
		Annotations: map[string]string{
			"output":      template,
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				//opts.ValidateProjectID,
				opts.initClient(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), template),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context())
		},
	}
	cmd.Flags().StringVar(&opts.clusterName, "clusterName", , "Human-readable label that identifies the cluster.")	cmd.Flags().StringVar(&opts.groupId, "groupId", , "Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.")

	
	_ = cmd.MarkFlagRequired("clusterName")
	_ = cmd.MarkFlagRequired("groupId")

	return cmd
}
type GetSharedClusterBackupOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	groupId string
	clusterName string
	snapshotId string
}

func (opts *GetSharedClusterBackupOpts) initClient(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *GetSharedClusterBackupOpts) Run(ctx context.Context) error {
	params := &admin.GetSharedClusterBackupApiParams{
		GroupId: opts.groupId,
		ClusterName: opts.clusterName,
		SnapshotId: opts.snapshotId,
	}
	resp, _, err := opts.client.SharedTierSnapshotsApi.GetSharedClusterBackupWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func GetSharedClusterBackupBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := GetSharedClusterBackupOpts{}
	cmd := &cobra.Command{
		Use:     "<<use>>",
		// Aliases: []string{"?"},
		Short:   "Return One Snapshot for One M2 or M5 Cluster",
		Long:    fmt.Sprintf(usage.RequiredRole, "Project Read Only"), // how to tell?
		Args:    require.NoArgs,
		Annotations: map[string]string{
			"output":      template,
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				//opts.ValidateProjectID,
				opts.initClient(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), template),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "groupId", , "Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.")	cmd.Flags().StringVar(&opts.clusterName, "clusterName", , "Human-readable label that identifies the cluster.")	cmd.Flags().StringVar(&opts.snapshotId, "snapshotId", , "Unique 24-hexadecimal digit string that identifies the desired snapshot.")

	
	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("clusterName")
	_ = cmd.MarkFlagRequired("snapshotId")

	return cmd
}
type ListSharedClusterBackupsOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	groupId string
	clusterName string
}

func (opts *ListSharedClusterBackupsOpts) initClient(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *ListSharedClusterBackupsOpts) Run(ctx context.Context) error {
	params := &admin.ListSharedClusterBackupsApiParams{
		GroupId: opts.groupId,
		ClusterName: opts.clusterName,
	}
	resp, _, err := opts.client.SharedTierSnapshotsApi.ListSharedClusterBackupsWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func ListSharedClusterBackupsBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := ListSharedClusterBackupsOpts{}
	cmd := &cobra.Command{
		Use:     "<<use>>",
		// Aliases: []string{"?"},
		Short:   "Return All Snapshots for One M2 or M5 Cluster",
		Long:    fmt.Sprintf(usage.RequiredRole, "Project Read Only"), // how to tell?
		Args:    require.NoArgs,
		Annotations: map[string]string{
			"output":      template,
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				//opts.ValidateProjectID,
				opts.initClient(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), template),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "groupId", , "Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.")	cmd.Flags().StringVar(&opts.clusterName, "clusterName", , "Human-readable label that identifies the cluster.")

	
	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("clusterName")

	return cmd
}
