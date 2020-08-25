package hw05_parallel_execution //nolint:golint,stylecheck

import (
	"errors"
	"sync"
	"sync/atomic"
)

var (
	ErrErrorsLimitExceeded = errors.New("errors limit exceeded")
	errorsCount            int32
	maxErrorsCount         int32
)

type Task func() error

// Run starts tasks in N goroutines and stops its work when receiving M errors from tasks.
func Run(tasks []Task, goroutinesCount, m int) error {
	errorsCount = 0
	maxErrorsCount = int32(m)

	jobs := make(chan Task, len(tasks))
	quit := make(chan struct{})

L:
	for _, task := range tasks {
		select {
		case <-quit:
			break L
		default:
			jobs <- task
		}
	}
	close(jobs)

	var wg sync.WaitGroup
	for i := 0; i < goroutinesCount; i++ {
		wg.Add(1)
		go consume(jobs, &wg)
	}

	close(quit)
	wg.Wait()

	if maxErrorsCount >= 0 && atomic.LoadInt32(&errorsCount) >= maxErrorsCount {
		return ErrErrorsLimitExceeded
	}

	return nil
}

func consume(jobs <-chan Task, wg *sync.WaitGroup) {
	for task := range jobs {
		err := task()

		if maxErrorsCount < 0 {
			continue
		}

		if err == nil {
			continue
		}

		atomic.AddInt32(&errorsCount, 1)
		if atomic.LoadInt32(&errorsCount) >= maxErrorsCount {
			break
		}
	}
	wg.Done()
}
