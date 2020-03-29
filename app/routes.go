package app

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ilham13/Covid-2020/controllers"
)

// Route App
type Routes struct {
	Router *mux.Router
	DB     *sql.DB
}

// Initialize routing
func (a *Routes) Initialize() {
	a.Router = mux.NewRouter()
	a.initializeRoutes()
}

// Run port
func (a *Routes) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func (a *Routes) initializeRoutes() {
	summary := controllers.TotalSummaryController{}
	user := controllers.UserController{}

	// summary
	a.Router.HandleFunc("/summary", summary.GetList).Methods("GET")

	//users
	a.Router.HandleFunc("/users", user.GetUsers).Methods("GET")
	a.Router.HandleFunc("/user", user.CreateUser).Methods("POST")
	a.Router.HandleFunc("/user/{id:[0-9]+}", user.GetUser).Methods("GET")
	a.Router.HandleFunc("/user/{id:[0-9]+}", user.UpdateUser).Methods("PUT")
	a.Router.HandleFunc("/user/{id:[0-9]+}", user.DeleteUser).Methods("DELETE")
}
