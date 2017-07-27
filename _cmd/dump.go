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
	"fmt"

	inv "github.com/shyang107/go-twinvoices"
	"github.com/shyang107/go-twinvoices/util"
	"github.com/spf13/cobra"
)

var dfile string

// dumpCmd represents the dump command
var dumpCmd = &cobra.Command{
	Use:   "dump",
	Short: "Dump all records from databse",
	Long:  `Dump all recirds from database into .json file.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("dump called")
	},
}

func init() {
	RootCmd.AddCommand(dumpCmd)

	// dumpCmd.PersistentFlags().String("foo", "", "A help for foo")
	// dumpCmd.PersistentFlags().StringVar(&dfile,"file","","")

	dumpCmd.Flags().StringVarP(&dfile, "file", "f",
		inv.DefaultConfig.DumpFilename,
		util.Sf("config file (default is %v)", inv.DefaultConfig.DumpFilename))
}
