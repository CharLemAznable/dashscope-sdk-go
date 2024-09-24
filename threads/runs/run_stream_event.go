package runs

import (
	"github.com/CharLemAznable/dashscope-sdk-go/threads"
	"github.com/CharLemAznable/dashscope-sdk-go/threads/messages"
	"github.com/CharLemAznable/gfx/net/gclientx"
	"github.com/gogf/gf/v2/encoding/gjson"
)

type RunStreamEventType string

//goland:noinspection GoUnusedConst
const (
	ThreadCreated           RunStreamEventType = "thread.created"
	ThreadRunCreated        RunStreamEventType = "thread.run.created"
	ThreadRunQueued         RunStreamEventType = "thread.run.queued"
	ThreadRunInProgress     RunStreamEventType = "thread.run.in_progress"
	ThreadRunRequiresAction RunStreamEventType = "thread.run.requires_action"
	ThreadRunCompleted      RunStreamEventType = "thread.run.completed"
	ThreadRunFailed         RunStreamEventType = "thread.run.failed"
	ThreadRunCancelling     RunStreamEventType = "thread.run.cancelling"
	ThreadRunCanceled       RunStreamEventType = "thread.run.cancelled"
	ThreadRunExpired        RunStreamEventType = "thread.run.expired"
	ThreadRunStepCreated    RunStreamEventType = "thread.run.step.created"
	ThreadRunStepInProgress RunStreamEventType = "thread.run.step.in_progress"
	ThreadRunStepDelta      RunStreamEventType = "thread.run.step.delta"
	ThreadRunStepCompleted  RunStreamEventType = "thread.run.step.completed"
	ThreadRunStepFailed     RunStreamEventType = "thread.run.step.failed"
	ThreadRunStepCancelled  RunStreamEventType = "thread.run.step.cancelled"
	ThreadRunStepExpired    RunStreamEventType = "thread.run.step.expired"
	ThreadMessageCreated    RunStreamEventType = "thread.message.created"
	ThreadMessageInProgress RunStreamEventType = "thread.message.in_progress"
	ThreadMessageDelta      RunStreamEventType = "thread.message.delta"
	ThreadMessageCompleted  RunStreamEventType = "thread.message.completed"
	ThreadMessageIncomplete RunStreamEventType = "thread.message.incomplete"
	RunStreamError          RunStreamEventType = "error"
)

type RunStreamEvent interface {
	Id() string
	Event() RunStreamEventType
}

func newRunStreamEvent(origin *gclientx.Event) RunStreamEvent {
	switch RunStreamEventType(origin.Event) {
	case ThreadCreated:
		return newThreadEvent(origin)
	case ThreadRunCreated, ThreadRunQueued, ThreadRunInProgress,
		ThreadRunRequiresAction, ThreadRunCompleted, ThreadRunFailed,
		ThreadRunCancelling, ThreadRunCanceled, ThreadRunExpired:
		return newRunEvent(origin)
	case ThreadRunStepCreated, ThreadRunStepInProgress, ThreadRunStepCompleted,
		ThreadRunStepFailed, ThreadRunStepCancelled, ThreadRunStepExpired:
		return newRunStepEvent(origin)
	case ThreadRunStepDelta:
		return newRunStepDeltaEvent(origin)
	case ThreadMessageCreated, ThreadMessageCompleted,
		ThreadMessageInProgress, ThreadMessageIncomplete:
		return newMessageEvent(origin)
	case ThreadMessageDelta:
		return newMessageDeltaEvent(origin)
	case RunStreamError:
		return newErrorEvent(origin)
	default:
		return newUnknownEvent(origin)
	}
}

////////////////////////////////////////////////////////////////////////////////

func buildEventBase(origin *gclientx.Event) eventBase {
	return eventBase{
		id:    origin.Id,
		event: RunStreamEventType(origin.Event),
	}
}

type eventBase struct {
	id    string
	event RunStreamEventType
}

func (e *eventBase) Id() string {
	return e.id
}
func (e *eventBase) Event() RunStreamEventType {
	return e.event
}

////////////////////////////////////////////////////////////////////////////////

type ThreadEvent interface {
	RunStreamEvent
	Thread() threads.Thread
}

func newThreadEvent(origin *gclientx.Event) ThreadEvent {
	return &threadEvent{
		eventBase: buildEventBase(origin),
		thread:    threads.NewThreadFromJson(gjson.New(origin.Data)),
	}
}

type threadEvent struct {
	eventBase
	thread threads.Thread
}

func (e *threadEvent) Thread() threads.Thread {
	return e.thread
}

////////////////////////////////////////////////////////////////////////////////

type RunEvent interface {
	RunStreamEvent
	Run() Run
}

func newRunEvent(origin *gclientx.Event) RunEvent {
	return &runEvent{
		eventBase: buildEventBase(origin),
		run:       NewRunFromJson(gjson.New(origin.Data)),
	}
}

type runEvent struct {
	eventBase
	run Run
}

func (e *runEvent) Run() Run {
	return e.run
}

////////////////////////////////////////////////////////////////////////////////

type RunStepEvent interface {
	RunStreamEvent
	RunStep() RunStep
}

func newRunStepEvent(origin *gclientx.Event) RunStepEvent {
	return &runStepEvent{
		eventBase: buildEventBase(origin),
		runStep:   NewRunStepFromJson(gjson.New(origin.Data)),
	}
}

type runStepEvent struct {
	eventBase
	runStep RunStep
}

func (e *runStepEvent) RunStep() RunStep {
	return e.runStep
}

////////////////////////////////////////////////////////////////////////////////

type RunStepDeltaEvent interface {
	RunStreamEvent
	RunStepDelta() RunStepDelta
}

func newRunStepDeltaEvent(origin *gclientx.Event) RunStepDeltaEvent {
	return &runStepDeltaEvent{
		eventBase:    buildEventBase(origin),
		runStepDelta: NewRunStepDeltaFromJson(gjson.New(origin.Data)),
	}
}

type runStepDeltaEvent struct {
	eventBase
	runStepDelta RunStepDelta
}

func (e *runStepDeltaEvent) RunStepDelta() RunStepDelta {
	return e.runStepDelta
}

////////////////////////////////////////////////////////////////////////////////

type MessageEvent interface {
	RunStreamEvent
	Message() messages.Message
}

func newMessageEvent(origin *gclientx.Event) MessageEvent {
	return &messageEvent{
		eventBase: buildEventBase(origin),
		message:   messages.NewMessageFromJson(gjson.New(origin.Data)),
	}
}

type messageEvent struct {
	eventBase
	message messages.Message
}

func (e *messageEvent) Message() messages.Message {
	return e.message
}

////////////////////////////////////////////////////////////////////////////////

type MessageDeltaEvent interface {
	RunStreamEvent
	MessageDelta() messages.MessageDelta
}

func newMessageDeltaEvent(origin *gclientx.Event) MessageDeltaEvent {
	return &messageDeltaEvent{
		eventBase:    buildEventBase(origin),
		messageDelta: messages.NewMessageDeltaFromJson(gjson.New(origin.Data)),
	}
}

type messageDeltaEvent struct {
	eventBase
	messageDelta messages.MessageDelta
}

func (e *messageDeltaEvent) MessageDelta() messages.MessageDelta {
	return e.messageDelta
}

////////////////////////////////////////////////////////////////////////////////

type ErrorEvent interface {
	RunStreamEvent
	Error() string
}

func newErrorEvent(origin *gclientx.Event) ErrorEvent {
	return &errorEvent{
		eventBase: buildEventBase(origin),
		error:     origin.Data,
	}
}

type errorEvent struct {
	eventBase
	error string
}

func (e *errorEvent) Error() string {
	return e.error
}

////////////////////////////////////////////////////////////////////////////////

type UnknownEvent interface {
	RunStreamEvent
	Data() string
}

func newUnknownEvent(origin *gclientx.Event) UnknownEvent {
	return &unknownEvent{
		eventBase: buildEventBase(origin),
		data:      origin.Data,
	}
}

type unknownEvent struct {
	eventBase
	data string
}

func (e *unknownEvent) Data() string {
	return e.data
}
