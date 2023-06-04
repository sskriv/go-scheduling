package scheduler

import (
	"fmt"
	"sync"
)

type Scheduler struct {
	jobs []Job
	wg   *sync.WaitGroup
}

func New(j []Job) *Scheduler {
	var wg sync.WaitGroup
	return &Scheduler{j, &wg}
}

func (s *Scheduler) Run() {
	for i := 0; i < len(s.jobs); i++ {
		s.wg.Add(1)
		go s.jobs[i].do(s.wg)
	}
}

// Stop() waits until all scheduled jobs are finished
func (s *Scheduler) Stop(done chan bool) {
	for i := 0; i < len(s.jobs); i++ {
		s.jobs[i].terminate()
		fmt.Println("stop command send -> ", s.jobs[i].name)
		s.jobs[i].quit <- true
	}

	s.wg.Wait()
	done <- true
}
