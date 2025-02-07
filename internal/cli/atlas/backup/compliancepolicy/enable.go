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

package compliancepolicy

import (
	"context"
	"fmt"
	"net/mail"

	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/internal/flag"
	store "github.com/mongodb/mongodb-atlas-cli/internal/store/atlas"
	"github.com/mongodb/mongodb-atlas-cli/internal/usage"
	"github.com/spf13/cobra"
	atlasv2 "go.mongodb.org/atlas-sdk/v20230201004/admin"
)

type EnableOpts struct {
	cli.GlobalOpts
	cli.WatchOpts
	policy          *atlasv2.DataProtectionSettings
	store           store.CompliancePolicy
	authorizedEmail string
	EnableWatch     bool
}

var enableWatchTemplate = `Backup Compliance Policy enabled without any configuration. Run "atlas backups compliancepolicy --help" for configuration options.
`
var enableTemplate = `Backup Compliance Policy is being enabled without any configuration. Run "atlas backups compliancepolicy --help" for configuration options.
`

func (opts *EnableOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *EnableOpts) enableWatcher() (bool, error) {
	res, err := opts.store.DescribeCompliancePolicy(opts.ConfigProjectID())
	if err != nil {
		return false, err
	}
	opts.policy = res
	if res.GetState() == "" {
		return false, fmt.Errorf("could not access State field")
	}
	return (res.GetState() == active), nil
}

func (opts *EnableOpts) getEmptyCompliancePolicy() *atlasv2.DataProtectionSettings {
	policy := atlasv2.NewDataProtectionSettings()
	policy.SetAuthorizedEmail(opts.authorizedEmail)
	policy.SetProjectId(opts.ConfigProjectID())
	return policy
}

func (opts *EnableOpts) Run() error {
	if _, err := mail.ParseAddress(opts.authorizedEmail); err != nil {
		return err
	}
	emptyPolicy := opts.getEmptyCompliancePolicy()

	compliancePolicy, err := opts.store.UpdateCompliancePolicy(opts.ConfigProjectID(), emptyPolicy)
	opts.policy = compliancePolicy
	if err != nil {
		return err
	}
	if opts.EnableWatch {
		if err := opts.Watch(opts.enableWatcher); err != nil {
			return err
		}
		opts.Template = enableWatchTemplate
	}

	return opts.Print(opts.policy)
}

func EnableBuilder() *cobra.Command {
	opts := new(EnableOpts)
	cmd := &cobra.Command{
		Use:   "enable",
		Short: "Enable Backup Compliance Policy without any configuration.",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), enableTemplate),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVar(&opts.authorizedEmail, flag.AuthorizedEmail, "", usage.AuthorizedEmail)
	_ = cmd.MarkFlagRequired(flag.AuthorizedEmail)
	cmd.Flags().BoolVarP(&opts.EnableWatch, flag.EnableWatch, flag.EnableWatchShort, false, usage.EnableWatchDefault)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())
	return cmd
}
