package joyride

type Result struct {
	reads    []interface{}
	writes   []interface{}
	messages []interface{}
	next     Executable
}

func NewResult(reads ...interface{}) *Result { return &Result{reads: reads} }

func (this *Result) RequiredReads() []interface{}   { return this.reads }
func (this *Result) Execute() TaskResult            { return this }
func (this *Result) PendingWrites() []interface{}   { return this.writes }
func (this *Result) PendingMessages() []interface{} { return this.messages }
func (this *Result) SubsequentTask() Executable     { return this.next }

func (this *Result) AddRequiredReads(reads ...interface{}) { this.reads = append(this.reads, reads...) }
func (this *Result) AddPendingWrites(w ...interface{})     { this.writes = append(this.writes, w...) }
func (this *Result) AddPendingMessages(m ...interface{})   { this.messages = append(this.messages, m...) }
func (this *Result) SetSubsequentTask(next Executable)     { this.next = next }
