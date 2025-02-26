package store

import (
	"fmt"
	"github.com/WuKongIM/WuKongIM/pkg/wklog"
	"github.com/WuKongIM/WuKongIM/pkg/wkutil"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"io"
	"net"
	"net/http"
	"os"
	"strings"
	"time"
)

var AdStore *AdapterStore

func init() {
	AdStore = NewAdapterStore()
	AdStore.Info("AdStore init with host:" + AdStore.Host)
}

type AdapterStore struct {
	wklog.Log
	httpClient *http.Client
	Host       string
}

// NewAdapterStore  创建API
func NewAdapterStore() *AdapterStore {
	return &AdapterStore{
		Log: wklog.NewWKLog("AdStore"),
		httpClient: &http.Client{
			Transport: &http.Transport{
				DialContext: (&net.Dialer{
					Timeout:   2 * time.Second,
					KeepAlive: 5 * time.Second,
				}).DialContext,
				ForceAttemptHTTP2:     true,
				MaxIdleConns:          200,
				MaxIdleConnsPerHost:   200,
				IdleConnTimeout:       300 * time.Second,
				TLSHandshakeTimeout:   time.Second * 3,
				ResponseHeaderTimeout: 2 * time.Second,
				ExpectContinueTimeout: 1 * time.Second,
			},
		},
		Host: os.Getenv("WK_ADAPTER_HOST"),
	}
}

func (as *AdapterStore) post(path string, data any) ([]byte, error) {
	url := fmt.Sprintf("%s%s", as.Host, path)
	startTime := time.Now().UnixMilli()
	as.Debug("数据服务开始请求", zap.String("url", url))
	resp, err := as.httpClient.Post(url, "application/json", strings.NewReader(wkutil.ToJson(data)))
	as.Debug("数据服务请求结束 耗时", zap.Int64("mill", time.Now().UnixMilli()-startTime))
	if err != nil {
		return nil, errors.Wrap(err, "AdapterStore http post error 1,url:"+url)
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	if resp.StatusCode != 200 {
		return nil, errors.Wrapf(err, "AdapterStore http post error 2,url:%v,response:%v", url, resp.StatusCode)
	}

	// 读取响应体
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrapf(err, "AdapterStore http post error 3,url:%v"+url)
	}
	// 返回响应体内容
	return responseBody, nil
}
