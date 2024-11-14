package kafka

import "github.com/IBM/sarama"

// Header stores key and value for a record header. It may be used to store request id, other metadata.
type Header struct {
	Key   []byte
	Value []byte
}

type ProducerRecord struct {
	Topic   string
	Key     []byte
	Value   []byte
	Headers []Header
}

type RecordMetadata struct {
	Partition int32
	Offset    int64
	Topic     string
	Key       []byte
	Value     []byte
}

type RecordMetadataWithError struct {
	RecordMetadata *RecordMetadata
	Error          error
}

type ConsumerRecord struct {
	Key       []byte
	Value     []byte
	Topic     string
	Partition int32
	Offset    int64
	Headers   []Header

	groupSession sarama.ConsumerGroupSession
}
