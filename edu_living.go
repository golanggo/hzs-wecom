package wework

import (
	"github.com/golanggo/hzs-wecom/internal"
)

type GetUserAllLivingIdRequest struct {
	UserId string `json:"userid" validate:"required"`
	Cursor string `json:"cursor,omitempty"`
	Limit  int    `json:"limit"`
}

type GetUserAllLivingIdResponse struct {
	internal.BizResponse
	NextCursor   string   `json:"next_cursor"`
	LivingIdList []string `json:"livingid_list"`
}

// GetUserAllLivingId 获取老师直播ID列表
// https://open.work.weixin.qq.com/api/doc/90001/90143/93856
func (ww *weWork) GetUserAllLivingId(corpId uint, request GetUserAllLivingIdRequest) (resp GetUserAllLivingIdResponse) {
	if ok := validate.Struct(request); ok != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = ok.Error()
		return
	}
	_, err := ww.getRequest(corpId).SetBody(request).SetResult(&resp).
		Post("/cgi-bin/living/get_user_all_livingid")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}

type GetLivingInfoResponse struct {
	internal.BizResponse
	LivingInfo struct {
		Theme          string `json:"theme"`
		LivingStart    int    `json:"living_start"`
		LivingDuration int    `json:"living_duration"`
		AnchorUserId   string `json:"anchor_userid"`
		LivingRange    struct {
			PartyIds   []int    `json:"partyids"`
			GroupNames []string `json:"group_names"`
		} `json:"living_range"`
		ViewerNum     int    `json:"viewer_num"`
		CommentNum    int    `json:"comment_num"`
		OpenReplay    int    `json:"open_replay"`
		PushStreamURL string `json:"push_stream_url"`
	} `json:"living_info"`
}

// GetLivingInfo 获取直播详情
// https://open.work.weixin.qq.com/api/doc/90001/90143/93857
func (ww *weWork) GetLivingInfo(corpId uint, liveId string) (resp GetLivingInfoResponse) {
	_, err := ww.getRequest(corpId).SetResult(&resp).
		SetQueryParam("livingid", liveId).Get("/cgi-bin/school/living/get_living_info")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}

type GetWatchStatRequest struct {
	LivingId string `json:"livingid" validate:"required"`
	NextKey  string `json:"next_key"`
}

type GetWatchStatResponse struct {
	internal.BizResponse
	Ending     int    `json:"ending"`
	NextKey    string `json:"next_key"`
	StatInfoes struct {
		Students []struct {
			StudentUserid string `json:"student_userid"`
			ParentUserid  string `json:"parent_userid"`
			Partyids      []int  `json:"partyids"`
			WatchTime     int    `json:"watch_time"`
			EnterTime     int    `json:"enter_time"`
			LeaveTime     int    `json:"leave_time"`
			IsComment     int    `json:"is_comment"`
		} `json:"students"`
		Visitors []struct {
			Nickname  string `json:"nickname"`
			WatchTime int    `json:"watch_time"`
			EnterTime int    `json:"enter_time"`
			LeaveTime int    `json:"leave_time"`
			IsComment int    `json:"is_comment"`
		} `json:"visitors"`
	} `json:"stat_infoes"`
}

// GetWatchStat 获取观看直播统计
// https://open.work.weixin.qq.com/api/doc/90001/90143/93858
func (ww *weWork) GetWatchStat(corpId uint, request GetWatchStatRequest) (resp GetWatchStatResponse) {
	if ok := validate.Struct(request); ok != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = ok.Error()
		return
	}
	_, err := ww.getRequest(corpId).SetBody(request).SetResult(&resp).
		Post("/cgi-bin/school/living/get_watch_stat")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}

type GetUnWatchStatResponse struct {
	internal.BizResponse
	Ending   int    `json:"ending"`
	NextKey  string `json:"next_key"`
	StatInfo struct {
		Students []struct {
			StudentUserid string `json:"student_userid"`
			ParentUserid  string `json:"parent_userid"`
			Partyids      []int  `json:"partyids"`
		} `json:"students"`
	} `json:"stat_info"`
}

// GetUnWatchStat 获取未观看直播统计
// https://open.work.weixin.qq.com/api/doc/90001/90143/93859
func (ww *weWork) GetUnWatchStat(corpId uint, request GetWatchStatRequest) (resp GetUnWatchStatResponse) {
	if ok := validate.Struct(request); ok != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = ok.Error()
		return
	}
	_, err := ww.getRequest(corpId).SetBody(request).SetResult(&resp).
		Post("/cgi-bin/school/living/get_unwatch_stat")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}

// DeleteReplayData 删除直播回放
// https://open.work.weixin.qq.com/api/doc/90001/90143/93860
func (ww *weWork) DeleteReplayData(corpId uint, livingId string) (resp internal.BizResponse) {
	h := H{}
	h["livingid"] = livingId
	_, err := ww.getRequest(corpId).SetBody(h).SetResult(&resp).
		Post("/cgi-bin/living/delete_replay_data")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}

type LivingCreateRequest struct {
	AnchorUserID         string             `json:"anchor_userid" validate:"required"`
	Theme                string             `form:"theme" json:"theme"`
	LivingStart          uint32             `form:"living_start" json:"living_start"`
	LivingDuration       string             `form:"living_duration" json:"living_duration"`
	Type                 int                `form:"type" json:"type"`
	Description          string             `form:"description" json:"description"`
	AgentID              int                `form:"agent_id" json:"agent_id"`
	RemindTime           uint32             `form:"remind_time" json:"remind_time"`
	ActivityCoverMediaID string             `form:"activity_cover_media_id" json:"activity_cover_media_id"`
	ActivityDetail       ActivityDetailInfo `form:"activity_detail" json:"activity_detail"`
}
type ActivityDetailInfo struct {
	Description string   `json:"title"`
	ImageList   []string `json:"image_list"`
}

type LivingCreateResponse struct {
	internal.BizResponse
	LivingID string `json:"livingid"`
}

// CreateLiving 创建预约直播
// https://developer.work.weixin.qq.com/document/path/93717
func (ww *weWork) LivingCreate(corpId uint, request LivingCreateRequest) (resp LivingCreateResponse) {
	h := H{}
	h["anchor_userid"] = request.AnchorUserID
	h["theme"] = request.Theme
	h["living_start"] = request.LivingStart
	h["living_duration"] = request.LivingDuration
	h["type"] = request.Type
	h["description"] = request.Description
	h["agent_id"] = request.AgentID
	h["remind_time"] = request.RemindTime
	h["activity_cover_media_id"] = request.ActivityCoverMediaID
	h["activity_detail"] = request.ActivityDetail
	_, err := ww.getRequest(corpId).SetBody(h).SetResult(&resp).
		Post("/cgi-bin/living/create")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}

type GetLivingCodeCreateResponse struct {
	internal.BizResponse
	LivingCode string `json:"living_code"`
}

// GetLivingCode 获取微信观看直播凭证
// https://developer.work.weixin.qq.com/document/path/93641
func (ww *weWork) GetLivingCode(corpId uint, live_id, openid string) (resp GetLivingCodeCreateResponse) {
	h := H{}
	h["livingid"] = live_id
	h["openid"] = openid
	_, err := ww.getRequest(corpId).SetBody(h).SetResult(&resp).
		Post("/cgi-bin/living/get_living_code")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}

// LivingCancel 取消预约直播
// https://developer.work.weixin.qq.com/document/path/93718
func (ww *weWork) LivingCancel(corpId uint, livingID string) (resp internal.BizResponse) {
	p := H{"livingid": livingID}
	_, err := ww.getRequest(corpId).SetBody(p).SetResult(&resp).
		Post("/cgi-bin/living/cancel")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}

type LivingGetWatchStatResponse struct {
	internal.BizResponse
	Ending   int    `json:"ending,omitempty"`   // 结束标志
	NextKey  string `json:"next_key,omitempty"` // 下一页的键
	StatInfo struct {
		Users []struct {
			UserID    string `json:"userid,omitempty"`     // 用户ID
			WatchTime int    `json:"watch_time,omitempty"` // 观看时间
			IsComment int    `json:"is_comment,omitempty"` // 是否评论
			IsMic     int    `json:"is_mic,omitempty"`     // 是否开麦
		} `json:"users,omitempty"` // 用户统计信息
		ExternalUsers []struct {
			ExternalUserID string `json:"external_userid,omitempty"` // 外部用户ID
			Type           int    `json:"type,omitempty"`            // 用户类型
			Name           string `json:"name,omitempty"`            // 用户名
			WatchTime      int    `json:"watch_time,omitempty"`      // 观看时间
			IsComment      int    `json:"is_comment,omitempty"`      // 是否评论
			IsMic          int    `json:"is_mic,omitempty"`          // 是否开麦
		} `json:"external_users,omitempty"` // 外部用户统计信息
	} `json:"stat_info,omitempty"` // 统计信息
}

// LivingGetWatchStat 获取直播观看明细
// https://developer.work.weixin.qq.com/document/path/96836
func (ww *weWork) LivingGetWatchStat(corpId uint, liveId, nextKey string) (resp LivingGetWatchStatResponse) {
	p := H{"livingid": liveId, "next_key": nextKey}
	_, err := ww.getRequest(corpId).SetBody(p).SetResult(&resp).
		Post("/cgi-bin/living/get_watch_stat")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}

type LivingGetLivingInfoResponse struct {
	internal.BizResponse
	LivingInfo struct {
		Theme                 string `json:"theme"`
		LivingStart           int    `json:"living_start"`
		LivingDuration        int    `json:"living_duration"`
		Status                int    `json:"status"`
		ReserveStart          int    `json:"reserve_start"`
		ReserveLivingDuration int    `json:"reserve_living_duration"`
		Description           string `json:"description"`
		AnchorUserid          string `json:"anchor_userid"`
		MainDepartment        int    `json:"main_department"`
		ViewerNum             int    `json:"viewer_num"`
		CommentNum            int    `json:"comment_num"`
		MicNum                int    `json:"mic_num"`
		OpenReplay            int    `json:"open_replay"`
		ReplayStatus          int    `json:"replay_status"`
		Type                  int    `json:"type"`
		PushStreamUrl         string `json:"push_stream_url"`
		OnlineCount           int    `json:"online_count"`
		SubscribeCount        int    `json:"subscribe_count"`
	} `json:"living_info"`
}

// LivingGetLivingInfo 获取直播详情
// https://developer.work.weixin.qq.com/document/path/93635
func (ww *weWork) LivingGetLivingInfo(corpId uint, liveId string) (resp LivingGetLivingInfoResponse) {
	_, err := ww.getRequest(corpId).SetResult(&resp).
		SetQueryParam("livingid", liveId).Get("/cgi-bin/living/get_living_info")
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	}
	return
}
