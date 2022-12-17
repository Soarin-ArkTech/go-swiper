package main

import (
	"blackbox/repo"
	sshWorker "blackbox/worker/backup"
	"fmt"
)

func main() {
	// Factory to Create SQLite DB named blackbox.db
	db := repo.DBFactory("sqlite", "blackbox.db").ReturnDB()

	// newService := landing.BasicUI()

	services := db.LoadAllServices()

	for _, service := range services {
		done := sshWorker.SSHBackup(service)
		if done {
			sshWorker.SCPBackups(service)
		}
		fmt.Println("Backup completed for", service.GrabServiceName())
	}
}
