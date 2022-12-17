package services

import (
	"blackbox/transport"
	"fmt"
)

type Service struct {
	Hostname string `gorm:"primaryKey"`
	Address  string
	Username string
	Password string
	Command  []Command
}

type Command struct {
	ServiceID string
	Cmd       string
}

// Run through our Configured Commands
func (server Service) RunCommands(input transport.InputTaker) {
	for _, cmds := range server.Command {
		fmt.Println(cmds.Cmd)
		_, err := fmt.Fprintf(input.GetInputPipe(), "%s\n", cmds.Cmd)
		if err != nil {
			fmt.Printf("Could not run command %q for service %q due to the following error:\n%s", cmds.ServiceID, cmds.Cmd, err)
		}
	}
}

////////////////
/// SETTERS  ///
////////////////

func (server *Service) SetServiceName(name string) {
	server.Hostname = name
}

func (server *Service) SetAddress(address string) {
	server.Address = address
}

func (server *Service) SetUser(username string) {
	server.Username = username
}

func (server *Service) SetPassword(password string) {
	server.Password = password
}

func (server *Service) AddCommands(commands []string) {
	for _, cmd := range commands {
		fmt.Println(cmd)
		server.Command = append(server.Command, Command{Cmd: cmd})
	}
}

////////////////
/// GRABBERS ///
////////////////

func (server Service) GrabServiceName() string {
	return server.Hostname
}

func (server Service) GrabAddress() string {
	return server.Address
}

func (server Service) GrabUser() string {
	return server.Username
}

func (server Service) GrabPassword() string {
	return server.Password
}

