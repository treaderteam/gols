package lift

import (
	"net/http"
	"path"
)

func (i Instance) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	check := r.Method + " " + r.URL.Path
	for p, ro := range i.routes {
		pattern := ro.Method + " " + p
		if ok, err := path.Match(pattern, check); ok && err == nil {
			ro.serve(w, r)
			return
		}
		pattern = "OPTIONS " + p
		if ok, err := path.Match(pattern, check); ok && err == nil {
			ro.serve(w, r)
			return
		}

	}

	http.NotFound(w, r)
}
