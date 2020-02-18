package joyride

type (
	Reader interface {
		Read(...interface{})
	}
	Writer interface {
		Write(...interface{})
	}
	Dispatcher interface {
		Dispatch(...interface{})
	}
	TaskRunner interface {
		Run(RunnableTask)
	}
)

type RunnerOption func(*Runner)

type nopDispatcher struct{}

func (n nopDispatcher) Dispatch(...interface{}) {}

type nopReader struct{}

func (n nopReader) Read(...interface{}) {}

type nopWriter struct{}

func (n nopWriter) Write(...interface{}) {}


func WithReader(reader Reader) RunnerOption {
	return func(runner *Runner) {
		runner.reader = reader
	}
}

func WithWriter(writer Writer) RunnerOption {
	return func(runner *Runner) {
		runner.writer = writer
	}
}

func WithDispatcher(dispatcher Dispatcher) RunnerOption {
	return func(runner *Runner) {
		runner.dispatcher = dispatcher
	}
}

type Runner struct {
	reader     Reader
	writer     Writer
	dispatcher Dispatcher
}

func NewRunner(options ...RunnerOption) Runner {
	runner := Runner{
		reader:     nopReader{},
		writer:     nopWriter{},
		dispatcher: nopDispatcher{},
	}
	for _, option := range options {
		option(&runner)
	}
	return runner
}

func (this Runner) Run(task RunnableTask) {
	if task == nil {
		return
	}

	this.reader.Read(task.Reads()...)
	task.Run()
	this.writer.Write(task.Writes()...)
	this.dispatcher.Dispatch(task.Messages()...)

	this.Run(task.Next())
}
