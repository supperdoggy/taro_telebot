package structs

import "time"

type TaroMeaning struct {
	ID string `json:"id" bson:"_id"`
	Value string `json:"value" bson:"value"`
	Created time.Time `json:"created" bson:"created"`
	Updated time.Time `json:"updated" bson:"updated"`
}
