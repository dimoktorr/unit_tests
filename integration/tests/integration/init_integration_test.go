//go:build integration

package integration

import (
	v1 "github.com/dimoktorr/unit_tests/integration/pkg/api/v1"
	"github.com/dimoktorr/unit_tests/integration/tests/dockertest"
	"github.com/dimoktorr/unit_tests/integration/tests/helpers"
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestIntegrationTransactionsServerSuite(t *testing.T) {
	suite.Run(t, new(IntegrationServerSuite))
}

type IntegrationServerSuite struct {
	suite.Suite
	cfg               *ConfigTest
	exampleGrpcClient v1.ExampleServiceClient
	repoHelper        *helpers.PostgresHelper
}

func (s *IntegrationServerSuite) SetupTest() {
	s.T().Logf("SetupTest")

	cfg, err := NewConfigTest(s.T(), &ConfigTestOption{
		DockerTestConnect: &dockertest.Connect{
			Redis:    true,
			Postgres: true,
			Kafka:    true,
		},
	})
	if err != nil {
		s.T().Fatalf("new config-1 test: %s", err)
	}
	s.cfg = cfg

	s.exampleGrpcClient = v1.NewExampleServiceClient(cfg.conn)
}

func (s *IntegrationServerSuite) TearDownTest() {
	s.T().Log("TearDownTest")
	s.cfg.Close()
}
