package main

import (
	"database/sql"
	"fmt"
	"os"
)


func connectDB() (*sql.DB, error) {
  var db_name string = os.Getenv("DB")
  var db_host string = os.Getenv("DB_HOST")
  var db_port string = os.Getenv("DB_PORT")
  var db_username string = os.Getenv("DB_USER_NAME")
  var db_password string = os.Getenv("DB_PASSWORD")

  db,err := sql.Open("mysql",fmt.Sprintf("%s:%s@(%s:%s)/%s?parseTime=true",db_username,db_password,db_host,db_port,db_name))
  
  return db, err;
}

