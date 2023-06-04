package scheduler

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type Job struct {
	state    atomic.Uint32
	exit     atomic.Uint32
	name     string
	function func()
	duration time.Duration
	quit     chan bool
}

func NewJob(name string, f func(), d time.Duration) Job {
	return Job{
		name:     name,
		duration: d,
		function: f,
		quit:     make(chan bool),
	}
}

func (j *Job) lock() {
	j.state.Store(1)
}

func (j *Job) unlock() {
	j.state.Store(0)
}

func (j *Job) terminate() {
	j.exit.Store(1)
}

// check if job task is ready to run
func (j *Job) ready() bool {
	return j.exit.Load() == 0 && j.state.Load() == 0
}

func (j *Job) do(wg *sync.WaitGroup) {
	defer wg.Done()

	ticker := time.NewTicker(j.duration)
	for {
		select {
		case <-j.quit:
			fmt.Println("stop command received-> ", j.name)
			ticker.Stop()
			return
		case <-ticker.C:
			if j.ready() {
				j.lock()
				j.function() // exec job task
				j.unlock()
			}
		}
	}
}
