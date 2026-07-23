package main

import (
	"log"

	"flyte-league/database"
)

func main() {
	if err := database.OpenDatabase("./flyte.sqlite3?_foreign_keys=on"); err != nil {
		log.Fatal(err)
	}
	defer database.CloseDatabase()
}
