package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"short-url/util"
)

type App struct {
	Router      *mux.Router
	Middlewares *Middleware
}

type CreateShort struct {
	URL        string `json:"url"`
	Expiration int64  `json:"expiration"`
}

func (app *App) init() {
	log.SetFlags(log.LstdFlags)
	app.Router = mux.NewRouter()
	app.Middlewares = &Middleware{}
	app.initRouter()
}

func (app *App) initRouter() {
	app.Router.Use(app.Middlewares.Logging)
	app.Router.Use(app.Middlewares.Recover)
	app.Router.HandleFunc("/api/short", app.createShortUrl).Methods("POST")
	app.Router.HandleFunc("/api/short/{id}", app.getShortInfo).Methods("GET")
	app.Router.HandleFunc("/{url:[a-zA-Z0-9]{1,11}}", app.redirect)
}

func (app *App) createShortUrl(w http.ResponseWriter, r *http.Request) {
	var createShort = CreateShort{}
	if err := json.NewDecoder(r.Body).Decode(&createShort); err != nil {
		return
	}
	defer r.Body.Close()
	w.WriteHeader(http.StatusOK)
}

func (app *App) getShortInfo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(vars["id"]))
}

func (app *App) redirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://www.baidu.com", http.StatusFound)
}

func (app *App) run(addr string) {
	log.Fatal(http.ListenAndServe(addr, app.Router))
}

func responseWithError(w http.ResponseWriter, err error) {
	switch e := err.(type) {
	case Error:
		ajaxReturn := util.AjaxReturn(-1, e.Error(), nil)
		responseWithJson(w, e.Status(), ajaxReturn.JsonBytes())
	default:
		ajaxReturn := util.AjaxReturn(-1, http.StatusText(http.StatusInternalServerError), nil)
		responseWithJson(w, http.StatusInternalServerError, ajaxReturn.JsonBytes())
	}
}

func responseWithJson(w http.ResponseWriter, code int, respJson []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(respJson)
}
