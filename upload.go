package main

import (
	"fmt"
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
		fmt.Println("You have to provide a local and remote path")
		os.Exit(1)
	}

	token, err := getToken()
	if err != nil {
		fmt.Println("Authentication required")
		os.Exit(1)
	}

	addr := os.Getenv("CLAWIO_CLI_LOCALSTOREDATA_ADDR")

	fd, err := os.Open(args[0])
	if err != nil {
		fmt.Println("Cannot read local file: " + err.Error())
		os.Exit(1)
	}

	defer fd.Close()

	c := &http.Client{}
	req, err := http.NewRequest("PUT", "http://"+path.Join(addr, args[1]), fd)
	if err != nil {
		fmt.Println("Cannot created upload request: " + err.Error())
		os.Exit(1)
	}

	req.Header.Add("Content-Type", "application/octet-stream")
	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Add("CIO-Checksum", checksumFlag)

	res, err := c.Do(req)
	if err != nil {
		fmt.Println("Upload failed: " + err.Error())
		os.Exit(1)
	}

	defer res.Body.Close()

	if res.StatusCode == 412 {
		fmt.Printf("Object %s was corrupted during upload and server did not save it\n", args[0])
		os.Exit(1)
	}

	fmt.Println("Uploaded " + args[0] + " to " + args[1])
}
