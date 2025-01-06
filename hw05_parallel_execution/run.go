package hw05parallelexecution

import (
	"errors"
	"fmt"
	"runtime"
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

	fmt.Println("Кол-во каналов в начале:", runtime.NumGoroutine())

	// Запуск n горутин
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			// Получаем информацию о текущем вызове
			pc, _, _, _ := runtime.Caller(1)
			fn := runtime.FuncForPC(pc)
			// Получаем ID горутины
			gid := getGoroutineID()

			for task := range taskCh {
				fmt.Printf("%d) Горутина ID: %d, Функция: %s\n", i, gid, fn.Name())

				if err := task(); err != nil {
					mu.Lock()
					errCount++
					fmt.Printf("%d) Горутина ID: %d, Ошибка: %q, Кол-во ошибок: %d \n", i, gid, err, errCount)
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
	fmt.Println("Кол-во каналов в конце:", runtime.NumGoroutine())

	if errCount >= m {
		return ErrErrorsLimitExceeded
	}

	return nil
}

func getGoroutineID() uint64 {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	// Строка будет выглядеть как "goroutine 1 [running]: ..."
	// Извлекаем ID горутины из строки
	var id uint64
	fmt.Sscanf(string(buf[:n]), "goroutine %d", &id)
	return id
}
