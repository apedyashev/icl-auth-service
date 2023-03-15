package main

import (
	"icl-auth/database"
)

func main() {
	database.Connect()
	database.Migrate()

	router := InitRouter()
	router.Run(":80")
}
