package joyride

import (
	"testing"

	"github.com/smartystreets/assertions/should"
	"github.com/smartystreets/gunit"
)

func TestTaskFixture(t *testing.T) {
	gunit.Run(new(TaskFixture), t)
}

type TaskFixture struct {
	*gunit.Fixture

	task     *Task
	messages []interface{}
}

func (this *TaskFixture) Setup() {
	this.task = NewTask()
	this.messages = []interface{}{1, 2, 3}
}

func (this *TaskFixture) TestReadStateMaintained() {
	this.task.Read(this.messages...)
	this.So(this.task.Reads(), should.Resemble, this.messages)
}
func (this *TaskFixture) TestWriteStateMaintained() {
	this.task.Write(this.messages...)
	this.So(this.task.Writes(), should.Resemble, this.messages)
}
func (this *TaskFixture) TestMessageStateMaintained() {
	this.task.Dispatch(this.messages...)
	this.So(this.task.Messages(), should.Resemble, this.messages)
}
func (this *TaskFixture) TestNextStateMaintained() {
	next := NewTask()
	this.task.Chain(next)
	this.So(this.task.Next(), should.Equal, next)
}

func (this *TaskFixture) TestExecuteNoOperation() {
	this.task.Read(this.messages...)
	this.task.Run() // no op
	this.task.Write(this.messages...)
	this.task.Dispatch(this.messages...)

	this.So(this.task.Reads(), should.Resemble, this.messages)
	this.So(this.task.Writes(), should.Resemble, this.messages)
	this.So(this.task.Messages(), should.Resemble, this.messages)
}

func (this *TaskFixture) TestFunctionalOptions() {
	reads := []interface{}{1, 2, 3}
	writes := []interface{}{4, 5, 6}
	messages := []interface{}{7, 8, 9}

	task := NewTask(Read(reads...), Write(writes...), Dispatch(messages...))

	this.So(task.Reads(), should.Resemble, reads)
	this.So(task.Writes(), should.Resemble, writes)
	this.So(task.Messages(), should.Resemble, messages)
}
