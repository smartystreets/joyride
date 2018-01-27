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

type Runner struct {
	reader     Reader
	writer     Writer
	dispatcher Dispatcher
}

func NewRunner(reader Reader, writer Writer, dispatcher Dispatcher) Runner {
	return Runner{reader: reader, writer: writer, dispatcher: dispatcher}
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
