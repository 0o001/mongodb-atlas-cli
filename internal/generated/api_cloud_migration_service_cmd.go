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

type CreateLinkTokenOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client admin.APIClient
	orgId string
}

func (opts *CreateLinkTokenOpts) initClient(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *CreateLinkTokenOpts) Run(ctx context.Context) error {
	params := &admin.CreateLinkTokenApiParams{
		OrgId: opts.orgId,
	}
	resp, _, err := opts.client.CloudMigrationServiceApi.CreateLinkTokenWithParams(ctx, params)
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

const CreateLinkTokenTemplate = "<<some template>>"

func CreateLinkTokenBuilder() cobra.Command {
	opts := CreateLinkTokenOpts{}
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
				opts.InitOutput(cmd.OutOrStdout(), CreateLinkTokenTemplate),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context())
		},
	}
	cmd.Flags().StringVar(&opts.orgId, "orgId", "", "usage description")
	_ = cmd.MarkFlagRequired("orgId")

	return cmd
}
type CreatePushMigrationOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client admin.APIClient
	groupId string
}

func (opts *CreatePushMigrationOpts) initClient(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *CreatePushMigrationOpts) Run(ctx context.Context) error {
	params := &admin.CreatePushMigrationApiParams{
		GroupId: opts.groupId,
	}
	resp, _, err := opts.client.CloudMigrationServiceApi.CreatePushMigrationWithParams(ctx, params)
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

const CreatePushMigrationTemplate = "<<some template>>"

func CreatePushMigrationBuilder() cobra.Command {
	opts := CreatePushMigrationOpts{}
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
				opts.InitOutput(cmd.OutOrStdout(), CreatePushMigrationTemplate),
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
type CutoverMigrationOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client admin.APIClient
	groupId string
	liveMigrationId string
}

func (opts *CutoverMigrationOpts) initClient(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *CutoverMigrationOpts) Run(ctx context.Context) error {
	params := &admin.CutoverMigrationApiParams{
		GroupId: opts.groupId,
		LiveMigrationId: opts.liveMigrationId,
	}
	_, err := opts.client.CloudMigrationServiceApi.CutoverMigrationWithParams(ctx, params)
	if err != nil {
		return err
	}

	return opts.Print(nil)
}

const CutoverMigrationTemplate = "<<some template>>"

func CutoverMigrationBuilder() cobra.Command {
	opts := CutoverMigrationOpts{}
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
				opts.InitOutput(cmd.OutOrStdout(), CutoverMigrationTemplate),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "groupId", "", "usage description")
	_ = cmd.MarkFlagRequired("groupId")
	cmd.Flags().StringVar(&opts.liveMigrationId, "liveMigrationId", "", "usage description")
	_ = cmd.MarkFlagRequired("liveMigrationId")

	return cmd
}
type DeleteLinkTokenOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client admin.APIClient
	orgId string
}

func (opts *DeleteLinkTokenOpts) initClient(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *DeleteLinkTokenOpts) Run(ctx context.Context) error {
	params := &admin.DeleteLinkTokenApiParams{
		OrgId: opts.orgId,
	}
	resp, _, err := opts.client.CloudMigrationServiceApi.DeleteLinkTokenWithParams(ctx, params)
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

const DeleteLinkTokenTemplate = "<<some template>>"

func DeleteLinkTokenBuilder() cobra.Command {
	opts := DeleteLinkTokenOpts{}
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
				opts.InitOutput(cmd.OutOrStdout(), DeleteLinkTokenTemplate),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context())
		},
	}
	cmd.Flags().StringVar(&opts.orgId, "orgId", "", "usage description")
	_ = cmd.MarkFlagRequired("orgId")

	return cmd
}
type GetPushMigrationOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client admin.APIClient
	groupId string
	liveMigrationId string
}

func (opts *GetPushMigrationOpts) initClient(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *GetPushMigrationOpts) Run(ctx context.Context) error {
	params := &admin.GetPushMigrationApiParams{
		GroupId: opts.groupId,
		LiveMigrationId: opts.liveMigrationId,
	}
	resp, _, err := opts.client.CloudMigrationServiceApi.GetPushMigrationWithParams(ctx, params)
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

const GetPushMigrationTemplate = "<<some template>>"

func GetPushMigrationBuilder() cobra.Command {
	opts := GetPushMigrationOpts{}
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
				opts.InitOutput(cmd.OutOrStdout(), GetPushMigrationTemplate),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "groupId", "", "usage description")
	_ = cmd.MarkFlagRequired("groupId")
	cmd.Flags().StringVar(&opts.liveMigrationId, "liveMigrationId", "", "usage description")
	_ = cmd.MarkFlagRequired("liveMigrationId")

	return cmd
}
type GetValidationStatusOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client admin.APIClient
	groupId string
	validationId string
}

func (opts *GetValidationStatusOpts) initClient(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *GetValidationStatusOpts) Run(ctx context.Context) error {
	params := &admin.GetValidationStatusApiParams{
		GroupId: opts.groupId,
		ValidationId: opts.validationId,
	}
	resp, _, err := opts.client.CloudMigrationServiceApi.GetValidationStatusWithParams(ctx, params)
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

const GetValidationStatusTemplate = "<<some template>>"

func GetValidationStatusBuilder() cobra.Command {
	opts := GetValidationStatusOpts{}
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
				opts.InitOutput(cmd.OutOrStdout(), GetValidationStatusTemplate),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "groupId", "", "usage description")
	_ = cmd.MarkFlagRequired("groupId")
	cmd.Flags().StringVar(&opts.validationId, "validationId", "", "usage description")
	_ = cmd.MarkFlagRequired("validationId")

	return cmd
}
type ListSourceProjectsOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client admin.APIClient
	orgId string
}

func (opts *ListSourceProjectsOpts) initClient(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *ListSourceProjectsOpts) Run(ctx context.Context) error {
	params := &admin.ListSourceProjectsApiParams{
		OrgId: opts.orgId,
	}
	resp, _, err := opts.client.CloudMigrationServiceApi.ListSourceProjectsWithParams(ctx, params)
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

const ListSourceProjectsTemplate = "<<some template>>"

func ListSourceProjectsBuilder() cobra.Command {
	opts := ListSourceProjectsOpts{}
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
				opts.InitOutput(cmd.OutOrStdout(), ListSourceProjectsTemplate),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context())
		},
	}
	cmd.Flags().StringVar(&opts.orgId, "orgId", "", "usage description")
	_ = cmd.MarkFlagRequired("orgId")

	return cmd
}
type ValidateMigrationOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client admin.APIClient
	groupId string
}

func (opts *ValidateMigrationOpts) initClient(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *ValidateMigrationOpts) Run(ctx context.Context) error {
	params := &admin.ValidateMigrationApiParams{
		GroupId: opts.groupId,
	}
	resp, _, err := opts.client.CloudMigrationServiceApi.ValidateMigrationWithParams(ctx, params)
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

const ValidateMigrationTemplate = "<<some template>>"

func ValidateMigrationBuilder() cobra.Command {
	opts := ValidateMigrationOpts{}
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
				opts.InitOutput(cmd.OutOrStdout(), ValidateMigrationTemplate),
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
