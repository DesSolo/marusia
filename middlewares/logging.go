package middlewares

import (
	"context"
	"log"

	"github.com/DesSolo/marusia"
)

func Logging(next marusia.HandlerFunc) marusia.HandlerFunc {
	return func(req *marusia.Request, resp *marusia.Response, ctx context.Context) {
		log.Printf("userID: %s text: %s", req.Session.UserID, req.Request.OriginalUtterance)
		next(req, resp, ctx)
	}
}
