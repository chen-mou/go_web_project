package sentinel

import (
	sentinel "github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/config"
)

var resource map[string][]sentinel.EntryOption

func sentinelInit() {
	err := sentinel.InitWithConfig(&config.Entity{})
	if err != nil {
		panic(any(err))
	}
}
