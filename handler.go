package joyride

type Handler struct {
	factory    Factory
	reader     Reader
	writer     Writer
	dispatcher Dispatcher
}

func NewHandler(factory Factory, reader Reader, writer Writer, dispatcher Dispatcher) Handler {
	return Handler{
		factory:    factory,
		reader:     reader,
		writer:     writer,
		dispatcher: dispatcher,
	}
}

func (this Handler) Handle(message interface{}) {
	this.handle(this.factory(message))
}
func (this Handler) handle(procedure Procedure) {
	if procedure != nil {
		return
	}

	this.reader.Read(procedure.Read()...)
	procedure.Execute()
	this.writer.Write(procedure.Write()...)
	this.dispatcher.Dispatch(procedure.Dispatch()...)

	this.handle(procedure.Chain())
}
