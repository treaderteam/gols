package lift

import (
	"net/url"
)

func parseQueryParams(exp strmap, real url.Values) (result map[string]string, err bool) {
	result = make(strmap)
	for v := range exp {
		value := real.Get(v)
		if len(value) < 1 {
			err = true
			return
		}
		result[v] = value
	}
	return
}
