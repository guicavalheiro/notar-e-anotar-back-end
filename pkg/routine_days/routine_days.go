package routine_days

import (
	"context"
	"time"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
)

type RoutineDays struct {
	// DefaultModel adds _id, created_at and updated_at fields to the Model
	mgm.DefaultModel `bson:",inline"`
	Date             time.Time     `json:"date" bson:"date"`
	Chores           []interface{} `json:"chores" bson:"chores"`
}

func create(ctx context.Context, routineDays *RoutineDays) (RoutineDays, error) {
	coll := mgm.Coll(routineDays)

	err := coll.CreateWithCtx(ctx, routineDays)

	if err != nil {
		return RoutineDays{}, err
	}

	return *routineDays, nil
}

func delete(ctx context.Context, id interface{}) error {
	routineDay := &RoutineDays{}
	coll := mgm.Coll(routineDay)

	err := coll.FindByIDWithCtx(ctx, id, routineDay)
	if err != nil {
		return err
	}

	err = coll.DeleteWithCtx(ctx, routineDay)
	if err != nil {
		return err
	}

	return nil
}

func findAll(ctx context.Context) ([]RoutineDays, error) {
	coll := mgm.Coll(&RoutineDays{})

	result, err := coll.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	routineDays := make([]RoutineDays, 0)
	err = result.All(ctx, &routineDays)
	if err != nil {
		return nil, err
	}

	return routineDays, nil
}

func findById(ctx context.Context, id interface{}) (RoutineDays, error) {
	routineDay := &RoutineDays{}
	coll := mgm.Coll(routineDay)

	err := coll.FindByIDWithCtx(ctx, id, routineDay)
	if err != nil {
		return RoutineDays{}, err
	}

	return *routineDay, nil
}

func update(ctx context.Context, id interface{}, newRoutineDays *RoutineDays) (RoutineDays, error) {
	routineDay := &RoutineDays{}
	coll := mgm.Coll(routineDay)

	err := coll.FindByIDWithCtx(ctx, id, routineDay)
	if err != nil {
		return RoutineDays{}, err
	}

	routineDay.Chores = newRoutineDays.Chores

	err = coll.UpdateWithCtx(ctx, routineDay)
	if err != nil {
		return RoutineDays{}, err
	}

	return *routineDay, nil
}
