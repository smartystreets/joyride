package joyride

type MultiTask struct {
	tasks []Procedure
}

func NewMultiTask(tasks ...Procedure) *MultiTask {
	var filtered []Procedure

	for _, task := range tasks {
		if task != nil {
			filtered = append(filtered, task)
		}
	}

	return &MultiTask{tasks: filtered}
}

func (this *MultiTask) Read() (reads []interface{}) {
	for _, task := range this.tasks {
		reads = append(reads, task.Read())
	}
	return reads
}

func (this *MultiTask) Execute() {
	for _, task := range this.tasks {
		task.Execute()
	}
}

func (this *MultiTask) Write() (writes []interface{}) {
	for _, task := range this.tasks {
		writes = append(writes, task.Write())
	}
	return writes
}

func (this *MultiTask) Dispatch() (messages []interface{}) {
	for _, task := range this.tasks {
		messages = append(messages, task.Dispatch())
	}
	return messages

}

func (this *MultiTask) Next() Procedure {
	var tasks []Procedure
	for _, task := range this.tasks {
		tasks = append(tasks, task.Next())
	}

	return NewMultiTask(tasks...)
}
