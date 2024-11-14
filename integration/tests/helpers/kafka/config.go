package kafka

import (
	"crypto/tls"
	"time"

	"github.com/IBM/sarama"
)

func NewDefaultProducerConfig(clientID, brokers string, saslConfig SaslConfig) ProducerConfig {
	return ProducerConfig{
		ClientID:         clientID,
		BootstrapServers: brokers,
		SaslConfig:       saslConfig,
	}
}

type Acks string

const (
	AcksNoResponse    = "NO_RESPONSE"
	AcksWaitForLeader = "WAIT_FOR_LEADER"
	AcksWaitForAll    = "WAIT_FOR_ALL"
)

func (a Acks) toSaramaAcks() sarama.RequiredAcks {
	switch a {
	case AcksNoResponse:
		return sarama.NoResponse
	case AcksWaitForLeader:
		return sarama.WaitForLocal
	case AcksWaitForAll:
		return sarama.WaitForAll
	default:
		return sarama.WaitForLocal
	}
}

type SaslConfig struct {
	SaslUsername  string               `env:"SASL_USERNAME"`
	SaslPassword  string               `env:"SASL_PASSWORD"`
	SaslEnable    bool                 `env:"SASL_ENABLE" envDefault:"true"`
	SaslHandshake bool                 `env:"SASL_HANDSHAKE" envDefault:"true"`
	SaslMechanism sarama.SASLMechanism `env:"SASL_MECHANISM" envDefault:"PLAIN"`
	ScramClient   func() sarama.SCRAMClient
	TokenProvider sarama.AccessTokenProvider
}

type TlsConfig struct {
	Enable bool
	Config *tls.Config
}

type ProducerConfig struct {
	Version string `env:"KAFKA_PRODUCER_VERSION" envDefault:"3.2.0"`
	// client.id - An id string to pass to the server when making requests.
	ClientID string `env:"KAFKA_PRODUCER_CLIENT_ID" validate:"required"`
	// bootstrap.servers - Full set of brokers with "host:port" format.
	BootstrapServers string `env:"KAFKA_PRODUCER_BROKERS" validate:"required"`
	// acks - One of parameters [AcksNoResponse, AcksWaitForLeader, AcksWaitForAll].
	// When producer send message to kafka we need to wait while brokers send response that message has been received:
	// AcksNoResponse (0) - don't wait any broker (message may get lost).
	// AcksWaitForLeader (1) - wait until leader send response that message has been received.
	// AcksWaitForAll (-1) -  wait until leader and all replicas send response that message has been received.
	Acks Acks `env:"KAFKA_PRODUCER_ACKS" envDefault:"NO_RESPONSE" validate:"oneof=NO_RESPONSE WAIT_FOR_LEADER WAIT_FOR_ALL"`
	// max.in.flight.requests.per.connection - The maximum number of unacknowledged requests the client will send
	// on a single connection before blocking.
	MaxInFlightRequestsPerConnection int `env:"KAFKA_PRODUCER_MAX_IN_FLIGHT_REQUESTS_PER_CONNECTION" envDefault:"5"`
	// enable.idempotence - When set to 'true', the producer will ensure that exactly one copy of each message is
	// written in the stream.
	EnableIdempotence bool `env:"KAFKA_PRODUCER_ENABLE_IDEMPOTENCE" envDefault:"false"`
	// request.timeout.ms - Controls the maximum amount of time the client will wait for the response.
	RequestTimeoutMs time.Duration `env:"KAFKA_PRODUCER_REQUEST_TIMEOUT_MS" envDefault:"10s"`
	// retry.backoff.ms - The amount of time to wait before attempting to retry a failed request to a given partition
	RetryBackoffMs time.Duration `env:"KAFKA_PRODUCER_RETRY_BACKOFF_MS" envDefault:"100ms"`
	// retries - The total number of times to retry sending a message
	Retries int `env:"KAFKA_PRODUCER_RETRIES" envDefault:"3"`
	// linger.ms - This setting accomplishes this by adding a small amount of artificial delay—that is, rather than
	// immediately sending out a record, the producer will wait for up to the given delay to allow other records
	// to be sent so that the sends can be batched together
	LingerMs time.Duration `env:"KAFKA_PRODUCER_LINGER_MS" envDefault:"1s"`
	// max.request.size - The maximum size of a request in bytes.
	MaxRequestSize int `env:"KAFKA_PRODUCER_MAX_REQUEST_SIZE" envDefault:"1048576"`
	// max messages in local queue to flush
	// 0 - infinity by default
	MaxMessages int        `env:"KAFKA_PRODUCER_MAX_MESSAGES" envDefault:"0"`
	SaslConfig  SaslConfig `envPrefix:"KAFKA_PRODUCER_"`
	TlsConfig   TlsConfig
}

type ResetOffset string

const (
	OffsetLatest   ResetOffset = "LATEST"
	OffsetEarliest ResetOffset = "EARLIEST"
)

func (ro ResetOffset) toSaramaOffsets() int64 {
	switch ro {
	case OffsetEarliest:
		return sarama.OffsetOldest
	case OffsetLatest:
		return sarama.OffsetNewest
	default:
		return sarama.OffsetNewest
	}
}

type ConsumerConfig struct {
	Version string `env:"KAFKA_CONSUMER_VERSION" envDefault:"3.2.0"`
	// client.id - An id string to pass to the server when making requests.
	ClientID string `env:"KAFKA_CONSUMER_CLIENT_ID" validate:"required"`
	// bootstrap.servers - Full set of brokers with "host:port" format.
	BootstrapServers string `env:"KAFKA_CONSUMER_BROKERS" validate:"required"`
	// group.id - ConsumerGroup identifier. The offsets change for the group.
	GroupID string `env:"KAFKA_CONSUMER_GROUP_ID" validate:"required"`
	// auto.offset.reset - One of parameters [OffsetEarliest, OffsetLatest].
	// When ConsumerGroup starts reading messages from partition which already exists messages:
	// OffsetEarliest mean ConsumerGroup will read all messages from partition.
	// OffsetLatest mean ConsumerGroup will read only new messages.
	AutoOffsetReset ResetOffset `env:"KAFKA_CONSUMER_AUTO_OFFSET_RESET" validate:"required,oneof=EARLIEST LATEST"`
	// auto.commit.interval.ms - The frequency in milliseconds that the consumer offsets are auto-committed
	// if EnableAutoCommit is true.
	AutoCommitIntervalMs time.Duration `env:"KAFKA_CONSUMER_AUTO_COMMIT_INTERVAL_MS" envDefault:"5s"`
	// fetch.min.bytes - The minimum amount of data the server should return for a fetch request.
	FetchMinBytes int32 `env:"KAFKA_CONSUMER_FETCH_MIN_BYTES" envDefault:"1"`
	// fetch.max.bytes - The maximum number of bytes we will return for a fetch request.
	FetchMaxBytes int32 `env:"KAFKA_CONSUMER_FETCH_MAX_BYTES" envDefault:"52428800"`
	// fetch.max.wait.ms - The maximum amount of time the server will block before answering the fetch request
	// if there isn't sufficient data to immediately satisfy the requirement given by fetch.min.bytes.
	FetchMaxWaitMs time.Duration `env:"KAFKA_CONSUMER_FETCH_MAX_WAIT_MS" envDefault:"500ms"`
	//session.timeout.ms - The timeout used to detect client failures.
	SessionTimeoutMs time.Duration `env:"KAFKA_SESSION_TIMEOUT_MS" envDefault:"10s"`
	// heartbeat.interval.ms - The expected time between heartbeats to the consumer coordinator. Used to ensure that
	// the consumer's session stays active and to facilitate rebalancing when new consumers join or leave the group.
	HeartbeatIntervalMs time.Duration `env:"KAFKA_HEARTBEAT_INTERVAL_MS" envDefault:"3s"`
	SaslConfig          SaslConfig    `envPrefix:"KAFKA_CONSUMER_"`
}
