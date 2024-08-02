package main

import (
	"log"
	"static-api/cmd/api"
	"static-api/db"
)

func main() {
	db.Init()

	server := api.NewAPIServer(":9000", nil)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
