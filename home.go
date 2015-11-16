package main

import (
	pb "github.com/clawio/service.localstore.meta/proto"
	log "github.com/fatih/color"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"os"
)

var homeCmd = &cobra.Command{
	Use:   "home",
	Short: "Create user home directory",
	Run:   home,
}

func home(cmd *cobra.Command, args []string) {

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

	in := &pb.HomeReq{}
	in.AccessToken = token

	ctx := context.Background()

	_, err = c.Home(ctx, in)
	if err != nil {
		log.Red("Cannot create homedir: " + err.Error())
		os.Exit(1)
	}

	log.Green("Home directory created")
}
