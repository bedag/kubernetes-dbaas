package dpool

import (
	"fmt"
	"github.com/ory/dockertest/v3"
)

var dockerPool *dockertest.Pool

func GetDockerPool() (*dockertest.Pool, error) {
	var err error
	if dockerPool == nil {
		// Spin up sqlserver container
		dockerPool, err = dockertest.NewPool("")
		if err != nil {
			return nil, fmt.Errorf("could not connect to docker: %s", err)
		}
	}
	return dockerPool, nil
}