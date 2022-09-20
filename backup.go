package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func createFolders(in Services) {

	//Check and Make Backups Folder
	_, err := os.ReadDir("./backups/")
	if err != nil {
		os.Mkdir("backups", 0777)
	}

	// Create individual host folders under backup folder
	os.Mkdir("backups/"+in.GrabHostname(), 0777)

}

func createFiles(in Services) *os.File {

	// Define our time format
	dateParsed := parseDate()

	// Create our file for SCP to copy into
	file, err := os.Create("./backups/" + in.GrabHostname() + "/" + in.GrabHostname() + "-" + dateParsed + ".txt")
	if err != nil {
		log.Fatalf("We were unable to create the files locally for SCP to copy into! Error: ", err)
	}

	return file
}

func runCMDs(cmdlist []string, stdin io.WriteCloser) {

	// Iterate over commands and give 2.0s for latency and everything to be groovy
	for _, cmd := range cmdlist {
		fmt.Fprintf(stdin, "%s\n", cmd)
		time.Sleep(time.Millisecond * 2500)
	}

}

func LaunchBackup(in Services) error {

	// I know this is inefficient to repeatedly check & create folders/files but leave me alone I'll fix it
	createFolders(in)
	file := createFiles(in)

	// Initiate our SSH Connections
	conn, stdin, err := SSHConnect(in)
	if err != nil {
		log.Fatal(err)
	}

	// Execute our methods that hold our commands to send thru SSH
	runCMDs(in.backupCMD(), stdin)

	// SCP Backups
	err = SCPBackups(conn, in, file)
	if err != nil {
		log.Fatalf("SCP backups was not a success.")
	}

	return err
}
