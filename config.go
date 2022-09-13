package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

func loadConf() []BackupList {
	// Open JSON Config
	jsonFile, err := os.Open("config.json")
	if err != nil {
		log.Fatalln("Couldn't open config.json!")
	}

	// Close at the End
	defer jsonFile.Close()

	// Turn Read JSON into Bytes
	byteMe, _ := ioutil.ReadAll(jsonFile)

	// Buffer the output from the Json Formatting Func
	var jsonBuff bytes.Buffer

	// This is the struct we'll unmarshal into
	var toStruct []BackupList
	// Insert Formatted JSON to our Buffer
	err = json.Indent(&jsonBuff, byteMe, "", "   ")
	if err != nil {
		log.Fatalln("FUCK we couldn't parse the JSON!")
	}

	//Take our parsed JSON and install to structs
	_ = json.Unmarshal(jsonBuff.Bytes(), &toStruct)
	return toStruct
}
