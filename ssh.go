package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/bramvdbogaerde/go-scp"
	"golang.org/x/crypto/ssh"
)

func vyBackup(vy BackupList) {

	// Get Date & Time for Filenames
	date := time.Now()
	dateParsed := date.Format("01-02-2006")

	// Form New Connection & Ready to Input Commands
	stdin, conn, err := sshNew(vy)
	if err != nil {
		log.Fatalln("PANIC: Unable to Create SSH Connection: ", err)
	}

	// mmmyeah commands
	vyCmds := []string{
		"source /opt/vyatta/etc/functions/script-template",
		"run show conf commands >" + vy.Hostname + "-" + dateParsed + ".txt",
		" ",
	}

	// Go through commands & execute
	for k, cmd := range vyCmds {
		_, err := fmt.Fprintf(stdin, "%s\n", cmd)
		if err != nil {
			log.Fatal(err)
		}

		if k == len(vyCmds) {
			time.Sleep(time.Second * 2)
			break
		}
	}

	// Create SCP Client using Existing SSH Connection
	scpClient, err := scp.NewClientBySSH(conn)
	if err != nil {
		log.Fatal("FUCK We're unable to create our SCP Client! \n", err)
	}

	// Create the .txt for SCP to copy into
	formattedConfig, err := os.Create("./backups/" + vy.Hostname + "/" + vy.Hostname + "-" + dateParsed + ".txt")
	if err != nil {
		fmt.Println("We were unable to create the local .txt file for " + vy.Hostname)
	}

	// honestly idk what this does, i got lucky by putting it here and somehow worked. note to self -> research contexts l8r
	ctx := context.Background()

	// SCP Configs Over then Safely Close
	err = scpClient.CopyFromRemote(ctx, formattedConfig, "/home/"+vy.User+"/"+vy.Hostname+"-"+dateParsed+".txt")
	if err != nil {
		fmt.Printf("Unable to Download via SCP, error reason: %v", err)
	}

	scpClient.Close()
}

// Start a new SSH connection and return client & stdinpipe
func sshNew(vy BackupList) (io.WriteCloser, *ssh.Client, error) {

	// Allow us to connect even if host is not in known_hosts file
	hostKeyCallback := ssh.InsecureIgnoreHostKey()

	// Cient Config to load into SSH
	config := &ssh.ClientConfig{
		User: vy.User,
		Auth: []ssh.AuthMethod{
			ssh.Password(vy.Passwd),
		},

		HostKeyCallback: hostKeyCallback,
	}

	// Connection Info
	conn, err := ssh.Dial("tcp", vy.Address, config)
	if err != nil {
		log.Fatalf("\nERROR: We were unable to connect to %v... ", vy.Hostname+"\n ")
	}

	// Create SSH Session
	session, err := conn.NewSession()
	if err != nil {
		log.Fatal(err)
	}

	// StdinPipe for sending input (commands)
	sshIn, err := session.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}

	// Output for session
	//session.Stdout = os.Stdout
	//session.Stderr = os.Stderr

	// Shell starts a login shell on the remote host.
	session.Shell()

	// Modify our Terminal SSH Mode (its default rn, but putting this here if I want to add flags)
	sshMode := ssh.TerminalModes{
		ssh.ECHO: 0,
	}

	// Make it so I can go deep inside VyOS
	session.RequestPty("vbash --posix -s", 1920, 1080, sshMode)

	return sshIn, conn, err
}
