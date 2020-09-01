package common

var LogDataJob chan *LogData

type LogData struct {
	Topic string `json:"topic"` // index索引库
	Data string`json:"data"`
}

