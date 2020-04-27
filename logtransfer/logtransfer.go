package main

import (
	"gopkg.in/ini.v1"
	"log"
	"logtransfer/conf"
	"logtransfer/es"
	"logtransfer/kafkaConsumer"
	"sync"
)

// log transfer
// send data from kafka to ElasticSearch

func main() {
	// 0 加载配置文件
	cfg := new(conf.TransConf)
	err := ini.MapTo(cfg, "./conf/conf.ini")
	if err != nil {
		log.Panicf("can't resolve configure for file, err:%v", err)
	}

	// 1. 初始化
	kafkaConsumer.Init(cfg.KafkaConf.Addr)
	es.Init(cfg.ESConf.Addr, cfg.ESConf.ChanMaxSize)
	// 2. 从kafka取日志数据
	var wg sync.WaitGroup
	wg.Add(2)
	go kafkaConsumer.Reader(cfg.KafkaConf.Topic)
	// 3. 发往ES
	go es.SendToEs(cfg.KafkaConf.Topic)
	wg.Wait()

}
