package main

import (
	"go_bootcamp/H8-swagger/database"

	"go_bootcamp/H8-swagger/routers"
)

func main() {
	var PORT = ":8080"
	database.StartDB()
	routers.StartServer().Run(PORT)

}
