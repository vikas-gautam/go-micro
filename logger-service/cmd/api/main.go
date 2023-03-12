package main

import (
	"context"
	"fmt"
	"log"
	"log-service/data"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	webPort  = "9092"
	rpcPort  = "5001"
	mongoURL = "mongodb://mongo:27017/logs"
	gRpcPort = "50001"
)

var client *mongo.Client

type Config struct {
	Models data.Models
}

func main() {
	// connect to mongo
	mongoClient, err := connectToMongo()
	log.Println("Printing error: ", err)
	log.Println("Printing client mongo: ", mongoClient)
	if err != nil {
		log.Println("error returned by connectToMongo func", err)
		log.Panic(err)
	}
	client = mongoClient

	// create a context in order to disconnect
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	// close connection
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			log.Println("error returned while closing connection", err)
			panic(err)
		}
	}()

	app := Config{
		Models: data.New(client),
	}

	// start web server
	log.Println("Starting logger service on port: ", webPort)
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Println("error starting logging server", err)
		log.Panic(err)
	}

}

func connectToMongo() (*mongo.Client, error) {
	// create connection options
	clientOptions := options.Client().ApplyURI(mongoURL)

	clientOptions.SetAuth(options.Credential{
		Username: "admin",
		Password: "password",
	})

	// connect
	c, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Println("Error connecting:", err)
		return nil, err
	}

	// mongoConnect does not validate URI so we need to validate this using ping
	err = c.Ping(context.TODO(), nil)
	if err != nil {
		log.Println("Error while doing ping to db:", err)
		return nil, err
	}

	return c, nil
}
