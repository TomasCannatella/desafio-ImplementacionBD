package main

import (
	"app/internal/application"
	"fmt"
	"os"

	"github.com/go-sql-driver/mysql"
)

func main() {
	// env
	// ...

	// app
	// - config
	cfg := &application.ConfigApplicationDefault{
		Db: &mysql.Config{
			User:   os.Getenv("DB_USER"),
			Passwd: os.Getenv("DB_PASSWORD"),
			Net:    "tcp",
			Addr:   os.Getenv("DB_HOST"),
			DBName: os.Getenv("DB_NAME"),
		},
		Addr: "127.0.0.1:8080",
	}
	app := application.NewApplicationDefault(cfg)
	// - set up
	err := app.SetUp()
	if err != nil {
		fmt.Println(err)
		return
	}
	// - run
	err = app.Run()
	if err != nil {
		fmt.Println(err)
		return
	}
}
