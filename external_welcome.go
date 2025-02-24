package wework

import (
	"github.com/golanggo/hzs-wecom/internal"
)

type WelcomeMessage struct {
	WelcomeCode string              `json:"welcome_code"`
	Text        ExternalText        `json:"text,omitempty"`
	Image       ExternalImage       `json:"image,omitempty"`
	Link        ExternalLink        `json:"link,omitempty"`
	Miniprogram ExternalMiniprogram `json:"miniprogram,omitempty"`
	Video       ExternalVideo       `json:"video,omitempty"`
	File        ExternalFile        `json:"file,omitempty"`
	Agentid     int                 `json:"agentid"`
	Notify      int                 `json:"notify"`
}

// SendWelcomeMsg 发送新客户欢迎语
// https://open.work.weixin.qq.com/api/doc/90001/90143/92599
func (ww *weWork) SendWelcomeMsg(corpId uint, msg WelcomeMessage) (resp internal.BizResponse) {
	if ok := validate.Struct(msg); ok != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = ok.Error()
		return
	}
	_, err := ww.getRequest(corpId).SetBody(msg).SetResult(&resp).
		Post("/cgi-bin/externalcontact/send_welcome_msg")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}
