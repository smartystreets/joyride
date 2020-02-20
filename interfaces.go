package joyride

type MessageHandler interface {
	Handle(...interface{})
	HandleMessage(interface{}) bool
	Run()
}

type (
	RequiredReads interface{ RequiredReads() []interface{} }
	Executable    interface{ Execute() TaskResult }
)

type TaskResult interface {
	PendingWrites() []interface{}
	PendingMessages() []interface{}
	SubsequentTask() Executable
}

type (
	TaskRunner        interface{ Run(Executable) }
	StorageReader     interface{ Read(...interface{}) }
	StorageWriter     interface{ Write(...interface{}) }
	MessageDispatcher interface{ Dispatch(...interface{}) }
)
