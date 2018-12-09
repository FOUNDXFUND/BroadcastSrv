package main

import (
	"github.com/lonnng/nano"
	"github.com/lonnng/nano/component"
	"strings"
	"github.com/lonnng/nano/serialize/json"
	"gokuEx/broadcast/src"
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
	api2 "gokuEx/broadcast/api"
)

// nano中有四种消息类型的消息
// 分别是请求(Request), 响应(Response),通知(Notify)和推送(Push)
// 客户端发起Request到服务器端,服务器端处理后会给其返回响应Response;
// Notify是客户端发给服务端的通知,也就是不需要服务端给予回复的请求
// Push是服务端主动给客户端推送消息的类型


func main() {
	nano.SetSerializer(json.NewSerializer())
	gameMgr := src.NewGameManager()
	nano.Register(gameMgr,
		component.WithName("game"),
		component.WithNameFunc(strings.ToLower),
	)

	pipeline := nano.NewPipeline()
	nano.EnableDebug()
	log.SetFlags(log.LstdFlags | log.Llongfile)
	nano.SetWSPath("/ws")
	http.Handle("/web/", http.StripPrefix("/web/", http.FileServer(http.Dir("test/web"))))
	nano.SetCheckOriginFunc(func(_ *http.Request) bool { return true })
	go nano.ListenWS(":3250", nano.WithPipeline(pipeline))

	// inner http api
	api := &api2.Api{GameMgr:gameMgr}
	router := gin.Default()
	router.POST("/broadcast/single", api.BroadCastSingle)
	router.Run(":3000")
}


