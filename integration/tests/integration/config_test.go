// nolint: goerr113 // its ok
package integration

import (
	"context"
	"errors"
	"fmt"
	"github.com/dimoktorr/unit_tests/integration/internal/pkg/api"
	v1 "github.com/dimoktorr/unit_tests/integration/pkg/api/v1"
	"github.com/dimoktorr/unit_tests/integration/tests/dockertest"
	"net"
	"os"
	"path/filepath"
	"strconv"
	"testing"
	"time"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"

	"google.golang.org/grpc/credentials/insecure"

	"github.com/prometheus/client_golang/prometheus"
	"google.golang.org/grpc"
	"gopkg.in/yaml.v3"
)

type ConfigTestOption struct {
	DockerTestConnect *dockertest.Connect
}

type ConfigTest struct {
	t           *testing.T
	isLocal     bool
	dockerTest  *dockertest.Resource
	ctx         context.Context
	ctxClose    context.CancelFunc
	serverHost  string
	grpcPort    string
	conn        *grpc.ClientConn
	redisClient redis.UniversalClient
	pgClient    *pgxpool.Pool
	scanApi     *pgxscan.API
}

func NewConfigTest(t *testing.T, opt *ConfigTestOption) (*ConfigTest, error) {
	t.Helper()
	s := ConfigTest{t: t}

	s.ctx, s.ctxClose = context.WithCancel(context.Background())

	s.serverHost = "127.0.0.1"
	s.grpcPort = "5390"

	t.Log("[init] local mode")
	s.isLocal = true
	if err := s.initLocal(opt); err != nil {
		return nil, fmt.Errorf("local init: %w", err)
	}

	conn, err := newGrpcClient(s.serverHost, s.grpcPort)
	if err != nil {
		return nil, fmt.Errorf("repository new grpc client: %w", err)
	}
	s.conn = conn

	return &s, nil
}

func (s *ConfigTest) initLocal(opt *ConfigTestOption) error {
	dockerTest, err := dockertest.NewResource()
	if err != nil {
		return fmt.Errorf("dockertest.NewResource: %w", err)
	}
	s.dockerTest = dockerTest

	if err = dockerTest.Connect(opt.DockerTestConnect); err != nil {
		return fmt.Errorf("dockertest connect: %w", err)
	}
	s.t.Log("[dockertest] connect")

	if err = setLocalEnv(dockerTest, opt); err != nil {
		return fmt.Errorf("setEnv: %w", err)
	}
	s.t.Log("[setEnv] done")

	if err := s.exampleStart(); err != nil {
		return fmt.Errorf("example app start failed: %w", err)
	}

	WaitUpGrpcServer(s.t, s.ctx, s.serverHost, s.grpcPort)

	return nil
}

func setLocalEnv(dockerTest *dockertest.Resource, opt *ConfigTestOption) error {
	path, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("os.Getwd: %w", err)
	}

	envs, err := containerEnvs(path)
	if err != nil {
		return fmt.Errorf("containerEnvs: %w", err)
	}

	for k, v := range envs {
		_ = os.Setenv(k, fmt.Sprintf("%v", v))
	}

	_ = os.Setenv("DATABASE_POSTGRES_DSN", dockerTest.PostgresDSN())
	_ = os.Setenv("DATABASE_POSTGRES_MIGRATION_PATH", "file://"+filepath.Join(path, "./../../migrations/files/postgres"))

	_ = os.Setenv("KAFKA_CONSUMER_ENABLE", strconv.FormatBool(opt.DockerTestConnect.Kafka))
	_ = os.Setenv("KAFKA_CONSUMER_BROKERS", dockerTest.KafkaDSN())
	_ = os.Setenv("KAFKA_CONSUMER_SASL_ENABLE", "false")

	_ = os.Setenv("KAFKA_PRODUCER_ENABLE", strconv.FormatBool(opt.DockerTestConnect.Kafka))
	_ = os.Setenv("KAFKA_PRODUCER_BROKERS", dockerTest.KafkaDSN())
	_ = os.Setenv("KAFKA_PRODUCER_SASL_ENABLE", "false")

	_ = os.Setenv("REDIS_HOSTS", dockerTest.RedisDSN())

	return nil
}

func (s *ConfigTest) exampleStart() error {
	server := grpc.NewServer()

	serverApi := &api.Server{}
	server.RegisterService(&v1.ExampleService_ServiceDesc, serverApi)

	l, err := NewTCPListener(s.serverHost, s.grpcPort)
	if err != nil {
		return err
	}

	go func() {
		s.t.Log("[server] start")
		server.Serve(l)
	}()

	go func() {
		<-s.ctx.Done()
		server.GracefulStop()
		s.t.Log("[server] stop")
	}()

	return nil
}

func NewTCPListener(host, port string) (net.Listener, error) {
	addr := host + ":" + port
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, fmt.Errorf("failed to listen on %q: %w", addr, err)
	}

	return l, nil
}

func (s *ConfigTest) Close() {
	if s == nil {
		return
	}
	if s.ctxClose != nil {
		s.ctxClose()
	}
	if s.dockerTest != nil {
		s.dockerTest.Close()
	}
	if s.isLocal {
		prometheus.DefaultRegisterer = prometheus.NewRegistry()
		WaitDownGrpcServer(s.t, context.Background(), s.serverHost, s.grpcPort)
	}
}

func containerEnvs(path string) (map[string]any, error) {
	valuesIncrData, err := os.ReadFile(filepath.Join(path, "./../../k8s/values-incr.yaml"))
	if err != nil {
		return nil, fmt.Errorf("os.ReadFile: %w", err)
	}

	valuesIncr := make(map[string]any)
	err = yaml.Unmarshal(valuesIncrData, &valuesIncr)
	if err != nil {
		return nil, fmt.Errorf("yaml.Unmarshal: %w", err)
	}

	containers, ok := valuesIncr["containers"].([]any)
	if !ok {
		return nil, errors.New("block containers not exist")
	}

	var containerExample map[string]any
	for _, container := range containers {
		if containerT, ok := container.(map[string]any); ok {
			if containerT["name"].(string) == "example" {
				containerExample = containerT
				break
			}
		}
	}

	if containerExample == nil {
		return nil, errors.New("block container example not exist")
	}

	envs, ok := containerExample["env"].(map[string]any)
	if !ok {
		return nil, errors.New("block container example env not exist")
	}

	return envs, nil
}

func WaitUpGrpcServer(t *testing.T, ctx context.Context, serverHost, grpcPort string) {
	t.Helper()

	timeout := 100 * time.Millisecond
	for {
		if ctx.Err() != nil {
			return
		}
		conn, err := net.DialTimeout("tcp", net.JoinHostPort(serverHost, grpcPort), timeout)
		if err != nil {
			fmt.Println("Connecting error:", err)
			continue
		}
		if conn != nil {
			conn.Close()
			break
		}
	}
}

func WaitDownGrpcServer(t *testing.T, ctx context.Context, serverHost, grpcPort string) {
	t.Helper()

	timeout := 100 * time.Millisecond
	for {
		if ctx.Err() != nil {
			return
		}
		_, err := net.DialTimeout("tcp", net.JoinHostPort(serverHost, grpcPort), timeout)
		if err == nil {
			continue
		}
		break
	}
}

func newGrpcClient(host, port string, opts ...grpc.DialOption) (*grpc.ClientConn, error) {
	opts = append(
		[]grpc.DialOption{
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		},
		opts...,
	)
	return grpc.NewClient(net.JoinHostPort(host, port), opts...)
}
