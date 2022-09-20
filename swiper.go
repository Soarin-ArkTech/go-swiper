package main

import (
	"fmt"
	"sync"
)

func main() {
	// WIP, groundwork for extra services laid.

	// Load Config
	config := loadConf()

	// Create wait groups for GoRoutines
	var wg sync.WaitGroup
	wg.Add(len(config))

	// Loop over config for v.VyOS Index and start a goroutine for each
	for _, v := range config {
		go func(v VyOS) {
			defer wg.Done()
			err := LaunchBackup(v)
			if err != nil {
				fmt.Println("Backup Routine Failed!")
			}

		}(v.VyOS)
	}

	// Wait for all goroutines to close
	wg.Wait()

	// I really hope this is the outcome!
	fmt.Println("Backup complete!")
}
