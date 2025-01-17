package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	outStream := in

	// Цикл проходит по каждому этапу в списке функций stages
	for _, stage := range stages {
		outStream = runStage(outStream, done, stage) // Это означает, что выходной канал текущего этапа становится входным каналом для следующего этапа.
	}
	return outStream
}

func runStage(in In, done In, stage Stage) Out {
	out := make(Bi)

	go func() {
		defer close(out)

		for {
			select {
			case <-done: // Сигнальный канал done используется для остановки пайплайна.
				return
			case value, ok := <-in:
				if !ok {
					return
				}
				// Запускаем стейдж. Передает канал с данными в стейдж, который читает из него данные и возвращает результат.
				stageOut := stage(valueToChannel(value, done))

				// Передаем результат в выходной канал
				select {
				case out <- <-stageOut:
				case <-done:
					return
				}
			}
		}
	}()

	return out
}

// Создает новый канал, в который отправляется значение value
func valueToChannel(value interface{}, done <-chan interface{}) In {
	channel := make(Bi)

	go func() {
		defer close(channel)

		select {
		case channel <- value:
		case <-done: // Сигнальный канал done используется для остановки пайплайна.
			return
		}
	}()
	return channel
}
