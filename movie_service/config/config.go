package config

import "os"

var (
	CONSUL_SERVER         = "consul:8500"
	MYSQL_SERVER          = "dbtrabalho:3306"
	MYSQL_USER            = "moviews"
	MYSQL_PASSWORD        = "123456"
	MYSQL_DATABASE        = "moviedb"
	CONSUL_SERVICE_NAME   = "fiap.micro.netflix.movie"
	PORT                  = 8088
	CONSUL_SEARCH_STRINGS = []string{"movie", "movies", "movie_service"}
	HOSTNAME, _           = os.Hostname()
)
