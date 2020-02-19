package joyride

type RunnableTask interface {
	Reads() []interface{}
	Execute()
	Writes() []interface{}
	Messages() []interface{}
	Next() RunnableTask
}

type MessageHandler interface {
	Handle(...interface{})
	HandleMessage(interface{}) bool
	Run()
}

type StorageReader interface {
	Read(...interface{})
}

type StorageWriter interface {
	Write(...interface{})
}

type MessageDispatcher interface {
	Dispatch(...interface{})
}

type TaskRunner interface {
	Run(RunnableTask)
}
