package repo

import (
	"fmt"

	"gorm.io/gorm"
)

type SQLDatabase interface {
	ReturnDB() SQL
}

type SQL struct {
	Data *gorm.DB
}

// type DataPayload interface {
// 	ReadTable() DataPayload
// }

// Fix Later
func DBFactory(dbType string, name string) SQL {
	switch dbType {
	case "sqlite":
		sql := SQL{}
		sql.LoadDB(name)
		return sql

	default:
		fmt.Println("Returned empty DB struct.")
		return SQL{}
	}
}
