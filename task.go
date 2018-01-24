package joyride

type (
	ExecutableTask interface {
		Reads() []interface{}
		Execute()
		Writes() []interface{}
		Messages() []interface{}
		Next() ExecutableTask
	}
	Task struct {
		reads    []interface{}
		writes   []interface{}
		messages []interface{}
		next     ExecutableTask
	}
	Option func(*Task)
)

func Read(items ...interface{}) Option     { return func(this *Task) { this.Read(items...) } }
func Write(items ...interface{}) Option    { return func(this *Task) { this.Write(items...) } }
func Dispatch(items ...interface{}) Option { return func(this *Task) { this.Dispatch(items...) } }

func New(options ...Option) *Task {
	this := &Task{}

	for _, option := range options {
		option(this)
	}

	return this
}

func (this *Task) Reads() []interface{}    { return this.reads }
func (this *Task) Execute()                {}
func (this *Task) Writes() []interface{}   { return this.writes }
func (this *Task) Messages() []interface{} { return this.messages }
func (this *Task) Next() ExecutableTask    { return this.next }

func (this *Task) Read(items ...interface{})     { this.reads = append(this.reads, items...) }
func (this *Task) Write(items ...interface{})    { this.writes = append(this.writes, items...) }
func (this *Task) Dispatch(items ...interface{}) { this.messages = append(this.messages, items...) }
func (this *Task) Chain(next ExecutableTask)     { this.next = next }
