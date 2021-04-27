package routine_plans

import (
	"context"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
)

type RoutinePlan struct {
	// DefaultModel adds _id, created_at and updated_at fields to the Model
	mgm.DefaultModel `bson:",inline"`
	NumberOfWeeks    int         `json:"numberOfWeeks" bson:"numberOfWeeks"`
	UserID           interface{} `json:"user_id" bson:"user_id"`
	// Subjects         []subjects.Subject `json:"subjects" bson:"subjects"`
	// WeekRoutines     []interface{}      `json:"week_routines" bson:"week_routines"`
}

func create(ctx context.Context, routinePlan *RoutinePlan) (RoutinePlan, error) {
	coll := mgm.Coll(routinePlan)

	err := coll.CreateWithCtx(ctx, routinePlan)

	if err != nil {
		return RoutinePlan{}, err
	}

	return *routinePlan, nil
}

func delete(ctx context.Context, id interface{}) error {
	routinePlan := &RoutinePlan{}
	coll := mgm.Coll(routinePlan)

	err := coll.FindByIDWithCtx(ctx, id, routinePlan)
	if err != nil {
		return err
	}

	err = coll.DeleteWithCtx(ctx, routinePlan)
	if err != nil {
		return err
	}

	return nil
}

func findAll(ctx context.Context) ([]RoutinePlan, error) {
	coll := mgm.Coll(&RoutinePlan{})

	result, err := coll.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	routinePlans := make([]RoutinePlan, 0)
	err = result.All(ctx, &routinePlans)
	if err != nil {
		return nil, err
	}

	return routinePlans, nil
}

func findById(ctx context.Context, id interface{}) (RoutinePlan, error) {
	routinePlan := &RoutinePlan{}
	coll := mgm.Coll(routinePlan)

	err := coll.FindByIDWithCtx(ctx, id, routinePlan)
	if err != nil {
		return RoutinePlan{}, err
	}

	return *routinePlan, nil
}

func update(ctx context.Context, id interface{}, newRoutinePlan *RoutinePlan) (RoutinePlan, error) {
	routinePlan := &RoutinePlan{}
	coll := mgm.Coll(routinePlan)

	err := coll.FindByIDWithCtx(ctx, id, routinePlan)
	if err != nil {
		return RoutinePlan{}, err
	}

	routinePlan.NumberOfWeeks = newRoutinePlan.NumberOfWeeks
	routinePlan.UserID = newRoutinePlan.UserID
	// routinePlan.Subjects = newRoutinePlan.Subjects
	// TODO week routines

	err = coll.UpdateWithCtx(ctx, routinePlan)
	if err != nil {
		return RoutinePlan{}, err
	}

	return *routinePlan, nil
}
