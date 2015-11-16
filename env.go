package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var envCmd = &cobra.Command{
	Use:   "env",
	Short: "Print env vars used by the cli",
	Run:   env,
}

func env(cmd *cobra.Command, args []string) {
	fmt.Println("clawio_cli_auth_addr\t" + os.Getenv("CLAWIO_CLI_AUTHADDR"))
	fmt.Println("clawio_cli_localstoremeta_addr\t" + os.Getenv("CLAWIO_CLI_LOCALSTOREMETA_ADDR"))
	fmt.Println("clawio_cli_localstoredata_addr\t" + os.Getenv("CLAWIO_CLI_LOCALSTOREDATA_ADDR"))
}
