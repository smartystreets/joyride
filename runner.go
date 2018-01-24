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
	builder    func(...interface{}) ExecutableTask
	reader     Reader
	writer     Writer
	dispatcher Dispatcher
}

func NewRunner(builder func(...interface{}) ExecutableTask, reader Reader, writer Writer, dispatcher Dispatcher) Runner {
	return Runner{
		builder:    builder,
		reader:     reader,
		writer:     writer,
		dispatcher: dispatcher,
	}
}

func (this Runner) Run(messages ...interface{}) {
	this.run(this.builder(messages...))
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
