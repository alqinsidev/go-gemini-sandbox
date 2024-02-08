package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"

	"github.com/spf13/viper"
)

func InitMySQL(viper *viper.Viper) *sql.DB {
	host := viper.GetString("MYSQL_HOST")
	port := viper.GetString("MYSQL_PORT")
	username := viper.GetString("MYSQL_USER")
	password := viper.GetString("MYSQL_PASSWORD")
	dbName := viper.GetString("MYSQL_DB_NAME")

	dsn := fmt.Sprintf(`%s:%s@tcp(%s:%s)/%s`, username, password, host, port, dbName)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("DB Conn error: %s", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("DB Conn error: %s", err)
	}

	return db
}
