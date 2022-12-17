package landing

import (
	"blackbox/services"
	"bufio"
	"fmt"
	"os"
)

func BasicUI() services.Service {
	var service services.Service

	fmt.Println("Shitty Interface for Backup Service Creation")
	service.SetServiceName(createPrompt("Service Name: "))
	service.SetAddress(createPrompt("Connection Address: "))
	service.SetUser(createPrompt("Username: "))
	service.SetPassword(createPrompt("Password: "))
	service.AddCommands(createCommandsPrompt())

	return service
}

func createPrompt(prompt string) string {
	fmt.Print(prompt)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	return scanner.Text()
}

func createCommandsPrompt() []string {
	var cmds []string

	fmt.Println("\nPlease input the commands you wish to execute on the remote host.")
	fmt.Println("You can add multiple, when you are finished just type DONE")

	for {
		switch cmd := createPrompt("Command: "); {
		case cmd == "DONE":
			return cmds

		default:
			cmds = append(cmds, cmd)
		}
	}
}
