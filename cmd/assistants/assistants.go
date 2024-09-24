package main

import (
	"context"
	"github.com/CharLemAznable/dashscope-sdk-go/assistants"
	"github.com/CharLemAznable/dashscope-sdk-go/tools"
	"github.com/gogf/gf/v2/frame/g"
)

func main() {
	ctx := context.TODO()

	assistant, err := assistants.Create(ctx, &assistants.AssistantParam{
		Model: "qwen-max",
		Tools: []tools.Tool{
			tools.NewFunctionTool("get_current_datetime", "查询当前时间", map[string]interface{}{}),
			tools.NewCodeInterpreterTool(),
		},
	})
	g.Dump(assistant, err)
	assistantId := assistant.GetId()
	if assistantId == "" {
		return
	}

	defer func() {
		status, err := assistants.Delete(ctx, assistantId)
		g.Dump(status, err)
	}()

	assistant, err = assistants.Update(ctx, assistantId, &assistants.AssistantParam{
		Name: "test_assistant",
	})
	g.Dump(assistant, err)

	assistantLs, err := assistants.List(ctx)
	g.Dump(assistantLs, err)

	assistant, err = assistants.Retrieve(ctx, assistantId)
	g.Dump(assistant, err)
}
