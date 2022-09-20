package main

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

func loadConf() []BackupList {

	// Open it UP G
	conf, err := openConf()
	if err != nil {
		log.Fatal(err)
	}

	// Create Variable Instance of BackupList Slice for our JSON
	var OurList []BackupList

	// Spit Formatted JSON out to Struct
	err = json.Unmarshal(conf, &OurList)
	if err != nil {

		log.Fatal(err)
	}

	return OurList
}

func openConf() ([]byte, error) {
	// Load File
	BareFile, err := os.Open("config.json")
	if err != nil {
		log.Fatal("Couldn't find config.json!")
	}
	defer BareFile.Close()

	// Read to Bytes
	ConfBytes, err := io.ReadAll(BareFile)
	if err != nil {
		log.Fatal(err)
	}

	return ConfBytes, err
}
