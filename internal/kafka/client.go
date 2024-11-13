package kafka

import (
	"log"

	"github.com/IBM/sarama"
)

type Client struct {
	Producer sarama.SyncProducer
	Consumer sarama.Consumer
	Brokers  []string
}

// NewClient initializes a new Kafka client
func NewClient(brokers []string) *Client {
	config := sarama.NewConfig()

	// Producer config
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true

	// Consumer config
	config.Consumer.Return.Errors = true

	// Create a producer
	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		log.Fatalf("Failed to create Kafka producer: %v", err)
	}

	// Create a consumer
	consumer, err := sarama.NewConsumer(brokers, config)
	if err != nil {
		log.Fatalf("Failed to create Kafka consumer: %v", err)
	}

	log.Println("Kafka client initialized successfully")

	return &Client{
		Producer: producer,
		Consumer: consumer,
		Brokers:  brokers,
	}
}

// Close closes the Kafka client (producer and consumer)
func (c *Client) Close() {
	if err := c.Producer.Close(); err != nil {
		log.Printf("Failed to close Kafka producer: %v", err)
	}

	if err := c.Consumer.Close(); err != nil {
		log.Printf("Failed to close Kafka consumer: %v", err)
	}
}
