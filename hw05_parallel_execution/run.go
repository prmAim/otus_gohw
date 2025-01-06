package hw05parallelexecution

import (
	"errors"
	"sync"
)

var (
	ErrErrorsLimitExceeded = errors.New("errors limit exceeded")
	ErrErrorsLimitWorker   = errors.New("errors limit workers")
)

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	if m <= 0 {
		return ErrErrorsLimitExceeded
	}
	if n <= 0 {
		return ErrErrorsLimitWorker
	}

	wg := sync.WaitGroup{}

	errCount := 0
	mu := sync.Mutex{}

	taskCh := make(chan Task)

	// Запуск n горутин
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for task := range taskCh {
				if err := task(); err != nil {
					mu.Lock()
					errCount++
					mu.Unlock()
				}
			}
		}()
	}

	// Передача задач в канал
	for _, task := range tasks {
		mu.Lock()
		if errCount >= m {
			mu.Unlock()
			break
		}
		mu.Unlock()
		taskCh <- task
	}
	close(taskCh)

	wg.Wait()

	if errCount >= m {
		return ErrErrorsLimitExceeded
	}

	return nil
}
