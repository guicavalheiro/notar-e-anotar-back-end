package tasks

import (
	"context"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
)

type Tasks struct {
	// DefaultModel adds _id, created_at and updated_at fields to the Model
	mgm.DefaultModel `bson:",inline"`
	Name             string `json:"name" bson:"name"`
	Description      string `json:"description" bson:"description"`
	ParentRating     int    `json:"parent_rating" bson:"parent_rating"`
	ChildrenRating   int    `json:"children_rating" bson:"children_rating"`
	Challenge        bool   `json:"challenge" bson:"challenge"`
}

func create(ctx context.Context, task *Tasks) (Tasks, error) {
	coll := mgm.Coll(task)

	err := coll.CreateWithCtx(ctx, task)

	if err != nil {
		return Tasks{}, err
	}

	return *task, nil
}

func Delete(ctx context.Context, id interface{}) error {
	task := &Tasks{}
	coll := mgm.Coll(task)

	err := coll.FindByIDWithCtx(ctx, id, task)
	if err != nil {
		return err
	}

	err = coll.DeleteWithCtx(ctx, task)
	if err != nil {
		return err
	}

	return nil
}

func findAll(ctx context.Context) ([]Tasks, error) {
	coll := mgm.Coll(&Tasks{})

	result, err := coll.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	task := make([]Tasks, 0)
	err = result.All(ctx, &task)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func findById(ctx context.Context, id interface{}) (Tasks, error) {
	task := &Tasks{}
	coll := mgm.Coll(task)

	err := coll.FindByIDWithCtx(ctx, id, task)
	if err != nil {
		return Tasks{}, err
	}

	return *task, nil
}

func Update(ctx context.Context, id interface{}, newTask *Tasks) (Tasks, error) {
	oldTask := &Tasks{}
	coll := mgm.Coll(oldTask)

	err := coll.FindByIDWithCtx(ctx, id, oldTask)
	if err != nil {
		return Tasks{}, err
	}

	oldTask.Name = newTask.Name
	oldTask.Description = newTask.Description
	oldTask.ParentRating = newTask.ParentRating
	oldTask.ChildrenRating = newTask.ChildrenRating
	oldTask.Challenge = newTask.Challenge

	err = coll.UpdateWithCtx(ctx, oldTask)
	if err != nil {
		return Tasks{}, err
	}

	return *oldTask, nil
}
