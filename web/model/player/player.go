package player


type Player struct {
	Name string `json:"name"`
	Sign string `json:"sign"`
}

func GetPlayer(name, sign string) Player {
	return Player{
		Name: name,
		Sign: sign,
	}
}
