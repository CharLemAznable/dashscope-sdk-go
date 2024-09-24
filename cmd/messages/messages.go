package main

import (
	"context"
	"github.com/CharLemAznable/dashscope-sdk-go/common"
	"github.com/CharLemAznable/dashscope-sdk-go/threads"
	"github.com/CharLemAznable/dashscope-sdk-go/threads/messages"
	"github.com/gogf/gf/v2/frame/g"
)

func main() {
	ctx := context.TODO()

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

	message, err := messages.Create(ctx, threadId, &messages.MessageParam{
		Role:    "user",
		Content: "你好",
	})
	g.Dump(message, err)
	messageId := message.GetId()
	if messageId == "" {
		return
	}

	message, err = messages.Update(ctx, threadId, messageId, &common.UpdateMetadataParam{
		Metadata: map[string]string{
			"key": "value",
		},
	})
	g.Dump(message, err)

	message, err = messages.Retrieve(ctx, threadId, messageId)
	g.Dump(message, err)

	messageLs, err := messages.List(ctx, threadId)
	g.Dump(messageLs, err)
}
