package discovery

import (
	consul "github.com/hashicorp/consul/api"
)

// Exec creates a consul entry then queries it
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
