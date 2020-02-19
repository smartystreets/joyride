package joyride

type RunnerOption func(*Runner)

func WithStorageReader(reader StorageReader) RunnerOption {
	return func(this *Runner) { this.reader = reader }
}

func WithStorageWriter(writer StorageWriter) RunnerOption {
	return func(this *Runner) { this.writer = writer }
}

func WithMessageDispatcher(dispatcher MessageDispatcher) RunnerOption {
	return func(this *Runner) { this.dispatcher = dispatcher }
}
