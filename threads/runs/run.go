package runs

import (
	"github.com/CharLemAznable/dashscope-sdk-go/common"
	"github.com/CharLemAznable/dashscope-sdk-go/tools"
	"github.com/gogf/gf/v2/encoding/gjson"
)

type Run interface {
	GetId() string
	GetObject() string
	GetCreatedAt() int64
	GetThreadId() string
	GetAssistantId() string
	GetStatus() RunStatus
	GetRequiredAction() RequiredAction
	GetLastError() LastError
	GetExpiresAt() *int64
	GetStartedAt() *int64
	GetCancelledAt() *int64
	GetFailedAt() *int64
	GetCompletedAt() *int64
	GetIncompleteDetails() IncompleteDetails
	GetModel() string
	GetInstructions() string
	GetTools() []tools.Tool
	GetMetadata() map[string]string
	GetUsage() Usage
	GetTemperature() *float64
	GetMaxPromptTokens() *int64
	GetMaxCompletionTokens() *int64
	GetTruncationStrategy() TruncationStrategy
	GetToolChoice() interface{}
	GetResponseFormat() interface{}
}

type RunStatus string

//goland:noinspection ALL
const (
	RunQueued         RunStatus = "queued"
	RunInProgress     RunStatus = "in_progress"
	RunRequiresAction RunStatus = "requires_action"
	RunCancelling     RunStatus = "cancelling"
	RunCancelled      RunStatus = "cancelled"
	RunFailed         RunStatus = "failed"
	RunCompleted      RunStatus = "completed"
	RunExpired        RunStatus = "expired"
)

func NewRunFromJson(json *gjson.Json) Run {
	return &run{
		Id:                  json.Get("id").String(),
		Object:              json.Get("object").String(),
		CreatedAt:           json.Get("created_at").Int64(),
		ThreadId:            json.Get("thread_id").String(),
		AssistantId:         json.Get("assistant_id").String(),
		Status:              RunStatus(json.Get("status").String()),
		RequiredAction:      NewRequiredActionFromJson(json.GetJson("required_action")),
		LastError:           NewLastErrorFromJson(json.GetJson("last_error")),
		ExpiresAt:           common.VarInt64(json.Get("expires_at")),
		StartedAt:           common.VarInt64(json.Get("started_at")),
		CancelledAt:         common.VarInt64(json.Get("cancelled_at")),
		FailedAt:            common.VarInt64(json.Get("failed_at")),
		CompletedAt:         common.VarInt64(json.Get("completed_at")),
		IncompleteDetails:   NewIncompleteDetailsFromJson(json.GetJson("incomplete_details")),
		Model:               json.Get("model").String(),
		Instructions:        json.Get("instructions").String(),
		Tools:               tools.FromJsons(json.GetJsons("tools")),
		Metadata:            json.Get("metadata").MapStrStr(),
		Usage:               NewUsageFromJson(json.GetJson("usage")),
		Temperature:         common.VarFloat64(json.Get("temperature")),
		MaxPromptTokens:     common.VarInt64(json.Get("max_prompt_tokens")),
		MaxCompletionTokens: common.VarInt64(json.Get("max_completion_tokens")),
		TruncationStrategy:  NewTruncationStrategyFromJson(json.GetJson("truncation_strategy")),
		ToolChoice:          json.Get("tool_choice").Val(),
		ResponseFormat:      json.Get("response_format").Val(),
	}
}

type run struct {
	Id                  string             `json:"id"`
	Object              string             `json:"object"`
	CreatedAt           int64              `json:"created_at"`
	ThreadId            string             `json:"thread_id"`
	AssistantId         string             `json:"assistant_id"`
	Status              RunStatus          `json:"status"`
	RequiredAction      RequiredAction     `json:"required_action"`
	LastError           LastError          `json:"last_error"`
	ExpiresAt           *int64             `json:"expires_at"`
	StartedAt           *int64             `json:"started_at"`
	CancelledAt         *int64             `json:"cancelled_at"`
	FailedAt            *int64             `json:"failed_at"`
	CompletedAt         *int64             `json:"completed_at"`
	IncompleteDetails   IncompleteDetails  `json:"incomplete_details"`
	Model               string             `json:"model"`
	Instructions        string             `json:"instructions"`
	Tools               []tools.Tool       `json:"tools"`
	Metadata            map[string]string  `json:"metadata"`
	Usage               Usage              `json:"usage"`
	Temperature         *float64           `json:"temperature"`
	MaxPromptTokens     *int64             `json:"max_prompt_tokens"`
	MaxCompletionTokens *int64             `json:"max_completion_tokens"`
	TruncationStrategy  TruncationStrategy `json:"truncation_strategy"`
	ToolChoice          interface{}        `json:"tool_choice"`
	ResponseFormat      interface{}        `json:"response_format"`
}

func (r *run) GetId() string {
	return r.Id
}
func (r *run) GetObject() string {
	return r.Object
}
func (r *run) GetCreatedAt() int64 {
	return r.CreatedAt
}
func (r *run) GetThreadId() string {
	return r.ThreadId
}
func (r *run) GetAssistantId() string {
	return r.AssistantId
}
func (r *run) GetStatus() RunStatus {
	return r.Status
}
func (r *run) GetRequiredAction() RequiredAction {
	return r.RequiredAction
}
func (r *run) GetLastError() LastError {
	return r.LastError
}
func (r *run) GetExpiresAt() *int64 {
	return r.ExpiresAt
}
func (r *run) GetStartedAt() *int64 {
	return r.StartedAt
}
func (r *run) GetCancelledAt() *int64 {
	return r.CancelledAt
}
func (r *run) GetFailedAt() *int64 {
	return r.FailedAt
}
func (r *run) GetCompletedAt() *int64 {
	return r.CompletedAt
}
func (r *run) GetIncompleteDetails() IncompleteDetails {
	return r.IncompleteDetails
}
func (r *run) GetModel() string {
	return r.Model
}
func (r *run) GetInstructions() string {
	return r.Instructions
}
func (r *run) GetTools() []tools.Tool {
	return r.Tools
}
func (r *run) GetMetadata() map[string]string {
	return r.Metadata
}
func (r *run) GetUsage() Usage {
	return r.Usage
}
func (r *run) GetTemperature() *float64 {
	return r.Temperature
}
func (r *run) GetMaxPromptTokens() *int64 {
	return r.MaxPromptTokens
}
func (r *run) GetMaxCompletionTokens() *int64 {
	return r.MaxCompletionTokens
}
func (r *run) GetTruncationStrategy() TruncationStrategy {
	return r.TruncationStrategy
}
func (r *run) GetToolChoice() interface{} {
	return r.ToolChoice
}
func (r *run) GetResponseFormat() interface{} {
	return r.ResponseFormat
}

////////////////////////////////////////////////////////////////////////////////

type RequiredAction interface {
	GetType() string
	GetSubmitToolOutputs() InnerRequiredAction
}

func NewRequiredActionFromJson(json *gjson.Json) RequiredAction {
	if json.IsNil() {
		return nil
	}
	return &requiredAction{
		Type:              json.Get("type").String(),
		SubmitToolOutputs: NewInnerRequiredActionFromJson(json.GetJson("submit_tool_outputs")),
	}
}

type requiredAction struct {
	Type              string              `json:"type"`
	SubmitToolOutputs InnerRequiredAction `json:"submit_tool_outputs"`
}

func (r *requiredAction) GetType() string {
	return r.Type
}
func (r *requiredAction) GetSubmitToolOutputs() InnerRequiredAction {
	return r.SubmitToolOutputs
}

////////////////////////////////////////////////////////////////////////////////

type InnerRequiredAction interface {
	GetToolCalls() []tools.ToolCall
}

func NewInnerRequiredActionFromJson(json *gjson.Json) InnerRequiredAction {
	if json.IsNil() {
		return nil
	}
	return &innerRequiredAction{
		ToolCalls: tools.ToolCallsFromJsons(json.GetJsons("tool_calls")),
	}
}

type innerRequiredAction struct {
	ToolCalls []tools.ToolCall `json:"tool_calls"`
}

func (i *innerRequiredAction) GetToolCalls() []tools.ToolCall {
	return i.ToolCalls
}

////////////////////////////////////////////////////////////////////////////////

type LastError interface {
	GetCode() string
	GetMessage() string
}

func NewLastErrorFromJson(json *gjson.Json) LastError {
	if json.IsNil() {
		return nil
	}
	return &lastError{
		Code:    json.Get("code").String(),
		Message: json.Get("message").String(),
	}
}

type lastError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (l *lastError) GetCode() string {
	return l.Code
}
func (l *lastError) GetMessage() string {
	return l.Message
}

////////////////////////////////////////////////////////////////////////////////

type IncompleteDetails interface {
	GetReason() string
}

func NewIncompleteDetailsFromJson(json *gjson.Json) IncompleteDetails {
	if json.IsNil() {
		return nil
	}
	return &incompleteDetails{
		Reason: json.Get("reason").String(),
	}
}

type incompleteDetails struct {
	Reason string `json:"reason"`
}

func (i *incompleteDetails) GetReason() string {
	return i.Reason
}

////////////////////////////////////////////////////////////////////////////////

type Usage interface {
	GetPromptTokens() int64
	GetCompletionTokens() int64
	GetTotalTokens() int64
}

func NewUsageFromJson(json *gjson.Json) Usage {
	if json.IsNil() {
		return nil
	}
	return &usage{
		PromptTokens:     json.Get("prompt_tokens").Int64(),
		CompletionTokens: json.Get("completion_tokens").Int64(),
		TotalTokens:      json.Get("total_tokens").Int64(),
	}
}

type usage struct {
	PromptTokens     int64 `json:"prompt_tokens"`
	CompletionTokens int64 `json:"completion_tokens"`
	TotalTokens      int64 `json:"total_tokens"`
}

func (u *usage) GetPromptTokens() int64 {
	return u.PromptTokens
}
func (u *usage) GetCompletionTokens() int64 {
	return u.CompletionTokens
}
func (u *usage) GetTotalTokens() int64 {
	return u.TotalTokens
}

////////////////////////////////////////////////////////////////////////////////

type TruncationStrategy interface {
	GetType() string
	GetLastMessages() *int
}

//goland:noinspection GoUnusedExportedFunction
func NewTruncationStrategyLastMessages(lastMessages int) TruncationStrategy {
	return &truncationStrategy{
		Type:         "last_messages",
		LastMessages: common.Int(lastMessages),
	}
}

func NewTruncationStrategyFromJson(json *gjson.Json) TruncationStrategy {
	if json.IsNil() {
		return nil
	}
	return &truncationStrategy{
		Type:         json.Get("type").String(),
		LastMessages: common.VarInt(json.Get("last_messages")),
	}
}

type truncationStrategy struct {
	Type         string `json:"type"`
	LastMessages *int   `json:"last_messages"`
}

func (t *truncationStrategy) GetType() string {
	return t.Type
}
func (t *truncationStrategy) GetLastMessages() *int {
	return t.LastMessages
}
