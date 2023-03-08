package main

import (
	"github.com/RyhoBtw/3D-printer-api/api"
	"github.com/RyhoBtw/3D-printer-api/api/db"
)

func main() {
	db.ConnectToDatabase()
	api.HandleRequests()
}
