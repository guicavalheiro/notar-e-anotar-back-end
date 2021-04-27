package weekly_routines

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type WeeklyRoutinesRouter struct{}

func NewWeeklyRoutinesRouter() *WeeklyRoutinesRouter {
	return &WeeklyRoutinesRouter{}
}

func (u *WeeklyRoutinesRouter) RegisterRoutes(r *mux.Router) {
	r.Path("/").Methods(http.MethodGet).HandlerFunc(u.GetAll)
	r.Path("/{id}").Methods(http.MethodGet).HandlerFunc(u.Get)
	r.PathPrefix("/").Methods(http.MethodPost).HandlerFunc(u.Post)
	r.PathPrefix("/{id}").Methods(http.MethodPut).HandlerFunc(u.Put)
	r.Path("/{id}").Methods(http.MethodDelete).HandlerFunc(u.Delete)
}

func (u *WeeklyRoutinesRouter) Get(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	weeklyRoutine, err := findById(r.Context(), vars["id"])
	if err != nil {
		log.Printf("error finding weeklyRoutine by id: %s", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	bytes, err := json.Marshal(weeklyRoutine)
	if err != nil {
		log.Printf("error marshaling weeklyRoutine: %s", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.Write(bytes)
}

func (u *WeeklyRoutinesRouter) GetAll(rw http.ResponseWriter, r *http.Request) {
	weeklyRoutines, err := findAll(r.Context())
	if err != nil {
		log.Printf("error finding weeklyRoutine: %s", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	bytes, err := json.Marshal(weeklyRoutines)
	if err != nil {
		log.Printf("error marshaling weeklyRoutine: %s", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.Write(bytes)
}

func (u *WeeklyRoutinesRouter) Post(rw http.ResponseWriter, r *http.Request) {
	var weeklyRoutine WeeklyRoutine

	err := json.NewDecoder(r.Body).Decode(&weeklyRoutine)
	if err != nil {
		log.Printf("error decoding weeklyRoutine: %s", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	weeklyRoutine, err = create(r.Context(), &weeklyRoutine)
	if err != nil {
		log.Printf("error creating weeklyRoutine: %s", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	bytes, err := json.Marshal(weeklyRoutine)
	if err != nil {
		log.Printf("error marshaling weeklyRoutine: %s", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.Write(bytes)
}

func (u *WeeklyRoutinesRouter) Put(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var weeklyRoutine WeeklyRoutine

	err := json.NewDecoder(r.Body).Decode(&weeklyRoutine)
	if err != nil {
		log.Printf("error decoding weeklyRoutine: %s", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	weeklyRoutine, err = Update(r.Context(), vars["id"], &weeklyRoutine)
	if err != nil {
		log.Printf("error creating weeklyRoutine: %s", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	bytes, err := json.Marshal(weeklyRoutine)
	if err != nil {
		log.Printf("error marshaling weeklyRoutine: %s", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.Write(bytes)
}

func (u *WeeklyRoutinesRouter) Delete(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	err := Delete(r.Context(), vars["id"])
	if err != nil {
		log.Printf("error finding weeklyRoutine by id: %s", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
}
