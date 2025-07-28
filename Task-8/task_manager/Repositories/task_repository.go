package repositories

import (
	"context"
	domain "task_manager/Domain"
	"task_manager/Repositories/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type taskRepository struct {
	taskCollection *mongo.Collection
}

func NewTaskRepository(db *mongo.Database)domain.ITaskRepository{
	return &taskRepository{taskCollection: db.Collection("tasks")}
}

func(tr *taskRepository) CreateTask(ctx context.Context, task *domain.Task)(*domain.Task, error){
	taskMongo, err := models.DomainToMongoTask(task)
	if err != nil{
		return nil, err
	}
	taskMongo.ID = primitive.NewObjectID()

	_ , err = tr.taskCollection.InsertOne(ctx, taskMongo)

	return models.MongoToDomainTask(taskMongo), err


}

func(tr *taskRepository) GetTasks(ctx context.Context) ([]domain.Task, error){
	var taskMongoList []models.TaskMongoModel
	cursor, err := tr.taskCollection.Find(ctx, bson.M{})
	if err != nil{
		return nil , err
	}
	defer cursor.Close(ctx)
	if err = cursor.All(ctx, &taskMongoList); err!= nil{
		return nil,  err
	}
	var domainTasks []domain.Task
	for _, mongoTask := range taskMongoList {
		task := models.MongoToDomainTask(&mongoTask)
		domainTasks = append(domainTasks, *task)
	}

	return domainTasks, nil

}
func(tr *taskRepository)GetTaskByID(ctx context.Context, id string)(*domain.Task, error){
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil{
		return nil , err
	}
	var task models.TaskMongoModel
	err = tr.taskCollection.FindOne(ctx, bson.M{"_id": oid}).Decode(&task)
	return models.MongoToDomainTask(&task), err
}
func(tr *taskRepository)UpdateTask(ctx context.Context, id string, updatedTask *domain.Task)(*domain.Task, error){
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil{
		return nil,err
	}
	taskMongo, err := models.DomainToMongoTask(updatedTask)
	if err != nil{
		return nil, err
	}
	taskMongo.ID = oid
	_, err = tr.taskCollection.ReplaceOne(ctx, bson.M{"_id":oid}, taskMongo)
	return models.MongoToDomainTask(taskMongo), err
}
func(tr *taskRepository)DeleteTask (ctx context.Context, id string) error{
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil{
		return err
	}

	_, err = tr.taskCollection.DeleteOne(ctx, bson.M{"_id": oid})
	return err
}