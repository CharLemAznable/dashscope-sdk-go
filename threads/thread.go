package threads

import "github.com/gogf/gf/v2/encoding/gjson"

type Thread interface {
	GetId() string
	GetCreatedAt() int64
	GetMetadata() map[string]string
}

func NewThreadFromJson(json *gjson.Json) Thread {
	a := &thread{
		Id:        json.Get("id").String(),
		CreatedAt: json.Get("created_at").Int64(),
		Metadata:  json.Get("metadata").MapStrStr(),
	}
	return a
}

type thread struct {
	Id        string            `json:"id"`
	CreatedAt int64             `json:"created_at"`
	Metadata  map[string]string `json:"metadata"`
}

func (t *thread) GetId() string {
	return t.Id
}
func (t *thread) GetCreatedAt() int64 {
	return t.CreatedAt
}
func (t *thread) GetMetadata() map[string]string {
	return t.Metadata
}
