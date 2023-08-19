//go:generate go run ../gencode/main.go
package errmsg

type Resp struct {
	Code int
	Msg  string
}

const (
	TypeMsgSuccess = "成功"
	TypeMsgFailed  = "失败"
)
