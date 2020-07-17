package joyride

import "context"

type CompositeTask []Executable

func (this CompositeTask) RequiredReads() (reads []interface{}) {
	for _, task := range this {
		if reader, ok := task.(RequiredReads); ok {
			reads = append(reads, reader.RequiredReads()...)
		}
	}
	return reads
}

func (this CompositeTask) Execute(ctx context.Context) TaskResult {
	if len(this) == 0 {
		return nil
	}

	result := New()
	var executables CompositeTask

	for _, task := range this {
		if task == nil {
			continue
		}
		inner := task.Execute(ctx)
		if inner == nil {
			continue
		}
		result.AddPendingWrites(inner.PendingWrites()...)
		result.AddPendingMessages(inner.PendingMessages()...)
		subsequentTask := inner.SubsequentTask()
		if subsequentTask != nil {
			executables = append(executables, subsequentTask)
		}
	}
	if len(executables) > 0 {
		result.SetSubsequentTask(executables)
	}

	return result
}
