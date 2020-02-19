package joyride

type (
	RunnableTask interface {
		Reads() []interface{}
		Execute()
		Writes() []interface{}
		Messages() []interface{}
		Next() RunnableTask
	}
	Task struct {
		reads    []interface{}
		writes   []interface{}
		messages []interface{}
		next     RunnableTask
	}
	TaskOption func(*Task)
)

func WithPreparedRead(items ...interface{}) TaskOption {
	return func(this *Task) { this.PrepareRead(items...) }
}
func WithPreparedWrite(items ...interface{}) TaskOption {
	return func(this *Task) { this.PrepareWrite(items...) }
}
func WithPreparedDispatch(items ...interface{}) TaskOption {
	return func(this *Task) { this.PrepareDispatch(items...) }
}

func NewTask(options ...TaskOption) *Task {
	this := &Task{}

	for _, option := range options {
		option(this)
	}

	return this
}

func (this *Task) Reads() []interface{}    { return this.reads }
func (this *Task) Execute()                { /* no-op; embed the task into another struct and override this method */ }
func (this *Task) Writes() []interface{}   { return this.writes }
func (this *Task) Messages() []interface{} { return this.messages }
func (this *Task) Next() RunnableTask      { return this.next }

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
