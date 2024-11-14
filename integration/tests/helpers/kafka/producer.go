//go:generate mockgen -destination=./mock/producer_mock.go -package=mock  -source producer.go
package kafka

import (
	"context"
	"fmt"
	"strings"
	"sync"

	"github.com/IBM/sarama"
)

type callbackKey struct{}

type Callback func(*RecordMetadata, error)

type AsyncProducer interface {
	Successes() <-chan *sarama.ProducerMessage
	Errors() <-chan *sarama.ProducerError
	Input() chan<- *sarama.ProducerMessage
	AsyncClose()
}

type Producer struct {
	saramaProducer AsyncProducer
	wg             *sync.WaitGroup
}

func NewProducer(
	cfg ProducerConfig,
) (*Producer, error) {
	config := sarama.NewConfig()
	version, err := sarama.ParseKafkaVersion(cfg.Version)
	if err != nil {
		return nil, fmt.Errorf("invalid kafka version: %w", err)
	}
	config.Version = version

	config.ClientID = cfg.ClientID
	config.Producer.Timeout = cfg.RequestTimeoutMs
	config.Producer.Retry.Backoff = cfg.RetryBackoffMs
	config.Net.MaxOpenRequests = cfg.MaxInFlightRequestsPerConnection
	config.Producer.Idempotent = cfg.EnableIdempotence
	config.Producer.RequiredAcks = cfg.Acks.toSaramaAcks()
	config.Producer.Retry.Max = cfg.Retries
	config.Producer.Flush.Frequency = cfg.LingerMs
	config.Producer.Flush.Bytes = cfg.MaxRequestSize
	config.Producer.Flush.Messages = cfg.MaxMessages
	config.Producer.Return.Successes = true

	config.Net.SASL.Version = 1
	config.Net.SASL.Enable = cfg.SaslConfig.SaslEnable
	config.Net.SASL.Mechanism = cfg.SaslConfig.SaslMechanism
	config.Net.SASL.User = cfg.SaslConfig.SaslUsername
	config.Net.SASL.Password = cfg.SaslConfig.SaslPassword
	config.Net.SASL.Handshake = cfg.SaslConfig.SaslHandshake
	config.Net.SASL.SCRAMClientGeneratorFunc = cfg.SaslConfig.ScramClient

	config.Net.SASL.TokenProvider = cfg.SaslConfig.TokenProvider
	if cfg.TlsConfig.Enable && cfg.TlsConfig.Config != nil {
		config.Net.TLS.Enable = cfg.TlsConfig.Enable
		config.Net.TLS.Config = cfg.TlsConfig.Config
	}

	saramaProducer, err := sarama.NewAsyncProducer(strings.Split(cfg.BootstrapServers, ","), config)
	if err != nil {
		return nil, err
	}

	producer := &Producer{
		saramaProducer: saramaProducer,
		wg:             new(sync.WaitGroup),
	}

	producer.wg.Add(2)
	go producer.handleSuccesses()
	go producer.handleErrors()

	return producer, nil
}

func (p *Producer) handleSuccesses() {
	defer p.wg.Done()
	for message := range p.saramaProducer.Successes() {
		metadata, ok := message.Metadata.(map[any]func(sarama.ProducerError))
		if !ok {
			panic(fmt.Sprintf("casting metadata was not success, metadata: %#v", message.Metadata))
		}
		callback := metadata[callbackKey{}]
		callback(sarama.ProducerError{Msg: message, Err: nil})
	}
}

func (p *Producer) handleErrors() {
	defer p.wg.Done()
	for err := range p.saramaProducer.Errors() {
		metadata, ok := err.Msg.Metadata.(map[any]func(sarama.ProducerError))
		if !ok {
			panic(fmt.Sprintf("casting metadata was not success, metadata: %#v", err.Msg.Metadata))
		}
		callback := metadata[callbackKey{}]
		callback(sarama.ProducerError{Msg: err.Msg, Err: err.Err})
	}
}

func (p *Producer) Send(ctx context.Context, msg *ProducerRecord, callbacks ...Callback) <-chan *RecordMetadataWithError {
	ch := make(chan *RecordMetadataWithError, 1)

	saramaMessage := p.toSaramaMessage(msg)
	saramaMessage.Metadata = map[any]func(producerErr sarama.ProducerError){
		callbackKey{}: func(producerErr sarama.ProducerError) {
			recordMetadata := RecordMetadata{
				Partition: producerErr.Msg.Partition,
				Offset:    producerErr.Msg.Offset,
				Topic:     producerErr.Msg.Topic,
				Key:       msg.Key,
				Value:     msg.Value,
			}
			for _, cb := range callbacks {
				cb(&recordMetadata, producerErr.Err)
			}

			ch <- &RecordMetadataWithError{
				RecordMetadata: &recordMetadata,
				Error:          producerErr.Err,
			}
			close(ch)
		},
	}

	select {
	case <-ctx.Done():
		ch <- &RecordMetadataWithError{
			RecordMetadata: nil,
			Error:          ctx.Err(),
		}
		close(ch)
	case p.saramaProducer.Input() <- saramaMessage:
	}

	return ch
}

func (p *Producer) toSaramaMessage(msg *ProducerRecord) *sarama.ProducerMessage {
	var headers []sarama.RecordHeader
	for i := range msg.Headers {
		h := msg.Headers[i]

		headers = append(headers, sarama.RecordHeader(h))
	}

	result := &sarama.ProducerMessage{
		Headers: headers,
		Topic:   msg.Topic,
		Value:   sarama.StringEncoder(msg.Value),
	}

	// otherwise all messages with empty key will send to one partition
	if len(msg.Key) > 0 {
		result.Key = sarama.StringEncoder(msg.Key)
	}

	return result
}

func (p *Producer) Stop() error {
	p.saramaProducer.AsyncClose()
	p.wg.Wait()
	return nil
}
