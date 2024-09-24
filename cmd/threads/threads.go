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

	thread, err := threads.Create(ctx, &threads.ThreadParam{
		Messages: []*messages.MessageParam{
			{
				Role:    "user",
				Content: "你好",
			},
		},
	})
	g.Dump(thread, err)
	threadId := thread.GetId()
	if threadId == "" {
		return
	}

	defer func() {
		status, err := threads.Delete(ctx, threadId)
		g.Dump(status, err)
	}()

	thread, err = threads.Update(ctx, threadId, &common.UpdateMetadataParam{
		Metadata: map[string]string{
			"user": "abc123",
		},
	})
	g.Dump(thread, err)

	thread, err = threads.Retrieve(ctx, threadId)
	g.Dump(thread, err)
}
