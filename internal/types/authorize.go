package types

type AuthParams struct {
	ResponseType string `form:"response_type" binding:"required"`
	Redirect     string `form:"redirect" binding:"required"`
	Scope        string `form:"scope" binding:"required"`
	State        string `form:"state" binding:"required"`
}

func (a AuthParams) Code() string {
	return a.State + "code"
}
