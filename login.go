package main

import (
	pb "github.com/clawio/service.auth/proto"
	log "github.com/fatih/color"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"io/ioutil"
	"os"
	"os/user"
	"path"
)

var loginCmd = &cobra.Command{
	Use:   "login <username> <password>",
	Short: "Login into ClawIO",
	Run:   login,
}

func login(cmd *cobra.Command, args []string) {

	if len(args) != 2 {
		log.Red("You have to provide username and password")
		os.Exit(1)
	}

	addr := os.Getenv("CLAWIO_CLI_AUTHADDR")

	con, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Red("Cannot connect to server " + addr)
		os.Exit(1)
	}

	defer con.Close()

	c := pb.NewAuthClient(con)

	in := &pb.AuthRequest{}
	in.Username = args[0]
	in.Password = args[1]

	ctx := context.Background()

	res, err := c.Authenticate(ctx, in)
	if err != nil {
		if grpc.Code(err) == codes.Unauthenticated {
			log.Red("Invalid username or password")
			os.Exit(1)
		}
		log.Red("Cannot connect to server " + addr)
		os.Exit(1)
	}

	// Save token into $HOMR/.clawio/credentials
	u, err := user.Current()
	if err != nil {
		log.Red("Cannot access your home directory")
		os.Exit(1)
	}

	err = os.MkdirAll(path.Join(u.HomeDir, ".clawio"), 0755)
	if err != nil {
		log.Red("Cannot create $HOME/.clawio configuration directory")
		os.Exit(1)
	}

	err = ioutil.WriteFile(path.Join(u.HomeDir, ".clawio", "credentials"), []byte(res.Token), 0644)
	if err != nil {
		log.Red("Cannot save credentials into $HOME/.clawio/credentials")
		os.Exit(1)
	}

	log.Green("Logged in as " + in.Username)
	os.Exit(0)
}
