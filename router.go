package marusia

import (
	"fmt"
	"strings"
)

// DialogFunc ...
type DialogFunc func(resp *Response, req *Request) *Response

// DiaogRouter ...
type DiaogRouter struct {
	router         map[string]*DialogFunc
	defaultHandler *DialogFunc
	ignoreCase bool
}

// NewDiaogRouter ...
func NewDiaogRouter(ignoreCase bool) *DiaogRouter {
	return &DiaogRouter{
		router: make(map[string]*DialogFunc),
		ignoreCase: ignoreCase,
	}
}

// Register add dialog function
func (r *DiaogRouter) Register(token string, df DialogFunc) {
	if r.ignoreCase {
		token = strings.ToLower(token)
	}
	r.router[token] = &df
}

// RegisterDefault register default handler
// will be allways call when another router token not mathed
func (r *DiaogRouter) RegisterDefault(df DialogFunc) {
	r.defaultHandler = &df
}

// Select select dialog function by token name
func (r *DiaogRouter) Select(token string) (DialogFunc, error) {
	df, ok := r.router[token]
	if ok {
		return *df, nil
	}
	if r.defaultHandler != nil {
		return *r.defaultHandler, nil
	}
	return nil, fmt.Errorf("dialog endpoint not found by token: %s", token)
}
