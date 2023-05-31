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

type CreateOrganizationOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
}

func (opts *CreateOrganizationOpts) initClient(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *CreateOrganizationOpts) Run(ctx context.Context) error {
	params := &admin.CreateOrganizationApiParams{
	}
	resp, _, err := opts.client.OrganizationsApi.CreateOrganizationWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func CreateOrganizationBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := CreateOrganizationOpts{}
	cmd := &cobra.Command{
		Use:     "<<use>>",
		// Aliases: []string{"?"},
		Short:   "Create One Organization",
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


	

	return cmd
}
type CreateOrganizationInvitationOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	orgId string
}

func (opts *CreateOrganizationInvitationOpts) initClient(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *CreateOrganizationInvitationOpts) Run(ctx context.Context) error {
	params := &admin.CreateOrganizationInvitationApiParams{
		OrgId: opts.orgId,
	}
	resp, _, err := opts.client.OrganizationsApi.CreateOrganizationInvitationWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func CreateOrganizationInvitationBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := CreateOrganizationInvitationOpts{}
	cmd := &cobra.Command{
		Use:     "<<use>>",
		// Aliases: []string{"?"},
		Short:   "Invite One MongoDB Cloud User to Join One Atlas Organization",
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
	cmd.Flags().StringVar(&opts.orgId, "orgId", , "Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.")

	
	_ = cmd.MarkFlagRequired("orgId")

	return cmd
}
type DeleteOrganizationOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	orgId string
}

func (opts *DeleteOrganizationOpts) initClient(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *DeleteOrganizationOpts) Run(ctx context.Context) error {
	params := &admin.DeleteOrganizationApiParams{
		OrgId: opts.orgId,
	}
	resp, _, err := opts.client.OrganizationsApi.DeleteOrganizationWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func DeleteOrganizationBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := DeleteOrganizationOpts{}
	cmd := &cobra.Command{
		Use:     "<<use>>",
		// Aliases: []string{"?"},
		Short:   "Remove One Organization",
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
	cmd.Flags().StringVar(&opts.orgId, "orgId", , "Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.")

	
	_ = cmd.MarkFlagRequired("orgId")

	return cmd
}
type DeleteOrganizationInvitationOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	orgId string
	invitationId string
}

func (opts *DeleteOrganizationInvitationOpts) initClient(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *DeleteOrganizationInvitationOpts) Run(ctx context.Context) error {
	params := &admin.DeleteOrganizationInvitationApiParams{
		OrgId: opts.orgId,
		InvitationId: opts.invitationId,
	}
	resp, _, err := opts.client.OrganizationsApi.DeleteOrganizationInvitationWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func DeleteOrganizationInvitationBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := DeleteOrganizationInvitationOpts{}
	cmd := &cobra.Command{
		Use:     "<<use>>",
		// Aliases: []string{"?"},
		Short:   "Cancel One Organization Invitation",
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
	cmd.Flags().StringVar(&opts.orgId, "orgId", , "Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.")	cmd.Flags().StringVar(&opts.invitationId, "invitationId", , "Unique 24-hexadecimal digit string that identifies the invitation.")

	
	_ = cmd.MarkFlagRequired("orgId")
	_ = cmd.MarkFlagRequired("invitationId")

	return cmd
}
type GetOrganizationOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	orgId string
}

func (opts *GetOrganizationOpts) initClient(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *GetOrganizationOpts) Run(ctx context.Context) error {
	params := &admin.GetOrganizationApiParams{
		OrgId: opts.orgId,
	}
	resp, _, err := opts.client.OrganizationsApi.GetOrganizationWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func GetOrganizationBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := GetOrganizationOpts{}
	cmd := &cobra.Command{
		Use:     "<<use>>",
		// Aliases: []string{"?"},
		Short:   "Return One Organization",
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
	cmd.Flags().StringVar(&opts.orgId, "orgId", , "Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.")

	
	_ = cmd.MarkFlagRequired("orgId")

	return cmd
}
type GetOrganizationInvitationOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	orgId string
	invitationId string
}

func (opts *GetOrganizationInvitationOpts) initClient(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *GetOrganizationInvitationOpts) Run(ctx context.Context) error {
	params := &admin.GetOrganizationInvitationApiParams{
		OrgId: opts.orgId,
		InvitationId: opts.invitationId,
	}
	resp, _, err := opts.client.OrganizationsApi.GetOrganizationInvitationWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func GetOrganizationInvitationBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := GetOrganizationInvitationOpts{}
	cmd := &cobra.Command{
		Use:     "<<use>>",
		// Aliases: []string{"?"},
		Short:   "Return One Organization Invitation",
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
	cmd.Flags().StringVar(&opts.orgId, "orgId", , "Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.")	cmd.Flags().StringVar(&opts.invitationId, "invitationId", , "Unique 24-hexadecimal digit string that identifies the invitation.")

	
	_ = cmd.MarkFlagRequired("orgId")
	_ = cmd.MarkFlagRequired("invitationId")

	return cmd
}
type GetOrganizationSettingsOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	orgId string
}

func (opts *GetOrganizationSettingsOpts) initClient(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *GetOrganizationSettingsOpts) Run(ctx context.Context) error {
	params := &admin.GetOrganizationSettingsApiParams{
		OrgId: opts.orgId,
	}
	resp, _, err := opts.client.OrganizationsApi.GetOrganizationSettingsWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func GetOrganizationSettingsBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := GetOrganizationSettingsOpts{}
	cmd := &cobra.Command{
		Use:     "<<use>>",
		// Aliases: []string{"?"},
		Short:   "Return Settings for One Organization",
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
	cmd.Flags().StringVar(&opts.orgId, "orgId", , "Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.")

	
	_ = cmd.MarkFlagRequired("orgId")

	return cmd
}
type ListOrganizationInvitationsOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	orgId string
	username string
}

func (opts *ListOrganizationInvitationsOpts) initClient(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *ListOrganizationInvitationsOpts) Run(ctx context.Context) error {
	params := &admin.ListOrganizationInvitationsApiParams{
		OrgId: opts.orgId,
		Username: opts.username,
	}
	resp, _, err := opts.client.OrganizationsApi.ListOrganizationInvitationsWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func ListOrganizationInvitationsBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := ListOrganizationInvitationsOpts{}
	cmd := &cobra.Command{
		Use:     "<<use>>",
		// Aliases: []string{"?"},
		Short:   "Return All Organization Invitations",
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
	cmd.Flags().StringVar(&opts.orgId, "orgId", , "Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.")
	cmd.Flags().StringVar(&opts.username, "username", , "Email address of the user account invited to this organization. If you exclude this parameter, this resource returns all pending invitations.")
	
	_ = cmd.MarkFlagRequired("orgId")

	return cmd
}
type ListOrganizationProjectsOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	orgId string
	includeCount bool
	itemsPerPage int32
	pageNum int32
	name string
}

func (opts *ListOrganizationProjectsOpts) initClient(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *ListOrganizationProjectsOpts) Run(ctx context.Context) error {
	params := &admin.ListOrganizationProjectsApiParams{
		OrgId: opts.orgId,
		IncludeCount: opts.includeCount,
		ItemsPerPage: opts.itemsPerPage,
		PageNum: opts.pageNum,
		Name: opts.name,
	}
	resp, _, err := opts.client.OrganizationsApi.ListOrganizationProjectsWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func ListOrganizationProjectsBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := ListOrganizationProjectsOpts{}
	cmd := &cobra.Command{
		Use:     "<<use>>",
		// Aliases: []string{"?"},
		Short:   "Return One or More Projects in One Organization",
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
	cmd.Flags().StringVar(&opts.orgId, "orgId", , "Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.")
	cmd.Flags().StringVar(&opts.includeCount, "includeCount", true, "Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.")	cmd.Flags().StringVar(&opts.itemsPerPage, "itemsPerPage", 100, "Number of items that the response returns per page.")	cmd.Flags().StringVar(&opts.pageNum, "pageNum", 1, "Number of the page that displays the current set of the total objects that the response returns.")	cmd.Flags().StringVar(&opts.name, "name", , "Human-readable label of the project to use to filter the returned list. Performs a case-insensitive search for a project within the organization which is prefixed by the specified name.")
	
	_ = cmd.MarkFlagRequired("orgId")

	return cmd
}
type ListOrganizationUsersOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	orgId string
	includeCount bool
	itemsPerPage int32
	pageNum int32
}

func (opts *ListOrganizationUsersOpts) initClient(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *ListOrganizationUsersOpts) Run(ctx context.Context) error {
	params := &admin.ListOrganizationUsersApiParams{
		OrgId: opts.orgId,
		IncludeCount: opts.includeCount,
		ItemsPerPage: opts.itemsPerPage,
		PageNum: opts.pageNum,
	}
	resp, _, err := opts.client.OrganizationsApi.ListOrganizationUsersWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func ListOrganizationUsersBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := ListOrganizationUsersOpts{}
	cmd := &cobra.Command{
		Use:     "<<use>>",
		// Aliases: []string{"?"},
		Short:   "Return All MongoDB Cloud Users in One Organization",
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
	cmd.Flags().StringVar(&opts.orgId, "orgId", , "Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.")
	cmd.Flags().StringVar(&opts.includeCount, "includeCount", true, "Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.")	cmd.Flags().StringVar(&opts.itemsPerPage, "itemsPerPage", 100, "Number of items that the response returns per page.")	cmd.Flags().StringVar(&opts.pageNum, "pageNum", 1, "Number of the page that displays the current set of the total objects that the response returns.")
	
	_ = cmd.MarkFlagRequired("orgId")

	return cmd
}
type ListOrganizationsOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	includeCount bool
	itemsPerPage int32
	pageNum int32
	name string
}

func (opts *ListOrganizationsOpts) initClient(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *ListOrganizationsOpts) Run(ctx context.Context) error {
	params := &admin.ListOrganizationsApiParams{
		IncludeCount: opts.includeCount,
		ItemsPerPage: opts.itemsPerPage,
		PageNum: opts.pageNum,
		Name: opts.name,
	}
	resp, _, err := opts.client.OrganizationsApi.ListOrganizationsWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func ListOrganizationsBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := ListOrganizationsOpts{}
	cmd := &cobra.Command{
		Use:     "<<use>>",
		// Aliases: []string{"?"},
		Short:   "Return All Organizations",
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

	cmd.Flags().StringVar(&opts.includeCount, "includeCount", true, "Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.")	cmd.Flags().StringVar(&opts.itemsPerPage, "itemsPerPage", 100, "Number of items that the response returns per page.")	cmd.Flags().StringVar(&opts.pageNum, "pageNum", 1, "Number of the page that displays the current set of the total objects that the response returns.")	cmd.Flags().StringVar(&opts.name, "name", , "Human-readable label of the organization to use to filter the returned list. Performs a case-insensitive search for an organization that starts with the specified name.")
	

	return cmd
}
type RenameOrganizationOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	orgId string
}

func (opts *RenameOrganizationOpts) initClient(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *RenameOrganizationOpts) Run(ctx context.Context) error {
	params := &admin.RenameOrganizationApiParams{
		OrgId: opts.orgId,
	}
	resp, _, err := opts.client.OrganizationsApi.RenameOrganizationWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func RenameOrganizationBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := RenameOrganizationOpts{}
	cmd := &cobra.Command{
		Use:     "<<use>>",
		// Aliases: []string{"?"},
		Short:   "Rename One Organization",
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
	cmd.Flags().StringVar(&opts.orgId, "orgId", , "Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.")

	
	_ = cmd.MarkFlagRequired("orgId")

	return cmd
}
type UpdateOrganizationInvitationOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	orgId string
}

func (opts *UpdateOrganizationInvitationOpts) initClient(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *UpdateOrganizationInvitationOpts) Run(ctx context.Context) error {
	params := &admin.UpdateOrganizationInvitationApiParams{
		OrgId: opts.orgId,
	}
	resp, _, err := opts.client.OrganizationsApi.UpdateOrganizationInvitationWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func UpdateOrganizationInvitationBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := UpdateOrganizationInvitationOpts{}
	cmd := &cobra.Command{
		Use:     "<<use>>",
		// Aliases: []string{"?"},
		Short:   "Update One Organization Invitation",
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
	cmd.Flags().StringVar(&opts.orgId, "orgId", , "Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.")

	
	_ = cmd.MarkFlagRequired("orgId")

	return cmd
}
type UpdateOrganizationInvitationByIdOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	orgId string
	invitationId string
}

func (opts *UpdateOrganizationInvitationByIdOpts) initClient(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *UpdateOrganizationInvitationByIdOpts) Run(ctx context.Context) error {
	params := &admin.UpdateOrganizationInvitationByIdApiParams{
		OrgId: opts.orgId,
		InvitationId: opts.invitationId,
	}
	resp, _, err := opts.client.OrganizationsApi.UpdateOrganizationInvitationByIdWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func UpdateOrganizationInvitationByIdBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := UpdateOrganizationInvitationByIdOpts{}
	cmd := &cobra.Command{
		Use:     "<<use>>",
		// Aliases: []string{"?"},
		Short:   "Update One Organization Invitation by Invitation ID",
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
	cmd.Flags().StringVar(&opts.orgId, "orgId", , "Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.")	cmd.Flags().StringVar(&opts.invitationId, "invitationId", , "Unique 24-hexadecimal digit string that identifies the invitation.")

	
	_ = cmd.MarkFlagRequired("orgId")
	_ = cmd.MarkFlagRequired("invitationId")

	return cmd
}
type UpdateOrganizationSettingsOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	orgId string
}

func (opts *UpdateOrganizationSettingsOpts) initClient(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *UpdateOrganizationSettingsOpts) Run(ctx context.Context) error {
	params := &admin.UpdateOrganizationSettingsApiParams{
		OrgId: opts.orgId,
	}
	resp, _, err := opts.client.OrganizationsApi.UpdateOrganizationSettingsWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func UpdateOrganizationSettingsBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := UpdateOrganizationSettingsOpts{}
	cmd := &cobra.Command{
		Use:     "<<use>>",
		// Aliases: []string{"?"},
		Short:   "Update Settings for One Organization",
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
	cmd.Flags().StringVar(&opts.orgId, "orgId", , "Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.")

	
	_ = cmd.MarkFlagRequired("orgId")

	return cmd
}
