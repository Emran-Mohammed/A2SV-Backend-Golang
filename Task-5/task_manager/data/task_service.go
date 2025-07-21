package data

import (
	"context"
	"task_manager/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskManager struct{
	TaskCollection *mongo.Collection
	
}

func NewTaskManager(client *mongo.Client) *TaskManager{
	return &TaskManager{TaskCollection: client.Database("tasksdb").Collection("tasks")}
}

func (t *TaskManager) GetTasks(ctx context.Context) ([]models.Task, error){
	var books []models.Task
	cursor, err := t.TaskCollection.Find(ctx, bson.M{})
	if err != nil{
		return nil , err
	}
	defer cursor.Close(ctx)
	if err = cursor.All(ctx, &books); err!= nil{
		return nil,  err
	}
	return books, nil
	
}

func (t* TaskManager) GetTaskByID(ctx context.Context, id string)(*models.Task, error){
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil{
		return nil , err
	}
	var task models.Task
	err = t.TaskCollection.FindOne(ctx, bson.M{"_id": oid}).Decode(&task)
	return &task, err
}

func (t * TaskManager) UpdateTask(ctx context.Context, id string, updatedTask models.Task)( *mongo.UpdateResult, error){
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil{
		return nil, err
	}
	updatedTask.ID = oid
	return t.TaskCollection.ReplaceOne(ctx, bson.M{"_id":oid}, updatedTask)
	
}

func (t *TaskManager) DeleteTask (ctx context.Context, id string) error{
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil{
		return err
	}

	_, err = t.TaskCollection.DeleteOne(ctx, bson.M{"_id": oid})
	return err
}

func (t *TaskManager) CreateTask(ctx context.Context, task models.Task)(*mongo.InsertOneResult, error) {
	task.ID = primitive.NewObjectID()
	result, err := t.TaskCollection.InsertOne(ctx, task)
	return result, err

}