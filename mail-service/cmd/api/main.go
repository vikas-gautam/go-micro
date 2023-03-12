package main

import (
	"fmt"
	"log"
	"net/http"
)

const webPort = 9093

type Config struct{}

func main() {
	app := Config{}

	log.Printf("Starting mailer service on port %v\n", webPort)

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
