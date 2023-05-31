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

type CreateCustomDatabaseRoleOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	groupId string
}

func (opts *CreateCustomDatabaseRoleOpts) initClient(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *CreateCustomDatabaseRoleOpts) Run(ctx context.Context) error {
	params := &admin.CreateCustomDatabaseRoleApiParams{
		GroupId: opts.groupId,
	}
	resp, _, err := opts.client.CustomDatabaseRolesApi.CreateCustomDatabaseRoleWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func CreateCustomDatabaseRoleBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := CreateCustomDatabaseRoleOpts{}
	cmd := &cobra.Command{
		Use:     "<<use>>",
		// Aliases: []string{"?"},
		Short:   "Create One Custom Role",
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
	cmd.Flags().StringVar(&opts.groupId, "groupId", , "Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.")

	
	_ = cmd.MarkFlagRequired("groupId")

	return cmd
}
type DeleteCustomDatabaseRoleOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	groupId string
	roleName string
}

func (opts *DeleteCustomDatabaseRoleOpts) initClient(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *DeleteCustomDatabaseRoleOpts) Run(ctx context.Context) error {
	params := &admin.DeleteCustomDatabaseRoleApiParams{
		GroupId: opts.groupId,
		RoleName: opts.roleName,
	}
	_, err := opts.client.CustomDatabaseRolesApi.DeleteCustomDatabaseRoleWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(nil)
}

func DeleteCustomDatabaseRoleBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := DeleteCustomDatabaseRoleOpts{}
	cmd := &cobra.Command{
		Use:     "<<use>>",
		// Aliases: []string{"?"},
		Short:   "Remove One Custom Role from One Project",
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
	cmd.Flags().StringVar(&opts.groupId, "groupId", , "Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.")	cmd.Flags().StringVar(&opts.roleName, "roleName", , "Human-readable label that identifies the role for the request. This name must be unique for this custom role in this project.")

	
	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("roleName")

	return cmd
}
type GetCustomDatabaseRoleOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	groupId string
	roleName string
}

func (opts *GetCustomDatabaseRoleOpts) initClient(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *GetCustomDatabaseRoleOpts) Run(ctx context.Context) error {
	params := &admin.GetCustomDatabaseRoleApiParams{
		GroupId: opts.groupId,
		RoleName: opts.roleName,
	}
	resp, _, err := opts.client.CustomDatabaseRolesApi.GetCustomDatabaseRoleWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func GetCustomDatabaseRoleBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := GetCustomDatabaseRoleOpts{}
	cmd := &cobra.Command{
		Use:     "<<use>>",
		// Aliases: []string{"?"},
		Short:   "Return One Custom Role in One Project",
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
	cmd.Flags().StringVar(&opts.groupId, "groupId", , "Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.")	cmd.Flags().StringVar(&opts.roleName, "roleName", , "Human-readable label that identifies the role for the request. This name must be unique for this custom role in this project.")

	
	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("roleName")

	return cmd
}
type ListCustomDatabaseRolesOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	groupId string
}

func (opts *ListCustomDatabaseRolesOpts) initClient(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *ListCustomDatabaseRolesOpts) Run(ctx context.Context) error {
	params := &admin.ListCustomDatabaseRolesApiParams{
		GroupId: opts.groupId,
	}
	resp, _, err := opts.client.CustomDatabaseRolesApi.ListCustomDatabaseRolesWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func ListCustomDatabaseRolesBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := ListCustomDatabaseRolesOpts{}
	cmd := &cobra.Command{
		Use:     "<<use>>",
		// Aliases: []string{"?"},
		Short:   "Return All Custom Roles in One Project",
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
	cmd.Flags().StringVar(&opts.groupId, "groupId", , "Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.")

	
	_ = cmd.MarkFlagRequired("groupId")

	return cmd
}
type UpdateCustomDatabaseRoleOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	groupId string
	roleName string
}

func (opts *UpdateCustomDatabaseRoleOpts) initClient(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *UpdateCustomDatabaseRoleOpts) Run(ctx context.Context) error {
	params := &admin.UpdateCustomDatabaseRoleApiParams{
		GroupId: opts.groupId,
		RoleName: opts.roleName,
	}
	resp, _, err := opts.client.CustomDatabaseRolesApi.UpdateCustomDatabaseRoleWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func UpdateCustomDatabaseRoleBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := UpdateCustomDatabaseRoleOpts{}
	cmd := &cobra.Command{
		Use:     "<<use>>",
		// Aliases: []string{"?"},
		Short:   "Update One Custom Role in One Project",
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
	cmd.Flags().StringVar(&opts.groupId, "groupId", , "Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.")	cmd.Flags().StringVar(&opts.roleName, "roleName", , "Human-readable label that identifies the role for the request. This name must beunique for this custom role in this project.")

	
	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("roleName")

	return cmd
}
