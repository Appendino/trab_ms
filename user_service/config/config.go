package config

import "os"

var (
	CONSUL_SERVER         = "consul:8500"
	MYSQL_SERVER          = "dbtrabalho:3306"
	MYSQL_USER            = "userws"
	MYSQL_PASSWORD        = "123456"
	MYSQL_DATABASE        = "userdb"
	MOVIE_CONSUL          = "movie"
	MOVIE_ENDPOINT        = "/movie"
	MOVIE_PROTOCOL        = "http"
	CONSUL_SERVICE_NAME   = "fiap.micro.netflix"
	PORT                  = 8089
	CONSUL_SEARCH_STRINGS = []string{"user", "users", "user_service"}
	HOSTNAME, _           = os.Hostname()
)
