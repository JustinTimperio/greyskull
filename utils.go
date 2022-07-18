package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"strings"
)

// createFile opens and writes bytes to a path
func createFile(path string, fileBytes []byte, perms fs.FileMode) (err error) {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	os.Chmod(path, perms)

	_, err = file.Write(fileBytes)
	return err
}

// readFile reads a file from path and return bytes and error
func readFile(path string) (fileBytes []byte, err error) {
	return ioutil.ReadFile(path)
}

// pathExists returns whether the given file or directory exists
func pathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

// askForConfirmation asks the user for confirmation. A user must type in "yes" or "no" and
// then press enter. It has fuzzy matching, so "y", "Y", "yes", "YES", and "Yes" all count as
// confirmations. If the input is not recognized, it will ask again. The function does not return
// until it gets a valid response from the user.
func askForConfirmation(s string) bool {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("%s [y/n]: ", s)

		response, err := reader.ReadString('\n')
		if err != nil {
			return false
		}

		response = strings.ToLower(strings.TrimSpace(response))

		if response == "y" || response == "yes" {
			return true
		} else if response == "n" || response == "no" {
			return false
		}
	}
}
