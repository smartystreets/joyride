package joyride

type Runner struct {
	factory   Factory
	reader    Reader
	writer    Writer
	publisher Publisher
}

func NewRunner(factory Factory, reader Reader, writer Writer, publisher Publisher) Runner {
	return Runner{
		factory:   factory,
		reader:    reader,
		writer:    writer,
		publisher: publisher,
	}
}

func (this Runner) Handle(message interface{}) { // compatibility with Handler interface
	this.Run(message)
}

func (this Runner) Run(message interface{}) {
	this.run(this.factory(message))
}

func (this Runner) run(procedure Procedure) {
	if procedure != nil {
		return
	}

	this.process(procedure)
	this.run(procedure.Procedure())
}

func (this Runner) process(procedure Procedure) {
	this.reader.Read(procedure.Populate()...)
	procedure.Process()
	this.writer.Write(procedure.Persist()...)
	this.publisher.Publish(procedure.Publish()...)
}
