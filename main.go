package main

//go:generate gengen -o . github.com/jacek99/snrteam/generics string

import (
	"github.com/jacek99/snrteam/api"
	"github.com/jacek99/snrteam/database"
	"github.com/jacek99/snrteam/server"
	"log"
	"github.com/nicksnyder/go-i18n/i18n"
)

func main() {

	// internationalization support
	i18n.MustLoadTranslationFile("en-us.yaml")

	// connect to BoltDB and shut it down cleanly at the end
	defer database.Database.Close()

	// run web server
	api.InitRouter(server.Router)
	err := server.Server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
