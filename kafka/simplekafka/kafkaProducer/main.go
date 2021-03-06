package main

import (
	"fmt"

	"github.com/Shopify/sarama"
)

func main() {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	msg := &sarama.ProducerMessage{}

	msg.Topic = "test_kafka01"
	msg.Value = sarama.StringEncoder("this is a test message")
	client, err := sarama.NewSyncProducer([]string{"192.168.126.150:9092"}, config)
	if err != nil {
		fmt.Println("producer closed , err:", err)
		return
	}
	defer client.Close()

	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		fmt.Println("send message failed,err:", err)
		return
	}
	fmt.Printf("pid:%v offset:%v\n", pid, offset)

}
