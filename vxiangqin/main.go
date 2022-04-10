package main

import (
	"Go-Spider/vxiangqin/engine"
	"Go-Spider/vxiangqin/persist"
	"Go-Spider/vxiangqin/scheduler"
	"Go-Spider/vxiangqin/xaingqin/parser"
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
