package ctxvalues

import (
	"context"
	"github.com/eurofurence/reg-mail-service/internal/repository/config"
)

const ContextMap = "map"

const ContextRequestId = "requestid"
const ContextApiToken = "apitoken"

func CreateContextWithValueMap(ctx context.Context) context.Context {
	// this is so we can add values to our context, like ... I don't know ... the http status from the response!
	contextMap := make(map[string]string)

	ctx = context.WithValue(ctx, ContextMap, contextMap)
	return ctx
}

func valueOrDefault(ctx context.Context, key string, defaultValue string) string {
	contextMapUntyped := ctx.Value(ContextMap)
	if contextMapUntyped == nil {
		return defaultValue
	}
	contextMap := contextMapUntyped.(map[string]string)

	if val, ok := contextMap[key]; ok {
		return val
	} else {
		return defaultValue
	}
}

func setValue(ctx context.Context, key string, value string) {
	contextMapUntyped := ctx.Value(ContextMap)
	if contextMapUntyped != nil {
		contextMap := contextMapUntyped.(map[string]string)
		contextMap[key] = value
	}
}

func RequestId(ctx context.Context) string {
	return valueOrDefault(ctx, ContextRequestId, "00000000")
}

func SetRequestId(ctx context.Context, requestId string) {
	setValue(ctx, ContextRequestId, requestId)
}

func HasApiToken(ctx context.Context) bool {
	v := valueOrDefault(ctx, ContextApiToken, "")
	return v == config.FixedApiToken()
}

func SetApiToken(ctx context.Context, apiToken string) {
	setValue(ctx, ContextApiToken, apiToken)
}
