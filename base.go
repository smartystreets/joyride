package joyride

import "context"

type Base struct {
	reads    []interface{}
	writes   []interface{}
	messages []interface{}
	next     Executable
}

func New(reads ...interface{}) *Base { return &Base{reads: reads} }

func (this *Base) RequiredReads() []interface{}         { return this.reads }
func (this *Base) Execute(_ context.Context) TaskResult { return this }
func (this *Base) PendingWrites() []interface{}         { return this.writes }
func (this *Base) PendingMessages() []interface{}       { return this.messages }
func (this *Base) SubsequentTask() Executable           { return this.next }

func (this *Base) AddRequiredReads(r ...interface{})   { this.reads = append(this.reads, r...) }
func (this *Base) AddPendingWrites(w ...interface{})   { this.writes = append(this.writes, w...) }
func (this *Base) AddPendingMessages(m ...interface{}) { this.messages = append(this.messages, m...) }
func (this *Base) SetSubsequentTask(next Executable)   { this.next = next }
