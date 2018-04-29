package lift

import (
	"net/http"
	"strings"
)

type strmap map[string]string

func parseHeaders(exp strmap, real http.Header) (result map[string]string, err bool) {
	result = make(strmap)
	for v := range exp {
		value := real.Get(v)
		if len(value) < 1 {
			v = strings.ToLower(v)
			value := real.Get(v)
			if len(value) < 1 {
				err = true
				return
			}
		}
		result[v] = value
	}

	return result, false
}
