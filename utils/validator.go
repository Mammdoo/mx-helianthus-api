/*

LICENSE:  MIT
Author:   sine
Email:    sinerwr@gmail.com

*/

package utils

import (
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)

func ValidatePostJSON(req *http.Request) []byte {
	if req.Header.Get("Content-Type") != "application/json" {
		rsp, _ := json.Marshal(&ResponseJSON{Code: 0x00000020, Data: "Only Accept application/json"})
		return rsp
	}
	return nil
}

func ValidatorMissingRequest(rw http.ResponseWriter) {
	rsp, _ := json.Marshal(&ResponseJSON{Code: 0x00000030, Data: "Bad Request Body"})
	HttpResponse("json", rw, rsp)
}

type HandlerFunc func(http.ResponseWriter, *http.Request)

func AuthMiddlerware(next HandlerFunc) HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		if (req.Header.Get("Authorization") != "Bearer aGVsaWFudGh1cwo=") {
			rsp, _ := json.Marshal(&ResponseJSON{Code: 0x00000010, Data: "token error"})
			HttpResponse("json", rw, rsp)
			return
		}
		next(rw, req)
	}
}

func GetRouteName(req *http.Request, name string) string {
	return mux.Vars(req)[name]
}