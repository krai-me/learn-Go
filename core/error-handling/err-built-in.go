package main

import (
	"fmt"
	"os"
)

// The 'os' package provides functions for interacting with the operating system.
// It includes functionalities like opening files, creating directories, and handling file system errors.

// The 'os.Open' function attempts to open a file.
// If the file cannot be opened (e.g., it does not exist), it returns an error.
// This error is part of Go's built-in error handling and provides details about the issue.

func main() {
	val, err := os.Open("nonexistentfile.txt")
	if err != nil {
		// If an error occurs, 'err' will contain details about the failure
		fmt.Println("Error:", err.Error())
	} else {
		// If no error occurs, 'val' will be a file descriptor
		fmt.Println(val)
	}
}
