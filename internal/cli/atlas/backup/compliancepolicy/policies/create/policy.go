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

package create

import (
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/mongodb/mongodb-atlas-cli/internal/telemetry"
	atlasv2 "go.mongodb.org/atlas-sdk/v20230201004/admin"
)

const (
	onDemandPolicyItem  = "on-demand policy"
	scheduledPolicyItem = "scheduled policy"
)

var policiesListTemplate = `Your existing policies

ID	FREQUENCY INTERVAL	FREQUENCY TYPE	RETENTION
{{- range .ScheduledPolicyItems}}
{{.Id}}	{{if eq .FrequencyType "hourly"}}{{.FrequencyInterval}}{{else}}-{{end}}	{{.FrequencyType}}	{{.RetentionValue}} {{.RetentionUnit}}
{{- end}}
{{if .OnDemandPolicyItem}}{{.OnDemandPolicyItem.Id}}	-	{{.OnDemandPolicyItem.FrequencyType}}	{{.OnDemandPolicyItem.RetentionValue}} {{.OnDemandPolicyItem.RetentionUnit}}{{end}}
`

func (opts *CreateOpts) askPolicyOptions(compliancePolicy *atlasv2.DataProtectionSettings) (bool, error) {

	if doExistingPoliciesExist(compliancePolicy) {
		opts.printExistingPolicies(compliancePolicy)
	}

	q := newPolicyItemSelectionQuestion()
	var policyType string
	if err := telemetry.TrackAskOne(q, &policyType); err != nil {
		return false, err
	}
	return policyType == onDemandPolicyItem, nil
}

func doExistingPoliciesExist(compliancePolicy *atlasv2.DataProtectionSettings) bool {
	return compliancePolicy.HasOnDemandPolicyItem() || len(compliancePolicy.GetScheduledPolicyItems()) > 0
}

func (opts *CreateOpts) printExistingPolicies(compliancePolicy *atlasv2.DataProtectionSettings) {
	opts.Template = policiesListTemplate
	opts.Print(compliancePolicy)
}

func newPolicyItemSelectionQuestion() survey.Prompt {
	options := []string{"on-demand policy", "scheduled policy"}

	return &survey.Select{
		Message: "Choose which policy you'd like to create",
		Help:    "With a on-demand policy you can take a snapshot whenever you want. Scheduled policies take snapshots at regular intervals, for example every month.",
		Options: options,
		Filter: func(filter string, _ string, i int) bool {
			filter = strings.ToLower(filter)
			return strings.HasPrefix(options[i], filter)
		},
	}
}
