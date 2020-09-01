package es

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"time"
	"z.cn/logtransferDemo/common"
	"z.cn/logtransferDemo/conf"
)

var (
	client *elastic.Client

)

func init(){
	if err := InitElasticConnect(conf.GetEsConf()) ;err != nil{
		fmt.Println("init elastic connect failed ,err",err)
		return
	}
}

func InitElasticConnect(esConf conf.EsConf) (err error){
	common.LogDataJob = make(chan *common.LogData,esConf.Jobsize)
	client,err = elastic.NewClient(elastic.SetURL(esConf.Addrs...))
	if err != nil{
		fmt.Println("init elastic connect failed ,err",err)
		return
	}
	//启动发送es的自动任务
	go StartSendLogDataToElasticJob()
	return nil
}

func StartSendLogDataToElasticJob(){
	if client == nil {
		err := InitElasticConnect(conf.GetEsConf())
		if err!=nil{
			fmt.Println("init elastic connect failed ,err",err)
			return
		}
	}
	for  {
		select {
			case msg := <-common.LogDataJob:
				res,err := client.Index().Index(msg.Topic).BodyJson(msg).Do(context.Background())
				if err != nil {
					fmt.Println("add doc failed ,err",err)
					continue
				}
				fmt.Printf("add doc success ; return data : %#v",res)
		default:
			time.Sleep(time.Second)
		}
	}
}