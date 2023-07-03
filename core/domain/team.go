package domain

type Team struct {
	ID      int       `json:"id"`
	Name    string    `json:"name"`
	Coach   string    `json:"coach"`
	Players []*Player `json:"players,omitempty"`
}
