package joyride

type Runner struct {
	reader     StorageReader
	writer     StorageWriter
	dispatcher MessageDispatcher
}

func NewRunner(options ...RunnerOption) Runner {
	runner := Runner{
		reader:     nopIO{},
		writer:     nopIO{},
		dispatcher: nopIO{},
	}
	for _, option := range options {
		option(&runner)
	}
	return runner
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
