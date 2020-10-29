package joyride

import (
	"context"
	"reflect"
)

type CompositeTask []Executable

func (this CompositeTask) RequiredReads() (reads []interface{}) {
	for _, task := range this {
		reader, ok := task.(RequiredReads)
		if !ok || isNil(reader) {
			continue
		}
		reads = append(reads, reader.RequiredReads()...)
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
		if task == nil || isNil(task) {
			continue
		}

		inner := task.Execute(ctx)
		if inner == nil || isNil(inner) {
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

func isNil(v interface{}) bool {
	switch reflect.TypeOf(v).Kind() {
	case reflect.Chan, reflect.Func, reflect.Map, reflect.Ptr, reflect.UnsafePointer, reflect.Interface, reflect.Slice:
		return reflect.ValueOf(v).IsNil()
	default:
		return v == nil
	}
}
