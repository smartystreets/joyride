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
func (this *BaseProcedure) Execute()                { panic("not implemented") }
func (this *BaseProcedure) Write() []interface{}    { return this.Writes }
func (this *BaseProcedure) Dispatch() []interface{} { return this.Messages }
func (this *BaseProcedure) Continue() Procedure     { return this.Next }
