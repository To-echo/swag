package model3

// Response 响应结构体
type Response struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Msg     string      `json:"msg"`
	TraceId string      `json:"trace_id"`
}
