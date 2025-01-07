package hw05parallelexecution

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"sync/atomic"
	"testing"
	"time"

	//nolint:depguard
	"github.com/bxcodec/faker/v3"
	//nolint:depguard
	"github.com/stretchr/testify/require"
	"go.uber.org/goleak"
)

func TestRun(t *testing.T) {
	defer goleak.VerifyNone(t)

	t.Run("if were errors in first M tasks, than finished not more N+M tasks", func(t *testing.T) {
		tasksCount := 50
		tasks := make([]Task, 0, tasksCount)

		var runTasksCount int32

		for i := 0; i < tasksCount; i++ {
			err := fmt.Errorf("error from task %d", i)
			tasks = append(tasks, func() error {
				time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)))
				atomic.AddInt32(&runTasksCount, 1)
				return err
			})
		}

		workersCount := 10
		maxErrorsCount := 23
		err := Run(tasks, workersCount, maxErrorsCount)

		require.Truef(t, errors.Is(err, ErrErrorsLimitExceeded), "actual err - %v", err)
		require.LessOrEqual(t, runTasksCount, int32(workersCount+maxErrorsCount), "extra tasks were started")
	})

	t.Run("tasks without errors", func(t *testing.T) {
		tasksCount := 50
		tasks := make([]Task, 0, tasksCount)

		var runTasksCount int32
		var sumTime time.Duration

		for i := 0; i < tasksCount; i++ {
			taskSleep := time.Millisecond * time.Duration(rand.Intn(100))
			sumTime += taskSleep

			tasks = append(tasks, func() error {
				time.Sleep(taskSleep)
				atomic.AddInt32(&runTasksCount, 1)
				return nil
			})
		}

		workersCount := 5
		maxErrorsCount := 1

		start := time.Now()
		err := Run(tasks, workersCount, maxErrorsCount)
		elapsedTime := time.Since(start)
		require.NoError(t, err)

		require.Equal(t, runTasksCount, int32(tasksCount), "not all tasks were completed")
		require.LessOrEqual(t, int64(elapsedTime), int64(sumTime/2), "tasks were run sequentially?")
	})
}

func TestOfError(t *testing.T) {
	testsData := []struct {
		inputTasksCount     int
		inputWorkersCount   int
		inputMaxErrorsCount int
		expectedError       error
	}{
		{
			inputTasksCount: 10, inputWorkersCount: 0, inputMaxErrorsCount: 2,
			expectedError: errors.New("errors limit workers"),
		},
		{
			inputTasksCount: 10, inputWorkersCount: 4, inputMaxErrorsCount: 0,
			expectedError: errors.New("errors limit exceeded"),
		},
	}

	for _, tc := range testsData {
		tc := tc
		t.Run(tc.expectedError.Error(), func(t *testing.T) {
			tasks := make([]Task, 0, tc.inputTasksCount)

			for i := 0; i < tc.inputTasksCount; i++ {
				err := fmt.Errorf("error from task %d", i)
				tasks = append(tasks, func() error {
					return err
				})
			}

			err := Run(tasks, tc.inputWorkersCount, tc.inputMaxErrorsCount)
			require.Equal(t, tc.expectedError, err)
		})
	}
}

type TestStruct struct {
	TasksCount     int `faker:"oneof:1, 2, 4, 8, 16"`
	WorkersCount   int `faker:"oneof:1,2,3,4,5,6,7,8,9"`
	MaxErrorsCount int `faker:"oneof:1,2,3,4,5,6,7,8,9"`
}

func generateDataOfTest() TestStruct {
	var strFaker TestStruct
	err := faker.FakeData(&strFaker)
	if err != nil {
		fmt.Println("Ошибка генерации данных:", err)
	}
	fmt.Printf("Сгенерированные данные: %+v\n", strFaker) // Вывод сгенерированных данны
	return strFaker
}

func TestGenerateOfError(t *testing.T) {
	defer goleak.VerifyNone(t)

	for i := 0; i < 100; i++ {
		t.Run("generate the tests data "+strconv.Itoa(i), func(t *testing.T) {
			dataOftest := generateDataOfTest()
			tasks := make([]Task, 0, dataOftest.TasksCount)

			var runTasksCount int32

			for i := 0; i < dataOftest.TasksCount; i++ {
				err := fmt.Errorf("error from task %d", i)
				tasks = append(tasks, func() error {
					time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)))
					atomic.AddInt32(&runTasksCount, 1)
					return err
				})
			}

			Run(tasks, dataOftest.WorkersCount, dataOftest.MaxErrorsCount)

			require.LessOrEqual(
				t, runTasksCount, int32(dataOftest.WorkersCount+dataOftest.MaxErrorsCount),
				"extra tasks were started")
		})
	}
}
