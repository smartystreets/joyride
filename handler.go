package joyride

import "errors"

type MessageHandler interface {
	Handle(...interface{})
	HandleMessage(interface{}) bool
	Run()
}

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
