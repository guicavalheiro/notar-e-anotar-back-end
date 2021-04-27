package tasks

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type TasksRouter struct{}

func NewTaskRouter() *TasksRouter {
	return &TasksRouter{}
}

func (u *TasksRouter) RegisterRoutes(r *mux.Router) {
	r.Path("/").Methods(http.MethodGet).HandlerFunc(u.GetAll)
	r.Path("/{id}").Methods(http.MethodGet).HandlerFunc(u.Get)
	r.PathPrefix("/").Methods(http.MethodPost).HandlerFunc(u.Post)
	r.PathPrefix("/{id}").Methods(http.MethodPut).HandlerFunc(u.Put)
	r.Path("/{id}").Methods(http.MethodDelete).HandlerFunc(u.Delete)
}

func (u *TasksRouter) Get(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	task, err := findById(r.Context(), vars["id"])
	if err != nil {
		log.Printf("error finding task by id: %s", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	bytes, err := json.Marshal(task)
	if err != nil {
		log.Printf("error marshaling task: %s", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.Write(bytes)
}

func (u *TasksRouter) GetAll(rw http.ResponseWriter, r *http.Request) {
	tasks, err := findAll(r.Context())
	if err != nil {
		log.Printf("error finding tasks: %s", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	bytes, err := json.Marshal(tasks)
	if err != nil {
		log.Printf("error marshaling tasks: %s", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.Write(bytes)
}

func (u *TasksRouter) Post(rw http.ResponseWriter, r *http.Request) {
	var task Tasks

	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		log.Printf("error decoding task: %s", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	task, err = create(r.Context(), &task)
	if err != nil {
		log.Printf("error creating task: %s", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	bytes, err := json.Marshal(task)
	if err != nil {
		log.Printf("error marshaling task: %s", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.Write(bytes)
}

func (u *TasksRouter) Put(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var task Tasks

	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		log.Printf("error decoding task: %s", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	task, err = Update(r.Context(), vars["id"], &task)
	if err != nil {
		log.Printf("error creating task: %s", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	bytes, err := json.Marshal(task)
	if err != nil {
		log.Printf("error marshaling task: %s", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.Write(bytes)
}

func (u *TasksRouter) Delete(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	err := Delete(r.Context(), vars["id"])
	if err != nil {
		log.Printf("error finding task by id: %s", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
}
