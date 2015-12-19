package consul

import (
	"github.com/satori/go.uuid"
	consul"github.com/hashicorp/consul/api"

)

type ConsulRegistrator struct {
	Config *consul.Config
}

// Register announce the presence of a new service
// to Consul discovery server.
func (cr *ConsulRegistrator) Register(service string, port int) error {
	if cr.Config == nil {
		cr.Config = consul.DefaultConfig()
	}
	client, err := consul.NewClient(cr.Config)
	if err != nil {
		return err
	}

	err = client.Agent().ServiceRegister(&consul.AgentServiceRegistration{
			ID: uuid.NewV4().String(),
			Name: service,
			Port: port,
	})

	if err != nil {
		return err
	}
	return nil
}

