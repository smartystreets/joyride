package joyride

import (
	"context"
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

	ctx     context.Context
	task    *TracingTask
	io      *FakeExternalIO
	handler *ExampleHandler
	runner  Runner
}

func (this *JoyrideFixture) Setup() {
	this.io = &FakeExternalIO{}
	this.ctx = context.Background()
	this.runner = NewRunner(WithStorageReader(this.io), WithStorageWriter(this.io), WithMessageDispatcher(this.io))
	this.task = NewTracingTask()
	this.handler = NewExampleHandler(this.runner, this.task)
}

func (this *JoyrideFixture) TestRunnerDropsNilExecutablesWithoutPanicking() {
	this.So(func() { this.runner.Run(this.ctx, nil) }, should.NotPanic)
}

func (this *JoyrideFixture) TestCompositeTaskDropsNilResultFromExecutableWithoutPanicking() {
	this.handler = NewExampleHandler(this.runner, NewNilResultTask())
	this.So(func() { this.handler.Handle(this.ctx, 42) }, should.NotPanic)
}

func (this *JoyrideFixture) TestHandlerClearsTasksAfterExecuted() {
	this.handler.Handle(this.ctx, 42)
	this.handler.Handle(this.ctx, 42)

	this.So(this.task.executed, should.HaveLength, 1)
}

func (this *JoyrideFixture) TestMessageHandled_TaskExecuted() {
	this.handler.Handle(this.ctx, 42)

	this.So(this.handler.handled, should.Resemble, []interface{}{42})
	this.So(this.io.reads, should.Resemble, this.task.reads)
	this.So(this.io.writes, should.Resemble, this.task.writes)
	this.So(this.io.messages, should.Resemble, this.task.messages)
	this.So(this.task.Times(), should.BeChronological)
	this.So(this.handler.handleContext, should.Equal, this.ctx)
	this.So(this.task.executeContext, should.Equal, this.ctx)
	this.So(this.io.readContext, should.Equal, this.ctx)
	this.So(this.io.writeContext, should.Equal, this.ctx)
	this.So(this.io.dispatchContext, should.Equal, this.ctx)
}
func (this *JoyrideFixture) TestCannotHandleMessage_Panic() {
	this.handler.canHandle = false

	this.So(func() { this.handler.Handle(this.ctx, 42) }, should.Panic)

	this.So(this.handler.handled, should.Resemble, []interface{}{42})
	this.So(this.io.reads, should.BeEmpty)
	this.So(this.io.writes, should.BeEmpty)
	this.So(this.io.messages, should.BeEmpty)
}
func (this *JoyrideFixture) TestChainedTasksAreExecuted() {
	next := NewTracingTask()
	this.task.next = next

	this.handler.Handle(this.ctx, 42)

	if this.So(next.executed, should.HaveLength, 1) {
		this.So(next.executed[0].IsZero(), should.BeFalse)
	}
	this.So(next.Times(), should.BeChronological)
}
func (this *JoyrideFixture) TestAddedTasksAreExecuted() {
	next := NewTracingTask()
	this.handler.Add(next)

	this.handler.Handle(this.ctx, 42)

	if this.So(next.executed, should.HaveLength, 1) {
		this.So(next.executed[0].IsZero(), should.BeFalse)
	}
	this.So(next.Times(), should.BeChronological)
}

///////////////////////////////////////////////////////////////

type ExampleHandler struct {
	*Handler

	handleContext context.Context
	handled       []interface{}
	canHandle     bool
}

func NewExampleHandler(runner TaskRunner, task Executable) *ExampleHandler {
	this := &ExampleHandler{canHandle: true}
	this.Handler = NewHandler(this, runner, task)
	return this
}

func (this *ExampleHandler) HandleMessage(ctx context.Context, message interface{}) bool {
	this.handleContext = ctx
	this.handled = append(this.handled, message)
	// Note: Normally, data from the message would now be provided
	// to the tasks already registered with the handler.
	return this.canHandle
}

//////////////////////////////////////////////////////////////

type TracingTask struct {
	*Base
	initialized    time.Time
	read           time.Time
	executed       []time.Time
	executeContext context.Context
}

func NewTracingTask() *TracingTask {
	base := New()
	base.AddRequiredReads(1, 2, 3)
	base.AddPendingWrites("4", "5", 6.0)
	base.AddPendingMessages(7, "eight", 9, true)
	return &TracingTask{
		Base:        base,
		initialized: time.Now().UTC(),
	}
}

func (this *TracingTask) Times() []time.Time {
	return append([]time.Time{
		this.initialized,
		this.read,
	}, this.executed...)
}

func (this *TracingTask) RequiredReads() []interface{} {
	this.read = time.Now().UTC()
	return this.Base.RequiredReads()
}
func (this *TracingTask) Execute(ctx context.Context) TaskResult {
	this.executeContext = ctx
	this.executed = append(this.executed, time.Now().UTC())
	return this.Base.Execute(ctx)
}

/////////////////////////////////////////////////////////////

type NilResultTask struct{ *Base }

func NewNilResultTask() *NilResultTask {
	return &NilResultTask{Base: New()}
}

func (this *NilResultTask) Execute(_ context.Context) TaskResult {
	return nil
}

/////////////////////////////////////////////////////////////

type FakeExternalIO struct {
	readContext     context.Context
	writeContext    context.Context
	dispatchContext context.Context

	reads    []interface{}
	writes   []interface{}
	messages []interface{}
}

func (this *FakeExternalIO) Read(ctx context.Context, items ...interface{}) {
	this.readContext = ctx
	this.reads = append(this.reads, items...)
}
func (this *FakeExternalIO) Write(ctx context.Context, items ...interface{}) {
	this.writeContext = ctx
	this.writes = append(this.writes, items...)
}
func (this *FakeExternalIO) Dispatch(ctx context.Context, items ...interface{}) {
	this.dispatchContext = ctx
	this.messages = append(this.messages, items...)
}
