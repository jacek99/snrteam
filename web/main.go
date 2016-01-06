package main

import (
	"net/http"
	"time"
	"os"
	"fmt"
	"log"
	"path"
	"github.com/boltdb/bolt"
	"github.com/gin-gonic/gin"
)

// configure via environment variables to be Docker friendly
const ENV_HTTP_PORT = "HTTP_PORT"
const ENV_DB_FOLDER = "ENV_DB_FOLDER"

func init_db() *bolt.DB {
	db_folder := os.Getenv(ENV_DB_FOLDER)
	if db_folder == "" {
		db_folder = "."
	}
	db_path := path.Join(db_folder, "snrteam.db")
	db, err := bolt.Open(db_path, 0600, &bolt.Options{Timeout: 10 * time.Second})
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func init_gin() (*http.Server, *gin.Engine) {
	router := gin.Default()

	port := os.Getenv(ENV_HTTP_PORT)
	if port == "" {
		 port = "8080"
	}

	s := &http.Server{
		Addr:           fmt.Sprintf(":%s",port),
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	return s, router
}

func init_router(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World")
	})
}


func main() {

	// connect to BoltDB and shut it down cleanly at the end
	db := init_db()
	defer db.Close()

	// run web server
	server, router := init_gin()
	init_router(router)

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}