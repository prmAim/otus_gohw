package hw06pipelineexecution

type (
	In  = <-chan interface{} // Входной канал
	Out = In                 // Выходной канал
	Bi  = chan interface{}   // Двусторонний канал
)

// Стейдж - функция, принимающая канал на чтение и отдающая канал на чтение, внутри в горутине берущая данные из входного канала, выполняющая полезную работу и отдающая результат в выходной канал:
type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	outStream := make(Bi) // Создаем выходной канал для результата

	go func() {
		defer close(outStream)

		for i := range in { // перебор всех входных данных (пока не закроется канал)
			select {
			case <-done: // канал на прекращение работы
				return
			case outStream <- i:
			}
		}
	}()
	return outStream
}
