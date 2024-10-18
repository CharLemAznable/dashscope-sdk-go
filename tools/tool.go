package tools

import "github.com/gogf/gf/v2/encoding/gjson"

type Tool interface {
	GetType() string
}

func FromJsons(jsons []*gjson.Json) (tools []Tool) {
	tools = make([]Tool, 0, len(jsons))
	for _, json := range jsons {
		if json.Get("type").String() == "function" {
			tools = append(tools, NewFunctionToolFromJson(json))
		} else if json.Get("type").String() == "code_interpreter" {
			tools = append(tools, NewCodeInterpreterToolFromJson(json))
		}
	}
	return
}

////////////////////////////////////////////////////////////////////////////////

type FunctionTool interface {
	Tool
	GetFunction() *FunctionDefinition
}

type FunctionDefinition struct {
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	Parameters  map[string]interface{} `json:"parameters"`
}

func NewFunctionTool(name, description string, parameters map[string]interface{}) FunctionTool {
	return &functionTool{
		Type: "function",
		Function: &FunctionDefinition{
			Name:        name,
			Description: description,
			Parameters:  parameters,
		},
	}
}

func NewFunctionToolFromJson(json *gjson.Json) FunctionTool {
	return &functionTool{
		Type: json.Get("type").String(),
		Function: &FunctionDefinition{
			Name:        json.Get("function.name").String(),
			Description: json.Get("function.description").String(),
			Parameters:  json.Get("function.parameters").Map(),
		},
	}
}

type functionTool struct {
	Type     string              `json:"type"`
	Function *FunctionDefinition `json:"function"`
}

func (f *functionTool) GetType() string {
	return f.Type
}
func (f *functionTool) GetFunction() *FunctionDefinition {
	return f.Function
}

////////////////////////////////////////////////////////////////////////////////

type CodeInterpreterTool interface {
	Tool
}

func NewCodeInterpreterTool() CodeInterpreterTool {
	return &codeInterpreterTool{
		Type: "code_interpreter",
	}
}

func NewCodeInterpreterToolFromJson(json *gjson.Json) CodeInterpreterTool {
	return &codeInterpreterTool{
		Type: json.Get("type").String(),
	}
}

type codeInterpreterTool struct {
	Type string `json:"type"`
}

func (c *codeInterpreterTool) GetType() string {
	return c.Type
}
