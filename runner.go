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

func WithReader(reader Reader) RunnerOption {
	return func(this *Runner) { this.reader = reader }
}
func WithWriter(writer Writer) RunnerOption {
	return func(this *Runner) { this.writer = writer }
}
func WithDispatcher(dispatcher Dispatcher) RunnerOption {
	return func(this *Runner) { this.dispatcher = dispatcher }
}

type Runner struct {
	reader     Reader
	writer     Writer
	dispatcher Dispatcher
}

func NewRunner(options ...RunnerOption) Runner {
	runner := Runner{
		reader:     nopIO{},
		writer:     nopIO{},
		dispatcher: nopIO{},
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
	task.Execute()
	this.writer.Write(task.Writes()...)
	this.dispatcher.Dispatch(task.Messages()...)

	this.Run(task.Next())
}

type nopIO struct{}

func (nopIO) Dispatch(...interface{}) {}
func (nopIO) Read(...interface{})     {}
func (nopIO) Write(...interface{})    {}
