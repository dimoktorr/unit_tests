package dockertest

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/IBM/sarama"
	"github.com/dimoktorr/unit_tests/integration/tests/helpers/kafka"
	"github.com/go-zookeeper/zk"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
)

func (r *Resource) KafkaDSN() string {
	if len(r.kafkaDSN) == 0 {
		return "nosetkafkahost:9000"
	}
	return r.kafkaDSN
}

type ResourceZookeeper struct {
}

type ResourceKafka struct {
}

func (r *Resource) ConnectZookeeper(conf *ResourceZookeeper) error {
	networkExist := false
	networks, err := r.pool.Client.ListNetworks()
	if err != nil {
		return fmt.Errorf("could not get a network: %s", err)
	}
	for i, network := range networks {
		if network.Name == "zookeeper_kafka_network" {
			r.network = &networks[i]
			networkExist = true
			break
		}
	}
	if !networkExist {
		network, err := r.pool.Client.CreateNetwork(docker.CreateNetworkOptions{Name: "zookeeper_kafka_network"})
		if err != nil {
			return fmt.Errorf("could not create a network to zookeeper and kafka: %s", err)
		}
		r.network = network
	}

	opt := &dockertest.RunOptions{
		Repository:   "wurstmeister/zookeeper",
		Name:         "zookeepertest",
		Hostname:     "zookeepertest",
		NetworkID:    r.network.ID,
		ExposedPorts: []string{"2181"},
		PortBindings: map[docker.Port][]docker.PortBinding{"2181/tcp": {{HostPort: "2181"}}},
	}

	if zookeeperExist, ok := r.pool.ContainerByName(opt.Name); ok {
		_ = r.pool.Purge(zookeeperExist)
	}

	zookeeperResource, err := r.pool.RunWithOptions(opt)
	if err != nil {
		return fmt.Errorf("could not start zookeeper: %s", err)
	}
	r.zookeeper = zookeeperResource

	if err := r.zookeeper.Expire(uint(600)); err != nil {
		return fmt.Errorf("expire docker zookeeper container: %s", err)
	}

	dsn := fmt.Sprintf("localhost:%s", r.zookeeper.GetPort("2181/tcp"))
	r.zookeeperDSN = dsn

	conn, _, err := zk.Connect([]string{dsn}, 10*time.Second)
	if err != nil {
		return fmt.Errorf("could not connect zookeeper: %s", err)
	}
	defer conn.Close()

	return r.pool.Retry(func() error {
		switch conn.State() {
		case zk.StateHasSession, zk.StateConnected:
			return nil
		default:
			return errors.New("not yet connected")
		}
	})
}

func (r *Resource) ConnectKafka(conf *ResourceKafka) error {
	opt := &dockertest.RunOptions{
		Repository: "wurstmeister/kafka",
		Name:       "kafkatest",
		Hostname:   "kafkatest",
		NetworkID:  r.network.ID,
		Env: []string{
			"KAFKA_CREATE_TOPICS=domain.test:1:1:compact",
			"KAFKA_ADVERTISED_LISTENERS=INSIDE://kafkatest:9092,OUTSIDE://localhost:9093",
			"KAFKA_LISTENER_SECURITY_PROTOCOL_MAP=INSIDE:PLAINTEXT,OUTSIDE:PLAINTEXT",
			"KAFKA_LISTENERS=INSIDE://0.0.0.0:9092,OUTSIDE://0.0.0.0:9093",
			"KAFKA_ZOOKEEPER_CONNECT=zookeepertest:2181",
			"KAFKA_INTER_BROKER_LISTENER_NAME=INSIDE",
			"KAFKA_OFFSETS_TOPIC_REPLICATION_FACTOR=1",
		},
		ExposedPorts: []string{"9093/tcp"},
		PortBindings: map[docker.Port][]docker.PortBinding{
			"9093/tcp": {{HostIP: "localhost", HostPort: "9093/tcp"}},
		},
	}

	if kafkaExist, ok := r.pool.ContainerByName(opt.Name); ok {
		_ = r.pool.Purge(kafkaExist)
	}

	kafkaRes, err := r.pool.RunWithOptions(opt)
	if err != nil {
		return fmt.Errorf("could not start kafka: %s", err)
	}
	r.kafka = kafkaRes

	if err := r.kafka.Expire(uint(600)); err != nil {
		return fmt.Errorf("expire docker kafka container: %s", err)
	}

	dsn := fmt.Sprintf("localhost:%s", r.kafka.GetPort("9093/tcp"))
	r.kafkaDSN = dsn

	return r.pool.Retry(func() error {
		if err != nil {
			return fmt.Errorf("failed log init: %s", err)
		}

		producer, err := kafka.NewProducer(kafka.ProducerConfig{
			BootstrapServers:                 dsn,
			Version:                          sarama.V3_3_2_0.String(),
			ClientID:                         "payment",
			Acks:                             "WAIT_FOR_ALL",
			MaxInFlightRequestsPerConnection: 5,
			RequestTimeoutMs:                 10 * time.Second,
			RetryBackoffMs:                   100 * time.Millisecond,
			Retries:                          1,
			LingerMs:                         1 * time.Second,
			MaxRequestSize:                   1048576,
			SaslConfig: kafka.SaslConfig{
				SaslEnable: false,
			},
		})
		if err != nil {
			return err
		}
		defer func() {
			_ = producer.Stop()
			time.Sleep(1 * time.Second)
		}()

		send := <-producer.Send(context.Background(), &kafka.ProducerRecord{
			Topic: "domain.test",
			Key:   []byte("any-key"),
			Value: []byte("Hello World"),
		})
		if send.Error != nil {
			return send.Error
		}

		return nil
	})
}
