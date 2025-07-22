package data

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"

	"context"
	"errors"
	"task_manager/config"
	"task_manager/models"
)



type UserManager struct {
	UserCollection *mongo.Collection
}

func NewUserManager(client *mongo.Client)*UserManager{
	return &UserManager{UserCollection: client.Database("taskdb").Collection("users")}
}


func (u *UserManager) Register(ctx context.Context, user models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil{
		return err
	}
	count, err := u.UserCollection.CountDocuments(ctx, bson.M{"username": user.Username})
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("username already taken")
	}
	if user.Role == "" {
	user.Role = models.RoleUser
	}
	user.ID = primitive.NewObjectID()
	user.Password = string(hashedPassword)
	_, err = u.UserCollection.InsertOne(ctx, user)
	return err

}
func (u *UserManager)Login(ctx context.Context, user models.User)(*string,error){

	// var jwtSecret = []byte("my_jwt_secret")

	var existingUser models.User

	err := u.UserCollection.FindOne(ctx, bson.M{"username": user.Username}).Decode(&existingUser)
    if err != nil {
        return nil, errors.New("invalid username or password")
    }
	
	if bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(user.Password)) != nil{
		return nil, errors.New("invalid username or password")

	}

	token:=jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":existingUser.ID.Hex(),
		"username":existingUser.Username,
		"role":existingUser.Role,
		"exp":time.Now().Add(5 * time.Minute).Unix(),

	})

	jwtToken, err := token.SignedString(config.JWTSecret)
	if err != nil{
		return nil, err
	}
	return &jwtToken, nil

}