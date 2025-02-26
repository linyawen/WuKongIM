package adapter

// ComResp 调用wkAdapter rest api 时的通用返回结构
type ComResp[T any] struct {
	RequestId string // 请求id,用于追踪链路，后期整合apm产品，例如 skyWalking
	Path      string // 请求接口路径
	DATA      T      // 返回的业务数据对象
	Online    bool   // 是否本次调用是否上线，适合试验阶段使用
	Code      int    // 业务成功 200, 异常 500
	Message   string // 异常描述
}

func (c *ComResp[T]) Success() bool {
	return c.Code == 200
}

func (c *ComResp[T]) Error() bool {
	return !c.Success()
}
