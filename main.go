package main

import (
	"flyte-league/database"
	"log"
)

func main() {
	if err := database.OpenDatabase("./flyte.sqlite3"); err != nil {
		log.Fatal(err)
	}
	defer database.CloseDatabase()
}
