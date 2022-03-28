package tool

import "github.com/chen-mou/gf/frame/g"

type Result g.Map

func (Result) Success(msg string, data *map[string]any) g.Map {
	return g.Map{
		"code": 200,
		"msg":  msg,
		"data": data,
	}
}

func (Result) Fail(code int, msg string) g.Map {
	return g.Map{
		"code": code,
		"msg":  msg,
		"data": nil,
	}
}
