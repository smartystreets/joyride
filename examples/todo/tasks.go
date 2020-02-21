package main

import "github.com/smartystreets/joyride/v2"

type ListTODOsTask struct {
	*joyride.Base
	query   *LoadTODOsFromStorage
	context *ListTODOs
}

func NewListTODOsTask(context *ListTODOs) *ListTODOsTask {
	storage := &LoadTODOsFromStorage{}
	return &ListTODOsTask{
		Base:    joyride.New(storage),
		query:   storage,
		context: context,
	}
}

func (this *ListTODOsTask) Execute() joyride.TaskResult {
	for _, record := range this.query.Results {
		this.context.Results = append(this.context.Results, TODO{
			Description: record.Description,
			Completed:   record.Completed,
		})
	}
	return this
}

//////////////////////////////////////////////////////////////////////////

type AddTODOTask struct {
	*joyride.Base
}

func NewAddTODOTask(context AddTODO) *AddTODOTask {
	operation := InsertTODOIntoStorage{Description: context.Description}
	base := joyride.New()
	base.AddPendingWrites(operation)
	return &AddTODOTask{Base: base}
}

//////////////////////////////////////////////////////////////////////////

type CompleteTODOTask struct {
	*joyride.Base
	description string
}

func NewCompleteTODOTask(context CompleteTODO) *CompleteTODOTask {
	result := joyride.New()
	result.AddPendingWrites(UpdateTODOInStorage{Description: context.Description})
	return &CompleteTODOTask{Base: result}
}
