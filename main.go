package main

import (
	"real-time-backend/backend/database"
	"real-time-backend/backend/server"
)

func main() {
	database.InitMainDB()
	server.LoadServer()
}
