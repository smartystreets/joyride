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
	this.runner = NewRunner(this.task.Initialize, this.io, this.io, this.io)
}

func (this *RunnerFixture) TestSkipNilTasks() {
	this.runner = NewRunner(func(interface{}) ExecutableTask {
		return nil
	}, this.io, this.io, this.io)

	this.So(func() { this.runner.Run(0) }, should.NotPanic)
}

func (this *RunnerFixture) TestRunner() {
	const message = "Hello, World!"

	this.runner.Run(message)

	this.So(this.task.initialMessage, should.Equal, message)
	this.So(this.io.reads, should.Resemble, this.task.reads)
	this.So(this.io.writes, should.Resemble, this.task.writes)
	this.So(this.io.messages, should.Resemble, this.task.messages)
	this.So(this.task.Times(), should.BeChronological)
}

func (this *RunnerFixture) TestNextTask() {
	next := &FakeTask{}
	this.task.next = next

	this.runner.Run("message")

	this.So(next.executed, should.NotEqual, time.Time{})
	this.So(next.Times(), should.BeChronological)
}

/////////////////////////////////////////////////////////////

type FakeTask struct {
	initialMessage                                             interface{}
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

func (this *FakeTask) Initialize(message interface{}) ExecutableTask {
	this.initialMessage = message
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
func (this *FakeTask) Next() ExecutableTask {
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
