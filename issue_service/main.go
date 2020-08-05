package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/appendino/trab_ms/issue_service/config"
	"github.com/appendino/trab_ms/issue_service/database"
	"github.com/appendino/trab_ms/issue_service/discovery"
	"github.com/appendino/trab_ms/issue_service/issue"
	_ "github.com/go-sql-driver/mysql"
)

const apiBasePath = "/issue"

func main() {
	database.SetupDatabase()
	discovery.RegisterService(config.CONSUL_SERVER, config.CONSUL_SERVICE_NAME,
		config.HOSTNAME, config.PORT, config.CONSUL_SEARCH_STRINGS)
	issue.SetupRoutes(apiBasePath)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(config.PORT), nil))
}
