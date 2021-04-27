package subjects

import (
	"context"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
)

type Subject struct {
	// DefaultModel adds _id, created_at and updated_at fields to the Model
	mgm.DefaultModel `bson:",inline"`
	Name             string `json:"name" bson:"name"`
}

func findByName(ctx context.Context, name string) (Subject, error) {
	coll := mgm.Coll(&Subject{})

	result, err := coll.Find(ctx, bson.M{"name": name})
	if err != nil {
		return Subject{}, err
	}

	subjects := make([]Subject, 0)
	err = result.All(ctx, &subjects)
	if err != nil {
		return Subject{}, err
	}

	return subjects[0], nil
}

func findAll(ctx context.Context) ([]Subject, error) {
	coll := mgm.Coll(&Subject{})

	result, err := coll.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	subjects := make([]Subject, 0)
	err = result.All(ctx, &subjects)
	if err != nil {
		return nil, err
	}

	return subjects, nil
}
