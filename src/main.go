package main

import (
	"awesomeProject/src/database"
	"awesomeProject/src/server"
)

func main() {
	database.StartDB()
	server := server.NewServer()

	server.Run()
}
