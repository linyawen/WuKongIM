package store

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/WuKongIM/WuKongIM/adapter"
	"github.com/WuKongIM/WuKongIM/adapter/util"
	"github.com/WuKongIM/WuKongIM/pkg/raft/types"
	"github.com/WuKongIM/WuKongIM/pkg/wkdb"
	"github.com/WuKongIM/WuKongIM/pkg/wkutil"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

// AppendMessages ，保存消息，维护累加messageSeq,并且记录 该频道的最大lastMessageSeq，lastAppendTime
// seq会话内自增
// return result, error: result.Id 为messageId,result.Index 为 MessageSeq
func (as *AdapterStore) AppendMessages(ctx context.Context, channelId string, channelType uint8, msgs []wkdb.Message) (proposeRespSet types.ProposeRespSet, retErr error) {
	defer func() {
		if r := recover(); r != nil {
			// 将 panic 的值转换为 error
			if e, ok := r.(error); ok {
				retErr = e // 如果是 error 类型，直接赋值
			} else {
				retErr = fmt.Errorf("panic: %v", r) // 否则封装为 error
			}
			as.Error("AdapterStore AppendMessages recovered from panic:", zap.Error(retErr), zap.Any("channelId", channelId), zap.Any("channelType", channelType), zap.Any("msgs", wkutil.ToJSON(msgs)))
		}
	}()

	as.Info("AdapterStore AppendMessages 入参", zap.Any("channelId", channelId), zap.Any("channelType", channelType), zap.Any("msgs", wkutil.ToJSON(msgs)))
	path := fmt.Sprintf("/wkadapter/store/appendMessages?channelId=%v&channelType=%v", channelId, channelType)
	jsonData, err := as.post(path, msgs)
	util.PanicIfErrorf(err, "AdapterStore AppendMessages post error,url:%v", path)

	var resp adapter.ComResp[types.ProposeRespSet]
	err = json.Unmarshal(jsonData, &resp)
	util.PanicIfErrorf(err, "AdapterStore AppendMessages json.Unmarshal error,url:%v,jsonData:%v", path, jsonData)

	if resp.Success() {
		return resp.DATA, nil
	} else {
		return nil, errors.Errorf("AdapterStore AppendMessages business error,url:%v,resp:%v", path, resp)
	}
}
