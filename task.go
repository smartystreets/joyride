package joyride

type (
	RunnableTask interface {
		Reads() []interface{}
		Run()
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
	Option func(*Task)
)

func Read(items ...interface{}) Option     { return func(this *Task) { this.Read(items...) } }
func Write(items ...interface{}) Option    { return func(this *Task) { this.Write(items...) } }
func Dispatch(items ...interface{}) Option { return func(this *Task) { this.Dispatch(items...) } }

func NewTask(options ...Option) *Task {
	this := &Task{}

	for _, option := range options {
		option(this)
	}

	return this
}

func (this *Task) Reads() []interface{}    { return this.reads }
func (this *Task) Run()                    { /* no-op; embed the task into another struct and override */ }
func (this *Task) Writes() []interface{}   { return this.writes }
func (this *Task) Messages() []interface{} { return this.messages }
func (this *Task) Next() RunnableTask      { return this.next }

func (this *Task) Read(items ...interface{})     { this.reads = append(this.reads, items...) }
func (this *Task) Write(items ...interface{})    { this.writes = append(this.writes, items...) }
func (this *Task) Dispatch(items ...interface{}) { this.messages = append(this.messages, items...) }
func (this *Task) Chain(next RunnableTask)       { this.next = next }
