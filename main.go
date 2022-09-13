package main

import (
	"fmt"
	"os"
	"sync"
)

func main() {

	// Load config.json
	backupList := loadConf()

	// Make worker group and set the value to # of routers
	var holdup sync.WaitGroup
	holdup.Add(len(backupList))

	// Go through Routers and Backup
	for k, v := range backupList {
		fmt.Printf("%v is backing up!\n", v.Hostname)

		// Make our directories if not made
		os.Mkdir(v.Hostname, 0777)

		// Anonymous function, each k starts a new goroutine for vyBackup
		go func(k int) {
			defer holdup.Done()
			vyBackup(backupList[k])
		}(k)
	}

	// Make sure to wait for all worker groups to close out
	holdup.Wait()

}

type VyOS struct {
	Hostname string `json:"hostname"`
	Address  string `json:"address"`
	User     string `json:"user"`
	Passwd   string `json:"passwd"`
}

type Router struct {
	VyOS `json:"vyOS"`
}

type BackupList struct {
	VyOS   `json:"vyOS"`
	Router `json:"Router"`
}
