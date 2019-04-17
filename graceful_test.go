package graceful

import (
	"sync"
	"syscall"
	"testing"
	"time"

	"github.com/jasonlvhit/gocron"
)

func TestWorkerKillSIGINT(t *testing.T) {
	testWorker(t, syscall.SIGINT)
}

func TestWorkerKillSIGTERM(t *testing.T) {
	testWorker(t, syscall.SIGTERM)
}

func testWorker(t *testing.T, signal syscall.Signal) {
	// Run Worker
	go Worker("test", task)

	// Sleep
	time.Sleep(100 * time.Millisecond)

	// Kill Worker
	if err := syscall.Kill(syscall.Getpid(), signal); err != nil {
		t.Errorf("kill graceful with '%v' error: %v", signal, err)
	}

	t.Logf("kill graceful with '%v' ok", signal)
}

func task(wg *sync.WaitGroup) {
	cron := gocron.NewScheduler()
	cron.Every(1).Seconds().Do(TaskWrapper, wg, func() {})
	go func() {
		<-cron.Start()
	}()
}

func TestTaskWrapper(t *testing.T) {
	timeout := time.After(3 * time.Second)
	done := make(chan bool)

	go func(done chan bool) {
		// run test
		wg := &sync.WaitGroup{}
		for i := 0; i < 5; i++ {
			TaskWrapper(wg, func() { t.Log("task", i) })
		}
		wg.Wait()

		// synced on time
		done <- true
	}(done)

	select {
	case <-timeout:
		t.Fatal("timeout error")
	case <-done:
	}
}
