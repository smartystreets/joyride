package joyride

type CompositeTask []Executable

func (this CompositeTask) RequiredReads() (reads []interface{}) {
	for _, task := range this {
		if reader, ok := task.(RequiredReads); ok {
			reads = append(reads, reader.RequiredReads()...)
		}
	}
	return reads
}

func (this CompositeTask) Execute() TaskResult {
	if len(this) == 0 {
		return nil
	}

	result := NewResult()
	var executables CompositeTask

	for _, task := range this {
		if task == nil {
			continue
		}
		inner := task.Execute()
		if inner == nil {
			continue
		}
		result.AddWrites(inner.PendingWrites()...)
		result.AddMessages(inner.PendingMessages()...)
		executables = append(executables, inner.SubsequentTask())
	}
	result.SetSubsequentTask(executables)

	return result
}
