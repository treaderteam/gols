package lift

import (
	"net/http"
	"net/url"
	"time"
)

// CORSResolver must maintain CORS headers
type CORSResolver interface {
	ResolveCORS(rw http.ResponseWriter) (status int)
}

// Resolver performs main logic in request
type Resolver interface {
	Resolve() (status int, response interface{}, err error)
}

// ErrorHandler perform handling server errors and must respond to request
type ErrorHandler interface {
	HandleError(err *error, rw http.ResponseWriter)
}

// HeadersResolver return required headers to add to response
type HeadersResolver interface {
	ResolveHeaders() http.Header
}

// Parametrer must return pointer to required params
type Parametrer interface {
	GetParams() Params
}

// Logger logs requests
type Logger interface {
	Log(status *int, method string, url *url.URL, start *time.Time)
}
