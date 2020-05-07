package joyride

import "context"

type Runner struct {
	reader     StorageReader
	writer     StorageWriter
	dispatcher MessageDispatcher
}

func NewRunner(options ...RunnerOption) Runner {
	runner := Runner{reader: nop{}, writer: nop{}, dispatcher: nop{}}
	for _, option := range options {
		option(&runner)
	}
	return runner
}

func (this Runner) Run(ctx context.Context, task Executable) {
	if task == nil {
		return
	}

	if reads, ok := task.(RequiredReads); ok {
		this.reader.Read(ctx, reads.RequiredReads()...)
	}
	result := task.Execute(ctx)
	if result == nil {
		return
	}
	this.writer.Write(ctx, result.PendingWrites()...)
	this.dispatcher.Dispatch(ctx, result.PendingMessages()...)
	this.Run(ctx, result.SubsequentTask())
}

type nop struct{}

func (nop) Dispatch(_ context.Context, _ ...interface{}) {}
func (nop) Read(_ context.Context, _ ...interface{})     {}
func (nop) Write(_ context.Context, _ ...interface{})    {}
