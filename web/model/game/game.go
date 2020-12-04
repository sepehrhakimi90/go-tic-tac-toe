package game

import (
	"github.com/sepehrhakimi90/go-tic-tac-toe/console/utils"
	"github.com/sepehrhakimi90/go-tic-tac-toe/web/model/board"
	"github.com/sepehrhakimi90/go-tic-tac-toe/web/model/player"
)

// Game is a type to manage game

type Game struct {
	Board          board.GameBoard  `json:"-"`
	Players        [2]player.Player `json:"Players"`
	Winner         int              `json:"winner"`
	CurrentPlayer  int              `json:"current_player"`
	CurrentMove    int              `json:"-"`
	AvailableMoves int              `json:"-"`
}

func (g *Game) checkWinner() {
	rules := utils.GetWinRules()
	cm := g.CurrentMove
	cs := g.Players[g.CurrentPlayer].Sign

	f := false
	for _, rs := range rules[cm] {
		f = false
		for _, s := range rs {
			if g.Board[s] != cs {
				f = true
			}
		}
		if f == false {
			g.Winner = g.CurrentPlayer
			break
		}
	}
}

// NextMove is a func to set next move of user
func (g *Game) NextMove(mv int) {
	p := g.Players[g.CurrentPlayer]
	g.Board[mv] = p.Sign
	g.CurrentMove = mv
	g.AvailableMoves--
	g.checkWinner()
	if g.CurrentPlayer == 0 {
		g.CurrentPlayer = 1
	} else {
		g.CurrentPlayer = 0
	}
}

// InitGame to setup game
func (g *Game) InitGame(firstPlayer, secondPlayer player.Player) {
	for i := 0; i < 9; i++ {
		g.Board[i] = ""
	}
	g.Players[0] = firstPlayer
	g.Players[1] = secondPlayer
	g.CurrentMove = -1
	g.Winner = -1
	g.AvailableMoves = 9
	g.CurrentPlayer = 0
}
