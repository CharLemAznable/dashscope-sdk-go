package client

import (
	"context"
	"github.com/gogf/gf/v2/encoding/gjson"
	"net/http"
)

func Request[T any](ctx context.Context, fromJson func(json *gjson.Json) T,
	method string, url string, data ...interface{}) (T, error) {
	bytes, err := Client(ctx).RequestBytes(ctx, method, url, data...)
	if err != nil {
		return any(nil).(T), err
	}
	return fromJson(gjson.New(bytes)), nil
}

func Get[T any](ctx context.Context, fromJson func(json *gjson.Json) T,
	url string, data ...interface{}) (T, error) {
	return Request(ctx, fromJson, http.MethodGet, url, data...)
}

func Post[T any](ctx context.Context, fromJson func(json *gjson.Json) T,
	url string, data ...interface{}) (T, error) {
	return Request(ctx, fromJson, http.MethodPost, url, data...)
}

func Delete[T any](ctx context.Context, fromJson func(json *gjson.Json) T,
	url string, data ...interface{}) (T, error) {
	return Request(ctx, fromJson, http.MethodDelete, url, data...)
}
