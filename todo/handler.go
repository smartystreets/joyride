package main

import (
	"github.com/smartystreets/joyride"
)

type Handler struct {
	*joyride.Handler
}

func NewHandler(runner joyride.TaskRunner) *Handler {
	this := &Handler{}
	this.Handler = joyride.NewHandler(this, runner)
	return this
}

func (this *Handler) HandleMessage(message interface{}) bool {
	switch message := message.(type) {
	case *ListTODOs:
		this.Add(NewListTODOsTask(message))
	case AddTODO:
		this.Add(NewAddTODOTask(message))
	case CompleteTODO:
		this.Add(NewCompleteTODOTask(message))
	default:
		return false
	}
	return true
}
