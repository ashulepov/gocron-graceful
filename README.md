# gocron-graceful
Graceful shutdown implementation for [gocron](https://github.com/jasonlvhit/gocron) based worker 

## Installing

```
go get github.com/ashulepov/gocron-graceful
```

## Using

```
package main

import (
	"log"
	"sync"

	"github.com/ashulepov/gocron-graceful"
	"github.com/jasonlvhit/gocron"
)

func main() {
	graceful.Worker("tasks", task)
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
```

## Testing
Unit-tests:


```
go test -v
```