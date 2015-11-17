package main

import (
	"github.com/spf13/cobra"
)

// TODO(labkode) Set log output to file in $HOME/.clawio/log

var childrenFlag bool
var checksumFlag string

func main() {

	var mainCmd = &cobra.Command{
		Use:   "clawio",
		Short: "ClawIO is a framework to test different tech stacks against sync protocols",
	}

	statCmd.Flags().BoolVar(&childrenFlag, "children", false, "retrieve children metadata")
	uploadCmd.Flags().StringVar(&checksumFlag, "checksum", "", "send client checksum")

	mainCmd.AddCommand(envCmd)
	mainCmd.AddCommand(loginCmd)
	mainCmd.AddCommand(logoutCmd)
	mainCmd.AddCommand(statCmd)
	mainCmd.AddCommand(rmCmd)
	mainCmd.AddCommand(homeCmd)
	mainCmd.AddCommand(mkdirCmd)
	mainCmd.AddCommand(uploadCmd)
	mainCmd.AddCommand(downloadCmd)
	mainCmd.AddCommand(mvCmd)

	mainCmd.Execute()
}
