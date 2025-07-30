package repositories_test

import (
	"context"
	"testing"

	"task_manager/Domain"
	"task_manager/Repositories"
	"task_manager/Repositories/models"

	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserRepositoryIntegrationSuite struct{
	suite.Suite
	ctx context.Context
	db *mongo.Database
	repo domain.IUserRepository
	client *mongo.Client
}

func (s *UserRepositoryIntegrationSuite) SetupSuite() {
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	s.Require().NoError(err)

	s.ctx = ctx
	s.client = client
	s.db = client.Database("taskstest")
	s.repo = repositories.NewUserRepository(s.db)
}

func (s *UserRepositoryIntegrationSuite) SetupTest() {
	_, err := s.db.Collection("users").DeleteMany(s.ctx, bson.M{})
	s.Require().NoError(err)
}

func (s *UserRepositoryIntegrationSuite) TearDownSuite() {
	err := s.client.Disconnect(s.ctx)
	s.Require().NoError(err)
}

func TestUserRepositoryIntegrationSuite(t *testing.T) {
	suite.Run(t, new(UserRepositoryIntegrationSuite))
}

func (s *UserRepositoryIntegrationSuite) TestRegister_Success() {
	user := &domain.User{
		Username:  "testname",
		Password:  "secret",
		Role: domain.RoleAdmin,
	}
	err := s.repo.Register(s.ctx, user)
	s.NoError(err)

	// Check directly from DB
	var userMongo models.UserMongoModel
	err = s.db.Collection("users").FindOne(s.ctx, bson.M{"username": "testname"}).Decode(&userMongo)
	s.NoError(err)
	s.Equal("testname", userMongo.Username)
}
func (s *UserRepositoryIntegrationSuite) TestLogin_Success() {
	// First insert user manually
	user := models.UserMongoModel{
		ID:        primitive.NewObjectID(),
		Username:  "testname",
		Password:  "hashed-secret",
		Role:      string(domain.RoleUser),
	}
	_, err := s.db.Collection("users").InsertOne(s.ctx, user)
	s.NoError(err)

	// Now test login
	domainUser := &domain.User{
		Username: "testname",
		Password: "any-pass", // Password is not validated here
	}
	result, err := s.repo.Login(s.ctx, domainUser)
	s.NoError(err)
	s.Equal("testname", result.Username)
}
func (s *UserRepositoryIntegrationSuite) TestRegister_UsernameTaken() {
	user := &domain.User{
		Username: "testname",
		Password: "123",
	}
	// First register
	err := s.repo.Register(s.ctx, user)
	s.NoError(err)

	// Second time should fail
	err = s.repo.Register(s.ctx, user)
	s.Error(err)
	s.Contains(err.Error(), "username already taken")
}
func (s *UserRepositoryIntegrationSuite) TestLogin_InvalidUser() {
	user := &domain.User{
		Username: "notfound",
		Password: "any",
	}
	_, err := s.repo.Login(s.ctx, user)
	s.Error(err)
	s.Contains(err.Error(), "invalid username")
}
