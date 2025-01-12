package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	// Цикл проходит по каждому этапу в списке функций stages
	for _, stage := range stages {
		// Обернули в анонимная функция fn, которая принимает входной канал in.
		// Вызывает текущий этап stage с входным каналом in и получает выходной канал out.
		fn := func(in In) Out {
			out := stage(in)
			outStream := make(Bi)

			go func() {
				defer close(outStream)

				for v := range out {
					select {
					case outStream <- v: // чтение канал (результат текущий этап stage)
					case <-done:
						return // Сигнальный канал done используется для остановки пайплайна.
					}
				}
			}()
			return outStream
		}

		in = fn(in) // Это означает, что выходной канал текущего этапа становится входным каналом для следующего этапа.
	}
	return in
}
