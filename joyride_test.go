package joyride

import (
	"testing"
	"time"

	"github.com/smartystreets/assertions/should"
	"github.com/smartystreets/gunit"
)

func TestJoyrideFixture(t *testing.T) {
	gunit.Run(new(JoyrideFixture), t)
}

type JoyrideFixture struct {
	*gunit.Fixture

	task    *TracingTask
	io      *FakeExternalIO
	handler *ExampleHandler
	runner  Runner
}

func (this *JoyrideFixture) Setup() {
	this.io = &FakeExternalIO{}
	this.runner = NewRunner(WithStorageReader(this.io), WithStorageWriter(this.io), WithMessageDispatcher(this.io))
	this.task = NewTracingTask()
	this.handler = NewExampleHandler(this.runner, this.task)
}

func (this *JoyrideFixture) TestMessageHandled_TaskExecuted() {
	this.handler.Handle(42)

	this.So(this.handler.handled, should.Resemble, []interface{}{42})
	this.So(this.io.reads, should.Resemble, this.task.reads)
	this.So(this.io.writes, should.Resemble, this.task.writes)
	this.So(this.io.messages, should.Resemble, this.task.messages)
	this.So(this.task.Times(), should.BeChronological)
}
func (this *JoyrideFixture) TestCannotHandleMessage_Panic() {
	this.handler.canHandle = false

	this.So(func() { this.handler.Handle(42) }, should.PanicWith, ErrUnknownType)

	this.So(this.handler.handled, should.Resemble, []interface{}{42})
	this.So(this.io.reads, should.BeEmpty)
	this.So(this.io.writes, should.BeEmpty)
	this.So(this.io.messages, should.BeEmpty)
}
func (this *JoyrideFixture) TestChainedTasksAreExecuted() {
	next := NewTracingTask()
	this.task.next = next

	this.handler.Handle(42)

	this.So(next.executed.IsZero(), should.BeFalse)
	this.So(next.Times(), should.BeChronological)
}
func (this *JoyrideFixture) TestAddedTasksAreExecuted() {
	next := NewTracingTask()
	this.handler.Add(next)

	this.handler.Handle(42)

	this.So(next.executed.IsZero(), should.BeFalse)
	this.So(next.Times(), should.BeChronological)
	this.So(this.handler.Tasks(), should.Resemble, []Executable{this.task, next})
}

///////////////////////////////////////////////////////////////

type ExampleHandler struct {
	*Handler

	handled   []interface{}
	canHandle bool
}

func NewExampleHandler(runner TaskRunner, task Executable) *ExampleHandler {
	this := &ExampleHandler{canHandle: true}
	this.Handler = NewHandler(this, runner, task)
	return this
}

func (this *ExampleHandler) HandleMessage(message interface{}) bool {
	this.handled = append(this.handled, message)
	// Note: Normally, data from the message would now be provided
	// to the tasks already registered with the handler.
	return this.canHandle
}

//////////////////////////////////////////////////////////////

type TracingTask struct {
	*Result
	initialized time.Time
	read        time.Time
	executed    time.Time
}

func NewTracingTask() *TracingTask {
	result := NewResult()
	result.AddReads(1, 2, 3)
	result.AddWrites("4", "5", 6.0)
	result.AddMessages(7, "eight", 9, true)
	return &TracingTask{
		Result:      result,
		initialized: time.Now().UTC(),
	}
}

func (this *TracingTask) Times() []time.Time {
	return []time.Time{
		this.initialized,
		this.read,
		this.executed,
	}
}

func (this *TracingTask) RequiredReads() []interface{} {
	this.read = time.Now().UTC()
	return this.Result.RequiredReads()
}
func (this *TracingTask) Execute() TaskResult {
	this.executed = time.Now().UTC()
	return this.Result.Execute()
}

/////////////////////////////////////////////////////////////

type FakeExternalIO struct {
	reads    []interface{}
	writes   []interface{}
	messages []interface{}
}

func (this *FakeExternalIO) Read(items ...interface{}) {
	this.reads = append(this.reads, items...)
}
func (this *FakeExternalIO) Write(items ...interface{}) {
	this.writes = append(this.writes, items...)
}
func (this *FakeExternalIO) Dispatch(items ...interface{}) {
	this.messages = append(this.messages, items...)
}
