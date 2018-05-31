package app

import (
	"log"
	"net/http"
	"widgetFactory/app/handler"
	"widgetFactory/app/repository"
	"widgetFactory/config"
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
	a.Post("/auth", a.CreateToken)

	a.Get("/widgets", handler.ValidateMiddleware(a.Auth, a.GetAllWidgets))
	a.Post("/widgets", handler.ValidateMiddleware(a.Auth, a.CreateWidget))
	a.Get("/widgets/{id:[0-9]+}", handler.ValidateMiddleware(a.Auth, a.GetWidget))
	a.Put("/widgets/{id:[0-9]+}", handler.ValidateMiddleware(a.Auth, a.UpdateWidget))

	a.Get("/users", handler.ValidateMiddleware(a.Auth, a.GetAllUsers))
	a.Get("/users/{id:[0-9]+}", handler.ValidateMiddleware(a.Auth, a.GetUser))
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
func (a *App) CreateToken(w http.ResponseWriter, r *http.Request) {
	handler.CreateToken(a.Auth, w, r)
}

/*
** Widgets Handlers
 */
func (a *App) GetAllWidgets(w http.ResponseWriter, r *http.Request) {
	handler.GetAllWidgets(a.WidgetDb, w, r)
}

func (a *App) GetWidget(w http.ResponseWriter, r *http.Request) {
	handler.GetWidget(a.WidgetDb, w, r)
}

func (a *App) UpdateWidget(w http.ResponseWriter, r *http.Request) {
	handler.UpdateWidget(a.WidgetDb, w, r)
}

func (a *App) CreateWidget(w http.ResponseWriter, r *http.Request) {
	handler.CreateWidget(a.WidgetDb, w, r)
}

/*
** Users Handlers
 */
func (a *App) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	handler.GetAllUsers(a.UserDb, w, r)
}

func (a *App) GetUser(w http.ResponseWriter, r *http.Request) {
	handler.GetUser(a.UserDb, w, r)
}

/*
** Run the app
 */
func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}