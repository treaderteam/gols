package get

import (
	"log"

	"gitlab.com/alexnikita/gols/lift"
)

// NewGetResolver create new instance
func NewGetResolver() lift.ParamResolver {
	return getResolver{
		params: lift.Params{
			QueryParams: &map[string]string{"payload": ""},
		},
	}
}

type getResolver struct {
	params lift.Params
}

func (g getResolver) Resolve(params lift.Params) (status int, resp interface{}, err error) {
	query := params.QueryParams
	payload := (*query)["payload"]
	log.Println(payload)
	return 200, payload, nil
}

func (g getResolver) GetParams() lift.Params {
	return g.params
}
