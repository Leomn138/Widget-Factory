package app

import (
	"log"
	"net/http"
	"github.com/Leomn138/Widget-Factory/app/handler"
	"github.com/Leomn138/Widget-Factory/app/repository"
	"github.com/Leomn138/Widget-Factory/config"
	"github.com/gorilla/mux"
)
type App struct {
	Router *mux.Router
	Auth *config.Auth
	WidgetDb *config.DBConfig
	UserDb *config.DBConfig
}

func (a *App) Initialize(Config *config.Config) {
	a.Auth = Config.Auth
	a.WidgetDb = Config.WidgetDBConfig
	a.UserDb = Config.UserDBConfig
	repository.CreateDatabaseIfNotExists(Config.UserDBConfig)
	repository.CreateDatabaseIfNotExists(Config.WidgetDBConfig)
	a.Router = mux.NewRouter()
	a.setRouters()
}

func (a *App) setRouters() {
	a.Post("/auth", CreateToken(logEntry, a.Auth))

	a.Get("/widgets", handler.ValidateMiddleware(a.Auth, GetAllWidgets(logEntry, a.WidgetDb)))


	a.Post("/widgets", handler.ValidateMiddleware(a.Auth, CreateWidget(logEntry, a.WidgetDb)))
	a.Get("/widgets/{id:[0-9]+}", handler.ValidateMiddleware(a.Auth, GetWidget(logEntry, a.WidgetDb)))
	a.Put("/widgets/{id:[0-9]+}", handler.ValidateMiddleware(a.Auth, UpdateWidget(logEntry, a.WidgetDb)))

	a.Get("/users", handler.ValidateMiddleware(a.Auth, GetAllUsers(logEntry, a.UserDb)))
	a.Get("/users/{id:[0-9]+}", handler.ValidateMiddleware(a.Auth, GetUser(logEntry, a.UserDb)))
}


func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}

func (a *App) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("PUT")
}

/*
** Auth Handlers
 */
var CreateToken = func(f http.HandlerFunc, auth *config.Auth) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		f(w, r)
		handler.CreateToken(auth, w, r)
	}
}

/*
** Widgets Handlers
 */
var GetAllWidgets = func(f http.HandlerFunc, dbConfig *config.DBConfig) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		f(w, r)
		handler.GetAllWidgets(dbConfig, w, r)
	}
}

var GetWidget = func(f http.HandlerFunc, dbConfig *config.DBConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		f(w, r)
		handler.GetWidget(dbConfig, w, r)
	}
}

var UpdateWidget = func(f http.HandlerFunc, dbConfig *config.DBConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		f(w, r)
		handler.UpdateWidget(dbConfig, w, r)
	}
}

var CreateWidget = func(f http.HandlerFunc, dbConfig *config.DBConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		f(w, r)
		handler.CreateWidget(dbConfig, w, r)
	}
}

/*
** Users Handlers
 */
var GetAllUsers = func(f http.HandlerFunc, dbConfig *config.DBConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		f(w, r)
		handler.GetAllUsers(dbConfig, w, r)
	}
}

var GetUser = func(f http.HandlerFunc, dbConfig *config.DBConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		f(w, r)
		handler.GetUser(dbConfig, w, r)
	}
}

func logEntry(w http.ResponseWriter, r *http.Request) {
	log.Print("New request from host: " + r.Host)
}

/*
** Run the app
 */
func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}