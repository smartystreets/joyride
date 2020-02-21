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

	result := New()
	var executables CompositeTask

	for _, task := range this {
		if task == nil {
			continue
		}
		inner := task.Execute()
		if inner == nil {
			continue
		}
		result.AddPendingWrites(inner.PendingWrites()...)
		result.AddPendingMessages(inner.PendingMessages()...)
		executables = append(executables, inner.SubsequentTask())
	}
	result.SetSubsequentTask(executables)

	return result
}
