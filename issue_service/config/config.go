package config

import "os"

var (
	CONSUL_SERVER             = "consul:8500"
	MYSQL_SERVER              = "dbtrabalho:3306"
	MYSQL_USER                = "issuews"
	MYSQL_PASSWORD            = "123456"
	MYSQL_DATABASE            = "issuedb"
	MOVIE_CONSUL              = "movie"
	MOVIE_ENDPOINT            = "/movie"
	USER_CONSUL               = "user"
	USER_ENDPOINT             = "/user"
	MOVIE_PROTOCOL            = "http"
	CONSUL_SERVICE_NAME       = "fiap.micro.netflix.issue"
	USER_CONSUL_SERVICE_NAME  = "fiap.micro.netflix.user"
	MOVIE_CONSUL_SERVICE_NAME = "fiap.micro.netflix.movie"
	PORT                      = 8090
	CONSUL_SEARCH_STRINGS     = []string{"issue", "issues", "issue_service"}
	HOSTNAME, _               = os.Hostname()
)
