package joyride

import "context"

type MessageHandler interface {
	Handle(ctx context.Context, messages ...interface{})
	HandleMessage(ctx context.Context, message interface{}) bool
	Run(ctx context.Context)
}

type (
	RequiredReads interface{ RequiredReads() []interface{} }
	Executable    interface {
		Execute(ctx context.Context) TaskResult
	}
)

type TaskResult interface {
	PendingWrites() []interface{}
	PendingMessages() []interface{}
	SubsequentTask() Executable
}

type (
	TaskRunner interface {
		Run(ctx context.Context, task Executable)
	}
	StorageReader interface {
		Read(ctx context.Context, reads ...interface{})
	}
	StorageWriter interface {
		Write(ctx context.Context, writes ...interface{})
	}
	MessageDispatcher interface {
		Dispatch(ctx context.Context, messages ...interface{})
	}
)
