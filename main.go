package main

import (
	"assignment8/database"
	"assignment8/router"
)

func main() {
	database.StartDB()
	r := router.StartApp()
	r.Run(":8080")
}
