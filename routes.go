/*

LICENSE:  MIT
Author:   sine
Email:    sinerwr@gmail.com

*/

package main

import (
	"helianthus/utils"
	"helianthus/cmdb"
	"helianthus/public"
	"github.com/gorilla/mux"
)

var httpHandler *mux.Router

func initRouteSub(r *mux.Router, path string) *mux.Router {
	return r.PathPrefix(path).Subrouter()
}

func routePublic(r *mux.Router) {
	r.Path("/").HandlerFunc(public.GetHealth).Methods("GET")
	r.Path("/version").HandlerFunc(public.GetCurrentVersion).Methods("GET")
}

func routeCMDB(r *mux.Router) {
	r.Path("/").HandlerFunc(public.GetHealth).Methods("GET")

	r.Path("/application").HandlerFunc(utils.AuthMiddlerware(cmdb.GetApplication)).Methods("GET")
	r.Path("/application").HandlerFunc(utils.AuthMiddlerware(cmdb.CreateApplication)).Methods("POST")

	r.Path("/application/{AppID}").HandlerFunc(utils.AuthMiddlerware(cmdb.GetApplicationByAppID)).Methods("GET")
	r.Path("/application/{AppID}").HandlerFunc(utils.AuthMiddlerware(cmdb.UpdateApplicationByAppID)).Methods("PUT")
}

func init() {
	httpHandler = mux.NewRouter()
	httpHandler.StrictSlash(true)

	routeRegisterPublic := initRouteSub(httpHandler, "/api/public")
	routeRegisterCMDB := initRouteSub(httpHandler, "/api/cmdb")

	routePublic(routeRegisterPublic)
	routeCMDB(routeRegisterCMDB)
}