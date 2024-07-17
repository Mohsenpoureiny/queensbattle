package integrationtest

import (
	"fmt"
	"os"
	"queensbattle/pkg/testhelper"
	"testing"
)

func TestMain(m *testing.M) {
	if !testhelper.IsIntegration() {
		return
	}
	pool := testhelper.StartDockerPool()
	redisRes := testhelper.StartDockerInstance(pool, "redis/redis-stack-server", "latest")

	redisRes.GetPort("6379/tcp")

	fmt.Println(redisRes.GetPort("6379/tcp"))
	fmt.Println("redis is up and running")

	exitCode := m.Run()

	redisRes.Close()
	os.Exit(exitCode)
}
