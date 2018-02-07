package joyride

type CompositeTask struct {
	tasks []Task
}

func NewCompositeTask(tasks ...Task) CompositeTask {
	var filtered []Task

	for _, task := range tasks {
		if task != nil {
			filtered = append(filtered, task)
		}
	}

	return CompositeTask{tasks: filtered}
}

func (this CompositeTask) Reads() (reads []interface{}) {
	for _, task := range this.tasks {
		reads = append(reads, task.Reads()...)
	}
	return reads
}

func (this CompositeTask) Run() {
	for _, task := range this.tasks {
		task.Run()
	}
}

func (this CompositeTask) Writes() (writes []interface{}) {
	for _, task := range this.tasks {
		writes = append(writes, task.Writes()...)
	}
	return writes
}

func (this CompositeTask) Messages() (messages []interface{}) {
	for _, task := range this.tasks {
		messages = append(messages, task.Messages()...)
	}
	return messages

}

func (this CompositeTask) Next() Task {
	var tasks []Task
	for _, task := range this.tasks {
		tasks = append(tasks, task.Next())
	}

	return NewCompositeTask(tasks...)
}
