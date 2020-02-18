package main

import "github.com/smartystreets/joyride"

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
	//case CompleteTODO:
	//	this.Add(NewCompleteTODOTask(message.ID))
	default:
		return false
	}
	return true
}

type ListTODOsTask struct {
	*joyride.Task
	storage *SelectTODOs
	context *ListTODOs
}


func NewListTODOsTask(context *ListTODOs) *ListTODOsTask {
	this := &ListTODOsTask{
		Task: joyride.NewTask(),
		storage: &SelectTODOs{},
		context: context,
	}
	this.Read(this.storage)
	return this
}

func (this *ListTODOsTask) Run() {
	for _, record := range this.storage.Results {
		this.context.Results = append(this.context.Results, TODO{
			Description: record.Description,
			Completed:   record.Completed,
		})
	}
}

type AddTODOTask struct {
	*joyride.Task
	description string
}

func NewAddTODOTask(context AddTODO) *AddTODOTask {
	return &AddTODOTask{
		Task:        joyride.NewTask(),
		description: context.Description,
	}
}

func (this *AddTODOTask) Run() {
	this.Write(InsertTODO{Description:this.description})
}

type CompleteTODOTask struct {
	*joyride.Task
}
