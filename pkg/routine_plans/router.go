package routine_plans

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type RoutinePlanRouter struct{}

func NewRoutinePlanRouter() *RoutinePlanRouter {
	return &RoutinePlanRouter{}
}

func (u *RoutinePlanRouter) RegisterRoutes(r *mux.Router) {
	r.Path("/").Methods(http.MethodGet).HandlerFunc(u.GetAll)
	r.Path("/{id}").Methods(http.MethodGet).HandlerFunc(u.Get)
	r.PathPrefix("/").Methods(http.MethodPost).HandlerFunc(u.Post)
	r.PathPrefix("/{id}").Methods(http.MethodPut).HandlerFunc(u.Put)
	r.Path("/{id}").Methods(http.MethodDelete).HandlerFunc(u.Delete)
}

func (u *RoutinePlanRouter) Get(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	routinePlan, err := findById(r.Context(), vars["id"])
	if err != nil {
		log.Printf("error finding routine plan by id: %s", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	bytes, err := json.Marshal(routinePlan)
	if err != nil {
		log.Printf("error marshaling routine plan: %s", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.Write(bytes)
}

func (u *RoutinePlanRouter) GetAll(rw http.ResponseWriter, r *http.Request) {
	routinePlans, err := findAll(r.Context())
	if err != nil {
		log.Printf("error finding routine plans: %s", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	bytes, err := json.Marshal(routinePlans)
	if err != nil {
		log.Printf("error marshaling routine plans: %s", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.Write(bytes)
}

func (u *RoutinePlanRouter) Post(rw http.ResponseWriter, r *http.Request) {
	var routinePlan RoutinePlan

	err := json.NewDecoder(r.Body).Decode(&routinePlan)
	if err != nil {
		log.Printf("error decoding routine plan: %s", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	routinePlan, err = create(r.Context(), &routinePlan)
	if err != nil {
		log.Printf("error creating routine plan: %s", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	bytes, err := json.Marshal(routinePlan)
	if err != nil {
		log.Printf("error marshaling routine plan: %s", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.Write(bytes)
}

func (u *RoutinePlanRouter) Put(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var routinePlan RoutinePlan

	err := json.NewDecoder(r.Body).Decode(&routinePlan)
	if err != nil {
		log.Printf("error decoding routine plan: %s", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	routinePlan, err = update(r.Context(), vars["id"], &routinePlan)
	if err != nil {
		log.Printf("error creating routine plan: %s", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	bytes, err := json.Marshal(routinePlan)
	if err != nil {
		log.Printf("error marshaling routine plan: %s", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.Write(bytes)
}

func (u *RoutinePlanRouter) Delete(rw http.ResponseWriter, r *http.Request) {
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
