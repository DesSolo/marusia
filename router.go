package marusia

import (
	"context"
	"fmt"
	"strings"
)

// HandlerFunc ...
type HandlerFunc func(req *Request, resp *Response, ctx context.Context)

// DialogRouter ...
type DialogRouter struct {
	router         map[string]HandlerFunc
	middlewares    []func(HandlerFunc) HandlerFunc
	defaultHandler HandlerFunc
	ignoreCase     bool
}

// NewDialogRouter ...
func NewDialogRouter(ignoreCase bool) *DialogRouter {
	return &DialogRouter{
		router:     make(map[string]HandlerFunc),
		ignoreCase: ignoreCase,
	}
}

// Register add dialog function
func (r *DialogRouter) Register(token string, hf HandlerFunc) {
	if r.ignoreCase {
		token = strings.ToLower(token)
	}

	r.router[token] = hf
}

// RegisterDefault register default handler
// will be allways call when another router token not matched
func (r *DialogRouter) RegisterDefault(hf HandlerFunc) {
	r.defaultHandler = hf
}

func (r *DialogRouter) Use(f func(next HandlerFunc) HandlerFunc) {
	r.middlewares = append(r.middlewares, f)
}

// Select select dialog function by token name
func (r *DialogRouter) Select(token string) (HandlerFunc, error) {
	hf, ok := r.router[token]
	if ok {
		h := chain(r.middlewares, hf)
		return h, nil
	}

	if r.defaultHandler != nil {
		return r.defaultHandler, nil
	}

	return nil, fmt.Errorf("dialog endpoint not found by token: %s", token)
}

func chain(middlewares []func(HandlerFunc) HandlerFunc, endpoint HandlerFunc) HandlerFunc {
	if len(middlewares) == 0 {
		return endpoint
	}

	h := middlewares[len(middlewares)-1](endpoint)
	for i := len(middlewares) - 2; i >= 0; i-- {
		h = middlewares[i](h)
	}

	return h
}
