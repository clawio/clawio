package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"os/user"
	"path"
)

var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Logout from ClawIO",
	Run:   logout,
}

func logout(cmd *cobra.Command, args []string) {

	u, err := user.Current()
	if err != nil {
		fmt.Println("Cannot access home directory")
		os.Exit(1)
	}

	err = os.RemoveAll(path.Join(u.HomeDir, ".clawio", "credentials"))
	if err != nil {
		fmt.Println("Cannot remove login credentials: ", err)
		os.Exit(1)
	}

	fmt.Println("Logged out")
}
