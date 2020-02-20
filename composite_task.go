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

func (this CompositeTask) Execute() (result TaskResult) {
	if len(this) == 0 {
		return result
	}

	var executables CompositeTask

	for _, task := range this {
		if task == nil {
			continue
		}
		inner := task.Execute()
		result.PendingWrites = append(result.PendingWrites, inner.PendingWrites...)
		result.PendingMessages = append(result.PendingMessages, inner.PendingMessages...)
		executables = append(executables, inner.SubsequentTask)
	}
	result.SubsequentTask = executables
	return result
}
