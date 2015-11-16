package main

import (
	pb "github.com/clawio/service.localstore.meta/proto"
	log "github.com/fatih/color"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"os"
)

var rmCmd = &cobra.Command{
	Use:   "rm <path>",
	Short: "Remove a resource",
	Run:   rm,
}

func rm(cmd *cobra.Command, args []string) {

	if len(args) != 1 {
		log.Red("You have to provide a path")
		os.Exit(1)
	}

	token, err := getToken()
	if err != nil {
		log.Red("Authentication required")
		os.Exit(1)
	}

	addr := os.Getenv("CLAWIO_CLI_LOCALSTOREMETA_ADDR")

	con, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Red("Cannot connect to server " + addr)
		os.Exit(1)
	}

	defer con.Close()

	c := pb.NewLocalClient(con)

	in := &pb.RmReq{}
	in.AccessToken = token
	in.Path = args[0]

	ctx := context.Background()

	_, err = c.Rm(ctx, in)
	if err != nil {
		log.Red("Cannot remove resource: " + err.Error())
		os.Exit(1)
	}

	log.Green("Removed " + args[0])
}
