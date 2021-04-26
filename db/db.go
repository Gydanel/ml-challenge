package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"ml-challenge/config"
)

var db *sql.DB

func Init() {
	c := config.GetConfig()
	db, err := sql.Open(c.GetString("db.type"), c.GetString("db.name"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}

func Migrate() {
	sqlStmt := `
	CREATE TABLE satellites (
	    id integer not null primary key, 
	    name text, 
	    x double,
	    y double
	);
	delete from satellites;
	CREATE TABLE messages (
	    id integer not null primary key, 
	    satellite_id  FOREIGN KEY(satellite_id) REFERENCES satellites(id), 
	    message long text
	);
	delete from messages;
	`
	db = GetDB()
	_, err := db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}
}

func GetDB() *sql.DB {
	return db
}
