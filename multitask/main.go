package main

import (
	"Go-Reptile/multitask/engine"
	"Go-Reptile/multitask/scheduler"
	"Go-Reptile/multitask/zhenai/parser"
)

func main() {
	concurrentEngine := engine.ConcurrentEngine{
		Scheduler:   &scheduler.SimpleScheduler{},
		WorkerCount: 10,
	}
	concurrentEngine.Run(engine.Request{
		Url:        "https://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
