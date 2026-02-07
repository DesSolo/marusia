package middlewares

import (
	"context"
	"log/slog"

	"github.com/DesSolo/marusia"
)

func Logging(next marusia.HandlerFunc) marusia.HandlerFunc {
	return func(ctx context.Context, req *marusia.Request, resp *marusia.Response) {
		slog.InfoContext(ctx, "user",
			"userID", req.Session.UserID,
			"sessionID", req.Session.SessionID,
		)

		next(ctx, req, resp)
	}
}
