package main

import (
	"github.com/spf13/cobra"
)

// TODO(labkode) Set log output to file in $HOME/.clawio/log

var childrenFlag bool

func main() {

	var mainCmd = &cobra.Command{
		Use:   "claw",
		Short: "ClawIO is a framework to test different tech stacks against sync protocols",
	}

	statCmd.Flags().BoolVar(&childrenFlag, "children", false, "retrieve children metadata")

	mainCmd.AddCommand(envCmd)
	mainCmd.AddCommand(loginCmd)
	mainCmd.AddCommand(logoutCmd)
	mainCmd.AddCommand(statCmd)
	mainCmd.AddCommand(rmCmd)
	mainCmd.AddCommand(homeCmd)
	mainCmd.AddCommand(mkdirCmd)
	mainCmd.AddCommand(uploadCmd)
	mainCmd.AddCommand(downloadCmd)

	mainCmd.Execute()
}
