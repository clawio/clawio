package main

import (
	"fmt"
	pb "github.com/clawio/clawio/proto/metadata"
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
		fmt.Println("Authentication required")
		os.Exit(1)
	}

	addr := os.Getenv("CLAWIO_CLI_META_ADDR")

	con, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Cannot connect to server " + addr)
		os.Exit(1)
	}

	defer con.Close()

	c := pb.NewMetaClient(con)

	in := &pb.HomeReq{}
	in.AccessToken = token

	ctx := context.Background()

	_, err = c.Home(ctx, in)
	if err != nil {
		fmt.Println("Cannot create homedir: " + err.Error())
		os.Exit(1)
	}

	fmt.Println("Home directory created")
}
