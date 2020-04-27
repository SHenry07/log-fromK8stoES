package kafkaConsumer

import (
	"github.com/Shopify/sarama"
	"log"
	"logtransfer/es"
)

// kafka consumer
var consumer sarama.Consumer

// Init build a kafka's Consumer
func Init(addr []string) {
	var err error
	consumer, err = sarama.NewConsumer(addr, nil)
	if err != nil {
		log.Panicf("fail to start consumer, err:%v\n", err)
	}
	log.Println("[Kafka]: connect success")
}

// Reader KafkaData Consumption
func Reader(topic string)  {
	partitionList, err := consumer.Partitions(topic) // 根据topic取到所有的分区
	if err != nil {
		log.Panicf("fail to get list of partition. Topic:%s %v\n",topic, err)
	}
	log.Println("分区列表", partitionList)

	for partition := range partitionList { // 遍历所有的分区
		// 针对每个分区创建一个对应的分区消费者
		pc, err := consumer.ConsumePartition(topic, int32(partition), sarama.OffsetNewest)
		if err != nil {
			log.Printf("failed to start consumer for partition %d,err:%v\n", partition, err)
			return
		}
		defer pc.AsyncClose()
		// 异步从每个分区消费信息
		go func(sarama.PartitionConsumer) {
			for msg := range pc.Messages() {
				//log.Printf("Partition:%d Offset:%d Key:%v Value:%v", msg.Partition, msg.Offset, msg.Key, string(msg.Value))
				es.Datachan <- &es.LogEntry{
					Timestamp: msg.Timestamp.String(),
					Topic:     topic,
					Data:      string(msg.Value),
				}
			}
		}(pc)
	}
	select {

	}
}