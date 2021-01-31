package user

import (
	"log"

	"github.com/cagox/ocfsocial/app/util/config"
	"github.com/cagox/ocfsocial/app/util/crypto"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const collectionName = "ocfUsers"

//GetUsers returns a list of all the users in the database.
func GetUsers() []User {
	var users []User

	collection := config.Config.MongoClient.Database(config.Config.DatabaseName).Collection(collectionName)

	cursor, err := collection.Find(config.Config.MongoContext, bson.D{})
	if err != nil {
		config.Config.Logger.Println(err) //TODO: Add proper error handling.
	}

	for cursor.Next(config.Config.MongoContext) {
		result := User{}
		err := cursor.Decode(result)
		if err != nil {
			config.Config.Logger.Println(err)
		}

		users = append(users, result)

	}

	return users
}

//AreThereAnyUsers checks to see if the database has any users or not.
func AreThereAnyUsers() bool {
	collection := config.Config.MongoClient.Database(config.Config.DatabaseName).Collection(collectionName)
	count, err := collection.CountDocuments(config.Config.MongoContext, bson.D{})

	if err != nil {
		config.Config.Logger.Println(err)
	}

	if count == 0 {
		return false
	}
	return true
}

//GetUserByEmail grabs a user object from the database based on the email address.
func GetUserByEmail(email string) *User {
	collection := config.Config.MongoClient.Database(config.Config.DatabaseName).Collection(collectionName)

	var user *User
	user = new(User)

	err := collection.FindOne(config.Config.MongoContext, bson.D{{"email", email}}).Decode(user)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil
		}
		log.Fatal(err)
	}
	return user
}

//SetPassword sets the password on the user object
func (user *User) SetPassword(password string) {
	user.PasswordHash = crypto.HashPassword(password)
}

//Authenticate allows the login method to make sure we have the right person.
func (user *User) Authenticate(password string) bool {
	return crypto.ComparePassword(password, user.PasswordHash)
}

//CreateUserFromForm creates a new user object from the data provided via a CreateUserForm object.
func CreateUserFromForm(newUser CreateUserForm) *User {
	user := &User{Email: newUser.Email, PasswordHash: crypto.HashPassword(newUser.Password)}
	return user
}

//IsEmailUnique lets you verify if a user exists in the database already. False means they are there.ss
func IsEmailUnique(email string) bool {
	user := GetUserByEmail(email)

	if user != nil {
		return false
	}
	return true
}

//InsertUser adds the user to the database.
func InsertUser(user *User) error {
	//TODO: Verify uniqueness and add error handling.
	collection := config.Config.MongoClient.Database(config.Config.DatabaseName).Collection(collectionName)

	_, err := collection.InsertOne(config.Config.MongoContext, user)
	if err != nil {
		return err
	}
	return nil
}

//TODO: add code to update an existing user.
//TODO: Review and decide if I need to entirely rewrite this page.
