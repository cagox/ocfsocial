package user

import (
	"encoding/gob"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

var Version = 1

//User is meant to hold user related information in the Database.
type User struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Email        string
	PasswordHash string
	IsAdmin      bool
	Timestamp    time.Time
	LastUpdated  time.Time
}

//CreateUserForm is a struct to facilitate creating user objects.
type CreateUserForm struct {
	Email    string
	Password string
}

func init() {
	gob.Register(User{})
}
