package joyride

type BaseProcedure struct {
	Reads    []interface{}
	Writes   []interface{}
	Messages []interface{}
	Next     Procedure
}

func NewProcedure(reads ...interface{}) *BaseProcedure {
	return &BaseProcedure{Reads: reads}
}
func (this *BaseProcedure) Read() []interface{}     { return this.Reads }
func (this *BaseProcedure) Execute()                { /* noop */ }
func (this *BaseProcedure) Write() []interface{}    { return this.Writes }
func (this *BaseProcedure) Dispatch() []interface{} { return this.Messages }
func (this *BaseProcedure) Continue() Procedure     { return this.Next }

func (this *BaseProcedure) AppendReads(reads ...interface{}) *BaseProcedure {
	this.Reads = append(this.Reads, reads...)
	return this
}
func (this *BaseProcedure) AppendWrites(writes ...interface{}) *BaseProcedure {
	this.Reads = append(this.Writes, writes...)
	return this
}
func (this *BaseProcedure) AppendMessages(messages ...interface{}) *BaseProcedure {
	this.Messages = append(this.Messages, messages...)
	return this
}
func (this *BaseProcedure) AppendNext(next Procedure) *BaseProcedure {
	this.Next = next
	return this
}
