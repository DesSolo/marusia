package marusia

import (
	"fmt"
	"strings"
)

// HandlerFunc ...
type HandlerFunc func(req *Request) *Response

// DialogRouter ...
type DialogRouter struct {
	router         map[string]HandlerFunc
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

// Select select dialog function by token name
func (r *DialogRouter) Select(token string) (HandlerFunc, error) {
	hf, ok := r.router[token]
	if ok {
		return hf, nil
	}

	if r.defaultHandler != nil {
		return r.defaultHandler, nil
	}

	return nil, fmt.Errorf("dialog endpoint not found by token: %s", token)
}
