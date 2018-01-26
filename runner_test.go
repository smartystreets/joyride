package joyride

import (
	"testing"
	"time"

	"github.com/smartystreets/assertions/should"
	"github.com/smartystreets/clock"
	"github.com/smartystreets/gunit"
)

func TestRunnerFixture(t *testing.T) {
	gunit.Run(new(RunnerFixture), t)
}

type RunnerFixture struct {
	*gunit.Fixture
	task   *FakeTask
	io     *ExternalIO
	runner Runner
}

func (this *RunnerFixture) Setup() {
	this.task = NewFakeTask()
	this.io = &ExternalIO{}
	this.runner = NewRunner(this.io, this.io, this.io)
}
func (this *RunnerFixture) Build(messages ...interface{}) RunnableTask {
	if this.task != nil {
		this.task.Initialize(messages...)
	}

	if this.task == nil {
		return nil
	}

	return this.task
}

func (this *RunnerFixture) TestSkipNilTasks() {
	this.task = nil
	this.So(func() { this.runner.Run(nil) }, should.NotPanic)
}

func (this *RunnerFixture) TestHandler() {
	const message1 = "Hello, World!"
	const message2 = 42
	this.task.Initialize(message1, message2)

	this.runner.Run(this.task)

	this.So(this.task.initializedMessages, should.Resemble, []interface{}{message1, message2})
	this.So(this.io.reads, should.Resemble, this.task.reads)
	this.So(this.io.writes, should.Resemble, this.task.writes)
	this.So(this.io.messages, should.Resemble, this.task.messages)
	this.So(this.task.Times(), should.BeChronological)
}

func (this *RunnerFixture) TestNextTask() {
	next := &FakeTask{}
	this.task.next = next

	this.runner.Run(this.task)

	this.So(next.executed, should.NotEqual, time.Time{})
	this.So(next.Times(), should.BeChronological)
}

/////////////////////////////////////////////////////////////

type FakeTask struct {
	initializedMessages                                        []interface{}
	reads, writes, messages                                    []interface{}
	initialized, read, executed, written, dispatched, nextTime time.Time
	next                                                       *FakeTask
}

func NewFakeTask() *FakeTask {
	return &FakeTask{
		reads:    []interface{}{1, 2, 3},
		writes:   []interface{}{"4", "5", 6.0},
		messages: []interface{}{7, "eight", 9, true},
	}
}

func (this *FakeTask) Times() []time.Time {
	return []time.Time{this.initialized, this.read, this.executed, this.written, this.dispatched, this.nextTime}
}

func (this *FakeTask) Initialize(messages ...interface{}) RunnableTask {
	this.initializedMessages = messages
	this.initialized = clock.UTCNow()
	return this
}
func (this *FakeTask) Reads() []interface{} {
	this.read = clock.UTCNow()
	return this.reads
}
func (this *FakeTask) Execute() {
	this.executed = clock.UTCNow()
}
func (this *FakeTask) Writes() []interface{} {
	this.written = clock.UTCNow()
	return this.writes
}
func (this *FakeTask) Messages() []interface{} {
	this.dispatched = clock.UTCNow()
	return this.messages
}
func (this *FakeTask) Next() RunnableTask {
	this.nextTime = clock.UTCNow()
	if this.next == nil {
		return nil // Go nil conversion quirks
	} else {
		return this.next
	}
}

/////////////////////////////////////////////////////////////

type ExternalIO struct{ reads, writes, messages []interface{} }

func (this *ExternalIO) Read(items ...interface{})     { this.reads = items }
func (this *ExternalIO) Write(items ...interface{})    { this.writes = items }
func (this *ExternalIO) Dispatch(items ...interface{}) { this.messages = items }
