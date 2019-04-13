package graceful

import (
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/jasonlvhit/gocron"
)

// Worker runs task worker
func Worker(name string, task func(wg *sync.WaitGroup)) {
	wg := &sync.WaitGroup{}

	task(wg)

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	doneChan := gocron.Start()
	log.Printf("%s: started", name)

	<-signalChan

	log.Printf("%s: shutting down...", name)
	doneChan <- true
	wg.Wait()

	log.Printf("%s: done", name)
}

// TaskWrapper is wrapper for graceful tasks
func TaskWrapper(wg *sync.WaitGroup, task func()) {
	wg.Add(1)
	task()
	wg.Done()
}
