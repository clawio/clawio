package main

import (
	log "github.com/fatih/color"
	"github.com/spf13/cobra"
	"io"
	"net/http"
	"os"
	"path"
)

var downloadCmd = &cobra.Command{
	Use:   "download <remotepath> <localpath>",
	Short: "Download an object",
	Run:   download,
}

func download(cmd *cobra.Command, args []string) {

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

	c := &http.Client{}
	req, err := http.NewRequest("GET", "http://"+path.Join(addr, args[0]), nil)
	if err != nil {
		log.Red("Cannot created download request: " + err.Error())
		os.Exit(1)
	}

	req.Header.Add("Authorization", "Bearer "+token)

	res, err := c.Do(req)
	if err != nil {
		log.Red("Download failed: " + err.Error())
		os.Exit(1)
	}

	if res.StatusCode != 200 {
		switch res.StatusCode {
		case 400:
			log.Red("Cannot download a container")
			os.Exit(1)
		}
	}

	fd, err := os.Create(args[1])
	if err != nil {
		log.Red("Cannot create local file: " + err.Error())
		os.Exit(1)
	}

	defer fd.Close()

	_, err = io.Copy(fd, res.Body)
	if err != nil {
		log.Red("Cannot download object: " + err.Error())
		os.Exit(1)
	}

	defer res.Body.Close()

	log.Green("Downloaded " + args[0] + " to " + args[1])
}
