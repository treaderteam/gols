package lift

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"gitlab.com/alexnikita/gols/gol"

	"gitlab.com/alexnikita/gols/lift/luftil"

	"gitlab.com/alexnikita/gols/lift/lifterr"
)

type instance struct {
	prefix string
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

func (ro *Route) serve(rw http.ResponseWriter, r *http.Request) {
	var (
		err         error
		response    interface{}
		res         []byte
		ps          Params
		clientError bool
		responded   bool
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
		}(&err)
	}

	if ro.DetailedLogger {
		defer func(req *http.Request, status *int, s *time.Time) {
			stat := luftil.ColorizeStatus(strconv.Itoa(*status))
			path := gol.Cyan(req.URL.Path)
			log.Printf("\n|%+v\n|%+v\n|%s [%s] [%s]\n", req.Header, req.URL.Query(), path, stat, time.Since(*s))
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
		if (*status) != 200 && (*status) != 204 {
			if clientError {
				http.Error(*writer, (*err).Error(), *status)
			} else if !responded {
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

	if ps.QueryParams != nil {
		for v := range *ps.QueryParams {
			value := r.URL.Query().Get(v)
			if len(value) < 1 {
				err = errors.New("not enough query params")
				clientError = true
				responseStatus = 400
				return
			}
			log.Printf("about to write query %s\n", time.Now())
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
			log.Printf("about to write header %s %s\n", value, time.Now())
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
			if err = json.Unmarshal(b, ps.Body); err != nil {
				log.Println(err)
				clientError = true
				responseStatus = 400
				err = fmt.Errorf("cannot parse input value, you probably sending data in incorrect format : %s", err.Error())
				return
			}
			break
		}
		r.Body.Close()
	}

	if ps.BodyRaw != nil {
		br := ps.BodyRaw
		*br = r.Body
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
		switch err.(type) {
		case lifterr.LiftClientError:
			responseStatus = 400
			clientError = true
			return
		default:
			return
		}
	}

	if !ro.DontMarshal {
		if res, err = json.Marshal(response); err != nil {
			responseStatus = 500
			return
		}
		rw.Write(res)
	} else {
		if response == nil {
			log.Println("write header")
			rw.WriteHeader(responseStatus)
		} else {
			res = response.([]byte)
			log.Println(len(res))
			rw.Write(res)
		}
	}

	responded = true
}
