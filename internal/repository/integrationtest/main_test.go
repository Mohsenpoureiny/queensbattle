package integrationtest

import (
	"fmt"
	"github.com/ory/dockertest/v3"
	"os"
	"queensbattle/internal/repository/redis"
	"queensbattle/pkg/testhelper"
	"testing"
)

var redisPort string

func TestMain(m *testing.M) {
	if !testhelper.IsIntegration() {
		return
	}
	pool := testhelper.StartDockerPool()
	redisRes := testhelper.StartDockerInstance(pool, "redis/redis-stack-server", "latest", func(resource *dockertest.Resource) error {
		port := resource.GetPort("6379/tcp")
		_, err := redis.NewRedisClient(fmt.Sprintf("localhost:%s", port))
		return err
	})

	redisPort = redisRes.GetPort("6379/tcp")

	exitCode := m.Run()

	redisRes.Close()
	os.Exit(exitCode)
}
