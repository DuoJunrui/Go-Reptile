package scheduler

import "Go-Reptile/multitask/engine"

type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (s *SimpleScheduler) ConfigWorkerChan(c chan engine.Request) {
	s.workerChan = c
}

func (s *SimpleScheduler) Submit(r engine.Request) {
	// 解决循环等待 假死
	go func() { s.workerChan <- r }()
}
