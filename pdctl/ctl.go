// Copyright 2016 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package pdctl

import (
	"github.com/pingcap/pd/pdctl/command"
	"github.com/spf13/cobra"
)

// CommandFlags are flags that used in all Commands
type CommandFlags struct {
	URL string
}

var (
	rootCmd = &cobra.Command{
		Use:   "pdctl",
		Short: "Placement Driver control",
	}
	commandFlags = CommandFlags{}
)

func init() {
	rootCmd.PersistentFlags().StringVarP(&commandFlags.URL, "pd", "u", "http://127.0.0.1:2379", "pd address")
	rootCmd.AddCommand(
		command.NewConfigCommand(),
		command.NewRegionCommand(),
		command.NewStoreCommand(),
		command.NewMemberCommand(),
		command.NewExitCommand(),
	)
	cobra.EnablePrefixMatching = true
}

// Start run Command
func Start(args []string) (string, error) {
	rootCmd.SetArgs(args)
	rootCmd.SilenceErrors = true
	rootCmd.SetUsageTemplate(command.UsageTemplate)
	if err := rootCmd.Execute(); err != nil {
		return rootCmd.UsageString(), err
	}
	return "", nil
}
