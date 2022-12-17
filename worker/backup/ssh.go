package sshWorker

import (
	filestore "blackbox/storage/local"
	transport "blackbox/transport"
	ssh "blackbox/transport/ssh"
	scp "blackbox/transport/ssh/scp"
	"fmt"
)

// Run Backup Job using SSH Transport
func SSHBackup(job transport.IConnectedJob) bool {
	// Connect to Host and Return a Pipe to Insert Commands Into
	sshCon := ssh.Connect(job)
	defer sshCon.Close()

	filestore.CreateBackupDir(sshCon)
	filestore.SetDirectory(sshCon)

	job.RunCommands(sshCon)

	return true
}

func SCPBackups(job transport.IConnectedJob) {
	conn := ssh.Connect(job)
	filestore.CreateBackupDir(conn)
	filestore.SetDirectory(conn)
	dir, file := filestore.ArchiveBackupDir(conn)
	scp.GetClient(conn).
		GetRemoteFile(file, fmt.Sprintf(
			"/home/%s/%s/%s", job.GrabUser(), dir, file,
		))
}
