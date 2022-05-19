package sentinel

import (
	sentinel "github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/base"
	"github.com/alibaba/sentinel-golang/core/circuitbreaker"
	"github.com/alibaba/sentinel-golang/core/config"
)

type Handler struct {
	Resource string
	Opt      sentinel.EntryOption
	Success  func()
	Error    func()
}

var resource map[string][]func(string, sentinel.EntryOption, func(), func())

func Register(name string, opt sentinel.EntryOption, error func(), success func()) {
	en, err := sentinel.Entry(name, sentinel.WithTrafficType(base.Inbound))
	if err != nil {
		error()
	} else {
		success()
		en.Exit()
	}
}

func AddRule(rule []*circuitbreaker.Rule) {
	_, err := circuitbreaker.LoadRules(rule)
}

func sentinelInit() {
	err := sentinel.InitWithConfig(&config.Entity{})
	if err != nil {
		panic(any(err))
	}
}
