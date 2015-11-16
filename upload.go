package main

import (
	log "github.com/fatih/color"
	"github.com/spf13/cobra"
	"net/http"
	"os"
	"path"
)

var uploadCmd = &cobra.Command{
	Use:   "upload <localpath> <remotepath>",
	Short: "Upload an object",
	Run:   upload,
}

func upload(cmd *cobra.Command, args []string) {

	if len(args) != 2 {
		log.Red("You have to provide a local and remote path")
		os.Exit(1)
	}

	token, err := getToken()
	if err != nil {
		log.Red("Authentication required")
		os.Exit(1)
	}

	addr := os.Getenv("CLAWIO_CLI_LOCALSTOREDATA_ADDR")

	fd, err := os.Open(args[0])
	if err != nil {
		log.Red("Cannot read local file: " + err.Error())
		os.Exit(1)
	}

	defer fd.Close()

	c := &http.Client{}
	req, err := http.NewRequest("PUT", "http://"+path.Join(addr, args[1]), fd)
	if err != nil {
		log.Red("Cannot created upload request: " + err.Error())
		os.Exit(1)
	}

	req.Header.Add("Content-Type", "application/octet-stream")
	req.Header.Add("Authorization", "Bearer "+token)

	res, err := c.Do(req)
	if err != nil {
		log.Red("Upload failed: " + err.Error())
		os.Exit(1)
	}

	defer res.Body.Close()

	log.Green("Uploaded " + args[0] + " to " + args[1])
}
