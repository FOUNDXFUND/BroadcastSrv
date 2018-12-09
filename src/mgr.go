package src
import (
	"github.com/lonnng/nano"
	"time"
	"github.com/lonnng/nano/session"
	"fmt"
)


func (mgr *GameManager) Join(s *session.Session, msg []byte) error {
	fmt.Println("player join the game!!")
	fakeUID := s.ID()
	s.Bind(fakeUID)
	s.Set(GameIDKey, mgr.game)
	s.Push(Route_Member , &AllMembers{Members: mgr.game.group.Members()})
	// 广播通知其他用户
	mgr.game.group.Broadcast(Route_NewUser, &NewUser{Content: fmt.Sprintf("New user: %d", s.ID())})
	mgr.game.group.Add(s)
	return s.Response(&PublicResponse{Result: "连接服务器成功"})
}

func (mgr *GameManager) GetUserSession(uid int64) (*session.Session, error) {
	return mgr.game.group.Member(uid)
}

func (mgr *GameManager) Message(s *session.Session, msg *UserMessage) error {
	if !s.HasKey(GameIDKey) {
		return fmt.Errorf("not join room yet")
	}
	g := s.Value(GameIDKey).(*Game)
	return g.group.Broadcast("onMessage", msg)
}

func NewGameManager() *GameManager {
	g := &Game {
		group: nano.NewGroup(fmt.Sprintf("game-%d", GameID)),
	}

	return &GameManager{
		game: g,
	}
}


func (mgr *GameManager) AfterInit() {
	session.Lifetime.OnClosed(func(s *session.Session) {
		if !s.HasKey(GameIDKey) {
			return
		}

		g := s.Value(GameIDKey).(*Game)
		g.group.Leave(s)
	})

	mgr.timer = nano.NewTimer(time.Minute, func() {
		println(fmt.Sprintf("【用户统计】 GameId=%d, Time=%s, Count=%d", GameID, time.Now().String(), mgr.game.group.Count()))
	})
}