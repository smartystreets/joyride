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

type Handler struct {
	builder    func(...interface{}) ExecutableTask
	reader     Reader
	writer     Writer
	dispatcher Dispatcher
}

func NewHandler(builder func(...interface{}) ExecutableTask, reader Reader, writer Writer, dispatcher Dispatcher) Handler {
	return Handler{
		builder:    builder,
		reader:     reader,
		writer:     writer,
		dispatcher: dispatcher,
	}
}

func (this Handler) Handle(messages ...interface{}) {
	this.run(this.builder(messages...))
}
func (this Handler) run(task ExecutableTask) {
	if task == nil {
		return
	}

	this.reader.Read(task.Reads()...)
	task.Execute()
	this.writer.Write(task.Writes()...)
	this.dispatcher.Dispatch(task.Messages()...)

	this.run(task.Next())
}
