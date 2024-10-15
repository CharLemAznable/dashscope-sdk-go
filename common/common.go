package common

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/samber/lo"
)

type UpdateMetadataParam struct {
	Metadata map[string]string `json:"metadata,omitempty"`
}

type DeletionStatus struct {
	Id      string `json:"id"`
	Deleted bool   `json:"deleted"`
	Object  string `json:"object"`
}

func NewDeletionStatusFromJson(json *gjson.Json) *DeletionStatus {
	return &DeletionStatus{
		Id:      json.Get("id").String(),
		Deleted: json.Get("deleted").Bool(),
		Object:  json.Get("object").String(),
	}
}

type ListParam struct {
	Limit *int    `json:"limit,omitempty"`
	Order *string `json:"order,omitempty"`
}

type ListResult[T any] struct {
	FirstId string `json:"first_id"`
	LastId  string `json:"last_id"`
	HasMore bool   `json:"has_more"`
	Data    []T    `json:"data"`
	Object  string `json:"object"`
}

func ListResultFromJson[T any](json *gjson.Json, dataItemMapping func(*gjson.Json) T) *ListResult[T] {
	return &ListResult[T]{
		FirstId: json.Get("first_id").String(),
		LastId:  json.Get("last_id").String(),
		HasMore: json.Get("has_more").Bool(),
		Data: lo.Map(json.GetJsons("data"),
			func(item *gjson.Json, _ int) T {
				return dataItemMapping(item)
			}),
		Object: json.Get("object").String(),
	}
}

func ListResultFunc[T any](dataItemMapping func(*gjson.Json) T) func(json *gjson.Json) *ListResult[T] {
	return func(json *gjson.Json) *ListResult[T] {
		return ListResultFromJson(json, dataItemMapping)
	}
}
