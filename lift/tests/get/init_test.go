package get_test

import (
	"log"
	"net/http"

	"gitlab.com/alexnikita/gols/lift"
	"gitlab.com/alexnikita/gols/lift/tests/get"
)

const (
	S_GET_PATH = "/get"
	GET_PORT   = ":8090"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	l := lift.New()
	getresolver := get.NewGetResolver()

	l.Register(lift.Route{
		Method:      "GET",
		Path:        S_GET_PATH,
		Resolver:    getresolver,
		Params:      getresolver,
		DontMarshal: true,
	})

	go http.ListenAndServe(GET_PORT, l)
}
