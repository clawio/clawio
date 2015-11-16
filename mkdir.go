package main

import (
	pb "github.com/clawio/service.localstore.meta/proto"
	log "github.com/fatih/color"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"os"
)

var mkdirCmd = &cobra.Command{
	Use:   "mkdir <path>",
	Short: "Creates a container",
	Run:   mkdir,
}

func mkdir(cmd *cobra.Command, args []string) {

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

	in := &pb.MkdirReq{}
	in.AccessToken = token
	in.Path = args[0]

	ctx := context.Background()

	_, err = c.Mkdir(ctx, in)
	if err != nil {
		log.Red("Cannot create container: " + err.Error())
		os.Exit(1)
	}

	log.Green("Created container " + args[0])
}
