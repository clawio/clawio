package main

import (
	"io/ioutil"
	"os/user"
	"path"
)

func getToken() (string, error) {

	u, err := user.Current()
	if err != nil {
		return "", err
	}

	token, err := ioutil.ReadFile(path.Join(u.HomeDir, ".clawio", "credentials"))
	if err != nil {
		return "", err
	}

	return string(token), nil
}
