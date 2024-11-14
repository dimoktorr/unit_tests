package dockertest

import (
	"fmt"

	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
)

type Resource struct {
	pool         *dockertest.Pool
	redis        *dockertest.Resource
	redisDSN     string
	postgres     *dockertest.Resource
	postgresDSN  string
	network      *docker.Network
	zookeeper    *dockertest.Resource
	zookeeperDSN string
	kafka        *dockertest.Resource
	kafkaDSN     string
}

type Connect struct {
	Redis    bool
	Postgres bool
	Kafka    bool
}

func NewResource() (*Resource, error) {
	pool, err := dockertest.NewPool("")
	if err != nil {
		return nil, fmt.Errorf("could not construct pool: %w", err)
	}

	// uses pool to try to connect to Docker
	if err = pool.Client.Ping(); err != nil {
		return nil, fmt.Errorf("could not connect to Docker: %w", err)
	}

	return &Resource{
		pool: pool,
	}, nil
}

func (r *Resource) Close() error {
	if r.redis != nil {
		err := r.redis.Close()
		if err != nil {
			return fmt.Errorf("close redis failed: %w", err)
		}
	}
	if r.postgres != nil {
		err := r.postgres.Close()
		if err != nil {
			return fmt.Errorf("close postgres failed: %w", err)
		}
	}
	return nil
}

func (r *Resource) Connect(con *Connect) error {
	if con.Redis {
		if err := r.ConnectRedis(&ResourceRedis{}); err != nil {
			return fmt.Errorf("connect redis: %w", err)
		}
	}

	if con.Postgres {
		if err := r.ConnectPostgres(&ResourcePostgres{
			User:     "example",
			Password: "example",
			DB:       "example",
			Schema:   "example",
		}); err != nil {
			return fmt.Errorf("connect postgres: %w", err)
		}
	}

	return nil
}
