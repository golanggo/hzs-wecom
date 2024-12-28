package logic

import "github.com/golanggo/hzs-wecom"

func CancelAuthEventLogic(data []byte, ww wework.IWeWork) {
	ww.Logger().Sugar().Info(string(data))
}
