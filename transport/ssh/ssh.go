package sshTransport

import (
	"blackbox/transport"
	"fmt"
	"io"
	"log"

	"golang.org/x/crypto/ssh"
)

type SSHConn struct {
	client  *ssh.Client
	session *ssh.Session
	inpipe  io.WriteCloser
}

// Create SSH Connection with Auth & StdInPipe
func Connect(service transport.IConnectCreds) SSHConn {
	ssh := SSHConn{}
	ssh.newClient(service)
	ssh.newSession()
	ssh.newInputPipe()
	return ssh
}

////////////////
/// SETTERS  ///
////////////////

// Create a new SSH Client with Provided Authentication
func (shell *SSHConn) newClient(sshConf transport.IConnectCreds) {
	client, err := ssh.Dial("tcp", sshConf.GrabAddress(), createConfig(sshConf))
	if err != nil {
		fmt.Println("We were unable to connect to the server via SSH. Error: ", err)
	}

	shell.client = client
}

func (shell *SSHConn) Close() {
	shell.client.Close()
	shell.session.Close()
	shell.inpipe.Close()
}

// Creates Client Configuration for Making Session
func createConfig(sshConf transport.IConnectCreds) *ssh.ClientConfig {
	return &ssh.ClientConfig{
		User: sshConf.GrabUser(),
		Auth: []ssh.AuthMethod{
			ssh.Password(sshConf.GrabPassword()),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
}

// Connect to our Service to Intended Destination
func (shell *SSHConn) newSession() {
	session, err := shell.client.NewSession()
	if err != nil {
		log.Fatal("Unable to initiate remote connection to the SSH host. Error:", err)
	}

	shell.session = session
}

// Return our Input Pipe from SSH Connection
func (shell *SSHConn) newInputPipe() {
	inPipe, err := shell.session.StdinPipe()
	if err != nil {
		log.Fatal("Unable to get Stdinpipe from Session. Error:", err)
	}

	// Shell starts a login shell on the remote host.
	err = shell.session.Shell()
	if err != nil {
		log.Fatal("Unable to initiate login shell on the SSH host. Error:", err)
	}

	shell.inpipe = inPipe
}

////////////////
/// GETTERS  ///
////////////////

func (shell SSHConn) GetClient() *ssh.Client {
	return shell.client
}

func (shell SSHConn) GetSession() *ssh.Session {
	return shell.session
}

func (shell SSHConn) GetInputPipe() io.WriteCloser {
	return shell.inpipe
}
