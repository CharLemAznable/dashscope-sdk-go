package messages

import (
	"github.com/CharLemAznable/dashscope-sdk-go/common"
	"github.com/gogf/gf/v2/encoding/gjson"
)

type Message interface {
	GetId() string
	GetObject() string
	GetCreatedAt() int64
	GetThreadId() string
	GetStatus() Status
	GetIncompleteDetails() interface{}
	GetCompletedAt() *int64
	GetIncompleteAt() *int64
	GetRole() Role
	GetContent() []Content
	GetAssistantId() *string
	GetRunId() *string
	GetFileIds() []string
	GetMetadata() map[string]string
}

type Status string

//goland:noinspection ALL
const (
	InProgress Status = "in_progress"
	Incomplete Status = "incomplete"
	Completed  Status = "completed"
)

type Role string

//goland:noinspection ALL
const (
	User      Role = "user"
	Assistant Role = "assistant"

	Bot        Role = "bot"
	System     Role = "system"
	Attachment Role = "attachment"
	Tool       Role = "tool"
)

func NewMessageFromJson(json *gjson.Json) Message {
	a := &message{
		Id:                json.Get("id").String(),
		Object:            json.Get("object").String(),
		CreatedAt:         json.Get("created_at").Int64(),
		ThreadId:          json.Get("thread_id").String(),
		Status:            Status(json.Get("status").String()),
		IncompleteDetails: json.Get("incomplete_details").Val(),
		CompletedAt:       common.VarInt64(json.Get("completed_at")),
		IncompleteAt:      common.VarInt64(json.Get("incomplete_at")),
		Role:              Role(json.Get("role").String()),
		Content:           ContentsFromJsons(json.GetJsons("content")),
		AssistantId:       common.VarString(json.Get("assistant_id")),
		RunId:             common.VarString(json.Get("run_id")),
		FileIds:           json.Get("file_ids").Strings(),
		Metadata:          json.Get("metadata").MapStrStr(),
	}
	return a
}

type message struct {
	Id                string            `json:"id"`
	Object            string            `json:"object"`
	CreatedAt         int64             `json:"created_at"`
	ThreadId          string            `json:"thread_id"`
	Status            Status            `json:"status"`
	IncompleteDetails interface{}       `json:"incomplete_details"`
	CompletedAt       *int64            `json:"completed_at"`
	IncompleteAt      *int64            `json:"incomplete_at"`
	Role              Role              `json:"role"`
	Content           []Content         `json:"content"`
	AssistantId       *string           `json:"assistant_id"`
	RunId             *string           `json:"run_id"`
	FileIds           []string          `json:"file_ids"`
	Metadata          map[string]string `json:"metadata"`
}

func (m *message) GetId() string {
	return m.Id
}
func (m *message) GetObject() string {
	return m.Object
}
func (m *message) GetCreatedAt() int64 {
	return m.CreatedAt
}
func (m *message) GetThreadId() string {
	return m.ThreadId
}
func (m *message) GetStatus() Status {
	return m.Status
}
func (m *message) GetIncompleteDetails() interface{} {
	return m.IncompleteDetails
}
func (m *message) GetCompletedAt() *int64 {
	return m.CompletedAt
}
func (m *message) GetIncompleteAt() *int64 {
	return m.IncompleteAt
}
func (m *message) GetRole() Role {
	return m.Role
}
func (m *message) GetContent() []Content {
	return m.Content
}
func (m *message) GetAssistantId() *string {
	return m.AssistantId
}
func (m *message) GetRunId() *string {
	return m.RunId
}
func (m *message) GetFileIds() []string {
	return m.FileIds
}
func (m *message) GetMetadata() map[string]string {
	return m.Metadata
}

////////////////////////////////////////////////////////////////////////////////

type MessageDelta interface {
	GetId() string
	GetObject() string
	GetDelta() Delta
}

func NewMessageDeltaFromJson(json *gjson.Json) MessageDelta {
	return &messageDelta{
		Id:     json.Get("id").String(),
		Object: json.Get("object").String(),
		Delta:  NewDeltaFromJson(json.GetJson("delta")),
	}
}

type messageDelta struct {
	Id     string `json:"id"`
	Object string `json:"object"`
	Delta  Delta  `json:"delta"`
}

func (m *messageDelta) GetId() string {
	return m.Id
}
func (m *messageDelta) GetObject() string {
	return m.Object
}
func (m *messageDelta) GetDelta() Delta {
	return m.Delta
}

////////////////////////////////////////////////////////////////////////////////

type Delta interface {
	GetRole() Role
	GetContent() []Content
	GetFileIds() []string
}

func NewDeltaFromJson(json *gjson.Json) Delta {
	if json.IsNil() {
		return nil
	}
	return &delta{
		Role:    Role(json.Get("role").String()),
		Content: ContentsFromJsons(json.GetJsons("content")),
		FileIds: json.Get("file_ids").Strings(),
	}
}

type delta struct {
	Role    Role      `json:"role"`
	Content []Content `json:"content"`
	FileIds []string  `json:"file_ids"`
}

func (d *delta) GetRole() Role {
	return d.Role
}
func (d *delta) GetContent() []Content {
	return d.Content
}
func (d *delta) GetFileIds() []string {
	return d.FileIds
}
