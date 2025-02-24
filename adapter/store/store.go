package store

import (
	"github.com/WuKongIM/WuKongIM/pkg/wklog"
)

var AdStore *AdapterStore

func init() {
	AdStore = NewAdapterStore()
	AdStore.Info("AdStore init")
}

type AdapterStore struct {
	wklog.Log
}

// NewAdapterStore  创建API
func NewAdapterStore() *AdapterStore {
	return &AdapterStore{
		Log: wklog.NewWKLog("AdStore"),
	}
}
