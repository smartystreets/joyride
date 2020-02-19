package joyride

type Task struct {
	reads    []interface{}
	writes   []interface{}
	messages []interface{}
	next     RunnableTask
}

func NewTask(options ...TaskOption) *Task {
	this := &Task{}

	for _, option := range options {
		option(this)
	}

	return this
}

func (this *Task) PrepareRead(items ...interface{}) {
	this.reads = append(this.reads, items...)
}
func (this *Task) PrepareWrite(items ...interface{}) {
	this.writes = append(this.writes, items...)
}
func (this *Task) PrepareDispatch(items ...interface{}) {
	this.messages = append(this.messages, items...)
}
func (this *Task) PrepareNextTask(next RunnableTask) {
	this.next = next
}

func (this *Task) Reads() []interface{} {
	return this.reads
}
func (this *Task) Execute() {
	/* no-op; embed *joyride.Task and override */
}
func (this *Task) Writes() []interface{} {
	return this.writes
}
func (this *Task) Messages() []interface{} {
	return this.messages
}
func (this *Task) Next() RunnableTask {
	return this.next
}
