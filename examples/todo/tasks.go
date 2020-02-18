package main

import "github.com/smartystreets/joyride/v2"

type ListTODOsTask struct {
	*joyride.Task
	storage *SelectTODOs
	context *ListTODOs
}

func NewListTODOsTask(context *ListTODOs) *ListTODOsTask {
	this := &ListTODOsTask{
		Task:    joyride.NewTask(),
		storage: &SelectTODOs{},
		context: context,
	}
	this.PrepareRead(this.storage)
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

//////////////////////////////////////////////////////////////////////////

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
	this.PrepareWrite(InsertTODO{Description: this.description})
}

//////////////////////////////////////////////////////////////////////////

type CompleteTODOTask struct {
	*joyride.Task
	description string
}

func NewCompleteTODOTask(context CompleteTODO) *CompleteTODOTask {
	return &CompleteTODOTask{
		Task:        joyride.NewTask(),
		description: context.Description,
	}
}

func (this *CompleteTODOTask) Run() {
	this.PrepareWrite(UpdateTODO{Description: this.description})
}
