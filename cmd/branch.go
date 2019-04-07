// Copyright © 2019 Souvik Maji <souvikmaji94@gmail.com>
//
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
	"log"

	"github.com/spf13/cobra"
	git "gopkg.in/libgit2/git2go.v27"
)

// branchCmd represents the branch command
var branchCmd = &cobra.Command{
	Use:   "branch branchName",
	Short: "close a branch",
	Long: `Closes a git branch. Which means:
	deletes local git branch
	tags the remote branch archive`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		delete(args[0])
	},
}

func init() {
	rootCmd.AddCommand(branchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// branchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// branchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func delete(branchName string) {
	repo, err := git.OpenRepository(".")
	if err != nil {
		log.Fatal(err)
	}

	branch, err := repo.LookupBranch(branchName, git.BranchLocal)
	if err != nil {
		log.Fatal(err)
	}

	if err := branch.Delete(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("deleted local branch ", branchName)
}
