package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Token struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	UserId string             `bson:"guid"`
}

func (t *Token) model() {

}
