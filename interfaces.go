package joyride

type MessageHandler interface {
	Handle(...interface{})
	HandleMessage(interface{}) bool
	Run()
}

type RunnableTask interface {
	Reads() []interface{}
	Execute()
	Writes() []interface{}
	Messages() []interface{}
	Next() RunnableTask
}

type (
	TaskRunner        interface{ Run(RunnableTask) }
	StorageReader     interface{ Read(...interface{}) }
	StorageWriter     interface{ Write(...interface{}) }
	MessageDispatcher interface{ Dispatch(...interface{}) }
)
