package middleware

import (
	"github.com/eurofurence/reg-mail-service/internal/repository/config"
	"github.com/eurofurence/reg-mail-service/web/util/ctlutil"
	"github.com/eurofurence/reg-mail-service/web/util/ctxvalues"
	"github.com/eurofurence/reg-mail-service/web/util/media"
	"net/http"
)

func fromApiTokenHeader(r *http.Request) string {
	return r.Header.Get(media.HeaderXApiKey)
}

func TokenValidator(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		// try api token first
		apiTokenValue := fromApiTokenHeader(r)
		if apiTokenValue != "" {
			// ignore jwt if set (may still need to pass it through to other service)
			if apiTokenValue == config.FixedApiToken() {
				ctxvalues.SetApiToken(ctx, apiTokenValue)
				next.ServeHTTP(w, r)
			} else {
				ctlutil.UnauthenticatedError(ctx, w, r, "invalid api token", "request supplied invalid api token, denying")
			}
			return
		}
	}

	return http.HandlerFunc(fn)
}
