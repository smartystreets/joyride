package main

import "github.com/smartystreets/joyride/v2"

type ListTODOsTask struct {
	*joyride.Task
	query   *SelectTODOs
	context *ListTODOs
}

func NewListTODOsTask(context *ListTODOs) *ListTODOsTask {
	return &ListTODOsTask{
		Task:    joyride.NewTask(),
		query:   &SelectTODOs{},
		context: context,
	}
}

func (this *ListTODOsTask) Read() (queries []interface{}) {
	return append(queries, this.query)
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

type AddTODOTask struct{ *joyride.Task }

func NewAddTODOTask(context AddTODO) *AddTODOTask {
	insert := InsertTODO{Description: context.Description}
	task := joyride.NewTask(joyride.WithPreparedWrite(insert))
	return &AddTODOTask{Task: task}
}

//////////////////////////////////////////////////////////////////////////

type CompleteTODOTask struct{ *joyride.Task }

func NewCompleteTODOTask(context CompleteTODO) *CompleteTODOTask {
	insert := UpdateTODO{Description: context.Description}
	task := joyride.NewTask(joyride.WithPreparedWrite(insert))
	return &CompleteTODOTask{Task: task}
}
