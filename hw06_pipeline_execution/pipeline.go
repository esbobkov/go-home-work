package hw06_pipeline_execution //nolint:golint,stylecheck

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in, done In, stages ...Stage) Out {
	for _, stage := range stages {
		stagesCh := make(Bi)

		go func(in In, out Bi) {
			defer close(out)
			for {
				select {
				case v, ok := <-in:
					if !ok {
						return
					}
					out <- v
				case <-done:
					return
				}
			}
		}(in, stagesCh)
		in = stage(stagesCh)
	}
	return in
}
