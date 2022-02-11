package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name     string             `json:"name" bson:"name"`
	Email    string             `json:"email" bson:"email"`
	Password string             `json:"-" bson:"password"`
	CartID   primitive.ObjectID `json:"cartID" bson:"cartID"`
}

type LoginUser struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name     string             `json:"name" bson:"name"`
	Email    string             `json:"email" bson:"email"`
	Password string             `json:"password" bson:"password"`
	CartID   primitive.ObjectID `json:"cartID" bson:"cartID"`
}

type UserInfo struct {
	Name  string `json:"name" bson:"name"`
	Email string `json:"email" bson:"email"`
}
