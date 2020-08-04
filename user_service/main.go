package main

import (
	"fmt"
	"net/http"

	"github.com/appendino/trab_ms/user_service/config"
	"github.com/appendino/trab_ms/user_service/database"
	"github.com/appendino/trab_ms/user_service/user"
	_ "github.com/go-sql-driver/mysql"
)

const apiBasePath = "/user"

func main() {
	database.SetupDatabase()
	user.SetupRoutes(apiBasePath)
	fmt.Println(config.CONSUL_SERVER)
	http.ListenAndServe(":8089", nil)
}
