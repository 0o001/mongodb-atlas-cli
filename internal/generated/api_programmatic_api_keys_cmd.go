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

type AddProjectApiKeyOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	groupId string
	apiUserId string
}

func (opts *AddProjectApiKeyOpts) initClient(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *AddProjectApiKeyOpts) Run(ctx context.Context) error {
	params := &admin.AddProjectApiKeyApiParams{
		GroupId: opts.groupId,
		ApiUserId: opts.apiUserId,
	}
	resp, _, err := opts.client.ProgrammaticAPIKeysApi.AddProjectApiKeyWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func AddProjectApiKeyBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := AddProjectApiKeyOpts{}
	cmd := &cobra.Command{
		Use:     "<<use>>",
		// Aliases: []string{"?"},
		Short:   "Assign One Organization API Key to One Project",
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
	cmd.Flags().StringVar(&opts.groupId, "groupId", , "Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.")	cmd.Flags().StringVar(&opts.apiUserId, "apiUserId", , "Unique 24-hexadecimal digit string that identifies this organization API key that you want to assign to one project.")

	
	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("apiUserId")

	return cmd
}
type CreateApiKeyOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	orgId string
}

func (opts *CreateApiKeyOpts) initClient(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *CreateApiKeyOpts) Run(ctx context.Context) error {
	params := &admin.CreateApiKeyApiParams{
		OrgId: opts.orgId,
	}
	resp, _, err := opts.client.ProgrammaticAPIKeysApi.CreateApiKeyWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func CreateApiKeyBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := CreateApiKeyOpts{}
	cmd := &cobra.Command{
		Use:     "<<use>>",
		// Aliases: []string{"?"},
		Short:   "Create One Organization API Key",
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
type CreateApiKeyAccessListOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	orgId string
	apiUserId string
	includeCount bool
	itemsPerPage int32
	pageNum int32
}

func (opts *CreateApiKeyAccessListOpts) initClient(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *CreateApiKeyAccessListOpts) Run(ctx context.Context) error {
	params := &admin.CreateApiKeyAccessListApiParams{
		OrgId: opts.orgId,
		ApiUserId: opts.apiUserId,
		IncludeCount: opts.includeCount,
		ItemsPerPage: opts.itemsPerPage,
		PageNum: opts.pageNum,
	}
	resp, _, err := opts.client.ProgrammaticAPIKeysApi.CreateApiKeyAccessListWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func CreateApiKeyAccessListBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := CreateApiKeyAccessListOpts{}
	cmd := &cobra.Command{
		Use:     "<<use>>",
		// Aliases: []string{"?"},
		Short:   "Create Access List Entries for One Organization API Key",
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
	cmd.Flags().StringVar(&opts.orgId, "orgId", , "Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.")	cmd.Flags().StringVar(&opts.apiUserId, "apiUserId", , "Unique 24-hexadecimal digit string that identifies this organization API key for which you want to create a new access list entry.")
	cmd.Flags().StringVar(&opts.includeCount, "includeCount", true, "Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.")	cmd.Flags().StringVar(&opts.itemsPerPage, "itemsPerPage", 100, "Number of items that the response returns per page.")	cmd.Flags().StringVar(&opts.pageNum, "pageNum", 1, "Number of the page that displays the current set of the total objects that the response returns.")
	
	_ = cmd.MarkFlagRequired("orgId")
	_ = cmd.MarkFlagRequired("apiUserId")

	return cmd
}
type CreateProjectApiKeyOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	groupId string
}

func (opts *CreateProjectApiKeyOpts) initClient(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *CreateProjectApiKeyOpts) Run(ctx context.Context) error {
	params := &admin.CreateProjectApiKeyApiParams{
		GroupId: opts.groupId,
	}
	resp, _, err := opts.client.ProgrammaticAPIKeysApi.CreateProjectApiKeyWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func CreateProjectApiKeyBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := CreateProjectApiKeyOpts{}
	cmd := &cobra.Command{
		Use:     "<<use>>",
		// Aliases: []string{"?"},
		Short:   "Create and Assign One Organization API Key to One Project",
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
type DeleteApiKeyOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	orgId string
	apiUserId string
}

func (opts *DeleteApiKeyOpts) initClient(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *DeleteApiKeyOpts) Run(ctx context.Context) error {
	params := &admin.DeleteApiKeyApiParams{
		OrgId: opts.orgId,
		ApiUserId: opts.apiUserId,
	}
	resp, _, err := opts.client.ProgrammaticAPIKeysApi.DeleteApiKeyWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func DeleteApiKeyBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := DeleteApiKeyOpts{}
	cmd := &cobra.Command{
		Use:     "<<use>>",
		// Aliases: []string{"?"},
		Short:   "Remove One Organization API Key",
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
	cmd.Flags().StringVar(&opts.orgId, "orgId", , "Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.")	cmd.Flags().StringVar(&opts.apiUserId, "apiUserId", , "Unique 24-hexadecimal digit string that identifies this organization API key.")

	
	_ = cmd.MarkFlagRequired("orgId")
	_ = cmd.MarkFlagRequired("apiUserId")

	return cmd
}
type DeleteApiKeyAccessListEntryOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	orgId string
	apiUserId string
	ipAddress string
}

func (opts *DeleteApiKeyAccessListEntryOpts) initClient(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *DeleteApiKeyAccessListEntryOpts) Run(ctx context.Context) error {
	params := &admin.DeleteApiKeyAccessListEntryApiParams{
		OrgId: opts.orgId,
		ApiUserId: opts.apiUserId,
		IpAddress: opts.ipAddress,
	}
	resp, _, err := opts.client.ProgrammaticAPIKeysApi.DeleteApiKeyAccessListEntryWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func DeleteApiKeyAccessListEntryBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := DeleteApiKeyAccessListEntryOpts{}
	cmd := &cobra.Command{
		Use:     "<<use>>",
		// Aliases: []string{"?"},
		Short:   "Remove One Access List Entry for One Organization API Key",
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
	cmd.Flags().StringVar(&opts.orgId, "orgId", , "Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.")	cmd.Flags().StringVar(&opts.apiUserId, "apiUserId", , "Unique 24-hexadecimal digit string that identifies this organization API key for which you want to remove access list entries.")	cmd.Flags().StringVar(&opts.ipAddress, "ipAddress", , "One IP address or multiple IP addresses represented as one CIDR block to limit requests to API resources in the specified organization. When adding a CIDR block with a subnet mask, such as 192.0.2.0/24, use the URL-encoded value %2F for the forward slash /.")

	
	_ = cmd.MarkFlagRequired("orgId")
	_ = cmd.MarkFlagRequired("apiUserId")
	_ = cmd.MarkFlagRequired("ipAddress")

	return cmd
}
type GetApiKeyOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	orgId string
	apiUserId string
}

func (opts *GetApiKeyOpts) initClient(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *GetApiKeyOpts) Run(ctx context.Context) error {
	params := &admin.GetApiKeyApiParams{
		OrgId: opts.orgId,
		ApiUserId: opts.apiUserId,
	}
	resp, _, err := opts.client.ProgrammaticAPIKeysApi.GetApiKeyWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func GetApiKeyBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := GetApiKeyOpts{}
	cmd := &cobra.Command{
		Use:     "<<use>>",
		// Aliases: []string{"?"},
		Short:   "Return One Organization API Key",
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
	cmd.Flags().StringVar(&opts.orgId, "orgId", , "Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.")	cmd.Flags().StringVar(&opts.apiUserId, "apiUserId", , "Unique 24-hexadecimal digit string that identifies this organization API key that  you want to update.")

	
	_ = cmd.MarkFlagRequired("orgId")
	_ = cmd.MarkFlagRequired("apiUserId")

	return cmd
}
type GetApiKeyAccessListOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	orgId string
	ipAddress string
	apiUserId string
}

func (opts *GetApiKeyAccessListOpts) initClient(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *GetApiKeyAccessListOpts) Run(ctx context.Context) error {
	params := &admin.GetApiKeyAccessListApiParams{
		OrgId: opts.orgId,
		IpAddress: opts.ipAddress,
		ApiUserId: opts.apiUserId,
	}
	resp, _, err := opts.client.ProgrammaticAPIKeysApi.GetApiKeyAccessListWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func GetApiKeyAccessListBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := GetApiKeyAccessListOpts{}
	cmd := &cobra.Command{
		Use:     "<<use>>",
		// Aliases: []string{"?"},
		Short:   "Return One Access List Entry for One Organization API Key",
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
	cmd.Flags().StringVar(&opts.orgId, "orgId", , "Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.")	cmd.Flags().StringVar(&opts.ipAddress, "ipAddress", , "One IP address or multiple IP addresses represented as one CIDR block to limit  requests to API resources in the specified organization. When adding a CIDR block with a subnet mask, such as  192.0.2.0/24, use the URL-encoded value %2F for the forward slash /.")	cmd.Flags().StringVar(&opts.apiUserId, "apiUserId", , "Unique 24-hexadecimal digit string that identifies this organization API key for  which you want to return access list entries.")

	
	_ = cmd.MarkFlagRequired("orgId")
	_ = cmd.MarkFlagRequired("ipAddress")
	_ = cmd.MarkFlagRequired("apiUserId")

	return cmd
}
type ListApiKeyAccessListsEntriesOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	orgId string
	apiUserId string
	includeCount bool
	itemsPerPage int32
	pageNum int32
}

func (opts *ListApiKeyAccessListsEntriesOpts) initClient(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *ListApiKeyAccessListsEntriesOpts) Run(ctx context.Context) error {
	params := &admin.ListApiKeyAccessListsEntriesApiParams{
		OrgId: opts.orgId,
		ApiUserId: opts.apiUserId,
		IncludeCount: opts.includeCount,
		ItemsPerPage: opts.itemsPerPage,
		PageNum: opts.pageNum,
	}
	resp, _, err := opts.client.ProgrammaticAPIKeysApi.ListApiKeyAccessListsEntriesWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func ListApiKeyAccessListsEntriesBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := ListApiKeyAccessListsEntriesOpts{}
	cmd := &cobra.Command{
		Use:     "<<use>>",
		// Aliases: []string{"?"},
		Short:   "Return All Access List Entries for One Organization API Key",
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
	cmd.Flags().StringVar(&opts.orgId, "orgId", , "Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.")	cmd.Flags().StringVar(&opts.apiUserId, "apiUserId", , "Unique 24-hexadecimal digit string that identifies this organization API key for which you want to return access list entries.")
	cmd.Flags().StringVar(&opts.includeCount, "includeCount", true, "Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.")	cmd.Flags().StringVar(&opts.itemsPerPage, "itemsPerPage", 100, "Number of items that the response returns per page.")	cmd.Flags().StringVar(&opts.pageNum, "pageNum", 1, "Number of the page that displays the current set of the total objects that the response returns.")
	
	_ = cmd.MarkFlagRequired("orgId")
	_ = cmd.MarkFlagRequired("apiUserId")

	return cmd
}
type ListApiKeysOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	orgId string
	includeCount bool
	itemsPerPage int32
	pageNum int32
}

func (opts *ListApiKeysOpts) initClient(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *ListApiKeysOpts) Run(ctx context.Context) error {
	params := &admin.ListApiKeysApiParams{
		OrgId: opts.orgId,
		IncludeCount: opts.includeCount,
		ItemsPerPage: opts.itemsPerPage,
		PageNum: opts.pageNum,
	}
	resp, _, err := opts.client.ProgrammaticAPIKeysApi.ListApiKeysWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func ListApiKeysBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := ListApiKeysOpts{}
	cmd := &cobra.Command{
		Use:     "<<use>>",
		// Aliases: []string{"?"},
		Short:   "Return All Organization API Keys",
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
type ListProjectApiKeysOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	groupId string
	includeCount bool
	itemsPerPage int32
	pageNum int32
}

func (opts *ListProjectApiKeysOpts) initClient(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *ListProjectApiKeysOpts) Run(ctx context.Context) error {
	params := &admin.ListProjectApiKeysApiParams{
		GroupId: opts.groupId,
		IncludeCount: opts.includeCount,
		ItemsPerPage: opts.itemsPerPage,
		PageNum: opts.pageNum,
	}
	resp, _, err := opts.client.ProgrammaticAPIKeysApi.ListProjectApiKeysWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func ListProjectApiKeysBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := ListProjectApiKeysOpts{}
	cmd := &cobra.Command{
		Use:     "<<use>>",
		// Aliases: []string{"?"},
		Short:   "Return All Organization API Keys Assigned to One Project",
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
	cmd.Flags().StringVar(&opts.includeCount, "includeCount", true, "Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.")	cmd.Flags().StringVar(&opts.itemsPerPage, "itemsPerPage", 100, "Number of items that the response returns per page.")	cmd.Flags().StringVar(&opts.pageNum, "pageNum", 1, "Number of the page that displays the current set of the total objects that the response returns.")
	
	_ = cmd.MarkFlagRequired("groupId")

	return cmd
}
type RemoveProjectApiKeyOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	groupId string
	apiUserId string
}

func (opts *RemoveProjectApiKeyOpts) initClient(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *RemoveProjectApiKeyOpts) Run(ctx context.Context) error {
	params := &admin.RemoveProjectApiKeyApiParams{
		GroupId: opts.groupId,
		ApiUserId: opts.apiUserId,
	}
	resp, _, err := opts.client.ProgrammaticAPIKeysApi.RemoveProjectApiKeyWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func RemoveProjectApiKeyBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := RemoveProjectApiKeyOpts{}
	cmd := &cobra.Command{
		Use:     "<<use>>",
		// Aliases: []string{"?"},
		Short:   "Unassign One Organization API Key from One Project",
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
	cmd.Flags().StringVar(&opts.groupId, "groupId", , "Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.")	cmd.Flags().StringVar(&opts.apiUserId, "apiUserId", , "Unique 24-hexadecimal digit string that identifies this organization API key that you want to unassign from one project.")

	
	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("apiUserId")

	return cmd
}
type UpdateApiKeyOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	orgId string
	apiUserId string
}

func (opts *UpdateApiKeyOpts) initClient(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *UpdateApiKeyOpts) Run(ctx context.Context) error {
	params := &admin.UpdateApiKeyApiParams{
		OrgId: opts.orgId,
		ApiUserId: opts.apiUserId,
	}
	resp, _, err := opts.client.ProgrammaticAPIKeysApi.UpdateApiKeyWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func UpdateApiKeyBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := UpdateApiKeyOpts{}
	cmd := &cobra.Command{
		Use:     "<<use>>",
		// Aliases: []string{"?"},
		Short:   "Update One Organization API Key",
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
	cmd.Flags().StringVar(&opts.orgId, "orgId", , "Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.")	cmd.Flags().StringVar(&opts.apiUserId, "apiUserId", , "Unique 24-hexadecimal digit string that identifies this organization API key you  want to update.")

	
	_ = cmd.MarkFlagRequired("orgId")
	_ = cmd.MarkFlagRequired("apiUserId")

	return cmd
}
type UpdateApiKeyRolesOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	groupId string
	apiUserId string
	pageNum int32
	itemsPerPage int32
	includeCount bool
}

func (opts *UpdateApiKeyRolesOpts) initClient(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *UpdateApiKeyRolesOpts) Run(ctx context.Context) error {
	params := &admin.UpdateApiKeyRolesApiParams{
		GroupId: opts.groupId,
		ApiUserId: opts.apiUserId,
		PageNum: opts.pageNum,
		ItemsPerPage: opts.itemsPerPage,
		IncludeCount: opts.includeCount,
	}
	resp, _, err := opts.client.ProgrammaticAPIKeysApi.UpdateApiKeyRolesWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func UpdateApiKeyRolesBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := UpdateApiKeyRolesOpts{}
	cmd := &cobra.Command{
		Use:     "<<use>>",
		// Aliases: []string{"?"},
		Short:   "Update Roles of One Organization API Key to One Project",
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
	cmd.Flags().StringVar(&opts.groupId, "groupId", , "Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.")	cmd.Flags().StringVar(&opts.apiUserId, "apiUserId", , "Unique 24-hexadecimal digit string that identifies this organization API key that you want to unassign from one project.")
	cmd.Flags().StringVar(&opts.pageNum, "pageNum", 1, "Number of the page that displays the current set of the total objects that the response returns.")	cmd.Flags().StringVar(&opts.itemsPerPage, "itemsPerPage", 100, "Number of items that the response returns per page.")	cmd.Flags().StringVar(&opts.includeCount, "includeCount", true, "Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.")
	
	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("apiUserId")

	return cmd
}
