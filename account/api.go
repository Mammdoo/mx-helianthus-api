/*

LICENSE:  MIT
Author:   sine
Email:    sinerwr@gmail.com

*/

package account

import (
	"net/http"
	"encoding/json"
	"helianthus/utils"
)

func CheckModuleLoaded(rw http.ResponseWriter, req *http.Request) {
	rsp, _ := json.Marshal(&utils.ResponseJSON{Code: 0, Data: "Loaded"})
	utils.HttpResponse("json", rw, rsp)
}

