package joyride

import (
	"log"
	"reflect"
)

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
			log.Panicf("[WARN] Handler of type [%s] unable to handle message of type [%s].",
				reflect.TypeOf(this.inner), reflect.TypeOf(message))
		}
	}

	this.inner.Run()
}

func (this *Handler) Add(task RunnableTask) {
	this.tasks = append(this.tasks, task)
}

func (this *Handler) Tasks() []RunnableTask {
	return this.tasks
}

func (this *Handler) Run() {
	this.runner.Run(NewCompositeTask(this.tasks...))
}
