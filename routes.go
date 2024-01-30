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
	r.path("/").HandlerFunc(public.GetHealth).Methods("GET")
	r.Path("/version").HandlerFunc(public.GetCurrentVersion).Methods("GET")
}

func routeCMDB(r *mux.Router) {
	r.Path("/application").HandlerFunc(cmdb.GetApplicationID).Methods("GET")
}

func init() {
	httpHandler = mux.NewRouter()
	httpHandler.StrictSlash(true)
	routeRegisterMain := initRouteSub(httpHandler, "/api")

	routeRegisterPublic := initRouteSub(routeRegisterMain, "/public")
	routeRegisterCMDB := initRouteSub(routeRegisterMain, "/cmdb")
	routePublic(routeRegisterPublic)
	routeCMDB(routeRegisterCMDB)
}