package conf

import (
	"errors"
	"fmt"
	"gopkg.in/ini.v1"
	"reflect"
)

var appconf AppConf

type AppConf struct {
	KafkaConf `ini:"kafka"`
	EsConf `ini:"es"`
}

type KafkaConf struct {
	Addrs []string `ini:"addrs"`
}

type EsConf struct {
	Addrs []string `ini:"addrs"`
	Jobsize int `ini:"jobsize"`
}

func init(){
	fmt.Println("start load conf")
	err := ini.MapTo(&appconf,"./conf/conf.ini")
	fmt.Println("end load conf",err,appconf)
}

func GetKafkaConf() KafkaConf{
	if reflect.DeepEqual(appconf,AppConf{}) {
		panic(errors.New("load conf failed "))
	}
	return appconf.KafkaConf
}

func GetEsConf() EsConf{
	if reflect.DeepEqual(appconf,AppConf{}) {
		panic(errors.New("load conf failed "))
	}
	return appconf.EsConf
}