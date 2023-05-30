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
	client admin.APIClient
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
	resp, _, err := opts.client.SharedTierSnapshotsApi.DownloadSharedClusterBackupWithParams(ctx, params)
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

const DownloadSharedClusterBackupTemplate = "<<some template>>"

func DownloadSharedClusterBackupBuilder() cobra.Command {
	opts := DownloadSharedClusterBackupOpts{}
	cmd := &cobra.Command{
		Use:     "<<use>>",
		Short:   "<<decription>>",
		Long:    fmt.Sprintf(usage.RequiredRole, "Project Read Only"), // how to tell?
		// Aliases: []string{"ls"},
		Args:    require.NoArgs,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				//opts.ValidateProjectID,
				opts.initClient(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), DownloadSharedClusterBackupTemplate),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context())
		},
	}
	cmd.Flags().StringVar(&opts.clusterName, "clusterName", "", "usage description")
	_ = cmd.MarkFlagRequired("clusterName")
	cmd.Flags().StringVar(&opts.groupId, "groupId", "", "usage description")
	_ = cmd.MarkFlagRequired("groupId")

	return cmd
}
type GetSharedClusterBackupOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client admin.APIClient
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
	resp, _, err := opts.client.SharedTierSnapshotsApi.GetSharedClusterBackupWithParams(ctx, params)
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

const GetSharedClusterBackupTemplate = "<<some template>>"

func GetSharedClusterBackupBuilder() cobra.Command {
	opts := GetSharedClusterBackupOpts{}
	cmd := &cobra.Command{
		Use:     "<<use>>",
		Short:   "<<decription>>",
		Long:    fmt.Sprintf(usage.RequiredRole, "Project Read Only"), // how to tell?
		// Aliases: []string{"ls"},
		Args:    require.NoArgs,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				//opts.ValidateProjectID,
				opts.initClient(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), GetSharedClusterBackupTemplate),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "groupId", "", "usage description")
	_ = cmd.MarkFlagRequired("groupId")
	cmd.Flags().StringVar(&opts.clusterName, "clusterName", "", "usage description")
	_ = cmd.MarkFlagRequired("clusterName")
	cmd.Flags().StringVar(&opts.snapshotId, "snapshotId", "", "usage description")
	_ = cmd.MarkFlagRequired("snapshotId")

	return cmd
}
type ListSharedClusterBackupsOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client admin.APIClient
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
	resp, _, err := opts.client.SharedTierSnapshotsApi.ListSharedClusterBackupsWithParams(ctx, params)
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

const ListSharedClusterBackupsTemplate = "<<some template>>"

func ListSharedClusterBackupsBuilder() cobra.Command {
	opts := ListSharedClusterBackupsOpts{}
	cmd := &cobra.Command{
		Use:     "<<use>>",
		Short:   "<<decription>>",
		Long:    fmt.Sprintf(usage.RequiredRole, "Project Read Only"), // how to tell?
		// Aliases: []string{"ls"},
		Args:    require.NoArgs,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				//opts.ValidateProjectID,
				opts.initClient(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), ListSharedClusterBackupsTemplate),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "groupId", "", "usage description")
	_ = cmd.MarkFlagRequired("groupId")
	cmd.Flags().StringVar(&opts.clusterName, "clusterName", "", "usage description")
	_ = cmd.MarkFlagRequired("clusterName")

	return cmd
}
