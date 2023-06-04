package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sskriv/go-scheduling/scheduler"
)

func main() {
	jobs := []scheduler.Job{
		scheduler.NewJob(
			"job A",
			func() { fmt.Println("job A") },
			time.Second,
		),
		scheduler.NewJob(
			"job B",
			func() {
				time.Sleep(5 * time.Second)
				fmt.Println("job B")
			},
			time.Second,
		),
	}

	s := scheduler.New(jobs)
	s.Run()

	// graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	done := make(chan bool, 1)
	s.Stop(done)
	<-done
	fmt.Println("scheduling server exiting.")
}
