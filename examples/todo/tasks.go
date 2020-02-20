package main

import "github.com/smartystreets/joyride/v2"

type ListTODOsTask struct {
	*joyride.Result
	query   *LoadTODOsFromStorage
	context *ListTODOs
}

func NewListTODOsTask(context *ListTODOs) *ListTODOsTask {
	storage := &LoadTODOsFromStorage{}
	return &ListTODOsTask{
		Result:  joyride.NewResult(storage),
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
	*joyride.Result
}

func NewAddTODOTask(context AddTODO) *AddTODOTask {
	result := joyride.NewResult()
	result.AddPendingWrites(InsertTODOIntoStorage{Description: context.Description})
	return &AddTODOTask{Result: result}
}

//////////////////////////////////////////////////////////////////////////

type CompleteTODOTask struct {
	*joyride.Result
	description string
}

func NewCompleteTODOTask(context CompleteTODO) *CompleteTODOTask {
	result := joyride.NewResult()
	result.AddPendingWrites(UpdateTODOInStorage{Description: context.Description})
	return &CompleteTODOTask{Result: result}
}
