package api


type BroadCastSingle struct {
	Uid int64 `json:"uid"`
	Route string `json:"route"`
}

type SrvMsg struct {
	Content string `json:"content"`
}