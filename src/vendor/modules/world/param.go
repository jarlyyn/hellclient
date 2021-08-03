package world

type RequiredParam struct {
	Name  string
	Desc  string
	Intro string
}

func NewRequiredParam() *RequiredParam {
	return &RequiredParam{}
}

type ParamsInfo struct {
	Params         map[string]string
	ParamComments  map[string]string
	RequiredParams []*RequiredParam
}
