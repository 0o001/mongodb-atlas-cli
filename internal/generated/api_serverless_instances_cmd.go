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

type CreateServerlessInstanceOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client admin.APIClient
	groupId string
}

func (opts *CreateServerlessInstanceOpts) initClient(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *CreateServerlessInstanceOpts) Run(ctx context.Context) error {
	params := &admin.CreateServerlessInstanceApiParams{
		GroupId: opts.groupId,
	}
	resp, _, err := opts.client.ServerlessInstancesApi.CreateServerlessInstanceWithParams(ctx, params)
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

const CreateServerlessInstanceTemplate = "<<some template>>"

func CreateServerlessInstanceBuilder() cobra.Command {
	opts := CreateServerlessInstanceOpts{}
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
				opts.InitOutput(cmd.OutOrStdout(), CreateServerlessInstanceTemplate),
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
type DeleteServerlessInstanceOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client admin.APIClient
	groupId string
	name string
}

func (opts *DeleteServerlessInstanceOpts) initClient(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *DeleteServerlessInstanceOpts) Run(ctx context.Context) error {
	params := &admin.DeleteServerlessInstanceApiParams{
		GroupId: opts.groupId,
		Name: opts.name,
	}
	resp, _, err := opts.client.ServerlessInstancesApi.DeleteServerlessInstanceWithParams(ctx, params)
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

const DeleteServerlessInstanceTemplate = "<<some template>>"

func DeleteServerlessInstanceBuilder() cobra.Command {
	opts := DeleteServerlessInstanceOpts{}
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
				opts.InitOutput(cmd.OutOrStdout(), DeleteServerlessInstanceTemplate),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "groupId", "", "usage description")
	_ = cmd.MarkFlagRequired("groupId")
	cmd.Flags().StringVar(&opts.name, "name", "", "usage description")
	_ = cmd.MarkFlagRequired("name")

	return cmd
}
type GetServerlessInstanceOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client admin.APIClient
	groupId string
	name string
}

func (opts *GetServerlessInstanceOpts) initClient(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *GetServerlessInstanceOpts) Run(ctx context.Context) error {
	params := &admin.GetServerlessInstanceApiParams{
		GroupId: opts.groupId,
		Name: opts.name,
	}
	resp, _, err := opts.client.ServerlessInstancesApi.GetServerlessInstanceWithParams(ctx, params)
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

const GetServerlessInstanceTemplate = "<<some template>>"

func GetServerlessInstanceBuilder() cobra.Command {
	opts := GetServerlessInstanceOpts{}
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
				opts.InitOutput(cmd.OutOrStdout(), GetServerlessInstanceTemplate),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "groupId", "", "usage description")
	_ = cmd.MarkFlagRequired("groupId")
	cmd.Flags().StringVar(&opts.name, "name", "", "usage description")
	_ = cmd.MarkFlagRequired("name")

	return cmd
}
type ListServerlessInstancesOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client admin.APIClient
	groupId string
	includeCount bool
	itemsPerPage int32
	pageNum int32
}

func (opts *ListServerlessInstancesOpts) initClient(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *ListServerlessInstancesOpts) Run(ctx context.Context) error {
	params := &admin.ListServerlessInstancesApiParams{
		GroupId: opts.groupId,
		IncludeCount: opts.includeCount,
		ItemsPerPage: opts.itemsPerPage,
		PageNum: opts.pageNum,
	}
	resp, _, err := opts.client.ServerlessInstancesApi.ListServerlessInstancesWithParams(ctx, params)
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

const ListServerlessInstancesTemplate = "<<some template>>"

func ListServerlessInstancesBuilder() cobra.Command {
	opts := ListServerlessInstancesOpts{}
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
				opts.InitOutput(cmd.OutOrStdout(), ListServerlessInstancesTemplate),
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

	return cmd
}
type UpdateServerlessInstanceOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client admin.APIClient
	groupId string
	name string
}

func (opts *UpdateServerlessInstanceOpts) initClient(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *UpdateServerlessInstanceOpts) Run(ctx context.Context) error {
	params := &admin.UpdateServerlessInstanceApiParams{
		GroupId: opts.groupId,
		Name: opts.name,
	}
	resp, _, err := opts.client.ServerlessInstancesApi.UpdateServerlessInstanceWithParams(ctx, params)
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

const UpdateServerlessInstanceTemplate = "<<some template>>"

func UpdateServerlessInstanceBuilder() cobra.Command {
	opts := UpdateServerlessInstanceOpts{}
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
				opts.InitOutput(cmd.OutOrStdout(), UpdateServerlessInstanceTemplate),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "groupId", "", "usage description")
	_ = cmd.MarkFlagRequired("groupId")
	cmd.Flags().StringVar(&opts.name, "name", "", "usage description")
	_ = cmd.MarkFlagRequired("name")

	return cmd
}
