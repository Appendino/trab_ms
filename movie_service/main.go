package main

import (
	"log"
	"net/http"
	"os"

	"github.com/appendino/trab_ms/movie_service/database"
	"github.com/appendino/trab_ms/movie_service/discovery"
	"github.com/appendino/trab_ms/movie_service/movie"
	_ "github.com/go-sql-driver/mysql"
)

const apiBasePath = "/movie"

func main() {
	database.SetupDatabase()
	route := movie.SetupRoutes(apiBasePath)
	// Registro do servico no Consul
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}
	discovery.RegisterService("consul:8500", "fiap.micro.netflix", hostname, 8088, []string{"movie", "movies", "movieservice"})
	http.ListenAndServe(":8088", route)
}
