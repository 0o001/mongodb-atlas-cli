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

type CreateOnlineArchiveOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client admin.APIClient
	groupId string
	clusterName string
}

func (opts *CreateOnlineArchiveOpts) initClient(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *CreateOnlineArchiveOpts) Run(ctx context.Context) error {
	params := &admin.CreateOnlineArchiveApiParams{
		GroupId: opts.groupId,
		ClusterName: opts.clusterName,
	}
	resp, _, err := opts.client.OnlineArchiveApi.CreateOnlineArchiveWithParams(ctx, params)
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

const CreateOnlineArchiveTemplate = "<<some template>>"

func CreateOnlineArchiveBuilder() cobra.Command {
	opts := CreateOnlineArchiveOpts{}
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
				opts.InitOutput(cmd.OutOrStdout(), CreateOnlineArchiveTemplate),
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
type DeleteOnlineArchiveOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client admin.APIClient
	groupId string
	archiveId string
	clusterName string
}

func (opts *DeleteOnlineArchiveOpts) initClient(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *DeleteOnlineArchiveOpts) Run(ctx context.Context) error {
	params := &admin.DeleteOnlineArchiveApiParams{
		GroupId: opts.groupId,
		ArchiveId: opts.archiveId,
		ClusterName: opts.clusterName,
	}
	resp, _, err := opts.client.OnlineArchiveApi.DeleteOnlineArchiveWithParams(ctx, params)
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

const DeleteOnlineArchiveTemplate = "<<some template>>"

func DeleteOnlineArchiveBuilder() cobra.Command {
	opts := DeleteOnlineArchiveOpts{}
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
				opts.InitOutput(cmd.OutOrStdout(), DeleteOnlineArchiveTemplate),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "groupId", "", "usage description")
	_ = cmd.MarkFlagRequired("groupId")
	cmd.Flags().StringVar(&opts.archiveId, "archiveId", "", "usage description")
	_ = cmd.MarkFlagRequired("archiveId")
	cmd.Flags().StringVar(&opts.clusterName, "clusterName", "", "usage description")
	_ = cmd.MarkFlagRequired("clusterName")

	return cmd
}
type DownloadOnlineArchiveQueryLogsOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client admin.APIClient
	groupId string
	clusterName string
	startDate int64
	endDate int64
	archiveOnly bool
}

func (opts *DownloadOnlineArchiveQueryLogsOpts) initClient(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *DownloadOnlineArchiveQueryLogsOpts) Run(ctx context.Context) error {
	params := &admin.DownloadOnlineArchiveQueryLogsApiParams{
		GroupId: opts.groupId,
		ClusterName: opts.clusterName,
		StartDate: opts.startDate,
		EndDate: opts.endDate,
		ArchiveOnly: opts.archiveOnly,
	}
	resp, _, err := opts.client.OnlineArchiveApi.DownloadOnlineArchiveQueryLogsWithParams(ctx, params)
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

const DownloadOnlineArchiveQueryLogsTemplate = "<<some template>>"

func DownloadOnlineArchiveQueryLogsBuilder() cobra.Command {
	opts := DownloadOnlineArchiveQueryLogsOpts{}
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
				opts.InitOutput(cmd.OutOrStdout(), DownloadOnlineArchiveQueryLogsTemplate),
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
	cmd.Flags().StringVar(&opts.startDate, "startDate", "", "usage description")
	cmd.Flags().StringVar(&opts.endDate, "endDate", "", "usage description")
	cmd.Flags().StringVar(&opts.archiveOnly, "archiveOnly", "", "usage description")

	return cmd
}
type GetOnlineArchiveOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client admin.APIClient
	groupId string
	archiveId string
	clusterName string
}

func (opts *GetOnlineArchiveOpts) initClient(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *GetOnlineArchiveOpts) Run(ctx context.Context) error {
	params := &admin.GetOnlineArchiveApiParams{
		GroupId: opts.groupId,
		ArchiveId: opts.archiveId,
		ClusterName: opts.clusterName,
	}
	resp, _, err := opts.client.OnlineArchiveApi.GetOnlineArchiveWithParams(ctx, params)
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

const GetOnlineArchiveTemplate = "<<some template>>"

func GetOnlineArchiveBuilder() cobra.Command {
	opts := GetOnlineArchiveOpts{}
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
				opts.InitOutput(cmd.OutOrStdout(), GetOnlineArchiveTemplate),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "groupId", "", "usage description")
	_ = cmd.MarkFlagRequired("groupId")
	cmd.Flags().StringVar(&opts.archiveId, "archiveId", "", "usage description")
	_ = cmd.MarkFlagRequired("archiveId")
	cmd.Flags().StringVar(&opts.clusterName, "clusterName", "", "usage description")
	_ = cmd.MarkFlagRequired("clusterName")

	return cmd
}
type ListOnlineArchivesOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client admin.APIClient
	groupId string
	clusterName string
	includeCount bool
	itemsPerPage int32
	pageNum int32
}

func (opts *ListOnlineArchivesOpts) initClient(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *ListOnlineArchivesOpts) Run(ctx context.Context) error {
	params := &admin.ListOnlineArchivesApiParams{
		GroupId: opts.groupId,
		ClusterName: opts.clusterName,
		IncludeCount: opts.includeCount,
		ItemsPerPage: opts.itemsPerPage,
		PageNum: opts.pageNum,
	}
	resp, _, err := opts.client.OnlineArchiveApi.ListOnlineArchivesWithParams(ctx, params)
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

const ListOnlineArchivesTemplate = "<<some template>>"

func ListOnlineArchivesBuilder() cobra.Command {
	opts := ListOnlineArchivesOpts{}
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
				opts.InitOutput(cmd.OutOrStdout(), ListOnlineArchivesTemplate),
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
	cmd.Flags().StringVar(&opts.includeCount, "includeCount", "", "usage description")
	cmd.Flags().StringVar(&opts.itemsPerPage, "itemsPerPage", "", "usage description")
	cmd.Flags().StringVar(&opts.pageNum, "pageNum", "", "usage description")

	return cmd
}
type UpdateOnlineArchiveOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client admin.APIClient
	groupId string
	archiveId string
	clusterName string
}

func (opts *UpdateOnlineArchiveOpts) initClient(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *UpdateOnlineArchiveOpts) Run(ctx context.Context) error {
	params := &admin.UpdateOnlineArchiveApiParams{
		GroupId: opts.groupId,
		ArchiveId: opts.archiveId,
		ClusterName: opts.clusterName,
	}
	resp, _, err := opts.client.OnlineArchiveApi.UpdateOnlineArchiveWithParams(ctx, params)
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

const UpdateOnlineArchiveTemplate = "<<some template>>"

func UpdateOnlineArchiveBuilder() cobra.Command {
	opts := UpdateOnlineArchiveOpts{}
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
				opts.InitOutput(cmd.OutOrStdout(), UpdateOnlineArchiveTemplate),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "groupId", "", "usage description")
	_ = cmd.MarkFlagRequired("groupId")
	cmd.Flags().StringVar(&opts.archiveId, "archiveId", "", "usage description")
	_ = cmd.MarkFlagRequired("archiveId")
	cmd.Flags().StringVar(&opts.clusterName, "clusterName", "", "usage description")
	_ = cmd.MarkFlagRequired("clusterName")

	return cmd
}
