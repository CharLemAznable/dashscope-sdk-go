package runs

import (
	"github.com/CharLemAznable/dashscope-sdk-go/common"
	"github.com/CharLemAznable/dashscope-sdk-go/tools"
	"github.com/gogf/gf/v2/encoding/gjson"
)

type RunStep interface {
	GetId() string
	GetObject() string
	GetCreatedAt() int64
	GetAssistantId() string
	GetThreadId() string
	GetRunId() string
	GetType() RunStepType
	GetStatus() RunStepStatus
	GetStepDetails() StepDetailBase
	GetLastError() LastError
	GetExpiredAt() *int64
	GetCancelledAt() *int64
	GetFailedAt() *int64
	GetCompletedAt() *int64
	GetMetadata() map[string]string
	GetUsage() Usage
}

type RunStepType string

//goland:noinspection ALL
const (
	RunStepMessageCreation RunStepType = "message_creation"
	RunStepToolCalls       RunStepType = "tool_calls"
)

type RunStepStatus string

//goland:noinspection ALL
const (
	RunStepInProgress RunStepStatus = "in_progress"
	RunStepCancelled  RunStepStatus = "cancelled"
	RunStepFailed     RunStepStatus = "failed"
	RunStepCompleted  RunStepStatus = "completed"
	RunStepExpired    RunStepStatus = "expired"
)

func NewRunStepFromJson(json *gjson.Json) RunStep {
	return &runStep{
		Id:          json.Get("id").String(),
		Object:      json.Get("object").String(),
		CreatedAt:   json.Get("created_at").Int64(),
		AssistantId: json.Get("assistant_id").String(),
		ThreadId:    json.Get("thread_id").String(),
		RunId:       json.Get("run_id").String(),
		Type:        RunStepType(json.Get("type").String()),
		Status:      RunStepStatus(json.Get("status").String()),
		StepDetails: NewStepDetailBaseFromJson(json.GetJson("step_details")),
		LastError:   NewLastErrorFromJson(json.GetJson("last_error")),
		ExpiredAt:   common.VarInt64(json.Get("expired_at")),
		CancelledAt: common.VarInt64(json.Get("cancelled_at")),
		FailedAt:    common.VarInt64(json.Get("failed_at")),
		CompletedAt: common.VarInt64(json.Get("completed_at")),
		Metadata:    json.Get("metadata").MapStrStr(),
		Usage:       NewUsageFromJson(json.GetJson("usage")),
	}
}

type runStep struct {
	Id          string            `json:"id"`
	Object      string            `json:"object"`
	CreatedAt   int64             `json:"created_at"`
	AssistantId string            `json:"assistant_id"`
	ThreadId    string            `json:"thread_id"`
	RunId       string            `json:"run_id"`
	Type        RunStepType       `json:"type"`
	Status      RunStepStatus     `json:"status"`
	StepDetails StepDetailBase    `json:"step_details"`
	LastError   LastError         `json:"last_error"`
	ExpiredAt   *int64            `json:"expired_at"`
	CancelledAt *int64            `json:"cancelled_at"`
	FailedAt    *int64            `json:"failed_at"`
	CompletedAt *int64            `json:"completed_at"`
	Metadata    map[string]string `json:"metadata"`
	Usage       Usage             `json:"usage"`
}

func (r *runStep) GetId() string {
	return r.Id
}
func (r *runStep) GetObject() string {
	return r.Object
}
func (r *runStep) GetCreatedAt() int64 {
	return r.CreatedAt
}
func (r *runStep) GetAssistantId() string {
	return r.AssistantId
}
func (r *runStep) GetThreadId() string {
	return r.ThreadId
}
func (r *runStep) GetRunId() string {
	return r.RunId
}
func (r *runStep) GetType() RunStepType {
	return r.Type
}
func (r *runStep) GetStatus() RunStepStatus {
	return r.Status
}
func (r *runStep) GetStepDetails() StepDetailBase {
	return r.StepDetails
}
func (r *runStep) GetLastError() LastError {
	return r.LastError
}
func (r *runStep) GetExpiredAt() *int64 {
	return r.ExpiredAt
}
func (r *runStep) GetCancelledAt() *int64 {
	return r.CancelledAt
}
func (r *runStep) GetFailedAt() *int64 {
	return r.FailedAt
}
func (r *runStep) GetCompletedAt() *int64 {
	return r.CompletedAt
}
func (r *runStep) GetMetadata() map[string]string {
	return r.Metadata
}
func (r *runStep) GetUsage() Usage {
	return r.Usage
}

////////////////////////////////////////////////////////////////////////////////

type RunStepDelta interface {
	GetId() string
	GetObject() string
	GetDelta() Delta
}

func NewRunStepDeltaFromJson(json *gjson.Json) RunStepDelta {
	return &runStepDelta{
		Id:     json.Get("id").String(),
		Object: json.Get("object").String(),
		Delta:  NewDeltaFromJson(json.GetJson("delta")),
	}
}

type runStepDelta struct {
	Id     string `json:"id"`
	Object string `json:"object"`
	Delta  Delta  `json:"delta"`
}

func (r *runStepDelta) GetId() string {
	return r.Id
}
func (r *runStepDelta) GetObject() string {
	return r.Object
}
func (r *runStepDelta) GetDelta() Delta {
	return r.Delta
}

////////////////////////////////////////////////////////////////////////////////

type Delta interface {
	GetStepDetails() StepDetailBase
}

func NewDeltaFromJson(json *gjson.Json) Delta {
	if json.IsNil() {
		return nil
	}
	return &delta{
		StepDetails: NewStepDetailBaseFromJson(json.GetJson("step_details")),
	}
}

type delta struct {
	StepDetails StepDetailBase `json:"step_details"`
}

func (d *delta) GetStepDetails() StepDetailBase {
	return d.StepDetails
}

////////////////////////////////////////////////////////////////////////////////

type StepDetailBase interface {
	GetType() string
}

func NewStepDetailBaseFromJson(json *gjson.Json) StepDetailBase {
	if json.IsNil() {
		return nil
	}
	switch RunStepType(json.Get("type").String()) {
	case RunStepMessageCreation:
		return NewStepMessageCreationFromJson(json)
	case RunStepToolCalls:
		return NewStepToolCallsFromJson(json)
	default:
		return nil
	}
}

////////////////////////////////////////////////////////////////////////////////

type StepMessageCreation interface {
	StepDetailBase
	GetMessageCreation() MessageCreation
}

func NewStepMessageCreationFromJson(json *gjson.Json) StepMessageCreation {
	if json.IsNil() {
		return nil
	}
	return &stepMessageCreation{
		Type:            json.Get("type").String(),
		MessageCreation: NewMessageCreationFromJson(json.GetJson("message_creation")),
	}
}

type stepMessageCreation struct {
	Type            string          `json:"type"`
	MessageCreation MessageCreation `json:"message_creation"`
}

func (s *stepMessageCreation) GetType() string {
	return s.Type
}
func (s *stepMessageCreation) GetMessageCreation() MessageCreation {
	return s.MessageCreation
}

////////////////////////////////////////////////////////////////////////////////

type MessageCreation interface {
	GetMessageId() string
}

func NewMessageCreationFromJson(json *gjson.Json) MessageCreation {
	if json.IsNil() {
		return nil
	}
	return &messageCreation{
		MessageId: json.Get("message_id").String(),
	}
}

type messageCreation struct {
	MessageId string `json:"message_id"`
}

func (m *messageCreation) GetMessageId() string {
	return m.MessageId
}

////////////////////////////////////////////////////////////////////////////////

type StepToolCalls interface {
	StepDetailBase
	GetToolCalls() []tools.ToolCall
}

func NewStepToolCallsFromJson(json *gjson.Json) StepToolCalls {
	if json.IsNil() {
		return nil
	}
	return &stepToolCalls{
		Type:      json.Get("type").String(),
		ToolCalls: tools.ToolCallsFromJsons(json.GetJsons("tool_calls")),
	}
}

type stepToolCalls struct {
	Type      string           `json:"type"`
	ToolCalls []tools.ToolCall `json:"tool_calls"`
}

func (s *stepToolCalls) GetType() string {
	return s.Type
}
func (s *stepToolCalls) GetToolCalls() []tools.ToolCall {
	return s.ToolCalls
}
