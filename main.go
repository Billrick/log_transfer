package main

import (
	"sync"
	"z.cn/logtransferDemo/conf"
	"z.cn/logtransferDemo/es"
	"z.cn/logtransferDemo/kafkaConsumer"
)

func main() {
	es.InitElasticConnect(conf.GetEsConf())
	kafkaConsumer.InitKafkaConnect("notify",conf.GetKafkaConf())
	var wg sync.WaitGroup
	wg.Add(1)
	wg.Wait()
}
