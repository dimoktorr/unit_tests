package dockertest

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"time"
)

func (r *Resource) PostgresDSN() string {
	return r.postgresDSN
}

type ResourcePostgres struct {
	User     string
	Password string
	DB       string
	Schema   string
}

func (r *Resource) ConnectPostgres(conf *ResourcePostgres) error {
	opt := &dockertest.RunOptions{
		Repository:   "postgres",
		Tag:          "15.2",
		Name:         "postgrestest",
		Hostname:     "postgrestest",
		ExposedPorts: []string{"5432"},
		PortBindings: map[docker.Port][]docker.PortBinding{"5432/tcp": {{HostPort: "5432"}}},
		Env: []string{
			"POSTGRES_USER=" + conf.User,
			"POSTGRES_PASSWORD=" + conf.Password,
			"POSTGRES_DB=" + conf.DB,
			"listen_addresses = '*'",
		},
	}

	if postgresExist, ok := r.pool.ContainerByName(opt.Name); ok {
		r.postgres = postgresExist
	} else {
		postgresResource, err := r.pool.RunWithOptions(opt, func(config *docker.HostConfig) {
			config.AutoRemove = true
			config.RestartPolicy = docker.RestartPolicy{Name: "no"}
		})
		if err != nil {
			r.Close()
			return fmt.Errorf("run docker postgres container: %s", err)
		}
		r.postgres = postgresResource
	}

	if err := r.postgres.Expire(uint(600)); err != nil {
		return fmt.Errorf("expire docker postgres container: %s", err)
	}

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s TimeZone=UTC sslmode=disable",
		"localhost", r.postgres.GetPort("5432/tcp"),
		conf.User, conf.Password, conf.DB,
	)
	r.postgresDSN = dsn

	return r.pool.Retry(func() error {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		cli, err := sql.Open("postgres", dsn)
		if err != nil {
			return fmt.Errorf("open postgres connection: %w", err)
		}

		if err := cli.PingContext(ctx); err != nil {
			return fmt.Errorf("ping postgres connection: %w", err)
		}

		if _, err := cli.ExecContext(ctx, `CREATE SCHEMA IF NOT EXISTS `+conf.Schema); err != nil {
			return fmt.Errorf("create postgres schema: %w", err)
		}

		return nil
	})
}
