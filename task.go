package joyride

type Task struct {
	Reads    []interface{}
	Writes   []interface{}
	Messages []interface{}
	NextTask Procedure
}

type TaskOption func(*Task)

func WithReads(reads ...interface{}) TaskOption {
	return func(this *Task) { this.Reads = append(this.Reads, reads...) }
}
func WithWrites(writes ...interface{}) TaskOption {
	return func(this *Task) { this.Writes = append(this.Writes, writes...) }
}
func WithMessages(messages ...interface{}) TaskOption {
	return func(this *Task) { this.Messages = append(this.Messages, messages...) }
}

func New(options ...TaskOption) *Task {
	procedure := &Task{}

	for _, option := range options {
		option(procedure)
	}

	return procedure
}

func (this *Task) Read() []interface{}     { return this.Reads }
func (this *Task) Execute()                { /* noop */ }
func (this *Task) Write() []interface{}    { return this.Writes }
func (this *Task) Dispatch() []interface{} { return this.Messages }
func (this *Task) Next() Procedure         { return this.NextTask }

func (this *Task) AppendReads(reads ...interface{}) {
	this.Reads = append(this.Reads, reads...)
}
func (this *Task) AppendWrites(writes ...interface{}) {
	this.Writes = append(this.Writes, writes...)
}
func (this *Task) AppendMessages(messages ...interface{}) {
	this.Messages = append(this.Messages, messages...)
}
