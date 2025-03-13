package configs

import (
	"context"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

func NewKafkaConnection(setting *Settings, topic string, partition int) (*kafka.Conn, error) {
	kafkaBrokerUrl := fmt.Sprintf("%v:%v", setting.KafkaHost, setting.KafkaPort)
	conn, err := kafka.DialLeader(context.Background(), setting.KafkaProtocol, kafkaBrokerUrl, topic, partition)

	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	return conn, err
}
