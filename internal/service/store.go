package service

import "github.com/WuKongIM/WuKongIM/pkg/cluster/store"
import adStore "github.com/WuKongIM/WuKongIM/adapter/store"

var Store *store.Store       // 存储相关接口
var AStore = adStore.AdStore //mongoDB 存储相关接口
