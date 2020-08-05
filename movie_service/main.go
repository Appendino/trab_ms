package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/appendino/trab_ms/movie_service/config"
	"github.com/appendino/trab_ms/movie_service/database"
	"github.com/appendino/trab_ms/movie_service/discovery"
	"github.com/appendino/trab_ms/movie_service/movie"
	_ "github.com/go-sql-driver/mysql"
)

const apiBasePath = "/movie"

func main() {
	database.SetupDatabase()
	route := movie.SetupRoutes(apiBasePath)
	discovery.RegisterService(config.CONSUL_SERVER, config.CONSUL_SERVICE_NAME,
		config.HOSTNAME, config.PORT, config.CONSUL_SEARCH_STRINGS)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(config.PORT), route))
}
