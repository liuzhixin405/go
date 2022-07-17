package main

import (
	"fmt"

	"github.com/Shopify/sarama"
)

func main() {
	consumer, err := sarama.NewConsumer([]string{"192.168.126.150:9092"}, nil)
	if err != nil {
		fmt.Printf("fail to start consumer, err: %v\n", err)
	}

	partitionList, err := consumer.Partitions("test_kafka01")
	if err != nil {
		fmt.Printf("fail to get list of partition, err: %v\n", err)
		return
	}
	fmt.Println("分区:", partitionList)

	for partition := range partitionList {
		pc, err := consumer.ConsumePartition("test_kafka01", int32(partition), sarama.OffsetNewest)
		if err != nil {
			fmt.Printf("fail to start consumer, err: %v\n", err)
			return
		}
		defer pc.AsyncClose()
		go func(sarama.PartitionConsumer) {
			for msg := range pc.Messages() {
				fmt.Printf("Partition:%d, Offset:%d, Key:%v, Value:%v\n", msg.Partition, msg.Offset, msg.Key, string(msg.Value))
			}
		}(pc)
		select {}
	}
}
