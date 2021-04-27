package weekly_routines

import (
	"context"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"tools.ages.pucrs.br/notar-e-anotar/notar-e-anotar-back-end/pkg/routine_days"
)

type WeeklyRoutine struct {
	mgm.DefaultModel `bson:",inline"`
	Year             primitive.DateTime         `json:"year" bson:"year"`
	Routine_plan_id  primitive.ObjectID         `json:"routine_plan_id" bson:"routine_plan_id"`
	Week_number      int32                      `json:"week_number" bson:"week_number"`
	Subject          primitive.ObjectID         `json:"subject" bson:"subject"`
	Routine_days     []routine_days.RoutineDays `json:"routine_days" bson:"routine_days"`
}

func create(ctx context.Context, weeklyRoutine *WeeklyRoutine) (WeeklyRoutine, error) {
	coll := mgm.Coll(weeklyRoutine)

	err := coll.CreateWithCtx(ctx, weeklyRoutine)

	if err != nil {
		return WeeklyRoutine{}, err
	}

	return *weeklyRoutine, nil
}

func Delete(ctx context.Context, id interface{}) error {
	weeklyRoutine := &WeeklyRoutine{}
	coll := mgm.Coll(weeklyRoutine)

	err := coll.FindByIDWithCtx(ctx, id, weeklyRoutine)
	if err != nil {
		return err
	}

	err = coll.DeleteWithCtx(ctx, weeklyRoutine)
	if err != nil {
		return err
	}

	return nil
}

func findAll(ctx context.Context) ([]WeeklyRoutine, error) {
	coll := mgm.Coll(&WeeklyRoutine{})

	result, err := coll.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	weeklyRoutine := make([]WeeklyRoutine, 0)
	err = result.All(ctx, &weeklyRoutine)
	if err != nil {
		return nil, err
	}

	return weeklyRoutine, nil
}

func findById(ctx context.Context, id interface{}) (WeeklyRoutine, error) {
	weeklyRoutine := &WeeklyRoutine{}
	coll := mgm.Coll(weeklyRoutine)

	err := coll.FindByIDWithCtx(ctx, id, weeklyRoutine)
	if err != nil {
		return WeeklyRoutine{}, err
	}

	return *weeklyRoutine, nil
}

func Update(ctx context.Context, id interface{}, newWeeklyRoutine *WeeklyRoutine) (WeeklyRoutine, error) {
	oldWeeklyRoutine := &WeeklyRoutine{}
	coll := mgm.Coll(oldWeeklyRoutine)

	err := coll.FindByIDWithCtx(ctx, id, oldWeeklyRoutine)
	if err != nil {
		return WeeklyRoutine{}, err
	}

	oldWeeklyRoutine.Year = newWeeklyRoutine.Year
	oldWeeklyRoutine.Routine_plan_id = newWeeklyRoutine.Routine_plan_id
	oldWeeklyRoutine.Week_number = newWeeklyRoutine.Week_number
	oldWeeklyRoutine.Subject = newWeeklyRoutine.Subject
	oldWeeklyRoutine.Routine_days = newWeeklyRoutine.Routine_days

	err = coll.UpdateWithCtx(ctx, oldWeeklyRoutine)
	if err != nil {
		return WeeklyRoutine{}, err
	}

	return *oldWeeklyRoutine, nil
}
