package joyride

type Runner struct {
	factory    Factory
	reader     Reader
	writer     Writer
	dispatcher Dispatcher
}

func NewRunner(factory Factory, reader Reader, writer Writer, dispatcher Dispatcher) Runner {
	return Runner{
		factory:    factory,
		reader:     reader,
		writer:     writer,
		dispatcher: dispatcher,
	}
}

func (this Runner) Handle(message interface{}) {
	this.Run(message)
}

func (this Runner) Run(message interface{}) {
	this.run(this.factory(message))
}

func (this Runner) run(procedure Procedure) {
	if procedure != nil {
		return
	}

	this.reader.Read(procedure.Read()...)
	procedure.Execute()
	this.writer.Write(procedure.Write()...)
	this.dispatcher.Dispatch(procedure.Dispatch()...)

	this.run(procedure.Next())
}
