package dockertest

import (
	"context"
	"fmt"
	"time"

	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"github.com/redis/go-redis/v9"
)

func (r *Resource) RedisDSN() string {
	return r.redisDSN
}

type ResourceRedis struct {
}

func (r *Resource) ConnectRedis(conf *ResourceRedis) error {
	opt := &dockertest.RunOptions{
		Repository:   "redis",
		Tag:          "7.2.3",
		Name:         "redistest",
		Hostname:     "redistest",
		ExposedPorts: []string{"6379"},
		PortBindings: map[docker.Port][]docker.PortBinding{"6379/tcp": {{HostPort: "6379"}}},
	}

	if redisExist, ok := r.pool.ContainerByName(opt.Name); ok {
		r.redis = redisExist
	} else {
		redisResource, err := r.pool.RunWithOptions(opt, func(config *docker.HostConfig) {
			config.AutoRemove = true
			config.RestartPolicy = docker.RestartPolicy{Name: "no"}
		})
		if err != nil {
			r.Close()
			return fmt.Errorf("run docker redis container: %s", err)
		}
		r.redis = redisResource
	}

	if err := r.redis.Expire(uint(600)); err != nil {
		return fmt.Errorf("expire docker redis container: %s", err)
	}

	dsn := fmt.Sprintf("localhost:%s", r.redis.GetPort("6379/tcp"))
	r.redisDSN = dsn

	return r.pool.Retry(func() error {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		cli := redis.NewUniversalClient(&redis.UniversalOptions{Addrs: []string{dsn}})

		if err := cli.Ping(ctx).Err(); err != nil {
			return fmt.Errorf("ping redis: %s", err)
		}

		return nil
	})
}
