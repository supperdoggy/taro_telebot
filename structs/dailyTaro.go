package structs

import "time"

type DailyTaro struct {
	UserID int64 `json:"user_id" bson:"user_id"`
	CardID string `json:"card_id" bson:"card_id"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
}
