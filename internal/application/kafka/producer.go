package kafka

import (
	"fmt"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

func NewKafkaProducer() *ckafka.Producer {
	configMap := &ckafka.ConfigMap{
		//no docker o bootstrap.servers:kafka:9092
		"bootstrap.servers":"localhost:9092",
	}
	p,err := ckafka.NewProducer(configMap)
	if err != nil {
		panic(err)
	}

	return p
}

//key é para o kafka escolher para qual particao quer enviar a msg
func  Publish(msg string, topic string, producer *ckafka.Producer, deliveryChan chan ckafka.Event) error {
	message := &ckafka.Message{
		TopicPartition: ckafka.TopicPartition{
			Topic: &topic,
			//envia a msg para qualquer particao do topico
			Partition: ckafka.PartitionAny,
		},
		Value: []byte(msg),
	}
	err := producer.Produce(message,deliveryChan)
	if err != nil {
		return err
	}
	return nil
}

//receber resultado da publicacao de msg no kafka
func DeliveryReport(deliveryChannel chan ckafka.Event) {
	for e := range deliveryChannel {
		//verificar o resutado da publicacao da msg
		switch ev := e.(type) {
		case *ckafka.Message:
			//erro na entrega da msg
			if ev.TopicPartition.Error != nil {
				fmt.Println("delivery msg failed: ",ev.TopicPartition)
			} else {
				fmt.Println("delivered msg to: ", ev.TopicPartition)
			}
		}
	}
}