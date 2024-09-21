package config

import "gorm.io/driver/mysql"
import "gorm.io/gorm"
import "log"

func MySQLConfig() (*gorm.DB, error) {

	var connector string

	connector = "root:root@tcp(localhost:3306)/jarta"
	
	database, err := gorm.Open(mysql.Open(connector), &gorm.Config{})
	if err != nil {
		log.Println(err);
	}
	return database, err
}