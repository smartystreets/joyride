package joyride

type CompositeProcedure struct {
	procedures []Procedure
}

func NewComposite(procedures ...Procedure) *CompositeProcedure {
	var filtered []Procedure

	for _, procedure := range procedures {
		if procedure != nil {
			filtered = append(filtered, procedure)
		}
	}

	return &CompositeProcedure{procedures: filtered}
}

func (this *CompositeProcedure) Read() (reads []interface{}) {
	for _, procedure := range this.procedures {
		reads = append(reads, procedure.Read())
	}
	return reads
}

func (this *CompositeProcedure) Execute() {
	for _, procedure := range this.procedures {
		procedure.Execute()
	}
}

func (this *CompositeProcedure) Write() (writes []interface{}) {
	for _, procedure := range this.procedures {
		writes = append(writes, procedure.Write())
	}
	return writes
}

func (this *CompositeProcedure) Dispatch() (messages []interface{}) {
	for _, procedure := range this.procedures {
		messages = append(messages, procedure.Dispatch())
	}
	return messages

}

func (this *CompositeProcedure) Next() Procedure {
	var procedures []Procedure
	for _, procedure := range this.procedures {
		procedures = append(procedures, procedure.Next())
	}

	return NewComposite(procedures...)
}
