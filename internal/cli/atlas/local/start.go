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

package local

import (
	"context"
	_ "embed"
	"os"
	"os/exec"
	"path"
	"time"

	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/internal/mongosh"
	"github.com/mongodb/mongodb-atlas-cli/internal/usage"
	"github.com/spf13/cobra"
)

var (
	//go:embed files/docker-compose.yml
	dockerComposeContents []byte

	//go:embed files/mms-config.json
	mmsConfigContents []byte

	//go:embed files/ca.pem
	CaContents []byte

	//go:embed files/seedRs.js
	SeedRsContents []byte

	//go:embed files/seedUser.js
	SeedUserContents []byte
)

const speed = 100 * time.Millisecond

type StartOpts struct {
	cli.OutputOpts
	cli.GlobalOpts
	debug bool
}

var startTemplate = `local environment started at {{.ConnectionString}}
`

func dumpTempFile(pattern string, contents []byte) (string, error) {
	f, err := os.CreateTemp("", pattern)
	if err != nil {
		return "", err
	}
	_, err = f.Write(contents)
	_ = f.Close()
	if err != nil {
		return f.Name(), err
	}
	return f.Name(), nil
}

func runDockerCompose(debug bool, args ...string) error {
	dockerComposeFilename, err := dumpTempFile("docker-compose", dockerComposeContents)
	if dockerComposeFilename != "" {
		defer os.Remove(dockerComposeFilename)
	}
	if err != nil {
		return err
	}
	cmdArgs := append([]string{"compose", "--compatibility", "-p", "docker", "-f", dockerComposeFilename}, args...)
	cmd := exec.Command("docker", cmdArgs...)
	cmd.Env = append(os.Environ(), "DOCKER_BUILDKIT=0")
	if debug {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin
	}
	return cmd.Run()
}

func mmsConfigPath() (string, error) {
	configHome, err := config.AtlasCLIConfigHome()
	if err != nil {
		return "", err
	}
	return path.Join(configHome, "mms-config.json"), nil
}

func dumpMmsConfig() error {
	mmsConfigfile, err := mmsConfigPath()
	if err != nil {
		return err
	}
	if _, err := os.Stat(mmsConfigfile); os.IsNotExist(err) {
		err = os.WriteFile(mmsConfigfile, mmsConfigContents, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}
func seed(script string) error {
	caFilename, err := dumpTempFile("ca", CaContents)
	if caFilename != "" {
		defer os.Remove(caFilename)
	}
	if err != nil {
		return err
	}

	return mongosh.Exec("mongodb://__system:keyfile@localhost:37017/admin?authSource=local&tls=true&tlsCAFile="+caFilename, "--eval", script)
}

func (opts *StartOpts) Run(_ context.Context) error {
	if err := dumpMmsConfig(); err != nil {
		return err
	}

	if err := runDockerCompose(opts.debug, "up", "-d"); err != nil {
		return err
	}

	for _, seedScriptContents := range []string{string(SeedRsContents), string(SeedUserContents)} {
		if err := seed(seedScriptContents); err != nil {
			return err
		}
	}

	return opts.Print(localData)
}

// atlas local start.
func StartBuilder() *cobra.Command {
	opts := &StartOpts{}
	cmd := &cobra.Command{
		Use:   "start",
		Short: "Starts a local instance.",
		Args:  require.NoArgs,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.InitOutput(cmd.OutOrStdout(), startTemplate),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context())
		},
	}

	cmd.Flags().BoolVarP(&opts.debug, flag.Debug, flag.DebugShort, false, usage.Debug)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	return cmd
}
