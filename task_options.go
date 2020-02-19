package joyride

type TaskOption func(*Task)

func WithPreparedRead(items ...interface{}) TaskOption {
	return func(this *Task) {
		this.PrepareRead(items...)
	}
}

func WithPreparedWrite(items ...interface{}) TaskOption {
	return func(this *Task) {
		this.PrepareWrite(items...)
	}
}

func WithPreparedDispatch(items ...interface{}) TaskOption {
	return func(this *Task) {
		this.PrepareDispatch(items...)
	}
}
