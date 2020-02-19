package joyride

type RunnerOption func(*Runner)

func WithReader(reader StorageReader) RunnerOption {
	return func(this *Runner) {
		this.reader = reader
	}
}

func WithWriter(writer StorageWriter) RunnerOption {
	return func(this *Runner) {
		this.writer = writer
	}
}

func WithDispatcher(dispatcher MessageDispatcher) RunnerOption {
	return func(this *Runner) {
		this.dispatcher = dispatcher
	}
}
