package messages

import (
	"context"
	"github.com/CharLemAznable/dashscope-sdk-go/common"
	"github.com/CharLemAznable/dashscope-sdk-go/internal/client"
	"github.com/gogf/gf/v2/util/gvalid"
)

type MessageParam struct {
	Role     string            `json:"role"`
	Content  string            `json:"content"`
	FileIds  []string          `json:"file_ids,omitempty"`
	Metadata map[string]string `json:"metadata,omitempty"`
}

func Create(ctx context.Context, threadId string, param *MessageParam) (Message, error) {
	if err := validateThreadId(ctx, threadId); err != nil {
		return nil, err
	}
	if err := validateMessageParam(ctx, param); err != nil {
		return nil, err
	}
	return client.Post(ctx, NewMessageFromJson, "/threads/"+threadId+"/messages", param)
}

func Retrieve(ctx context.Context, threadId string, messageId string) (Message, error) {
	if err := validateThreadId(ctx, threadId); err != nil {
		return nil, err
	}
	if err := validateMessageId(ctx, threadId); err != nil {
		return nil, err
	}
	return client.Get(ctx, NewMessageFromJson, "/threads/"+threadId+"/messages/"+messageId)
}

func Update(ctx context.Context, threadId string, messageId string, param *common.UpdateMetadataParam) (Message, error) {
	if err := validateThreadId(ctx, threadId); err != nil {
		return nil, err
	}
	if err := validateMessageId(ctx, threadId); err != nil {
		return nil, err
	}
	return client.Post(ctx, NewMessageFromJson, "/threads/"+threadId+"/messages/"+messageId, param)
}

func List(ctx context.Context, threadId string, param ...*common.ListParam) (*common.ListResult[Message], error) {
	if err := validateThreadId(ctx, threadId); err != nil {
		return nil, err
	}
	listParam := &common.ListParam{}
	if len(param) > 0 && param[0] != nil {
		listParam = param[0]
	}
	return client.Get(ctx, common.ListResultFunc(NewMessageFromJson), "/threads/"+threadId+"/messages", listParam)
}

func validateThreadId(ctx context.Context, threadId string) error {
	return gvalid.New().
		Rules("required").
		Messages("threadId is required").
		Data(threadId).Run(ctx)
}

var (
	messageValidRules = map[string]string{
		"Role":    "required|in:user,assistant",
		"Content": "required",
	}
	messageValidMessage = map[string]interface{}{
		"Role": map[string]string{
			"required": "role is required",
			"in":       "role must one of [user|assistant]",
		},
		"Content": "content is required",
	}
)

func validateMessageParam(ctx context.Context, param *MessageParam) error {
	return gvalid.New().
		Rules(messageValidRules).
		Messages(messageValidMessage).
		Data(param).Run(ctx)
}

func validateMessageId(ctx context.Context, messageId string) error {
	return gvalid.New().
		Rules("required").
		Messages("messageId is required").
		Data(messageId).Run(ctx)
}
