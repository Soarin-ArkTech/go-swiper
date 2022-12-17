package scpTransport

import (
	filestore "blackbox/storage/local"
	sshTransport "blackbox/transport/ssh"
	"context"
	"fmt"
	"time"

	"github.com/bramvdbogaerde/go-scp"
)

type SCPTransport struct {
	scp.Client
}

// Extract SCP Client from SSH Connection
func GetClient(client sshTransport.SSHConn) SCPTransport {
	scpClient, err := scp.NewClientBySSH(client.GetClient())
	if err != nil {
		fmt.Println("Unable to retrieve SCP client through SSH. Error: ", err)
	}
	return SCPTransport{scpClient}
}

// Initiate a File Download on the Remote end of the SSH session
func (transit SCPTransport) GetRemoteFile(localFile string, remoteFile string) {
	fmt.Println(localFile, remoteFile)
	time.Sleep(time.Second)
	err := transit.CopyFromRemote(context.Background(), filestore.CreateLocalFile(localFile), remoteFile)
	if err != nil {
		fmt.Println("Unable to download file from remote host. Error: ", err)
	}
}
