package lift_test

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)
import (
	"fmt"
	"io/ioutil"
	"net/http"

	"gitlab.com/alexnikita/gols/lift"
)

type User struct {
	Name string `json:"Name"`
	Info struct {
		Bio string `json:"Bio"`
	} `json:"Info"`
}

type getTest struct {
	params lift.Params
}

func (g getTest) Resolve() (status int, response interface{}, err error) {
	status = 200
	err = nil
	response = fmt.Sprintf("hello, %s %s!\n", (*g.params.QueryParams)["name"], (*g.params.QueryParams)["surname"])
	return
}

func (g getTest) GetParams() lift.Params {
	return g.params
}

type postTest struct {
	params lift.Params
}

func (p postTest) Resolve() (status int, response interface{}, err error) {
	par, _ := (p.params.Body).(*User)
	response = fmt.Sprintf("%s\n%s\n", par.Name, par.Info.Bio)
	status = 200
	return
}

func (p postTest) GetParams() lift.Params {
	return p.params
}

func TestPrepare(t *testing.T) {
	i := lift.New()
	g := getTest{}
	qp := map[string]string{"name": "name", "surname": "surname"}
	g.params = lift.Params{QueryParams: &qp}
	i.Register(lift.Route{
		Params:   g,
		Path:     "/test",
		Method:   "GET",
		Resolver: g,
	})

	u := User{}
	pt := postTest{}
	pt.params = lift.Params{Body: &u}
	i.Register(lift.Route{
		Params:   pt,
		Method:   "POST",
		Path:     "/posttest",
		Resolver: pt,
	})

	go http.ListenAndServe(":8080", i)

}

func TestHandlers(t *testing.T) {
	r, _ := http.Get("http://localhost:8080/test?name=tester&surname=testerov")
	defer r.Body.Close()
	res, _ := ioutil.ReadAll(r.Body)
	assert.Equal(t, 200, r.StatusCode)
	assert.Equal(t, "\"hello, tester testerov!\\n\"", string(res))
	t.Logf("%s\n", string(res))
}

func TestPodyParser(t *testing.T) {
	var mock User
	mock.Name = "test"
	mock.Info.Bio = "post test body"

	body, _ := json.Marshal(mock)

	r, _ := http.Post("http://localhost:8080/posttest", "application/json", bytes.NewReader(body))
	defer r.Body.Close()
	res, _ := ioutil.ReadAll(r.Body)
	assert.Equal(t, 200, r.StatusCode)
	assert.Equal(t, "\"test\\npost test body\\n\"", string(res))
	t.Logf("%s\n", string(res))
}
