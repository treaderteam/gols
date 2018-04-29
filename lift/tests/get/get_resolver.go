package get

import (
	"errors"

	"gitlab.com/alexnikita/gols/lift"
)

// NewGetResolver create new instance
func NewGetResolver() lift.ParamResolver {
	return getResolver{
		params: lift.Params{
			QueryParams: map[string]string{"payload": ""},
			Headers:     map[string]string{"Type": ""},
		},
	}
}

type getResolver struct {
	params lift.Params
}

func (g getResolver) Resolve(params lift.Params) (status int, resp interface{}, err error) {
	query := params.QueryParams
	payload := query["payload"]
	typ := params.Headers["Type"]
	if typ != "Test" {
		return 400, nil, errors.New("wrong type")
	}
	return 200, payload, nil
}

func (g getResolver) GetParams() lift.Params {
	return g.params
}
