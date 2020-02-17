package todo

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
	//case AddTODO:
	//	this.Add(NewAddTODOTask(message.ID, message.Description))
	//case CompleteTODO:
	//	this.Add(NewCompleteTODOTask(message.ID))
	default:
		return false
	}
	return true
}

type ListTODOsTask struct {
	*joyride.Task
	storage *ReadTODOs
	context *ListTODOs
}


func NewListTODOsTask(context *ListTODOs) *ListTODOsTask {
	this := &ListTODOsTask{
		storage: &ReadTODOs{},
		context: context,
	}
	this.Read(this.storage)
	return this
}

func (this *ListTODOsTask) Run() {
	this.context.Results = this.storage.Results
}

type AddTODOTask struct {
	*joyride.Task
}

type CompleteTODOTask struct {
	*joyride.Task
}
