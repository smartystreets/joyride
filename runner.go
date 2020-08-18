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

	this.performRequiredReads(ctx, task)

	result := task.Execute(ctx)
	if result == nil {
		return
	}

	this.performWrites(ctx, result)
	this.performDispatches(ctx, result)

	this.Run(ctx, result.SubsequentTask())
}

func (this Runner) performRequiredReads(ctx context.Context, task Executable) {
	reader, ok := task.(RequiredReads)
	if !ok {
		return
	}
	reads := reader.RequiredReads()
	if len(reads) > 0 {
		this.reader.Read(ctx, reads...)
	}
}

func (this Runner) performWrites(ctx context.Context, result TaskResult) {
	writes := result.PendingWrites()
	if len(writes) > 0 {
		this.writer.Write(ctx, writes...)
	}
}

func (this Runner) performDispatches(ctx context.Context, result TaskResult) {
	messages := result.PendingMessages()
	if len(messages) > 0 {
		this.dispatcher.Dispatch(ctx, messages...)
	}
}

type nop struct{}

func (nop) Dispatch(_ context.Context, _ ...interface{}) {}
func (nop) Read(_ context.Context, _ ...interface{})     {}
func (nop) Write(_ context.Context, _ ...interface{})    {}
