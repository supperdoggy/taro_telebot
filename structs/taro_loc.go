package structs

import "time"

type TaroLoc struct {
	ID      string    `json:"id" bson:"_id"`
	Value   string    `json:"value" bson:"value"`
	Created time.Time `json:"created" bson:"created"`
	Updated time.Time `bson:"updated" json:"updated"`
}
