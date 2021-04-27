package main

import (
	"net/http"

	"tools.ages.pucrs.br/notar-e-anotar/notar-e-anotar-back-end/pkg/routine_days"
	"tools.ages.pucrs.br/notar-e-anotar/notar-e-anotar-back-end/pkg/routine_plans"
	"tools.ages.pucrs.br/notar-e-anotar/notar-e-anotar-back-end/pkg/subjects"
	"tools.ages.pucrs.br/notar-e-anotar/notar-e-anotar-back-end/pkg/tasks"
	"tools.ages.pucrs.br/notar-e-anotar/notar-e-anotar-back-end/pkg/user"
	"tools.ages.pucrs.br/notar-e-anotar/notar-e-anotar-back-end/pkg/weekly_routines"
)

func main() {
	handlers := map[string]Router{
		UserRouter:          user.NewUserRouter(),
		RoutinePlanRouter:   routine_plans.NewRoutinePlanRouter(),
		TaskRouter:          tasks.NewTaskRouter(),
		SubjectRouter:       subjects.NewSubjectsRouter(),
		WeeklyRoutineRouter: weekly_routines.NewWeeklyRoutinesRouter(),
		RoutineDayRouter:    routine_days.NewRoutineDaysRouter(),
	}

	api := registerAPI(handlers)

	http.ListenAndServe(":8080", api)
}
