package routine_days

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type RoutineDaysRouter struct{}

func NewRoutineDaysRouter() *RoutineDaysRouter {
	return &RoutineDaysRouter{}
}

func (u *RoutineDaysRouter) RegisterRoutes(r *mux.Router) {
	r.Path("/").Methods(http.MethodGet).HandlerFunc(u.GetAll)
	r.Path("/{id}").Methods(http.MethodGet).HandlerFunc(u.Get)
	r.PathPrefix("/").Methods(http.MethodPost).HandlerFunc(u.Post)
	r.PathPrefix("/{id}").Methods(http.MethodPut).HandlerFunc(u.Put)
	r.Path("/{id}").Methods(http.MethodDelete).HandlerFunc(u.Delete)
}

func (u *RoutineDaysRouter) Get(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	routineDay, err := findById(r.Context(), vars["id"])
	if err != nil {
		log.Printf("error finding routine plan by id: %s", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	bytes, err := json.Marshal(routineDay)
	if err != nil {
		log.Printf("error marshaling routine plan: %s", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.Write(bytes)
}

func (u *RoutineDaysRouter) GetAll(rw http.ResponseWriter, r *http.Request) {
	routineDays, err := findAll(r.Context())
	if err != nil {
		log.Printf("error finding routine plans: %s", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	bytes, err := json.Marshal(routineDays)
	if err != nil {
		log.Printf("error marshaling routine plans: %s", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.Write(bytes)
}

func (u *RoutineDaysRouter) Post(rw http.ResponseWriter, r *http.Request) {
	var routineDay RoutineDays

	err := json.NewDecoder(r.Body).Decode(&routineDay)
	if err != nil {
		log.Printf("error decoding routine plan: %s", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	routineDay, err = create(r.Context(), &routineDay)
	if err != nil {
		log.Printf("error creating routine plan: %s", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	bytes, err := json.Marshal(routineDay)
	if err != nil {
		log.Printf("error marshaling routine plan: %s", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.Write(bytes)
}

func (u *RoutineDaysRouter) Put(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var routineDay RoutineDays

	err := json.NewDecoder(r.Body).Decode(&routineDay)
	if err != nil {
		log.Printf("error decoding routine plan: %s", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	routineDay, err = update(r.Context(), vars["id"], &routineDay)
	if err != nil {
		log.Printf("error creating routine plan: %s", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	bytes, err := json.Marshal(routineDay)
	if err != nil {
		log.Printf("error marshaling routine plan: %s", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.Write(bytes)
}

func (u *RoutineDaysRouter) Delete(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	err := delete(r.Context(), vars["id"])
	if err != nil {
		log.Printf("error finding routine plan by id: %s", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
}
