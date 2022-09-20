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

func parseDate() string {
	timeDate := time.Now()
	ParsedDate := timeDate.Format("01-02-2006")

	return ParsedDate
}

func SSHConnect(in Services) (*ssh.Client, io.WriteCloser, error) {

	// Pass our SSH Config
	config := SSHConfig(in)

	// Initiate Connection
	conn, err := ssh.Dial("tcp", in.GrabAddress(), config)
	if err != nil {
		log.Fatalf("\nERROR: We were unable to connect to %v... ", in, "\n ")
	}

	// Create Session & Input Pipe for Commands
	stdin, err := SSHSession(conn)
	if err != nil {
		log.Fatalf("We were unble to create the SSH Session. Error: %s", err)
	}

	return conn, stdin, err
}

func SSHSession(client *ssh.Client) (io.WriteCloser, error) {

	// Create SSH Session
	session, err := client.NewSession()
	if err != nil {
		log.Fatal(err)
	}

	// Close when done
	// defer session.Close()

	// StdinPipe for sending input (commands)
	sshIn, err := session.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}

	// Shell starts a login shell on the remote host.
	session.Shell()

	// Make it so I can go deep inside VyOS
	session.RequestPty("vbash --posix -s", 800, 600, ssh.TerminalModes{})

	// Output for Session [Debug]
	session.Stdout = os.Stdout
	session.Stderr = os.Stderr

	return sshIn, err
}

func SSHConfig(in Services) *ssh.ClientConfig {

	// Allow us to connect even if host is not in known_hosts file
	hostKeyCallback := ssh.InsecureIgnoreHostKey()

	// Cient Config to load into SSH
	config := &ssh.ClientConfig{
		User: in.GrabUsername(),
		Auth: []ssh.AuthMethod{
			ssh.Password(in.GrabPassword()),
		},

		HostKeyCallback: hostKeyCallback,
	}

	// Return pointer to our config
	return config
}

func SCPBackups(client *ssh.Client, in Services, file *os.File) error {

	// Get our date n format it
	dateParsed := parseDate()

	// Create new SCP Client over our Existing SSH Connection
	scpClient, err := scp.NewClientBySSH(client)
	if err != nil {
		log.Fatalf("We were unable to create the SCP connection. Error: %s", err)
	}

	// Create our context to allow timeouts and other features for our SCP functions
	context := context.Background()

	// Copy our file to the appropriate directory
	err = scpClient.CopyFromRemote(context, file, "/home/"+in.GrabUsername()+"/"+in.GrabHostname()+"-"+dateParsed+".txt")
	if err != nil {
		fmt.Println("The file could not be downloaded via SCP.")
	}

	defer scpClient.Close()

	return err
}
