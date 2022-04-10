package main

import (
	"Go-Spider/multitask/engine"
	"Go-Spider/multitask/scheduler"
	"Go-Spider/multitask/zhenai/parser"
)

func main() {
	concurrentEngine := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 10,
	}

	//concurrentEngine.Run(engine.Request{
	//	Url:        "https://www.zhenai.com/zhenghun",
	//	ParserFunc: parser.ParseCityList,
	//})

	concurrentEngine.Run(engine.Request{
		Url:        "https://www.zhenai.com/zhenghun/shanghai",
		ParserFunc: parser.ParseCity,
	})
}
