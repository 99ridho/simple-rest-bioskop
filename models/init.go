package models

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/99ridho/simple-rest-bioskop/config"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/gorp.v2"
)

var Dbm *gorp.DbMap

func init() {
	dbName := config.DbName
	dbHost := config.DbHost
	dbUsername := config.DbUsername
	dbPort := config.DbPort
	dbPassword := config.DbPassword
	dbUrl := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUsername, dbPassword, dbHost, dbPort, dbName)
	log.Println(dbUrl)

	db, err := sql.Open("mysql", dbUrl)
	if err != nil {
		panic(err)
	}

	Dbm = &gorp.DbMap{
		Db: db,
		Dialect: gorp.MySQLDialect{
			Engine:   "InnoDB",
			Encoding: "UTF8",
		},
	}

	Dbm.TraceOn("[gorm]", log.New(os.Stdout, "myapp:", log.Lmicroseconds))
	Dbm.AddTableWithName(MovieTheater{}, "movie_theaters").SetKeys(true, "ID")
	Dbm.AddTableWithName(Movie{}, "movies").SetKeys(true, "ID")
	Dbm.AddTableWithName(MovieSchedule{}, "movie_schedules").SetKeys(true, "ID")
	Dbm.TraceOff()
}

// Create tables
func CreateTables() error {
	if err := Dbm.DropTablesIfExists(); err != nil {
		return err
	}

	if err := Dbm.CreateTablesIfNotExists(); err != nil {
		return err
	}
	if err := Dbm.CreateIndex(); err != nil {
		return err
	}

	return nil
}
