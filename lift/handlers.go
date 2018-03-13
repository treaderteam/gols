package lift

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Instance instance

type Route route

type Params struct {
	QueryParams *map[string]string
	Headers     *map[string]string
	Body        interface{}
}

func (p Params) New() Params {
	return Params{
		QueryParams: new(map[string]string),
		Headers:     new(map[string]string),
		Body:        nil,
	}
}

type instance struct {
	routes map[string]Route
}

type route struct {
	Path            string
	Method          string
	Params          Parametrer
	Resolver        Resolver
	CORSResolver    CORSResolver
	ErrorHandler    ErrorHandler
	Logger          Logger
	ResponseHeaders HeadersResolver
	DetailedLogger  bool
	DontMarshal     bool
}

func New() Instance {
	return Instance{routes: make(map[string]Route)}
}

func (ro *Route) serve(rw http.ResponseWriter, r *http.Request) {
	var (
		err         error
		response    interface{}
		res         []byte
		ps          Params
		clientError bool
	)

	responseStatus := 500
	start := time.Now()

	if ro.ErrorHandler != nil {
		defer func(e *error) {
			if _e := recover(); _e != nil {
				*e = fmt.Errorf("panic happened %+v", _e)
				ro.ErrorHandler.HandleError(e, rw)
			} else if (*e) != nil {
				ro.ErrorHandler.HandleError(e, rw)
			}
		}(&err)
	} else {
		defer func(e *error) {
			if _e := recover(); _e != nil {
				log.Println(_e)
			}
			if (*e) != nil {
				log.Println(*e)
			}
		}(&err)
	}

	if ro.DetailedLogger {
		defer func(req *http.Request, status *int, s *time.Time) {
			log.Printf("|%+v\n|%+v\n|%s [%d] [%s]\n", req.Header, req.URL.Query(), req.URL.Path, *status, time.Since(*s))
		}(r, &responseStatus, &start)
	} else {
		if ro.Logger != nil {
			defer ro.Logger.Log(&responseStatus, r.Method, r.URL, &start)
		} else {
			defer func(status *int, method string, u *url.URL, s *time.Time) {
				log.Printf("%s [%s] %d %s\n", u.Path, method, *status, time.Since(*s))
			}(&responseStatus, r.Method, r.URL, &start)
		}
	}

	defer func(writer *http.ResponseWriter, method string, url *url.URL, s *time.Time, status *int, err *error) {
		log.Printf("%d \n", *status)
		if (*status) != 200 && (*status) != 204 {
			if clientError {
				http.Error(*writer, (*err).Error(), *status)
			} else {
				(*writer).WriteHeader(*status)
			}
		}
	}(&rw, r.Method, r.URL, &start, &responseStatus, &err)

	if r.Method == http.MethodOptions {
		if ro.CORSResolver != nil {
			responseStatus = ro.CORSResolver.ResolveCORS(rw)
			return
		}
		err = errors.New("got cors request, but resolver not specified")
		return
	}

	if ro.Params != nil {
		ps = ro.Params.GetParams()
	}

	defer r.Body.Close()

	if r.Method != ro.Method {
		responseStatus = http.StatusMethodNotAllowed
		clientError = true
		err = fmt.Errorf("method %s not allowed an such endpoint", r.Method)
		return
	}

	if ps.QueryParams != nil {
		for v := range *ps.QueryParams {
			value := r.URL.Query().Get(v)
			if len(value) < 1 {
				err = errors.New("not enough query params")
				clientError = true
				responseStatus = 400
				return
			}
			(*ps.QueryParams)[v] = value
		}
	}

	if ps.Headers != nil {
		for v := range *ps.Headers {
			value := r.Header.Get(v)
			if len(value) < 1 {
				v = strings.ToLower(v)
				value := r.Header.Get(v)
				if len(value) < 1 {
					clientError = true
					responseStatus = 400
					err = errors.New("not enough headers")
					return
				}
			}
			(*ps.Headers)[v] = value
		}
	}

	if ps.Body != nil {
		var b []byte
		if b, err = ioutil.ReadAll(r.Body); err != nil {
			return
		}

		switch ps.Body.(type) {
		case *[]byte:
			_p := ps.Body.(*[]byte)
			*_p = b
			break
		default:

			if err = json.NewDecoder(bytes.NewReader(b)).Decode(ps.Body); err != nil {
				return
			}
			break
		}

	}

	if ro.Resolver != nil {
		responseStatus, response, err = ro.Resolver.Resolve()
	}

	if ro.ResponseHeaders != nil {
		headers := ro.ResponseHeaders.ResolveHeaders()
		for n, v := range headers {
			rw.Header().Add(n, strings.Join(v, ","))
		}
	}

	if err != nil {
		responseStatus = 500
		return
	}

	if response == nil {
		responseStatus = 204
		return
	}

	if !ro.DontMarshal {
		if res, err = json.Marshal(response); err != nil {
			responseStatus = 500
			return
		}
	} else {
		res = response.([]byte)
	}

	rw.Write(res)
}

func (i *Instance) Register(r Route) {
	i.routes[r.Path] = r
}

func (i *Instance) Kindle() *http.ServeMux {
	mux := http.NewServeMux()
	return mux
}
