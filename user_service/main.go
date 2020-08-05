package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/appendino/trab_ms/user_service/config"
	"github.com/appendino/trab_ms/user_service/database"
	"github.com/appendino/trab_ms/user_service/discovery"
	"github.com/appendino/trab_ms/user_service/user"
	_ "github.com/go-sql-driver/mysql"
)

const apiBasePath = "/user"

func main() {
	database.SetupDatabase()
	discovery.RegisterService(config.CONSUL_SERVER, config.CONSUL_SERVICE_NAME,
		config.HOSTNAME, config.PORT, config.CONSUL_SEARCH_STRINGS)
	user.SetupRoutes(apiBasePath)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(config.PORT), nil))
}
