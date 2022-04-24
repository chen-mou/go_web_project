package tool

import "github.com/chen-mou/gf/frame/g"

var custom map[string]string

func Get(key string) string {
	return g.Config().Get(key).(string)
}
