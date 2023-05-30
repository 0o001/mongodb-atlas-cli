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

type CreatePipelineOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client admin.APIClient
	groupId string
}

func (opts *CreatePipelineOpts) initClient(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *CreatePipelineOpts) Run(ctx context.Context) error {
	params := &admin.CreatePipelineApiParams{
		GroupId: opts.groupId,
	}
	resp, _, err := opts.client.DataLakePipelinesApi.CreatePipelineWithParams(ctx, params)
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

const CreatePipelineTemplate = "<<some template>>"

func CreatePipelineBuilder() cobra.Command {
	opts := CreatePipelineOpts{}
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
				opts.InitOutput(cmd.OutOrStdout(), CreatePipelineTemplate),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "groupId", "", "usage description")
	_ = cmd.MarkFlagRequired("groupId")

	return cmd
}
type DeletePipelineOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client admin.APIClient
	groupId string
	pipelineName string
}

func (opts *DeletePipelineOpts) initClient(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *DeletePipelineOpts) Run(ctx context.Context) error {
	params := &admin.DeletePipelineApiParams{
		GroupId: opts.groupId,
		PipelineName: opts.pipelineName,
	}
	resp, _, err := opts.client.DataLakePipelinesApi.DeletePipelineWithParams(ctx, params)
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

const DeletePipelineTemplate = "<<some template>>"

func DeletePipelineBuilder() cobra.Command {
	opts := DeletePipelineOpts{}
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
				opts.InitOutput(cmd.OutOrStdout(), DeletePipelineTemplate),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "groupId", "", "usage description")
	_ = cmd.MarkFlagRequired("groupId")
	cmd.Flags().StringVar(&opts.pipelineName, "pipelineName", "", "usage description")
	_ = cmd.MarkFlagRequired("pipelineName")

	return cmd
}
type DeletePipelineRunDatasetOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client admin.APIClient
	groupId string
	pipelineName string
	pipelineRunId string
}

func (opts *DeletePipelineRunDatasetOpts) initClient(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *DeletePipelineRunDatasetOpts) Run(ctx context.Context) error {
	params := &admin.DeletePipelineRunDatasetApiParams{
		GroupId: opts.groupId,
		PipelineName: opts.pipelineName,
		PipelineRunId: opts.pipelineRunId,
	}
	resp, _, err := opts.client.DataLakePipelinesApi.DeletePipelineRunDatasetWithParams(ctx, params)
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

const DeletePipelineRunDatasetTemplate = "<<some template>>"

func DeletePipelineRunDatasetBuilder() cobra.Command {
	opts := DeletePipelineRunDatasetOpts{}
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
				opts.InitOutput(cmd.OutOrStdout(), DeletePipelineRunDatasetTemplate),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "groupId", "", "usage description")
	_ = cmd.MarkFlagRequired("groupId")
	cmd.Flags().StringVar(&opts.pipelineName, "pipelineName", "", "usage description")
	_ = cmd.MarkFlagRequired("pipelineName")
	cmd.Flags().StringVar(&opts.pipelineRunId, "pipelineRunId", "", "usage description")
	_ = cmd.MarkFlagRequired("pipelineRunId")

	return cmd
}
type GetPipelineOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client admin.APIClient
	groupId string
	pipelineName string
}

func (opts *GetPipelineOpts) initClient(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *GetPipelineOpts) Run(ctx context.Context) error {
	params := &admin.GetPipelineApiParams{
		GroupId: opts.groupId,
		PipelineName: opts.pipelineName,
	}
	resp, _, err := opts.client.DataLakePipelinesApi.GetPipelineWithParams(ctx, params)
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

const GetPipelineTemplate = "<<some template>>"

func GetPipelineBuilder() cobra.Command {
	opts := GetPipelineOpts{}
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
				opts.InitOutput(cmd.OutOrStdout(), GetPipelineTemplate),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "groupId", "", "usage description")
	_ = cmd.MarkFlagRequired("groupId")
	cmd.Flags().StringVar(&opts.pipelineName, "pipelineName", "", "usage description")
	_ = cmd.MarkFlagRequired("pipelineName")

	return cmd
}
type GetPipelineRunOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client admin.APIClient
	groupId string
	pipelineName string
	pipelineRunId string
}

func (opts *GetPipelineRunOpts) initClient(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *GetPipelineRunOpts) Run(ctx context.Context) error {
	params := &admin.GetPipelineRunApiParams{
		GroupId: opts.groupId,
		PipelineName: opts.pipelineName,
		PipelineRunId: opts.pipelineRunId,
	}
	resp, _, err := opts.client.DataLakePipelinesApi.GetPipelineRunWithParams(ctx, params)
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

const GetPipelineRunTemplate = "<<some template>>"

func GetPipelineRunBuilder() cobra.Command {
	opts := GetPipelineRunOpts{}
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
				opts.InitOutput(cmd.OutOrStdout(), GetPipelineRunTemplate),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "groupId", "", "usage description")
	_ = cmd.MarkFlagRequired("groupId")
	cmd.Flags().StringVar(&opts.pipelineName, "pipelineName", "", "usage description")
	_ = cmd.MarkFlagRequired("pipelineName")
	cmd.Flags().StringVar(&opts.pipelineRunId, "pipelineRunId", "", "usage description")
	_ = cmd.MarkFlagRequired("pipelineRunId")

	return cmd
}
type ListPipelineRunsOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client admin.APIClient
	groupId string
	pipelineName string
	includeCount bool
	itemsPerPage int32
	pageNum int32
	createdBefore time.Time
}

func (opts *ListPipelineRunsOpts) initClient(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *ListPipelineRunsOpts) Run(ctx context.Context) error {
	params := &admin.ListPipelineRunsApiParams{
		GroupId: opts.groupId,
		PipelineName: opts.pipelineName,
		IncludeCount: opts.includeCount,
		ItemsPerPage: opts.itemsPerPage,
		PageNum: opts.pageNum,
		CreatedBefore: opts.createdBefore,
	}
	resp, _, err := opts.client.DataLakePipelinesApi.ListPipelineRunsWithParams(ctx, params)
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

const ListPipelineRunsTemplate = "<<some template>>"

func ListPipelineRunsBuilder() cobra.Command {
	opts := ListPipelineRunsOpts{}
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
				opts.InitOutput(cmd.OutOrStdout(), ListPipelineRunsTemplate),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "groupId", "", "usage description")
	_ = cmd.MarkFlagRequired("groupId")
	cmd.Flags().StringVar(&opts.pipelineName, "pipelineName", "", "usage description")
	_ = cmd.MarkFlagRequired("pipelineName")
	cmd.Flags().StringVar(&opts.includeCount, "includeCount", "", "usage description")
	cmd.Flags().StringVar(&opts.itemsPerPage, "itemsPerPage", "", "usage description")
	cmd.Flags().StringVar(&opts.pageNum, "pageNum", "", "usage description")
	cmd.Flags().StringVar(&opts.createdBefore, "createdBefore", "", "usage description")

	return cmd
}
type ListPipelineSchedulesOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client admin.APIClient
	groupId string
	pipelineName string
}

func (opts *ListPipelineSchedulesOpts) initClient(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *ListPipelineSchedulesOpts) Run(ctx context.Context) error {
	params := &admin.ListPipelineSchedulesApiParams{
		GroupId: opts.groupId,
		PipelineName: opts.pipelineName,
	}
	resp, _, err := opts.client.DataLakePipelinesApi.ListPipelineSchedulesWithParams(ctx, params)
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

const ListPipelineSchedulesTemplate = "<<some template>>"

func ListPipelineSchedulesBuilder() cobra.Command {
	opts := ListPipelineSchedulesOpts{}
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
				opts.InitOutput(cmd.OutOrStdout(), ListPipelineSchedulesTemplate),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "groupId", "", "usage description")
	_ = cmd.MarkFlagRequired("groupId")
	cmd.Flags().StringVar(&opts.pipelineName, "pipelineName", "", "usage description")
	_ = cmd.MarkFlagRequired("pipelineName")

	return cmd
}
type ListPipelineSnapshotsOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client admin.APIClient
	groupId string
	pipelineName string
	includeCount bool
	itemsPerPage int32
	pageNum int32
	completedAfter time.Time
}

func (opts *ListPipelineSnapshotsOpts) initClient(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *ListPipelineSnapshotsOpts) Run(ctx context.Context) error {
	params := &admin.ListPipelineSnapshotsApiParams{
		GroupId: opts.groupId,
		PipelineName: opts.pipelineName,
		IncludeCount: opts.includeCount,
		ItemsPerPage: opts.itemsPerPage,
		PageNum: opts.pageNum,
		CompletedAfter: opts.completedAfter,
	}
	resp, _, err := opts.client.DataLakePipelinesApi.ListPipelineSnapshotsWithParams(ctx, params)
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

const ListPipelineSnapshotsTemplate = "<<some template>>"

func ListPipelineSnapshotsBuilder() cobra.Command {
	opts := ListPipelineSnapshotsOpts{}
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
				opts.InitOutput(cmd.OutOrStdout(), ListPipelineSnapshotsTemplate),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "groupId", "", "usage description")
	_ = cmd.MarkFlagRequired("groupId")
	cmd.Flags().StringVar(&opts.pipelineName, "pipelineName", "", "usage description")
	_ = cmd.MarkFlagRequired("pipelineName")
	cmd.Flags().StringVar(&opts.includeCount, "includeCount", "", "usage description")
	cmd.Flags().StringVar(&opts.itemsPerPage, "itemsPerPage", "", "usage description")
	cmd.Flags().StringVar(&opts.pageNum, "pageNum", "", "usage description")
	cmd.Flags().StringVar(&opts.completedAfter, "completedAfter", "", "usage description")

	return cmd
}
type ListPipelinesOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client admin.APIClient
	groupId string
}

func (opts *ListPipelinesOpts) initClient(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *ListPipelinesOpts) Run(ctx context.Context) error {
	params := &admin.ListPipelinesApiParams{
		GroupId: opts.groupId,
	}
	resp, _, err := opts.client.DataLakePipelinesApi.ListPipelinesWithParams(ctx, params)
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

const ListPipelinesTemplate = "<<some template>>"

func ListPipelinesBuilder() cobra.Command {
	opts := ListPipelinesOpts{}
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
				opts.InitOutput(cmd.OutOrStdout(), ListPipelinesTemplate),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "groupId", "", "usage description")
	_ = cmd.MarkFlagRequired("groupId")

	return cmd
}
type PausePipelineOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client admin.APIClient
	groupId string
	pipelineName string
}

func (opts *PausePipelineOpts) initClient(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *PausePipelineOpts) Run(ctx context.Context) error {
	params := &admin.PausePipelineApiParams{
		GroupId: opts.groupId,
		PipelineName: opts.pipelineName,
	}
	resp, _, err := opts.client.DataLakePipelinesApi.PausePipelineWithParams(ctx, params)
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

const PausePipelineTemplate = "<<some template>>"

func PausePipelineBuilder() cobra.Command {
	opts := PausePipelineOpts{}
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
				opts.InitOutput(cmd.OutOrStdout(), PausePipelineTemplate),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "groupId", "", "usage description")
	_ = cmd.MarkFlagRequired("groupId")
	cmd.Flags().StringVar(&opts.pipelineName, "pipelineName", "", "usage description")
	_ = cmd.MarkFlagRequired("pipelineName")

	return cmd
}
type ResumePipelineOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client admin.APIClient
	groupId string
	pipelineName string
}

func (opts *ResumePipelineOpts) initClient(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *ResumePipelineOpts) Run(ctx context.Context) error {
	params := &admin.ResumePipelineApiParams{
		GroupId: opts.groupId,
		PipelineName: opts.pipelineName,
	}
	resp, _, err := opts.client.DataLakePipelinesApi.ResumePipelineWithParams(ctx, params)
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

const ResumePipelineTemplate = "<<some template>>"

func ResumePipelineBuilder() cobra.Command {
	opts := ResumePipelineOpts{}
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
				opts.InitOutput(cmd.OutOrStdout(), ResumePipelineTemplate),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "groupId", "", "usage description")
	_ = cmd.MarkFlagRequired("groupId")
	cmd.Flags().StringVar(&opts.pipelineName, "pipelineName", "", "usage description")
	_ = cmd.MarkFlagRequired("pipelineName")

	return cmd
}
type TriggerSnapshotIngestionOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client admin.APIClient
	groupId string
	pipelineName string
}

func (opts *TriggerSnapshotIngestionOpts) initClient(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *TriggerSnapshotIngestionOpts) Run(ctx context.Context) error {
	params := &admin.TriggerSnapshotIngestionApiParams{
		GroupId: opts.groupId,
		PipelineName: opts.pipelineName,
	}
	resp, _, err := opts.client.DataLakePipelinesApi.TriggerSnapshotIngestionWithParams(ctx, params)
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

const TriggerSnapshotIngestionTemplate = "<<some template>>"

func TriggerSnapshotIngestionBuilder() cobra.Command {
	opts := TriggerSnapshotIngestionOpts{}
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
				opts.InitOutput(cmd.OutOrStdout(), TriggerSnapshotIngestionTemplate),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "groupId", "", "usage description")
	_ = cmd.MarkFlagRequired("groupId")
	cmd.Flags().StringVar(&opts.pipelineName, "pipelineName", "", "usage description")
	_ = cmd.MarkFlagRequired("pipelineName")

	return cmd
}
type UpdatePipelineOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client admin.APIClient
	groupId string
	pipelineName string
}

func (opts *UpdatePipelineOpts) initClient(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *UpdatePipelineOpts) Run(ctx context.Context) error {
	params := &admin.UpdatePipelineApiParams{
		GroupId: opts.groupId,
		PipelineName: opts.pipelineName,
	}
	resp, _, err := opts.client.DataLakePipelinesApi.UpdatePipelineWithParams(ctx, params)
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

const UpdatePipelineTemplate = "<<some template>>"

func UpdatePipelineBuilder() cobra.Command {
	opts := UpdatePipelineOpts{}
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
				opts.InitOutput(cmd.OutOrStdout(), UpdatePipelineTemplate),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "groupId", "", "usage description")
	_ = cmd.MarkFlagRequired("groupId")
	cmd.Flags().StringVar(&opts.pipelineName, "pipelineName", "", "usage description")
	_ = cmd.MarkFlagRequired("pipelineName")

	return cmd
}
