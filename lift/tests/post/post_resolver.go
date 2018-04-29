package post

import "gitlab.com/alexnikita/gols/lift"

// PostResolver exported type
type PostResolver postResolver

type postResolver struct {
	params lift.Params
}

// PostHelthcheck type for post test
type PostHelthcheck struct {
	Status      bool
	Description string
	Secret      string
}

type PostRequestHelthcheck struct {
	Secret string
}

// NewPostResolver func
func NewPostResolver() PostResolver {
	return PostResolver{
		params: lift.Params{
			Body: new(PostRequestHelthcheck),
		},
	}
}

// Resolve lift
func (PostResolver) Resolve(params lift.Params) (status int, response interface{}, err error) {
	status = 500

	req, ok := params.Body.(*PostRequestHelthcheck)
	if !ok {
		status = 400
		return
	}

	secret := req.Secret

	response = PostHelthcheck{
		Description: "allright",
		Secret:      secret,
		Status:      true,
	}

	status = 200

	return
}

// GetParameters lift
func (p PostResolver) GetParams() lift.Params {
	return p.params
}
