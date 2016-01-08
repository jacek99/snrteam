package main

//go:generate gengen -o . github.com/jacek99/snrteam/generics string

import (
	"github.com/jacek99/snrteam/api"
	"github.com/jacek99/snrteam/database"
	"github.com/jacek99/snrteam/server"
	"log"
)

func main() {

	// connect to BoltDB and shut it down cleanly at the end
	defer database.Database.Close()

	// run web server
	api.InitRouter(server.Router)
	err := server.Server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
