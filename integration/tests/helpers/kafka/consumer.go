//go:generate mockgen -destination=./mock/consumer_mock.go -package=mock  -source consumer.go
package kafka

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/IBM/sarama"
)

var ErrNoSubscribedTopics = errors.New("consumer is not subscribed to any topic")

type ConsumerGroup interface {
	Consume(ctx context.Context, topics []string, handler sarama.ConsumerGroupHandler) error
	Close() error
}

type Consumer struct {
	topics []string
	group  ConsumerGroup
}

func NewConsumer(
	cfg ConsumerConfig,
) (*Consumer, error) {
	config := sarama.NewConfig()
	version, err := sarama.ParseKafkaVersion(cfg.Version)
	if err != nil {
		return nil, fmt.Errorf("invalid kafka version: %w", err)
	}
	config.Version = version

	config.ClientID = cfg.ClientID
	config.Consumer.Fetch.Min = cfg.FetchMinBytes
	config.Consumer.Fetch.Max = cfg.FetchMaxBytes
	config.Consumer.MaxWaitTime = cfg.FetchMaxWaitMs
	config.Consumer.Group.Session.Timeout = cfg.SessionTimeoutMs
	config.Consumer.Group.Heartbeat.Interval = cfg.HeartbeatIntervalMs
	config.Consumer.Offsets.Initial = cfg.AutoOffsetReset.toSaramaOffsets()
	config.Consumer.Offsets.AutoCommit.Enable = true
	config.Consumer.Offsets.AutoCommit.Interval = cfg.AutoCommitIntervalMs
	config.Consumer.Return.Errors = true

	config.Net.SASL.Version = 1
	config.Net.SASL.Mechanism = cfg.SaslConfig.SaslMechanism
	config.Net.SASL.User = cfg.SaslConfig.SaslUsername
	config.Net.SASL.Password = cfg.SaslConfig.SaslPassword
	config.Net.SASL.Enable = cfg.SaslConfig.SaslEnable
	config.Net.SASL.Handshake = cfg.SaslConfig.SaslHandshake
	config.Net.SASL.SCRAMClientGeneratorFunc = cfg.SaslConfig.ScramClient

	group, err := sarama.NewConsumerGroup(strings.Split(cfg.BootstrapServers, ","), cfg.GroupID, config)
	if err != nil {
		return nil, err
	}

	go func() {
		for err := range group.Errors() {
			log.Default().Printf("kafka group %q error: %v", cfg.GroupID, err)
		}
	}()

	log.Default().Print("consumer up and running")

	return &Consumer{
		group: group,
	}, nil
}

func (c *Consumer) Consume(ctx context.Context) (<-chan *ConsumerRecord, func(), error) {
	if len(c.topics) == 0 {
		return nil, nil, ErrNoSubscribedTopics
	}

	incomingMessagesCh := make(chan *ConsumerRecord)

	go func() {
		defer close(incomingMessagesCh)

		for {
			handler := consumerGroupHandler{
				incomingEvents: incomingMessagesCh,
			}

			if err := c.group.Consume(ctx, c.topics, handler); err != nil {
				if errors.Is(err, sarama.ErrClosedConsumerGroup) {
					return
				}
				log.Default().Printf("consumer failed: %s", err.Error())
			}
			if ctx.Err() != nil {
				return
			}
		}
	}()

	ctxStop, cancel := context.WithCancel(ctx)
	go func() {
		<-ctxStop.Done()
		c.group.Close()
	}()

	return incomingMessagesCh, cancel, nil
}

func (c *Consumer) Subscribe(topic string) {
	if topic != "" {
		c.topics = append(c.topics, topic)
	}
}

type consumerGroupHandler struct {
	incomingEvents chan *ConsumerRecord
	logger         *log.Logger
}

func (c consumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error {
	return nil
}

func (c consumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error {
	return nil
}

func (c consumerGroupHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	defer log.Default().Printf("exit from ConsumeClaim")

	for {
		select {
		case msg, ok := <-claim.Messages():
			if !ok {
				return nil
			}
			c.incomingEvents <- toConsumerMessage(msg, sess)
		case <-sess.Context().Done():
			return nil
		}
	}
}

func toConsumerMessage(msg *sarama.ConsumerMessage, sess sarama.ConsumerGroupSession) *ConsumerRecord {
	var headers []Header
	for i := range msg.Headers {
		h := msg.Headers[i]

		if h == nil {
			continue
		}

		converted := Header(*h)
		headers = append(headers, converted)
	}

	return &ConsumerRecord{
		Key:          msg.Key,
		Value:        msg.Value,
		Topic:        msg.Topic,
		Partition:    msg.Partition,
		Offset:       msg.Offset,
		Headers:      headers,
		groupSession: sess,
	}
}

func (cr *ConsumerRecord) Ack() {
	cr.groupSession.MarkOffset(cr.Topic, cr.Partition, cr.Offset+1, "")
}
