package kafkaConsumer

import (
	"github.com/Shopify/sarama"
	"fmt"
	"z.cn/logtransferDemo/common"
	"z.cn/logtransferDemo/conf"
)



func InitKafkaConnect(topic string,conf conf.KafkaConf){
	consumer ,err := sarama.NewConsumer(conf.Addrs,nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	partitions , err := consumer.Partitions(topic)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(partitions)
	for _,p := range partitions {
		// 针对每个分区创建一个对应的分区消费者
		pc, err := consumer.ConsumePartition(topic, int32(p), sarama.OffsetNewest)
		if err != nil {
			fmt.Printf("failed to start consumer for partition %d,err:%v\n", p, err)
			return
		}
		//defer pc.AsyncClose()
		// 异步从每个分区消费信息
		go func(sarama.PartitionConsumer) {
			for msg := range pc.Messages() {
				fmt.Printf("read kafka msg : Partition:%d Offset:%d Key:%v Value:%v \n", msg.Partition, msg.Offset, msg.Key, string(msg.Value))
				data := &common.LogData{
					Topic: topic,
					Data: string(msg.Value),
				}
				common.LogDataJob <- data
				fmt.Printf("add es logdata job:%#v \n",*data)
			}
		}(pc)
	}
}