package filter

import (
	"github.com/eurofurence/reg-mail-service/web/util/ctlutil"
	"github.com/eurofurence/reg-mail-service/web/util/ctxvalues"
	"net/http"
)

func ApiToken(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		if ctxvalues.HasApiToken(ctx) {
			handler(w, r)
		} else {
			ctlutil.UnauthenticatedError(ctx, w, r, "incorrect api key was supplied", "anonymous access attempt")
		}
	}
}
