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
	init       func(...interface{}) ExecutableTask
	reader     Reader
	writer     Writer
	dispatcher Dispatcher
}

func NewRunner(init func(...interface{}) ExecutableTask, reader Reader, writer Writer, dispatcher Dispatcher) Runner {
	return Runner{
		init:       init,
		reader:     reader,
		writer:     writer,
		dispatcher: dispatcher,
	}
}

func (this Runner) Run(messages ...interface{}) {
	this.run(this.init(messages...))
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
