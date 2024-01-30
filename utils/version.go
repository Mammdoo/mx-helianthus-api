/*

LICENSE:  MIT
Author:   sine
Email:    sinerwr@gmail.com

*/

package utils

import (
	"net/http"
	"encoding/json"
)

func PublicCfgVersion(rw http.ResponseWriter, req *http.Request) {
	rsp, _ := json.Marshal(&ResponseJSON{Code: 0, Data: "[Success] testing"})
	HttpResponse("json", rw, rsp)
}