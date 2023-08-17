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
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/mongodb/mongodb-atlas-cli/internal/telemetry"
)

var retentionMessage = `Define the amount of time a snapshot is recoverable. 
For example: “15 days”, “1 month”, “12 hours”.
Use the format: number hour(s)/day(s)/month(s)
`

const (
	invalidFormatError      = "Could not parse input, please try again."
	retentionValueZeroError = "Retention value can not be 0. Please try again."
)

func (opts *CreateOpts) askForRetention() (string, int, error) {
	var retention string
	var retentionValue int
	var retentionUnit string
	fmt.Println(retentionMessage)
	for {
		q := newRetentionPeriodQuestion()
		err := telemetry.TrackAskOne(q, &retention)
		if err != nil {
			return "", -1, err
		}
		retentionValue, retentionUnit, err = convertRetentionString(retention)
		if err != nil {
			if err.Error() == invalidFormatError || err.Error() == retentionValueZeroError {
				fmt.Println(err.Error())
				continue
			} else {
				return "", -1, fmt.Errorf("encountered an error while converting input to data: %w", err)
			}
		}
		break
	}

	return retentionUnit, retentionValue, nil
}

func convertRetentionString(timeStr string) (int, string, error) {
	re := regexp.MustCompile(`^(\d+)\s*(hour|day|month|week)s?$`)
	matches := re.FindStringSubmatch(strings.ToLower(timeStr))

	if matches == nil || len(matches) < 3 {
		return -1, "", errors.New(invalidFormatError)
	}

	number, err := strconv.Atoi(matches[1])
	if err != nil {
		return -1, "", err
	}
	if number == 0 {
		return -1, "", errors.New(retentionValueZeroError)
	}

	unit := matches[2]
	unit = addSIfMissing(unit)

	return number, unit, nil
}

func addSIfMissing(str string) string {
	if !strings.HasSuffix(str, "s") {
		return str + "s"
	}
	return str
}

func newRetentionPeriodQuestion() survey.Prompt {
	return &survey.Input{
		Message: "For how long should we keep a snapshot?",
	}
}
