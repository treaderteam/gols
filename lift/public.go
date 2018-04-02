package lift

import "io"

// Instance lift instance
type Instance instance

// Route HTTP route
type Route route

// New creates new instance of lift
func New() Instance {
	return Instance{routes: make(map[string]Route)}
}

// NewWithPrefix creates new instance of lift with
// additional HTTP path prefix, which must be not trimmed
// This is mainly for more detailed logging
func NewWithPrefix(prefix string) Instance {
	return Instance{
		routes: make(map[string]Route),
		prefix: prefix,
	}
}

// Params HTTP params
type Params struct {
	QueryParams *map[string]string
	Headers     *map[string]string
	Body        interface{}
	BodyRaw     *io.ReadCloser
}

// New create new params
func (p Params) New() Params {
	return Params{
		QueryParams: new(map[string]string),
		Headers:     new(map[string]string),
		Body:        nil,
		BodyRaw:     nil,
	}
}

// Register register new lift route
func (i *Instance) Register(r Route) {
	path := r.Path
	if i.prefix != "" {
		path = i.prefix + path
	}
	i.routes[path] = r
}
