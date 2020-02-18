package main

import "github.com/smartystreets/joyride"

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
	this.Write(InsertTODO{Description: this.description})
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
	this.Write(UpdateTODO{Description: this.description})
}
