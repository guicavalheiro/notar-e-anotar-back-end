package user

import (
	"net/http"

	"github.com/gorilla/mux"
)

type UserRouter struct{}

func NewUserRouter() *UserRouter {
	return &UserRouter{}
}

func (u *UserRouter) RegisterRoutes(r *mux.Router) {
	r.Path("/").Methods(http.MethodGet).HandlerFunc(u.GetAll)
	r.Path("/{id}").Methods(http.MethodGet).HandlerFunc(u.Get)
	r.PathPrefix("/").Methods(http.MethodPut).HandlerFunc(u.Put)
	r.PathPrefix("/").Methods(http.MethodPost).HandlerFunc(u.Post)
}

func (u *UserRouter) Get(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	rw.Write([]byte(vars["id"]))
}

func (u *UserRouter) GetAll(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("list"))
}

func (u *UserRouter) Put(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("put"))
}

func (u *UserRouter) Post(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("post"))
}
