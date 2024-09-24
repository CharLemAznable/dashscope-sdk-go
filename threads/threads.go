package threads

import (
	"context"
	"github.com/CharLemAznable/dashscope-sdk-go/common"
	"github.com/CharLemAznable/dashscope-sdk-go/internal/client"
	"github.com/CharLemAznable/dashscope-sdk-go/threads/messages"
	"github.com/gogf/gf/v2/util/gvalid"
)

type ThreadParam struct {
	Messages []*messages.MessageParam `json:"messages,omitempty"`
	Metadata map[string]string        `json:"metadata,omitempty"`
}

func Create(ctx context.Context, param ...*ThreadParam) (Thread, error) {
	threadParam := &ThreadParam{}
	if len(param) > 0 && param[0] != nil {
		threadParam = param[0]
	}
	return client.Post(ctx, NewThreadFromJson, "/threads", threadParam)
}

func Retrieve(ctx context.Context, threadId string) (Thread, error) {
	if err := validateThreadId(ctx, threadId); err != nil {
		return nil, err
	}
	return client.Get(ctx, NewThreadFromJson, "/threads/"+threadId)
}

func Update(ctx context.Context, threadId string, param *common.UpdateMetadataParam) (Thread, error) {
	if err := validateThreadId(ctx, threadId); err != nil {
		return nil, err
	}
	return client.Post(ctx, NewThreadFromJson, "/threads/"+threadId, param)
}

func Delete(ctx context.Context, threadId string) (*common.DeletionStatus, error) {
	if err := validateThreadId(ctx, threadId); err != nil {
		return nil, err
	}
	return client.Delete(ctx, common.NewDeletionStatusFromJson, "/threads/"+threadId)
}

func validateThreadId(ctx context.Context, threadId string) error {
	return gvalid.New().
		Rules("required").
		Messages("threadId is required").
		Data(threadId).Run(ctx)
}
