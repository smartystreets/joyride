package main

import "github.com/smartystreets/joyride/v2"

type ListTODOsTask struct {
	query   *LoadTODOsFromStorage
	context *ListTODOs
}

func NewListTODOsTask(context *ListTODOs) *ListTODOsTask {
	return &ListTODOsTask{
		query:   &LoadTODOsFromStorage{},
		context: context,
	}
}

func (this *ListTODOsTask) RequiredReads() []interface{} {
	return []interface{}{this.query}
}

func (this *ListTODOsTask) Execute() joyride.TaskResult {
	for _, record := range this.query.Results {
		this.context.Results = append(this.context.Results, TODO{
			Description: record.Description,
			Completed:   record.Completed,
		})
	}
	return joyride.TaskResult{}
}

//////////////////////////////////////////////////////////////////////////

type AddTODOTask struct {
	description string
}

func NewAddTODOTask(context AddTODO) *AddTODOTask {
	return &AddTODOTask{description: context.Description}
}

func (this *AddTODOTask) Execute() joyride.TaskResult {
	return joyride.TaskResult{
		PendingWrites: []interface{}{InsertTODOIntoStorage{Description: this.description}},
	}
}

//////////////////////////////////////////////////////////////////////////

type CompleteTODOTask struct {
	description string
}

func NewCompleteTODOTask(context CompleteTODO) *CompleteTODOTask {
	return &CompleteTODOTask{description: context.Description}
}

func (this *CompleteTODOTask) Execute() joyride.TaskResult {
	return joyride.TaskResult{
		PendingWrites: []interface{}{UpdateTODOInStorage{Description: this.description}},
	}
}
