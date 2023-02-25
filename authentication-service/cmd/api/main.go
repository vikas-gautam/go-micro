package main

import (
	"database/sql"

	"authentication/data"
)

const webPort = "8080"

type Config struct {
	DB     *sql.DB
	Models data.Models
}

func main() {

}
