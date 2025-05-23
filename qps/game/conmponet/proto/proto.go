package proto

type GameRule struct {
	AddScores      []int `json:"addScores"`      //加注分
	BaseScore      int   `json:"baseScore"`      //底分
	Bureau         int   `json:"bureau"`         //局数
	CanEnter       bool  `json:"canEnter"`       //中途进人
	CanTrust       bool  `json:"canTrust"`       //允许托管
	CanWatch       bool  `json:"canWatch"`       //允许观战
	Cuopai         bool  `json:"cuopai"`         //高级 是否允许搓牌
	Fangzuobi      bool  `json:"fangzuobi"`      //防作弊
	Yuyin          bool  `json:"yuyin"`          //语音
	GameFrameType  int   `json:"gameFrameType"`  //游戏模式
	GameType       int   `json:"gameType"`       //游戏类型 牛牛 三公等
	MaxPlayerCount int   `json:"maxPlayerCount"` //最大人数
	MinPlayerCount int   `json:"minPlayerCount"` //最小人数
	MaxScore       int   `json:"maxScore"`       //最大加注分
	RoundType      int   `json:"roundType"`      //轮数
	PayDiamond     int   `json:"payDiamond"`     //房费
	PayType        int   `json:"payType"`        //支付方式 1 AA支付 2 赢家支付 3 我支付
	RoomType       int   `json:"roomType"`       // 1 正常房间 2 持续房间 3 百人房间
}

type GameType int
type SendCardType int
type GameFrameType int
type ScaleType int

const (
	PinSanZhang GameType = 1
	NiuNiu               = 2
	PaoDeKuai            = 3
	SanGong              = 4
	HongZhong            = 5
	DouGongNiu           = 8
)

type RoomMessageType int

const (
	UserReadyNotify             RoomMessageType = 301 // 用户准备的通知
	UserReadyPush                               = 401 // 用户准备的推送
	UserLeaveRoomNotify                         = 303 // 用户离开房间的通知
	UserLeaveRoomResponse                       = 403 //用户离开房间的回复
	UserLeaveRoomPush                           = 404 //用户离开房间的推送
	OtherUserEntryRoomPush                      = 402 // 用户进入房间的推送
	DismissPush                                 = 405 //房间解散的推送
	UserInfoChangePush                          = 406 //房间用户信息变化的推送
	UserChatNotify                              = 307 // 用户聊天通知
	UserChatPush                                = 407 // 用户聊天推送
	UserOffLinePush                             = 408 //用户掉线的推送
	DrawFinishedPush                            = 409 //开设的房间局数用完推送
	UserReconnectNotify                         = 312 //玩家断线重连
	UserReconnectPush                           = 412 //
	AskForDismissNotify                         = 313 //玩家请求解散房间
	AskForDismissPush                           = 413 //
	EndPush                                     = 414 //最终结果推送
	AskForDismissStatusNotify                   = 316
	AskForDismissStatusPush                     = 416
	GetRoomShowUserInfoNotify                   = 317 // 获取房间需要显示的玩家信息通知
	GetRoomShowUserInfoPush                     = 417 // 获取房间需要显示的玩家信息推送
	GetRoomSceneInfoNotify                      = 318 // 获取房间场景信息的通知
	GetRoomSceneInfoPush                        = 418 // 获取房间场景信息的推送
	GetRoomOnlineUserInfoNotify                 = 319 // 获取房间在线用户信息的通知
	GetRoomOnlineUserInfoPush                   = 419
	UserChangeSeatNotify                        = 320 //换座通知
	UserChangeSeatPush                          = 420
)

func UpdateUserInfoPush(roomId string) any {
	pushMsg := map[string]any{
		"roomID":     roomId,
		"pushRouter": "UpdateUserInfoPush",
	}
	return pushMsg
}

func UserLeaveRoomPushData(roomUserInfo *RoomUser) any {
	pushMsg := map[string]any{
		"type": UserLeaveRoomPush,
		"data": map[string]any{
			"roomUserInfo": roomUserInfo,
		},
		"pushRouter": "RoomMessagePush",
	}
	return pushMsg
}

func UserReadyPushData(chairID int) any {
	pushMsg := map[string]any{
		"type": UserReadyPush,
		"data": map[string]any{
			"chairID": chairID,
		},
		"pushRouter": "RoomMessagePush",
	}
	return pushMsg
}

func OtherUserEntryRoomPushData(roomUserInfo *RoomUser) any {
	pushMsg := map[string]any{
		"type": OtherUserEntryRoomPush,
		"data": map[string]any{
			"roomUserInfo": roomUserInfo,
		},
		"pushRouter": "RoomMessagePush",
	}
	return pushMsg
}
