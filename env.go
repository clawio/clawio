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
	fmt.Printf("export CLAWIO_CLI_AUTH_ADDR=%s\n", os.Getenv("CLAWIO_CLI_AUTH_ADDR"))
	fmt.Printf("export CLAWIO_CLI_META_ADDR=%s\n", os.Getenv("CLAWIO_CLI_META_ADDR"))
	fmt.Printf("export CLAWIO_CLI_DATA_ADDR=%s\n", os.Getenv("CLAWIO_CLI_DATA_ADDR"))
}
