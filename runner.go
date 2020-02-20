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

func (this Runner) Run(task Executable) {
	if task == nil {
		return
	}

	if reads, ok := task.(RequiredReads); ok {
		this.reader.Read(reads.RequiredReads()...)
	}
	result := task.Execute()
	this.writer.Write(result.PendingWrites...)
	this.dispatcher.Dispatch(result.PendingMessages...)
	this.Run(result.SubsequentTask)
}
