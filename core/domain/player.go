package domain

type Player struct {
	ID       int    `json:"id"`
	TeamID   int    `json:"team_id"`
	Name     string `json:"name"`
	Number   int    `json:"number"`
	Position string `json:"position"`
	Nation   string `json:"nation"`
}
