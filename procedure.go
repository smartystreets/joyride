package joyride

type BaseProcedure struct {
	Reads    []interface{}
	Writes   []interface{}
	Messages []interface{}
	Next     Procedure
}

type Option func(*BaseProcedure)

func WithReads(reads ...interface{}) Option {
	return func(this *BaseProcedure) { this.Reads = append(this.Reads, reads...) }
}
func WithWrites(writes ...interface{}) Option {
	return func(this *BaseProcedure) { this.Writes = append(this.Writes, writes...) }
}
func WithMessages(messages ...interface{}) Option {
	return func(this *BaseProcedure) { this.Messages = append(this.Messages, messages...) }
}

func NewProcedure(options ...Option) *BaseProcedure {
	procedure := &BaseProcedure{}

	for _, option := range options {
		option(procedure)
	}

	return procedure
}

func (this *BaseProcedure) Read() []interface{}     { return this.Reads }
func (this *BaseProcedure) Execute()                { /* noop */ }
func (this *BaseProcedure) Write() []interface{}    { return this.Writes }
func (this *BaseProcedure) Dispatch() []interface{} { return this.Messages }
func (this *BaseProcedure) Chain() Procedure        { return this.Next }
