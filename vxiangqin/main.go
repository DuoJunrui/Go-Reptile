package main

import (
	"Go-Reptile/vxiangqin/engine"
	"Go-Reptile/vxiangqin/scheduler"
	"Go-Reptile/vxiangqin/zhenai/parser"
)

func main() {
	concurrentEngine := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 10,
	}

	concurrentEngine.Run(engine.Request{
		Url:        "https://www.vxiangqin.com/city",
		ParserFunc: parser.ParseCityList,
	})

	//concurrentEngine.Run(engine.Request{
	//	Url:        "https://www.zhenai.com/zhenghun/shanghai",
	//	ParserFunc: parser.ParseCity,
	//})
}
