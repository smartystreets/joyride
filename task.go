package joyride

type (
	Task interface {
		Reads() []interface{}
		Run()
		Writes() []interface{}
		Messages() []interface{}
		Next() Task
	}
	DefaultTask struct {
		reads    []interface{}
		writes   []interface{}
		messages []interface{}
		next     Task
	}
	Option func(*DefaultTask)
)

func Read(items ...interface{}) Option     { return func(this *DefaultTask) { this.Read(items...) } }
func Write(items ...interface{}) Option    { return func(this *DefaultTask) { this.Write(items...) } }
func Dispatch(items ...interface{}) Option { return func(this *DefaultTask) { this.Dispatch(items...) } }

func NewTask(options ...Option) *DefaultTask {
	this := &DefaultTask{}

	for _, option := range options {
		option(this)
	}

	return this
}

func (this *DefaultTask) Reads() []interface{}    { return this.reads }
func (this *DefaultTask) Run()                    {}
func (this *DefaultTask) Writes() []interface{}   { return this.writes }
func (this *DefaultTask) Messages() []interface{} { return this.messages }
func (this *DefaultTask) Next() Task              { return this.next }

func (this *DefaultTask) Read(items ...interface{})  { this.reads = append(this.reads, items...) }
func (this *DefaultTask) Write(items ...interface{}) { this.writes = append(this.writes, items...) }
func (this *DefaultTask) Dispatch(items ...interface{}) {
	this.messages = append(this.messages, items...)
}
func (this *DefaultTask) Chain(next Task) { this.next = next }
