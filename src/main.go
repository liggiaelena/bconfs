package main

import "awesomeProject/src/server"

func main() {
	server := server.NewServer()

	server.Run()
}
