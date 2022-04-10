package engine

import "Go-Reptile/vxiangqin/model"

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
	ItemChan    chan interface{}
}

type Scheduler interface {
	ReadyNotifier
	Submit(Request)
	WorkerChan() chan Request
	Run()
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request) {

	out := make(chan ParseResult)
	e.Scheduler.Run()

	for i := 0; i < e.WorkerCount; i++ {
		createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}

	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}

	for {
		result := <-out
		for _, item := range result.Items {
			//判断item的数据类型，如果是model.Profile，将item加入到ItemChan中,因为save的时候只保存用户数据
			switch item.(type) {
			case model.Profile:
				go func() { e.ItemChan <- item }()
			}
		}
		for _, request := range result.Requests {
			e.Scheduler.Submit(request)
		}
	}
}

func createWorker(in chan Request, out chan ParseResult, ready ReadyNotifier) {
	go func() {
		for {
			//tell scheduler I`m ready
			ready.WorkerReady(in) // 1.2-1这里会出现循环等待，假死
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result // 1.2-2这里会出现循环等待，假死
		}
	}()
}
