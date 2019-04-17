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
	if err := syscall.Kill(syscall.Getpid(), syscall.SIGINT); err != nil {
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
