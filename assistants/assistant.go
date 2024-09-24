package assistants

import (
	"github.com/CharLemAznable/dashscope-sdk-go/common"
	"github.com/CharLemAznable/dashscope-sdk-go/tools"
	"github.com/gogf/gf/v2/encoding/gjson"
)

type Assistant interface {
	GetId() string
	GetCreatedAt() int64
	GetDescription() *string
	GetFileIds() []string
	GetInstructions() *string
	GetMetadata() map[string]string
	GetModel() string
	GetName() *string
	GetObject() string
	GetTools() []tools.Tool
}

func NewAssistantFromJson(json *gjson.Json) Assistant {
	a := &assistant{
		Id:           json.Get("id").String(),
		CreatedAt:    json.Get("created_at").Int64(),
		Description:  common.VarString(json.Get("description")),
		FileIds:      json.Get("file_ids").Strings(),
		Instructions: common.VarString(json.Get("instructions")),
		Metadata:     json.Get("metadata").MapStrStr(),
		Model:        json.Get("model").String(),
		Name:         common.VarString(json.Get("name")),
		Object:       json.Get("object").String(),
		Tools:        tools.FromJsons(json.GetJsons("tools")),
	}
	return a
}

type assistant struct {
	Id           string            `json:"id"`
	CreatedAt    int64             `json:"created_at"`
	Description  *string           `json:"description"`
	FileIds      []string          `json:"file_ids"`
	Instructions *string           `json:"instructions"`
	Metadata     map[string]string `json:"metadata"`
	Model        string            `json:"model"`
	Name         *string           `json:"name"`
	Object       string            `json:"object"`
	Tools        []tools.Tool      `json:"tools"`
}

func (a *assistant) GetId() string {
	return a.Id
}
func (a *assistant) GetCreatedAt() int64 {
	return a.CreatedAt
}
func (a *assistant) GetDescription() *string {
	return a.Description
}
func (a *assistant) GetFileIds() []string {
	return a.FileIds
}
func (a *assistant) GetInstructions() *string {
	return a.Instructions
}
func (a *assistant) GetMetadata() map[string]string {
	return a.Metadata
}
func (a *assistant) GetModel() string {
	return a.Model
}
func (a *assistant) GetName() *string {
	return a.Name
}
func (a *assistant) GetObject() string {
	return a.Object
}
func (a *assistant) GetTools() []tools.Tool {
	return a.Tools
}
