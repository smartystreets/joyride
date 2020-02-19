package main

import "github.com/smartystreets/joyride/v2"

type ListTODOsTask struct {
	*joyride.Task
	query   *SelectTODOs
	context *ListTODOs
}

func NewListTODOsTask(context *ListTODOs) *ListTODOsTask {
	this := &ListTODOsTask{
		Task:    joyride.NewTask(),
		query:   &SelectTODOs{},
		context: context,
	}
	this.PrepareRead(this.query)
	return this
}

func (this *ListTODOsTask) Execute() {
	for _, record := range this.query.Results {
		this.context.Results = append(this.context.Results, TODO{
			Description: record.Description,
			Completed:   record.Completed,
		})
	}
}

//////////////////////////////////////////////////////////////////////////

type AddTODOTask struct {
	*joyride.Task
}

func NewAddTODOTask(context AddTODO) *AddTODOTask {
	return &AddTODOTask{
		Task: joyride.NewTask(joyride.WithPreparedWrite(InsertTODO{Description: context.Description})),
	}
}

//////////////////////////////////////////////////////////////////////////

type CompleteTODOTask struct {
	*joyride.Task
}

func NewCompleteTODOTask(context CompleteTODO) *CompleteTODOTask {
	return &CompleteTODOTask{
		Task: joyride.NewTask(joyride.WithPreparedWrite(UpdateTODO{Description: context.Description})),
	}
}
