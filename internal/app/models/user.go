package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	GUID         string             `bson:"guid,omitempty"`
	RefreshToken string             `bson:"refresh_token,omitempty"`
}

func (u *User) model() {

}
