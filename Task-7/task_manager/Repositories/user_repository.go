package repositories

import (
	"context"
	"errors"
	domain "task_manager/Domain"
	"task_manager/Repositories/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepository struct {
	userCollection *mongo.Collection
}

func NewUserRepository(db *mongo.Database)domain.IUserRepository{
	return &userRepository{userCollection: db.Collection("users")}
}


func (ur *userRepository)Register(ctx context.Context, user *domain.User) error {
	userMongo, err := models.DomainToMongoUser(user)
	if err != nil{
		return err
	}
	count, err := ur.userCollection.CountDocuments(ctx, bson.M{"username": userMongo.Username})
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("username already taken")
	}
	if userMongo.Role == "" {
	userMongo.Role = string(domain.RoleUser)
	}
	userMongo.ID = primitive.NewObjectID()
	_, err = ur.userCollection.InsertOne(ctx, userMongo)
	return err

}
func (ur *userRepository)Login(ctx context.Context, user *domain.User) (*domain.User, error){
	userInput, err := models.DomainToMongoUser(user)
	if err != nil{
		return nil, err
	}
	var userMongo models.UserMongoModel
	err = ur.userCollection.FindOne(ctx, bson.M{"username": userInput.Username}).Decode(&userMongo)
    if err != nil {
        return nil, errors.New("invalid username or password")
    }
	return models.MongoToDomainUser(&userMongo), nil

}