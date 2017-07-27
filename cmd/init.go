// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"os"

	vp "github.com/shyang107/go-twinvoices"
	"github.com/shyang107/go-twinvoices/util"
	"github.com/urfave/cli"
)

// initCmd represents the init command
var initCmd = cli.Command{
	Name:        "init",
	Aliases:     []string{"i"},
	Usage:       "Empty and initialize the using database",
	Description: `Empty the database, and initializes database`,
	Action:      initAction,
}

func init() {
	util.PfBlue("init.init called\n")
	RootApp.Commands = append(RootApp.Commands, initCmd)
}

func initAction(c *cli.Context) error {
	util.PfBlue("init.initAction called\n")
	if err := vp.Initialdb(); err != nil {
		return err
	}
	os.Exit(0)
	return nil
}
