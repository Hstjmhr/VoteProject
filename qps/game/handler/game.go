package handler

import (
	"core/repo"
	"core/service"
	"encoding/json"
	"fmt"
	"framework/remote"
	"game/logic"
	"game/model/request"
	common "msqp"
	"msqp/biz"
)

type GameHandler struct {
	um          *logic.UnionManager
	userService *service.UserService
}

func (h *GameHandler) RoomMessageNotify(session *remote.Session, msg []byte) any {
	if len(session.GetUid()) <= 0 {
		return common.F(biz.InvalidUsers)
	}
	var req request.RoomMessageReq
	if err := json.Unmarshal(msg, &req); err != nil {
		return common.F(biz.RequestDataError)
	}
	roomId, ok := session.Get("roomId")
	if !ok {
		return common.F(biz.NotInRoom)
	}
	rm := h.um.GetRoomById(fmt.Sprintf("%v", roomId))
	if rm == nil {
		return common.F(biz.NotInRoom)
	}
	rm.RoomMessageHandler(session, req)

	return nil
}

func (h *GameHandler) GameMessageNotify(session *remote.Session, msg []byte) any {
	if len(session.GetUid()) <= 0 {
		return common.F(biz.InvalidUsers)
	}
	roomId, ok := session.Get("roomId")
	if !ok {
		return common.F(biz.NotInRoom)
	}
	rm := h.um.GetRoomById(fmt.Sprintf("%v", roomId))
	if rm == nil {
		return common.F(biz.NotInRoom)
	}
	rm.GameMessageHandle(session, msg)
	return nil
}

func NewGameHandler(r *repo.Manager, um *logic.UnionManager) *GameHandler {
	return &GameHandler{
		um:          um,
		userService: service.NewUserService(r),
	}
}
