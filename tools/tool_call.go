package tools

import "github.com/gogf/gf/v2/encoding/gjson"

type ToolCall interface {
	GetType() string
	GetId() string
}

func ToolCallsFromJsons(jsons []*gjson.Json) (toolCalls []ToolCall) {
	toolCalls = make([]ToolCall, 0, len(jsons))
	for _, json := range jsons {
		if json.Get("type").String() == "function" {
			toolCalls = append(toolCalls, NewFunctionToolCallFromJson(json))
		} else if json.Get("type").String() == "code_interpreter" {
			toolCalls = append(toolCalls, NewCodeInterpreterToolCallFromJson(json))
		}
	}
	return
}

////////////////////////////////////////////////////////////////////////////////

type FunctionToolCall interface {
	ToolCall
	GetFunction() *FunctionCall
}

type FunctionCall struct {
	Name      string `json:"name"`
	Arguments string `json:"arguments"`
	Output    string `json:"output"`
}

func NewFunctionToolCallFromJson(json *gjson.Json) FunctionToolCall {
	return &functionToolCall{
		Type: json.Get("type").String(),
		Id:   json.Get("id").String(),
		Function: &FunctionCall{
			Name:      json.Get("function.name").String(),
			Arguments: json.Get("function.arguments").String(),
			Output:    json.Get("function.output").String(),
		},
	}
}

type functionToolCall struct {
	Type     string        `json:"type"`
	Id       string        `json:"id"`
	Function *FunctionCall `json:"function"`
}

func (f *functionToolCall) GetType() string {
	return f.Type
}
func (f *functionToolCall) GetId() string {
	return f.Id
}
func (f *functionToolCall) GetFunction() *FunctionCall {
	return f.Function
}

////////////////////////////////////////////////////////////////////////////////

type CodeInterpreterToolCall interface {
	ToolCall
}

func NewCodeInterpreterToolCallFromJson(json *gjson.Json) CodeInterpreterToolCall {
	return &codeInterpreterToolCall{
		Type: json.Get("type").String(),
		Id:   json.Get("id").String(),
	}
}

type codeInterpreterToolCall struct {
	Type string `json:"type"`
	Id   string `json:"id"`
}

func (c *codeInterpreterToolCall) GetType() string {
	return c.Type
}
func (c *codeInterpreterToolCall) GetId() string {
	return c.Id
}
