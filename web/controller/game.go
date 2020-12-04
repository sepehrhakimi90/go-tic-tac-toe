package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sepehrhakimi90/go-tic-tac-toe/web/model/game"
	"github.com/sepehrhakimi90/go-tic-tac-toe/web/model/player"
	"github.com/sepehrhakimi90/go-tic-tac-toe/web/webutils"
	"net/http"
	"strconv"
)

func NewGame(c *gin.Context) {
	firstPlayer := player.GetPlayer(c.DefaultPostForm("FirstPlayer", "Player-1"), "X")
	secondPlayer := player.GetPlayer(c.DefaultPostForm("SecondPlayer", "Player-2"), "O")
	fmt.Println(firstPlayer)
	fmt.Println(secondPlayer)
	g := game.Game{}
	g.InitGame(firstPlayer, secondPlayer)
	session, err := webutils.Store.Get(c.Request, webutils.SESSION_COOKIE_NAME)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":"internal session issue"})
		return
	}
	session.Values["game"] = g
	session.Save(c.Request, c.Writer)
	c.JSON(http.StatusOK, gin.H{
		"Data": map[string]interface{}{
			"Players":g.Players,
		},
	})
}

func Move(c *gin.Context) {

	selectedCell, err := strconv.Atoi(c.PostForm("selectedCell"))
	if err != nil || (selectedCell < 0 && selectedCell > 8) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request parameter"})
		return
	}

	session, err := webutils.Store.Get(c.Request, webutils.SESSION_COOKIE_NAME)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":"internal session issue"})
		return
	}
	val := session.Values["game"]
	var g = &game.Game{}
	ok := true
	if g, ok = val.(*game.Game); !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error":"game retrive issue"})
		fmt.Println(g)
		return
	}

	g.NextMove(selectedCell)
	g.Board.Print()

	session.Values["game"] = g
	session.Save(c.Request, c.Writer)
	//if g.Winner != -1 {
	//	c.JSON(http.StatusOK, gin.H{"winner": g.Players[g.Winner].Name, "sign": g.Players[g.CurrentPlayer].Sign})
	//}
	c.JSON(http.StatusOK, gin.H{"data": g})
}

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}
