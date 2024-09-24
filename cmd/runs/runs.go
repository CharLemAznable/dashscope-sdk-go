package main

import (
	"context"
	"fmt"
	"github.com/CharLemAznable/dashscope-sdk-go/assistants"
	"github.com/CharLemAznable/dashscope-sdk-go/threads"
	"github.com/CharLemAznable/dashscope-sdk-go/threads/messages"
	"github.com/CharLemAznable/dashscope-sdk-go/threads/runs"
	"github.com/CharLemAznable/dashscope-sdk-go/tools"
	"github.com/gogf/gf/v2/frame/g"
)

func main() {
	ctx := context.TODO()

	assistant, err := assistants.Create(ctx, &assistants.AssistantParam{
		Model:        "qwen-turbo",
		Name:         "回答日常问题的机器人",
		Description:  "一个智能助手，解答用户的问题",
		Instructions: "请礼貌地回答用户的问题",
		Tools: []tools.Tool{
			tools.NewFunctionTool("get_current_datetime", "查询当前时间", map[string]interface{}{}),
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

	thread, err := threads.Create(ctx)
	g.Dump(thread, err)
	threadId := thread.GetId()
	if threadId == "" {
		return
	}

	defer func() {
		status, err := threads.Delete(ctx, threadId)
		g.Dump(status, err)
	}()

	////////////////////////////////////////////////////////////////////////////////

	message, err := messages.Create(ctx, threadId, &messages.MessageParam{
		Role:    "user",
		Content: "你好, 请介绍一下你自己",
	})
	g.Dump(message, err)
	messageId := message.GetId()
	if messageId == "" {
		return
	}

	run, err := runs.Create(ctx, threadId, &runs.RunParam{
		AssistantId: assistantId,
	})
	g.Dump(run, err)
	runId := run.GetId()
	if runId == "" {
		return
	}

	run, err = runs.Wait(ctx, threadId, runId)
	g.Dump(run.GetStatus())

	messageLs, err := messages.List(ctx, threadId)
	g.Dump(messageLs, err)

	////////////////////////////////////////////////////////////////////////////////

	message, err = messages.Create(ctx, threadId, &messages.MessageParam{
		Role:    "user",
		Content: "你有什么优点吗？",
	})
	g.Dump(message, err)
	messageId = message.GetId()
	if messageId == "" {
		return
	}

	stream, err := runs.CreateStream(ctx, threadId, &runs.RunParam{
		AssistantId: assistantId,
	})
	func(stream runs.RunStream, err error) {
		if err != nil {
			g.Dump(err)
			return
		}
		defer stream.Drain()
		for streamEvent := range stream.Event() {
			if messageEvent, ok := streamEvent.(runs.MessageDeltaEvent); ok {
				if len(messageEvent.MessageDelta().GetDelta().GetContent()) > 0 {
					content := messageEvent.MessageDelta().GetDelta().GetContent()[0]
					if textContent, ok := content.(messages.TextContent); ok {
						fmt.Println(textContent.GetText().Value)
					}
				}
			}
		}
		fmt.Println()
	}(stream, err)

	messageLs, err = messages.List(ctx, threadId)
	g.Dump(messageLs, err)

	////////////////////////////////////////////////////////////////////////////////

	message, err = messages.Create(ctx, threadId, &messages.MessageParam{
		Role:    "user",
		Content: "现在几点了？",
	})
	g.Dump(message, err)
	messageId = message.GetId()
	if messageId == "" {
		return
	}

	run, err = runs.Create(ctx, threadId, &runs.RunParam{
		AssistantId: assistantId,
	})
	g.Dump(run, err)
	runId = run.GetId()
	if runId == "" {
		return
	}

	run, err = runs.Wait(ctx, threadId, runId)
	g.Dump(run.GetStatus())

	if run.GetStatus() == runs.RunRequiresAction {
		g.Dump(run.GetRequiredAction().GetSubmitToolOutputs().GetToolCalls())
		if len(run.GetRequiredAction().GetSubmitToolOutputs().GetToolCalls()) > 0 {
			toolCallId := run.GetRequiredAction().GetSubmitToolOutputs().GetToolCalls()[0].GetId()
			toolOutput := tools.NewToolOutput(toolCallId, "{\"current_datetime\":\"2023-09-01 12:00:00\"}")
			stream, err = runs.SubmitStreamToolOutputs(ctx, threadId, runId, &runs.SubmitToolOutputsParam{
				ToolOutputs: []tools.ToolOutput{toolOutput},
			})
			func(stream runs.RunStream, err error) {
				if err != nil {
					g.Dump(err)
					return
				}
				defer stream.Drain()
				for streamEvent := range stream.Event() {
					if messageEvent, ok := streamEvent.(runs.MessageDeltaEvent); ok {
						if len(messageEvent.MessageDelta().GetDelta().GetContent()) > 0 {
							content := messageEvent.MessageDelta().GetDelta().GetContent()[0]
							if textContent, ok := content.(messages.TextContent); ok {
								fmt.Println(textContent.GetText().Value)
							}
						}
					}
				}
				fmt.Println()
			}(stream, err)

			messageLs, err = messages.List(ctx, threadId)
			g.Dump(messageLs, err)
		}
	}
}
