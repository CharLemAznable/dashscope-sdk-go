package tools

type ToolOutput interface {
	GetToolCallId() string
	GetOutput() string
}

func NewToolOutput(toolCallId string, output string) ToolOutput {
	return &toolOutput{
		ToolCallId: toolCallId,
		Output:     output,
	}
}

type toolOutput struct {
	ToolCallId string `json:"tool_call_id"`
	Output     string `json:"output"`
}

func (t *toolOutput) GetToolCallId() string {
	return t.ToolCallId
}
func (t *toolOutput) GetOutput() string {
	return t.Output
}
