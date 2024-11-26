package database

import (
	"gorm.io/gorm"
)

type Dbinstance struct {
	Db *gorm.DB
}

var DB Dbinstance

func ConnectDb() {
	gorm.Open(gorm_dialector, gorm_options)
}
