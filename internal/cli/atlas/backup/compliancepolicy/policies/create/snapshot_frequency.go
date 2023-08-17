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
)

func (opts *CreateOpts) askForSnapshotInterval(isOnDemand bool) (int, string, error) {
	if isOnDemand {
		// ondemand policy items have a frequencyInterval of 0.
		return 0, "ondemand", nil
	}

	q := newSnapshotIntervalQuestion()

	var snapshotType string
	if err := telemetry.TrackAskOne(q, &snapshotType); err != nil {
		return -1, "", err
	}

	snapshotInterval := convertResponse(snapshotType)

	return snapshotInterval, snapshotType, nil
}

func newSnapshotIntervalQuestion() survey.Prompt {
	options := []string{"daily", "weekly", "monthly", "hourly", "2 hours", "4 hours", "6 hours", "8 hours", "10 hours", "12 hours"}

	return &survey.Select{
		Message: "How often should we take a snapshot?",
		Options: options,
		Filter: func(filter string, _ string, i int) bool {
			filter = strings.ToLower(filter)
			return strings.HasPrefix(options[i], filter)
		},
	}
}

func convertResponse(response string) int {
	optionMap := map[string]int{
		"daily":    1,
		"weekly":   1,
		"monthly":  1,
		"hourly":   1,
		"2 hours":  2,
		"4 hours":  4,
		"6 hours":  6,
		"8 hours":  8,
		"10 hours": 10,
		"12 hours": 12,
	}
	return optionMap[response]
}
