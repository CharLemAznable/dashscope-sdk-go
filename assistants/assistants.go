package assistants

import (
	"context"
	"github.com/CharLemAznable/dashscope-sdk-go/common"
	"github.com/CharLemAznable/dashscope-sdk-go/internal/client"
	"github.com/CharLemAznable/dashscope-sdk-go/tools"
	"github.com/gogf/gf/v2/util/gvalid"
)

type AssistantParam struct {
	Model        string            `json:"model,omitempty" v:"required"`
	Name         string            `json:"name,omitempty"`
	Description  string            `json:"description,omitempty"`
	Instructions string            `json:"instructions,omitempty"`
	Tools        []tools.Tool      `json:"tools,omitempty"`
	FileIds      []string          `json:"file_ids,omitempty"`
	Metadata     map[string]string `json:"metadata,omitempty"`
}

func Create(ctx context.Context, param *AssistantParam) (Assistant, error) {
	if err := validateAssistantParam(ctx, param); err != nil {
		return nil, err
	}
	return client.Post(ctx, NewAssistantFromJson, "/assistants", param)
}

func Retrieve(ctx context.Context, assistantId string) (Assistant, error) {
	if err := validateAssistantId(ctx, assistantId); err != nil {
		return nil, err
	}
	return client.Get(ctx, NewAssistantFromJson, "/assistants/"+assistantId)
}

func Update(ctx context.Context, assistantId string, param *AssistantParam) (Assistant, error) {
	if err := validateAssistantId(ctx, assistantId); err != nil {
		return nil, err
	}
	return client.Post(ctx, NewAssistantFromJson, "/assistants/"+assistantId, param)
}

func Delete(ctx context.Context, assistantId string) (*common.DeletionStatus, error) {
	if err := validateAssistantId(ctx, assistantId); err != nil {
		return nil, err
	}
	return client.Delete(ctx, common.NewDeletionStatusFromJson, "/assistants/"+assistantId)
}

func List(ctx context.Context, param ...*common.ListParam) (*common.ListResult[Assistant], error) {
	listParam := &common.ListParam{}
	if len(param) > 0 && param[0] != nil {
		listParam = param[0]
	}
	return client.Get(ctx, common.ListResultFunc(NewAssistantFromJson), "/assistants", listParam)
}

func validateAssistantParam(ctx context.Context, param *AssistantParam) error {
	return gvalid.New().Data(param).Run(ctx)
}

func validateAssistantId(ctx context.Context, assistantId string) error {
	return gvalid.New().
		Rules("required").
		Messages("assistantId is required").
		Data(assistantId).Run(ctx)
}
