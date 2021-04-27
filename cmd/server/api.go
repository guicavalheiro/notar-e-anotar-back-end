package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Router interface {
	RegisterRoutes(r *mux.Router)
}

func registerAPI(handlers map[string]Router) http.Handler {
	r := mux.NewRouter()
	api := r.NewRoute().PathPrefix("/api").Subrouter()

	handlers[UserRouter].RegisterRoutes(api.NewRoute().PathPrefix(UserPath).Subrouter())
	handlers[RoutinePlanRouter].RegisterRoutes(api.NewRoute().PathPrefix(RoutinePlanPath).Subrouter())
	handlers[TaskRouter].RegisterRoutes(api.NewRoute().PathPrefix(TaskPath).Subrouter())
	handlers[SubjectRouter].RegisterRoutes(api.NewRoute().PathPrefix(SubjectPath).Subrouter())
	handlers[WeeklyRoutineRouter].RegisterRoutes(api.NewRoute().PathPrefix(WeeklyRoutinePath).Subrouter())
	handlers[RoutineDayRouter].RegisterRoutes(api.NewRoute().Subrouter().PathPrefix(RoutineDayPath).Subrouter())

	r.NewRoute().HandlerFunc(HealthHandler)

	return r
}

func HealthHandler(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("{ \"health\": \"up\" }"))
}
