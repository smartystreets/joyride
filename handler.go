package joyride

import (
	"context"
	"errors"
)

// This type is designed to be embedded into another type that implements the MessageHandler interface such that the
// only method necessary for the embedding type to implement is the remaining method of HandleMessage. In essence, this
// structure becomes, in C# terminology an, "abstract class" which needs some additional behavior in order to function
// properly. In addition, because this struct is embedded, the existing methods of Handle (and more especially) the
// Run method can be overridden by the embedding struct.
type Handler struct {
	runner TaskRunner
	inner  MessageHandler
	tasks  []Executable
}

func NewHandler(inner MessageHandler, runner TaskRunner, tasks ...Executable) *Handler {
	return &Handler{
		inner:  inner,
		runner: runner,
		tasks:  tasks,
	}
}

func (this *Handler) Handle(ctx context.Context, messages ...interface{}) {
	for _, message := range messages {
		if !this.inner.HandleMessage(ctx, message) {
			panic(ErrUnknownType)
		}
	}

	this.inner.Run(ctx)
}

func (this *Handler) Add(task Executable) {
	this.tasks = append(this.tasks, task)
}

func (this *Handler) Run(ctx context.Context) {
	this.runner.Run(ctx, CompositeTask(this.tasks))
	this.tasks = nil
}

func (this *Handler) Tasks() []Executable { return this.tasks }

var ErrUnknownType = errors.New("the handler does not understand the message type provided")
