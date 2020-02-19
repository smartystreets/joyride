package joyride

import "errors"

type MessageHandler interface {
	Handle(...interface{})
	HandleMessage(interface{}) bool
	Run()
}

// This type is designed to be embedded into another type that implements the MessageHandler interface such that the
// only method necessary for the embedding type to implement is the remaining method of HandleMessage. In essence, this
// structure becomes, in C# terminology an, "abstract class" which needs some additional behavior in order to function
// properly. In addition, because this struct is embedded, the existing methods of Handle (and more especially) the
// Run method can be overridden by the embedding struct.
type Handler struct {
	runner TaskRunner
	inner  MessageHandler
	tasks  []RunnableTask
}

func NewHandler(inner MessageHandler, runner TaskRunner, tasks ...RunnableTask) *Handler {
	return &Handler{inner: inner, runner: runner, tasks: tasks}
}

func (this *Handler) Handle(messages ...interface{}) {
	for _, message := range messages {
		if !this.inner.HandleMessage(message) {
			panic(ErrUnknownType)
		}
	}

	this.inner.Run()
}

func (this *Handler) Add(task RunnableTask) {
	this.tasks = append(this.tasks, task)
}

func (this *Handler) Run() {
	this.runner.Run(NewCompositeTask(this.tasks...))
}

func (this *Handler) Tasks() []RunnableTask { return this.tasks }

var ErrUnknownType = errors.New("the handler does not understand the message type provided")
