/*

LICENSE:  MIT
Author:   sine
Email:    sinerwr@gmail.com

*/

package public

import (
	"net/http"
	"encoding/json"
)

func GetCurrentVersion(rw http.ResponseWriter, req *http.Request) {
	rsp, _ := json.Marshal(&reponsePublic{Code: 0, Data: "0.1.0"})
	HttpResponse("json", rw, rsp)
}

func GetHealth(rw http.ResponseWriter, req *http.Request) {
	rw.WriteHeader(http.StatusNoContent)
}