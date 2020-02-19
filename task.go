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

func (this *Task) PrepareRead(v ...interface{})      { this.reads = append(this.reads, v...) }
func (this *Task) PrepareWrite(v ...interface{})     { this.writes = append(this.writes, v...) }
func (this *Task) PrepareDispatch(v ...interface{})  { this.messages = append(this.messages, v...) }
func (this *Task) PrepareNextTask(next RunnableTask) { this.next = next }

func (this *Task) Reads() []interface{}    { return this.reads }
func (this *Task) Execute()                { /* no-op; embed *joyride.Task and override */ }
func (this *Task) Writes() []interface{}   { return this.writes }
func (this *Task) Messages() []interface{} { return this.messages }
func (this *Task) Next() RunnableTask      { return this.next }
