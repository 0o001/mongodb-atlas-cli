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

type AcknowledgeAlertOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client admin.APIClient
	groupId string
	alertId string
}

func (opts *AcknowledgeAlertOpts) initClient(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *AcknowledgeAlertOpts) Run(ctx context.Context) error {
	params := &admin.AcknowledgeAlertApiParams{
		GroupId: opts.groupId,
		AlertId: opts.alertId,
	}
	resp, _, err := opts.client.AlertsApi.AcknowledgeAlertWithParams(ctx, params)
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

const AcknowledgeAlertTemplate = "<<some template>>"

func AcknowledgeAlertBuilder() cobra.Command {
	opts := AcknowledgeAlertOpts{}
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
				opts.InitOutput(cmd.OutOrStdout(), AcknowledgeAlertTemplate),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "groupId", "", "usage description")
	_ = cmd.MarkFlagRequired("groupId")
	cmd.Flags().StringVar(&opts.alertId, "alertId", "", "usage description")
	_ = cmd.MarkFlagRequired("alertId")

	return cmd
}
type GetAlertOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client admin.APIClient
	groupId string
	alertId string
}

func (opts *GetAlertOpts) initClient(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *GetAlertOpts) Run(ctx context.Context) error {
	params := &admin.GetAlertApiParams{
		GroupId: opts.groupId,
		AlertId: opts.alertId,
	}
	resp, _, err := opts.client.AlertsApi.GetAlertWithParams(ctx, params)
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

const GetAlertTemplate = "<<some template>>"

func GetAlertBuilder() cobra.Command {
	opts := GetAlertOpts{}
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
				opts.InitOutput(cmd.OutOrStdout(), GetAlertTemplate),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "groupId", "", "usage description")
	_ = cmd.MarkFlagRequired("groupId")
	cmd.Flags().StringVar(&opts.alertId, "alertId", "", "usage description")
	_ = cmd.MarkFlagRequired("alertId")

	return cmd
}
type ListAlertsOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client admin.APIClient
	groupId string
	includeCount bool
	itemsPerPage int32
	pageNum int32
	status string
}

func (opts *ListAlertsOpts) initClient(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *ListAlertsOpts) Run(ctx context.Context) error {
	params := &admin.ListAlertsApiParams{
		GroupId: opts.groupId,
		IncludeCount: opts.includeCount,
		ItemsPerPage: opts.itemsPerPage,
		PageNum: opts.pageNum,
		Status: opts.status,
	}
	resp, _, err := opts.client.AlertsApi.ListAlertsWithParams(ctx, params)
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

const ListAlertsTemplate = "<<some template>>"

func ListAlertsBuilder() cobra.Command {
	opts := ListAlertsOpts{}
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
				opts.InitOutput(cmd.OutOrStdout(), ListAlertsTemplate),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "groupId", "", "usage description")
	_ = cmd.MarkFlagRequired("groupId")
	cmd.Flags().StringVar(&opts.includeCount, "includeCount", "", "usage description")
	cmd.Flags().StringVar(&opts.itemsPerPage, "itemsPerPage", "", "usage description")
	cmd.Flags().StringVar(&opts.pageNum, "pageNum", "", "usage description")
	cmd.Flags().StringVar(&opts.status, "status", "", "usage description")

	return cmd
}
type ListAlertsByAlertConfigurationIdOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client admin.APIClient
	groupId string
	alertConfigId string
	includeCount bool
	itemsPerPage int32
	pageNum int32
}

func (opts *ListAlertsByAlertConfigurationIdOpts) initClient(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *ListAlertsByAlertConfigurationIdOpts) Run(ctx context.Context) error {
	params := &admin.ListAlertsByAlertConfigurationIdApiParams{
		GroupId: opts.groupId,
		AlertConfigId: opts.alertConfigId,
		IncludeCount: opts.includeCount,
		ItemsPerPage: opts.itemsPerPage,
		PageNum: opts.pageNum,
	}
	resp, _, err := opts.client.AlertsApi.ListAlertsByAlertConfigurationIdWithParams(ctx, params)
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

const ListAlertsByAlertConfigurationIdTemplate = "<<some template>>"

func ListAlertsByAlertConfigurationIdBuilder() cobra.Command {
	opts := ListAlertsByAlertConfigurationIdOpts{}
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
				opts.InitOutput(cmd.OutOrStdout(), ListAlertsByAlertConfigurationIdTemplate),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "groupId", "", "usage description")
	_ = cmd.MarkFlagRequired("groupId")
	cmd.Flags().StringVar(&opts.alertConfigId, "alertConfigId", "", "usage description")
	_ = cmd.MarkFlagRequired("alertConfigId")
	cmd.Flags().StringVar(&opts.includeCount, "includeCount", "", "usage description")
	cmd.Flags().StringVar(&opts.itemsPerPage, "itemsPerPage", "", "usage description")
	cmd.Flags().StringVar(&opts.pageNum, "pageNum", "", "usage description")

	return cmd
}
