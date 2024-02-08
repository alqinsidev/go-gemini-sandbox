package main

import (
	"log"

	"github.com/alqinsidev/go-gemini-sandbox/src/config"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/spf13/viper"
)

func main() {

	_ = config.InitViper()

	db := config.InitMySQL()

	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		log.Fatal("Err: ", err)
	}

	path := viper.GetString("MIGRATION_PATH")
	m, err := migrate.NewWithDatabaseInstance(path, "mysql", driver)
	if err != nil {
		log.Fatal("Err: ", err)
	}

	err = m.Up()
	if err != nil {
		log.Fatal("Err: ", err)
	}

}
