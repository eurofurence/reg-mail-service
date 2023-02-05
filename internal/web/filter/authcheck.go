package filter

import (
	"context"
	"fmt"
	aulogging "github.com/StephanHCB/go-autumn-logging"
	"github.com/eurofurence/reg-mail-service/internal/repository/config"
	"github.com/eurofurence/reg-mail-service/internal/web/util/ctlutil"
	"github.com/eurofurence/reg-mail-service/internal/web/util/ctxvalues"
	"net/http"
)

// checkInternalAdminRequestHeader is a temporary safety measure until we have 2FA for admins.
//
// enforce extra internal request header for admin requests (header blocked for external requests)
//
// TODO: remove this workaround
func checkInternalAdminRequestHeaderForGroup(ctx context.Context, r *http.Request, group string) bool {
	if group == config.OidcAdminGroup() {
		adminRequestHeaderValue := r.Header.Get("X-Admin-Request")
		if adminRequestHeaderValue != "available" {
			aulogging.Logger.Ctx(ctx).Warn().Print("X-Admin-Request header was not set correctly!")
			return false
		}
	}
	return true
}

func HasGroupOrApiToken(group string, handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		if ctxvalues.HasApiToken(ctx) || (ctxvalues.IsAuthorizedAsGroup(ctx, group) && checkInternalAdminRequestHeaderForGroup(ctx, r, group)) {
			handler(w, r)
		} else {
			culprit := ctxvalues.Subject(ctx)
			if culprit != "" {
				ctlutil.UnauthorizedError(ctx, w, r, "you are not authorized for this operation - the attempt has been logged", fmt.Sprintf("unauthorized access attempt for group %s by %s", group, culprit))
			} else {
				ctlutil.UnauthenticatedError(ctx, w, r, "you must be logged in for this operation", "anonymous access attempt")
			}
		}
	}
}
