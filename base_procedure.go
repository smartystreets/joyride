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
func (this *BaseProcedure) Populate() []interface{} { return this.Reads }
func (this *BaseProcedure) Process()                { panic("not implemented") }
func (this *BaseProcedure) Persist() []interface{}  { return this.Writes }
func (this *BaseProcedure) Publish() []interface{}  { return this.Messages }
func (this *BaseProcedure) Procedure() Procedure    { return this.Next }
