package main

import (
	"github.com/RyhoBtw/3D-printer-api/api"
	"github.com/RyhoBtw/3D-printer-api/api/database"
)

func main() {
	database.ConnectToDatabase()
	api.HandleRequests()
}
