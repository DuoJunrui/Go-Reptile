package main

import (
	"Go-Reptile/vxiangqin/engine"
	"Go-Reptile/vxiangqin/persist"
	"Go-Reptile/vxiangqin/scheduler"
	"Go-Reptile/vxiangqin/xaingqin/parser"
)

func main() {
	concurrentEngine := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 50,
		ItemChan:    persist.ItemSaver(),
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
