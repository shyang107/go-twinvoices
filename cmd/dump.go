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
	vp "github.com/shyang107/go-twinvoices"
	"github.com/shyang107/go-twinvoices/util"
	"github.com/urfave/cli"
)

var dfile string

// dumpCmd represents the dump command
var dumpCmd = cli.Command{
	Name:        "dump",
	Aliases:     []string{"d"},
	Usage:       "Dump all records from databse",
	Description: "Dump all recirds from database into .json file.",
	Action:      dumpAction,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "file,f",
			Usage: "specify the dump path",
		},
	},
}

func init() {
	util.Verbose = vp.Cfg.Verbose
	util.ColorsOn = vp.Cfg.ColorsOn
	util.PfBlue("dump.init called\n")
	RootApp.Commands = append(RootApp.Commands, dumpCmd)
}

func dumpAction(c *cli.Context) error {
	util.PfBlue("dump.dumpAction called\n")
	pstat("  > Dumping data from %q ...\n", vp.Cfg.DBfilename)
	return nil
}
