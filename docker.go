package testinfra

import (
	"fmt"
	"time"

	"gopkg.in/ory-am/dockertest.v3"
)

var pool *dockertest.Pool

type checkFunc func() error

func init() {
	var err error
	pool, err = dockertest.NewPool("")
	if err != nil {
		panic(err)
	}
}

func StartContainer(image, tag string, env []string) (*dockertest.Resource, error) {
	resource, err := pool.Run(image, tag, env)
	if err != nil {
		return nil, err
	}
	return resource, nil
}

func waitContainer(f checkFunc) error {
	pool.MaxWait = 30 * time.Second
	return pool.Retry(f)
}

func GetIP(resource *dockertest.Resource, publishedPort string) string {
	return resource.GetBoundIP(publishedPort)
}

func GetPort(resource *dockertest.Resource, publishedPort string) string {
	return resource.GetPort(publishedPort)
}

func GetIPPort(resource *dockertest.Resource, publishedPort string) string {
	return fmt.Sprintf("%s:%s", GetIP(resource, publishedPort), GetPort(resource, publishedPort))
}
