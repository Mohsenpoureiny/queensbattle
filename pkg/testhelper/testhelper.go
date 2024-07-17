package testhelper

import (
	"os"

	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"github.com/sirupsen/logrus"
)

func IsIntegration() bool {
	return os.Getenv("TEST_INTEGRATION") == "true"
}
func StartDockerPool() *dockertest.Pool {
	pool, err := dockertest.NewPool("")
	if err != nil {
		logrus.WithError(err).Fatal("Could not construct pool", err)
	}
	// uses pool to try to connect to Docker
	err = pool.Client.Ping()
	if err != nil {
		logrus.WithError(err).Fatal("Could not connect to Docker", err)
	}
	return pool
}
func StartDockerInstance(pool *dockertest.Pool, image, tag string, env ...string) *dockertest.Resource {
	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: image,
		Tag:        tag,
		Env:        env,
	}, func(config *docker.HostConfig) {
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{
			Name: "no",
		}
	})
	if err := resource.Expire(120); err != nil {
		logrus.WithError(err).Fatal("couldn't set the resource expration")
	}
	if err != nil {
		logrus.WithError(err).Fatal("Could not start resource", err)
	}
	return resource
}
