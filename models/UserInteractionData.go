package models

type UserInteractionData struct {
	TimStamp        string
	EventType       string `faker:"events"`
	UserId          string `faker:"user_id"`
	ProductId       string `faker:"prod_id"`
	SessionDuration string `faker:"sd"`
}
