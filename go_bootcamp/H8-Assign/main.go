package main

import (
	"go_bootcamp/H8-Assign/database"
	"go_bootcamp/H8-Assign/routers"
)

func main() {
	var PORT = ":8080"
	database.StartDB()
	routers.StartServer().Run(PORT)

}
