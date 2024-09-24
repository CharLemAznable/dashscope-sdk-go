package client

import (
	"context"
	"github.com/CharLemAznable/gfx/container/gvarx"
	"github.com/CharLemAznable/gfx/frame/gx"
	"github.com/CharLemAznable/gfx/net/gclientx"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gmutex"
)

const (
	loggerName = "dashscope"

	defaultBaseUrl = "https://dashscope.aliyuncs.com/api/v1"

	configKeyForUrl       = "dashscope.url"
	configKeyForApiKey    = "dashscope.apiKey"
	configKeyForWorkSpace = "dashscope.workSpace"

	headerAuthorization = "Authorization"
	headerWorkSpace     = "X-DashScope-WorkSpace"
)

var (
	logger      = g.Log(loggerName)
	clientVar   = gvar.New(nil, true)
	clientMutex = &gmutex.Mutex{}
)

func Client(ctx context.Context) (client *gclientx.Client) {
	if !clientVar.IsNil() {
		client = clientVar.Val().(*gclientx.Client)
		return
	}
	clientMutex.LockFunc(func() {
		if !clientVar.IsNil() {
			client = clientVar.Val().(*gclientx.Client)
			return
		}
		client = gx.Client().SetIntLog(logger).
			ContentJson().Prefix(getBaseUrl(ctx))
		if apiKey := getApiKey(ctx); apiKey != "" {
			client = client.Header(g.MapStrStr{
				headerAuthorization: "Bearer " + apiKey,
			})
		}
		if workSpace := getWorkSpace(ctx); workSpace != "" {
			client = client.Header(g.MapStrStr{
				headerWorkSpace: workSpace,
			})
		}
		clientVar.Set(client)
	})
	return
}

func getBaseUrl(ctx context.Context) string {
	return gvarx.DefaultIfEmpty(g.Cfg().MustGetWithEnv(ctx, configKeyForUrl), defaultBaseUrl).String()
}

func getApiKey(ctx context.Context) string {
	return g.Cfg().MustGetWithEnv(ctx, configKeyForApiKey).String()
}

func getWorkSpace(ctx context.Context) string {
	return g.Cfg().MustGetWithEnv(ctx, configKeyForWorkSpace).String()
}
