package structs

import "time"

type TaroPic struct {
	ID string `json:"id" bson:"_id"`
	Data []byte `json:"data" bson:"data"`
	Created time.Time `json:"created" bson:"created"`
	Updated time.Time `json:"updated" bson:"updated"`
}
