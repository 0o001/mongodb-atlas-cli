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

type GetOrganizationEventOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	orgId string
	eventId string
	includeRaw bool
}

func (opts *GetOrganizationEventOpts) initClient(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *GetOrganizationEventOpts) Run(ctx context.Context) error {
	params := &admin.GetOrganizationEventApiParams{
		OrgId: opts.orgId,
		EventId: opts.eventId,
		IncludeRaw: opts.includeRaw,
	}
	resp, _, err := opts.client.EventsApi.GetOrganizationEventWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func GetOrganizationEventBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := GetOrganizationEventOpts{}
	cmd := &cobra.Command{
		Use:     "<<use>>",
		// Aliases: []string{"?"},
		Short:   "Return One Event from One Organization",
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
	cmd.Flags().StringVar(&opts.orgId, "orgId", , "Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.")	cmd.Flags().StringVar(&opts.eventId, "eventId", , "Unique 24-hexadecimal digit string that identifies the event that you want to return. Use the [/events](#tag/Events/operation/listOrganizationEvents) endpoint to retrieve all events to which the authenticated user has access.")
	cmd.Flags().StringVar(&opts.includeRaw, "includeRaw", false, "Flag that indicates whether to include the raw document in the output. The raw document contains additional meta information about the event.")
	
	_ = cmd.MarkFlagRequired("orgId")
	_ = cmd.MarkFlagRequired("eventId")

	return cmd
}
type GetProjectEventOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	groupId string
	eventId string
	includeRaw bool
}

func (opts *GetProjectEventOpts) initClient(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *GetProjectEventOpts) Run(ctx context.Context) error {
	params := &admin.GetProjectEventApiParams{
		GroupId: opts.groupId,
		EventId: opts.eventId,
		IncludeRaw: opts.includeRaw,
	}
	resp, _, err := opts.client.EventsApi.GetProjectEventWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func GetProjectEventBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := GetProjectEventOpts{}
	cmd := &cobra.Command{
		Use:     "<<use>>",
		// Aliases: []string{"?"},
		Short:   "Return One Event from One Project",
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
	cmd.Flags().StringVar(&opts.groupId, "groupId", , "Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.")	cmd.Flags().StringVar(&opts.eventId, "eventId", , "Unique 24-hexadecimal digit string that identifies the event that you want to return. Use the [/events](#tag/Events/operation/listProjectEvents) endpoint to retrieve all events to which the authenticated user has access.")
	cmd.Flags().StringVar(&opts.includeRaw, "includeRaw", false, "Flag that indicates whether to include the raw document in the output. The raw document contains additional meta information about the event.")
	
	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("eventId")

	return cmd
}
type ListOrganizationEventsOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	orgId string
	includeCount bool
	itemsPerPage int32
	pageNum int32
	eventType []string
	includeRaw bool
	maxDate time.Time
	minDate time.Time
}

func (opts *ListOrganizationEventsOpts) initClient(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *ListOrganizationEventsOpts) Run(ctx context.Context) error {
	params := &admin.ListOrganizationEventsApiParams{
		OrgId: opts.orgId,
		IncludeCount: opts.includeCount,
		ItemsPerPage: opts.itemsPerPage,
		PageNum: opts.pageNum,
		EventType: opts.eventType,
		IncludeRaw: opts.includeRaw,
		MaxDate: opts.maxDate,
		MinDate: opts.minDate,
	}
	resp, _, err := opts.client.EventsApi.ListOrganizationEventsWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func ListOrganizationEventsBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := ListOrganizationEventsOpts{}
	cmd := &cobra.Command{
		Use:     "<<use>>",
		// Aliases: []string{"?"},
		Short:   "Return All Events from One Organization",
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
	cmd.Flags().StringVar(&opts.includeCount, "includeCount", true, "Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.")	cmd.Flags().StringVar(&opts.itemsPerPage, "itemsPerPage", 100, "Number of items that the response returns per page.")	cmd.Flags().StringVar(&opts.pageNum, "pageNum", 1, "Number of the page that displays the current set of the total objects that the response returns.")	cmd.Flags().StringVar(&opts.eventType, "eventType", , "Category of incident recorded at this moment in time.  **IMPORTANT**: The complete list of event type values changes frequently.")	cmd.Flags().StringVar(&opts.includeRaw, "includeRaw", false, "Flag that indicates whether to include the raw document in the output. The raw document contains additional meta information about the event.")	cmd.Flags().StringVar(&opts.maxDate, "maxDate", , "Date and time from when MongoDB Cloud stops returning events. This parameter uses the &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/ISO_8601\&quot; target&#x3D;\&quot;_blank\&quot; rel&#x3D;\&quot;noopener noreferrer\&quot;&gt;ISO 8601&lt;/a&gt; timestamp format in UTC.")	cmd.Flags().StringVar(&opts.minDate, "minDate", , "Date and time from when MongoDB Cloud starts returning events. This parameter uses the &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/ISO_8601\&quot; target&#x3D;\&quot;_blank\&quot; rel&#x3D;\&quot;noopener noreferrer\&quot;&gt;ISO 8601&lt;/a&gt; timestamp format in UTC.")
	
	_ = cmd.MarkFlagRequired("orgId")

	return cmd
}
type ListProjectEventsOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	groupId string
	includeCount bool
	itemsPerPage int32
	pageNum int32
	clusterNames []string
	eventType []string
	includeRaw bool
	maxDate time.Time
	minDate time.Time
}

func (opts *ListProjectEventsOpts) initClient(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *ListProjectEventsOpts) Run(ctx context.Context) error {
	params := &admin.ListProjectEventsApiParams{
		GroupId: opts.groupId,
		IncludeCount: opts.includeCount,
		ItemsPerPage: opts.itemsPerPage,
		PageNum: opts.pageNum,
		ClusterNames: opts.clusterNames,
		EventType: opts.eventType,
		IncludeRaw: opts.includeRaw,
		MaxDate: opts.maxDate,
		MinDate: opts.minDate,
	}
	resp, _, err := opts.client.EventsApi.ListProjectEventsWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func ListProjectEventsBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := ListProjectEventsOpts{}
	cmd := &cobra.Command{
		Use:     "<<use>>",
		// Aliases: []string{"?"},
		Short:   "Return All Events from One Project",
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
	cmd.Flags().StringVar(&opts.includeCount, "includeCount", true, "Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.")	cmd.Flags().StringVar(&opts.itemsPerPage, "itemsPerPage", 100, "Number of items that the response returns per page.")	cmd.Flags().StringVar(&opts.pageNum, "pageNum", 1, "Number of the page that displays the current set of the total objects that the response returns.")	cmd.Flags().StringVar(&opts.clusterNames, "clusterNames", , "Human-readable label that identifies the cluster.")	cmd.Flags().StringVar(&opts.eventType, "eventType", , "Category of incident recorded at this moment in time.  **IMPORTANT**: The complete list of event type values changes frequently.")	cmd.Flags().StringVar(&opts.includeRaw, "includeRaw", false, "Flag that indicates whether to include the raw document in the output. The raw document contains additional meta information about the event.")	cmd.Flags().StringVar(&opts.maxDate, "maxDate", , "Date and time from when MongoDB Cloud stops returning events. This parameter uses the &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/ISO_8601\&quot; target&#x3D;\&quot;_blank\&quot; rel&#x3D;\&quot;noopener noreferrer\&quot;&gt;ISO 8601&lt;/a&gt; timestamp format in UTC.")	cmd.Flags().StringVar(&opts.minDate, "minDate", , "Date and time from when MongoDB Cloud starts returning events. This parameter uses the &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/ISO_8601\&quot; target&#x3D;\&quot;_blank\&quot; rel&#x3D;\&quot;noopener noreferrer\&quot;&gt;ISO 8601&lt;/a&gt; timestamp format in UTC.")
	
	_ = cmd.MarkFlagRequired("groupId")

	return cmd
}
