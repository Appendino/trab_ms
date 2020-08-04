package discovery

import (
	"strconv"

	consul "github.com/hashicorp/consul/api"
)

func RegisterService(consulServer, serviceName, hostname string, port int, registerNames []string) error {
	config := consul.DefaultConfig()
	config.Address = consulServer

	cli, err := NewClient(config, hostname, serviceName, port)
	if err != nil {
		return err
	}

	if err := cli.Register(registerNames); err != nil {
		return err
	}

	return nil
}

func GetService(consulServer, serviceName, hostname string, port int, searchName string) (string, error) {
	config := consul.DefaultConfig()
	config.Address = consulServer

	cli, err := NewClient(config, hostname, serviceName, port)
	if err != nil {
		return "", err
	}

	entries, _, err := cli.Service(serviceName, searchName)
	if err != nil {
		return "", err
	}
	portStr := strconv.Itoa(entries[0].Service.Port)
	serverPort := entries[0].Service.Address + ":" + portStr
	return serverPort, nil
}
