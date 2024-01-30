/*

LICENSE:  MIT
Author:   sine
Email:    sinerwr@gmail.com

*/

package main

import (
	"net/http"
	// "github.com/gorilla/mux"
)

// var (
// 	Handler *mux.Router
// )

// func init() {
// 	Handler = mux.NewRouter()
// 	Handler.StrictSlash(true)
// 	routePublic()
// 	routeCMDB()
// }

func main() {
	http.ListenAndServe("0.0.0.0:2048", httpHandler)
}
