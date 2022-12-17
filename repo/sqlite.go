package repo

import (
	"fmt"
	"log"

	"blackbox/services"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (sql SQL) ReturnDB() SQL {
	return sql
}

// Load our DB
func (sql *SQL) LoadDB(name string) {
	database, err := gorm.Open(sqlite.Open(name), &gorm.Config{})
	if err != nil {
		log.Fatalln("Sorry scoobs, we were unable to open the SQLite file. Error: \n", err)
	}

	sql.Data = database
	sql.structureTables(services.Service{}, services.Command{})
}

// Load Specific Service
func (sql SQL) LoadService(name string) services.Service {
	var Result services.Service
	err := sql.Data.Where(&services.Service{}).First(&Result, "hostname = ?", name)
	if err != nil {
		fmt.Println("Failed to query database. Error: ", err)
	}

	return Result
}

// Load All Services
func (sql SQL) LoadAllServices() []services.Service {
	var Result []services.Service
	err := sql.Data.Preload(clause.Associations).Find(&Result).Error
	if err != nil {
		fmt.Println("Failed to query database. Error: ", err)
	}

	return Result
}

// Delete Servive from DB
func (sql *SQL) DeleteService(table interface{}) bool {
	err := sql.Data.Select(clause.Associations).Delete(table)
	if err != nil {
		fmt.Println("Failed to delete database. Error: ", err)
		return false
	}

	return true
}

func (sql *SQL) AddService(table interface{}) bool {
	err := sql.Data.Create(table).Error
	if err != nil {
		fmt.Println("We were unable to create the database. Error:\n", err)
		return false
	}

	return true
}

// Create our Empty Tables to Insert Data Into
func (sql *SQL) structureTables(tables ...interface{}) {
	for _, table := range tables {
		sql.Data.AutoMigrate(&table)
	}
}
