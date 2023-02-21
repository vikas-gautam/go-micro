package main

import (
	"fmt"
	"log"
	"net/http"
)

const webPort = 8080

type Config struct{}

func main() {
	app := Config{}

	log.Printf("Starting broker service on port %v\n", webPort)

	//define the http server
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%v", webPort),
		Handler: app.routes(),
	}

	//start the http server
	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
