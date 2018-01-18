package joyride

type Runner struct {
	factory    Factory
	reader     Reader
	writer     Writer
	dispatcher Dispatcher
}

func New(factory Factory, reader Reader, writer Writer, dispatcher Dispatcher) *Runner {
	return &Runner{
		factory:    factory,
		reader:     reader,
		writer:     writer,
		dispatcher: dispatcher,
	}
}

func (this *Runner) Handle(message interface{}) { // compatibility with Handler interface
	this.Run(message)
}

func (this *Runner) Run(message interface{}) {
	this.run(this.factory(message))
}

func (this *Runner) run(procedure Procedure) {
	if procedure == nil {
		return
	}

	defer procedure.Finally()
	this.reader.Read(procedure.Reads()...)
	procedure.Execute()
	this.writer.Write(procedure.Writes()...)
	this.dispatcher.Dispatch(procedure.Messages()...)

	this.run(procedure.Next())
}
