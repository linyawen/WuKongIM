package store

import (
	"context"
	"github.com/WuKongIM/WuKongIM/pkg/raft/types"
	"github.com/WuKongIM/WuKongIM/pkg/wkdb"
	"go.uber.org/zap"
)

func (as *AdapterStore) AppendMessages(ctx context.Context, channelId string, channelType uint8, msgs []wkdb.Message) (types.ProposeRespSet, error) {
	as.Info("AdapterStore AppendMessages begin", zap.Any("channelId", channelId), zap.Any("channelType", channelType), zap.Any("msg len", len(msgs)))
	//TODO 消息落库mongoDB

	resps := make([]*types.ProposeResp, 0, len(msgs))
	for _, msg := range msgs {
		resps = append(resps, &types.ProposeResp{
			Id:    uint64(msg.MessageID),
			Index: 99,
		})
	}

	return resps, nil
}
