package models

type UserInteractionData struct {
	TimeStamp       string
	EventType       string  `faker:"events"`
	UserId          string  `faker:"user_id"`
	ProductId       string  `faker:"prod_id"`
	SessionDuration uint    `faker:"sd"`
}
