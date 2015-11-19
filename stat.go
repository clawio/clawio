package main

import (
	"fmt"
	pb "github.com/clawio/clawio/proto/metadata"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"os"
	"text/tabwriter"
)

var statCmd = &cobra.Command{
	Use:   "stat <path>",
	Short: "Stat a resource",
	Run:   stat,
}

func stat(cmd *cobra.Command, args []string) {

	if len(args) != 1 {
		fmt.Println("You have to provide a path")
		os.Exit(1)
	}

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

	in := &pb.StatReq{}
	in.AccessToken = token
	in.Path = args[0]
	in.Children = childrenFlag

	ctx := context.Background()

	res, err := c.Stat(ctx, in)
	if err != nil {
		fmt.Println("Cannot stat resource: " + err.Error())
		os.Exit(1)
	}

	tabWriter := tabwriter.NewWriter(os.Stdout, 0, 4, 2, ' ', 0)
	defer tabWriter.Flush()

	fmt.Fprintln(tabWriter, "ID\tPath\tContainer\tSize\tModified\tPermissions\tETag\tMime\tChecksum")

	fmt.Fprintf(tabWriter, "%s\t%s\t%t\t%d\t%d\t%d\t%s\t%s\t%s\n",
		res.Id, res.Path, res.IsContainer, res.Size, res.Modified, res.Permissions, res.Etag, res.MimeType, res.Checksum)

	for _, child := range res.GetChildren() {
		fmt.Fprintf(tabWriter, "%s\t%s\t%t\t%d\t%d\t%d\t%s\t%s\t%s\n",
			child.Id, child.Path, child.IsContainer, child.Size, child.Modified, child.Permissions, child.Etag, child.MimeType, child.Checksum)
	}
}
