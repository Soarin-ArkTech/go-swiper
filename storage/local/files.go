package filestore

import (
	"blackbox/transport"
	"fmt"
	"io"
	"os"
	"time"

	"golang.org/x/crypto/ssh"
)

var dir string = "swiper-" + formatDate()
var file string = dir + ".tar.gz"

func ArchiveBackupDir(sshCon transport.InputTaker) (string, string) {
	fmt.Fprintf(sshCon.GetInputPipe(), "%s\n", fmt.Sprintf("tar -czf %s %s", file, "*"))
	return dir, file
}

func CreateBackupDir(pipe transport.InputTaker) error {
	_, err := fmt.Fprintf(pipe.GetInputPipe(), "%s\n", fmt.Sprintf("mkdir "+dir))
	return err
}

func SetDirectory(pipe transport.InputTaker) {
	fmt.Fprintf(pipe.GetInputPipe(), "%s\n", "cd "+dir)
}

func CreateLocalFile(name string) *os.File {
	local, err := os.Create(name)
	if err != nil {
		fmt.Printf("We were unable to create %q with the following error: \n%s", name, err)
	}

	return local
}

func CheckFileIntegrity(pipe io.WriteCloser, session *ssh.Session) {
	_, err := fmt.Fprintf(pipe, "%s\n", fmt.Sprintf("tar -tf "+file))
	if err != nil {
		fmt.Println("Unable to try to check Tar integritry. Error:\n", err)
	}

	bytes, _ := session.Output("ls")
	fmt.Println(bytes)
}

func formatDate() string {
	year, month, day := time.Now().Date()
	return fmt.Sprintf("%s-%v-%v", month, day, year)
}
