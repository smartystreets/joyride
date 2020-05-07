package main

import (
	"context"

	"github.com/smartystreets/joyride/v3"
)

type Handler struct {
	*joyride.Handler
}

func NewHandler(runner joyride.TaskRunner) *Handler {
	this := &Handler{}
	this.Handler = joyride.NewHandler(this, runner)
	return this
}

func (this *Handler) HandleMessage(ctx context.Context, message interface{}) bool {
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
