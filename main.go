package main

import (
	"log"
	"github.com/jacek99/snrteam/database"
	"github.com/jacek99/snrteam/server"
)


func main() {

	// connect to BoltDB and shut it down cleanly at the end
	defer database.Database.Close()

	// run web server
	err := server.Server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}