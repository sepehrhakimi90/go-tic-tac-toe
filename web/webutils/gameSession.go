package webutils

import (
	"encoding/gob"
	"github.com/gorilla/sessions"
	"github.com/sepehrhakimi90/go-tic-tac-toe/web/model/board"
	"github.com/sepehrhakimi90/go-tic-tac-toe/web/model/game"
	"github.com/sepehrhakimi90/go-tic-tac-toe/web/model/player"
	"os"
)

var Store *sessions.CookieStore

const (
	SESSION_COOKIE_NAME = "FAHsep"
)

type M map[string]interface{}

func init() {
	gob.Register(&player.Player{})
	gob.Register(&board.GameBoard{})
	gob.Register(&game.Game{})
	gob.Register(&M{})
	key := os.Getenv("SESSION_KEY")
	if key == "deflt-secret-key" {
		key = ""
	}
	Store = sessions.NewCookieStore([]byte(key))
	Store.Options.MaxAge = 3600
}
