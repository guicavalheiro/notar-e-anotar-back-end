package subjects

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type SubjectsRouter struct{}

func NewSubjectsRouter() *SubjectsRouter {
	return &SubjectsRouter{}
}

func (u *SubjectsRouter) RegisterRoutes(r *mux.Router) {
	r.Path("/").Methods(http.MethodGet).HandlerFunc(u.GetAll)
	r.Path("/{name}").Methods(http.MethodGet).HandlerFunc(u.Get)
}

func (u *SubjectsRouter) Get(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	subject, err := findByName(r.Context(), vars["name"])
	if err != nil {
		log.Printf("error finding subject by name: %s", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	bytes, err := json.Marshal(subject)
	if err != nil {
		log.Printf("error marshaling subject: %s", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rw.Write(bytes)
}

func (u *SubjectsRouter) GetAll(rw http.ResponseWriter, r *http.Request) {
	subjects, err := findAll(r.Context())
	if err != nil {
		log.Printf("error finding subjects: %s", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	bytes, err := json.Marshal(subjects)
	if err != nil {
		log.Printf("error marshaling subjects: %s", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rw.Write(bytes)
}
