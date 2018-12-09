package api

import (
	"github.com/gin-gonic/gin"
	"gokuEx/broadcast/src"
)

type Api struct {
	GameMgr *src.GameManager
}

func (a *Api) BroadCastSingle(c *gin.Context) {
	req := &BroadCastSingle{}
	if err := c.ShouldBindJSON(req);err != nil {
		panic(ErrWrongParam)
	}

	s,err := a.GameMgr.GetUserSession(req.Uid)
	if err != nil {
		panic(err)
	}

	//s.Push(ROUTE_Member , &AllMembers{Members: mgr.game.group.Members()})
	s.Push(src.Route_SrvMsg, &SrvMsg{"哈哈哈哈"})
}
