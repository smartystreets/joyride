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
	Runner interface {
		Run(Task)
	}
)

type DefaultRunner struct {
	reader     Reader
	writer     Writer
	dispatcher Dispatcher
}

func NewRunner(reader Reader, writer Writer, dispatcher Dispatcher) DefaultRunner {
	return DefaultRunner{reader: reader, writer: writer, dispatcher: dispatcher}
}
func (this DefaultRunner) Run(task Task) {
	if task == nil {
		return
	}

	this.reader.Read(task.Reads()...)
	task.Run()
	this.writer.Write(task.Writes()...)
	this.dispatcher.Dispatch(task.Messages()...)

	this.Run(task.Next())
}
