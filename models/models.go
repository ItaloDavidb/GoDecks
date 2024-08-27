package models

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Card struct {
	SetCode  string `json:"set_code"`
	Number   string `json:"number"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	JsonData string `json:"json_data"`
}

type UserCard struct {
	UserID   int    `json:"user_id"`
	SetCode  string `json:"set_code"`
	Number   string `json:"number"`
	Quantity int    `json:"quantity"`
}

type Deck struct {
	ID     int    `json:"id"`
	UserID int    `json:"user_id"`
	Name   string `json:"name"`
}

type DeckCard struct {
	DeckID   int    `json:"deck_id"`
	SetCode  string `json:"set_code"`
	Number   string `json:"number"`
	Quantity int    `json:"quantity"`
}
