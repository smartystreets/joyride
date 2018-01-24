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
	Builder interface {
		Build(...interface{}) ExecutableTask
	}
)

type Handler struct {
	builder    Builder
	reader     Reader
	writer     Writer
	dispatcher Dispatcher
}

func NewHandler(builder Builder, reader Reader, writer Writer, dispatcher Dispatcher) Handler {
	return Handler{
		builder:    builder,
		reader:     reader,
		writer:     writer,
		dispatcher: dispatcher,
	}
}

func (this Handler) Handle(messages ...interface{}) {
	this.run(this.builder.Build(messages...))
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
