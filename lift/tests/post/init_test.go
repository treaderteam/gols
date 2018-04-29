package post_test

import (
	"log"
	"net/http"

	"gitlab.com/alexnikita/gols/lift"
	"gitlab.com/alexnikita/gols/lift/tests/post"
)

const (
	PATH = "/post"
	PORT = ":8091"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	l := lift.New()
	resolver := post.NewPostResolver()

	l.Register(lift.Route{
		Method:   "POST",
		Path:     PATH,
		Resolver: resolver,
		Params:   resolver,
	})

	go http.ListenAndServe(PORT, l)
}
