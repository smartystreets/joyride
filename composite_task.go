package joyride

type CompositeTask struct {
	tasks []RunnableTask
}

func NewCompositeTask(tasks ...RunnableTask) RunnableTask {
	var filtered []RunnableTask

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

func (this CompositeTask) Next() RunnableTask {
	if len(this.tasks) == 0 {
		return nil
	}

	tasks := make([]RunnableTask, 0, len(this.tasks))
	for _, task := range this.tasks {
		tasks = append(tasks, task.Next())
	}

	return NewCompositeTask(tasks...)
}
