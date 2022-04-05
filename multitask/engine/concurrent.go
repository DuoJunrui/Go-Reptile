package engine

import (
	"log"
)

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
}

type Scheduler interface {
	Submit(Request)
	ConfigWorkerChan(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request) {

	in := make(chan Request)
	out := make(chan ParseResult)
	e.Scheduler.ConfigWorkerChan(in)

	for i := 0; i < e.WorkerCount; i++ {
		createWorker(in, out)
	}

	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}

	itemCount := 0
	for {
		result := <-out
		for _, item := range result.Items {
			log.Printf("Get #%d item: %v\n", itemCount, item)
			itemCount++
		}
		for _, request := range result.Requests {
			e.Scheduler.Submit(request)
		}
	}
}

func createWorker(in chan Request, out chan ParseResult) {
	go func() {
		for {
			request := <-in // 1.2-1这里会出现循环等待，假死
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result // 1.2-2这里会出现循环等待，假死
		}
	}()
}
