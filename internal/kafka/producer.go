package kafka

import (
	"log"

	"github.com/IBM/sarama"
)

func ProduceMessage(brokers []string, topic string, message []byte) error {
	config := sarama.NewConfig()
	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		return err
	}
	defer producer.Close()

	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(message),
	}

	_, _, err = producer.SendMessage(msg)
	if err != nil {
		log.Printf("Failed to send message: %v", err)
	}

	return err
}
