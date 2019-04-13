package main

import (
	"gocron-graceful"
	"log"
	"sync"

	"github.com/jasonlvhit/gocron"
)

func main() {
	graceful.Worker("example", task)
}

func task(wg *sync.WaitGroup) {
	cron := gocron.NewScheduler()
	cron.Every(1).Seconds().Do(graceful.TaskWrapper, wg, job1)
	cron.Every(2).Seconds().Do(graceful.TaskWrapper, wg, job2)
	go func() {
		<-cron.Start()
	}()
}

func job1() {
	log.Println("job 1")
}

func job2() {
	log.Println("job 2")
}
