package src

import (
	"github.com/lonnng/nano"
	"github.com/lonnng/nano/component"
)

const (
	GameID = 1
	GameIDKey  = "IDLE_GAME"
)

type (
	Game struct {
		group *nano.Group
	}

	GameManager struct {
		component.Base
		timer *nano.Timer
		game *Game
	}

	// UserMessage represents a message that user sent
	UserMessage struct {
		Name    string `json:"name"`
		Content string `json:"content"`
	}

	// NewUser message will be received when new user join room
	NewUser struct {
		Content string `json:"content"`
	}

	// AllMembers contains all members uid
	AllMembers struct {
		Members []int64 `json:"members"`
	}

	// JoinResponse represents the result of joining room
	PublicResponse struct {
		Code   int    `json:"code"`
		Result string `json:"result"`
	}
)
