package models

type User struct {
	ID        int        `gorm:"primaryKey"`
	Username  string     `json:"username"`
	Email     string     `json:"email"`
	Password  string     `json:"password"`
	UserCards []UserCard `gorm:"foreignKey:UserID"` // Relação com UserCard
	Decks     []Deck     `gorm:"foreignKey:UserID"` // Relação com Deck
}

type Card struct {
	SetCode  string `gorm:"primaryKey"`
	Number   string `gorm:"primaryKey"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	JsonData string `json:"json_data"`
	// Não precisa referenciar UserCard ou DeckCard
}

type UserCard struct {
	UserID   int    `gorm:"primaryKey;not null"`
	SetCode  string `gorm:"primaryKey;not null"`
	Number   string `gorm:"primaryKey;not null"`
	Quantity int    `json:"quantity"`
	User     User   `gorm:"foreignKey:UserID"`         // Relação com User
	Card     Card   `gorm:"foreignKey:SetCode,Number"` // Relação com Card
}

type Deck struct {
	ID        int        `json:"id"`
	UserID    int        `json:"user_id"`
	Name      string     `json:"name"`
	User      User       `gorm:"foreignKey:UserID"` // Relação com User
	DeckCards []DeckCard `gorm:"foreignKey:DeckID"` // Relação com DeckCard
}

type DeckCard struct {
	DeckID   int    `gorm:"primaryKey;not null"`
	SetCode  string `gorm:"primaryKey;not null"`
	Number   string `gorm:"primaryKey;not null"`
	Quantity int    `json:"quantity"`
	Deck     Deck   `gorm:"foreignKey:DeckID"`         // Relação com Deck
	Card     Card   `gorm:"foreignKey:SetCode,Number"` // Relação com Card
}
