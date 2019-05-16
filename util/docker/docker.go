package docker

import (
	"github.com/docker/docker/client"
)

type container struct {
	client *client.Client
}

var docker *container

func CreateClient() *container {
	return nil
}
