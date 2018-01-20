package joyride

type MultiTask struct {
	tasks []ExecutableTask
}

func NewMultiTask(tasks ...ExecutableTask) *MultiTask {
	var filtered []ExecutableTask

	for _, task := range tasks {
		if task != nil {
			filtered = append(filtered, task)
		}
	}

	return &MultiTask{tasks: filtered}
}

func (this *MultiTask) Reads() (reads []interface{}) {
	for _, task := range this.tasks {
		reads = append(reads, task.Reads()...)
	}
	return reads
}

func (this *MultiTask) Execute() {
	for _, task := range this.tasks {
		task.Execute()
	}
}

func (this *MultiTask) Writes() (writes []interface{}) {
	for _, task := range this.tasks {
		writes = append(writes, task.Writes()...)
	}
	return writes
}

func (this *MultiTask) Messages() (messages []interface{}) {
	for _, task := range this.tasks {
		messages = append(messages, task.Messages()...)
	}
	return messages

}

func (this *MultiTask) Next() ExecutableTask {
	var tasks []ExecutableTask
	for _, task := range this.tasks {
		tasks = append(tasks, task.Next())
	}

	return NewMultiTask(tasks...)
}
