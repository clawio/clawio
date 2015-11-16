package main

import (
	"fmt"
	pb "github.com/clawio/service.localstore.meta/proto"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"os"
)

var mvCmd = &cobra.Command{
	Use:   "mv <src> <dst>",
	Short: "Moves a resource",
	Run:   mv,
}

func mv(cmd *cobra.Command, args []string) {

	if len(args) != 2 {
		fmt.Println("You have to provide src and dst paths")
		os.Exit(1)
	}

	token, err := getToken()
	if err != nil {
		fmt.Println("Authentication required")
		os.Exit(1)
	}

	addr := os.Getenv("CLAWIO_CLI_LOCALSTOREMETA_ADDR")

	con, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Cannot connect to server " + addr)
		os.Exit(1)
	}

	defer con.Close()

	c := pb.NewLocalClient(con)

	in := &pb.MvReq{}
	in.AccessToken = token
	in.Src = args[0]
	in.Dst = args[1]

	ctx := context.Background()

	_, err = c.Mv(ctx, in)
	if err != nil {
		fmt.Println("Cannot mv resource: " + err.Error())
		os.Exit(1)
	}

	fmt.Println("Moved " + args[0] + " to " + args[1])
}
