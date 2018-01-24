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
)

type Runner struct {
	builder    Builder
	reader     Reader
	writer     Writer
	dispatcher Dispatcher
}

func NewRunner(init func(...interface{}) ExecutableTask, reader Reader, writer Writer, dispatcher Dispatcher) Runner {
	return Runner{
		builder:    BatchBuilder{callback: init},
		reader:     reader,
		writer:     writer,
		dispatcher: dispatcher,
	}
}

func (this Runner) Run(messages ...interface{}) {
	task := this.builder.Build(messages...)
	this.run(task)
}
func (this Runner) run(task ExecutableTask) {
	if task == nil {
		return
	}

	this.reader.Read(task.Reads()...)
	task.Execute()
	this.writer.Write(task.Writes()...)
	this.dispatcher.Dispatch(task.Messages()...)

	this.run(task.Next())
}

// compatibility with Handler interface
func (this Runner) Handle(messages ...interface{}) {
	this.Run(messages...)
}

/////////////////////////////////////////////////////////////////

type (
	Builder interface {
		Build(...interface{}) ExecutableTask
	}
	BatchBuilder struct {
		callback func(...interface{}) ExecutableTask
	}
)

func (this BatchBuilder) Build(messages ...interface{}) ExecutableTask {
	return this.callback(messages...)
}
