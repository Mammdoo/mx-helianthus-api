/*

LICENSE:  MIT
Author:   sine
Email:    sinerwr@gmail.com

*/

package cmdb

import (
	"net/http"
	"encoding/json"
	"helianthus/utils"
	"fmt"
)

func GetApplication(rw http.ResponseWriter, req *http.Request) {
	rst := utils.DB.Find(&application{})
	rsp, _ := json.Marshal(&utils.ResponseJSON{Code: 0, Data: rst })
	utils.HttpResponse("json", rw, rsp)
}

func GetApplicationByAppID(rw http.ResponseWriter, req *http.Request) {
	rst := utils.DB.Find(&application{AppID: utils.GetRouteName(req, "AppID")})
	fmt.Println(rst)
	rsp, _ := json.Marshal(&utils.ResponseJSON{Code: 0, Data: rst })
	utils.HttpResponse("json", rw, rsp)
}

func CreateApplication(rw http.ResponseWriter, req *http.Request) {
	rsp := utils.ValidatePostJSON(req)
	if rsp != nil {
		utils.HttpResponse("json", rw, rsp)
		return
	}

	var request requestApplication

	json.NewDecoder(req.Body).Decode(&request)
	
	if request.AppName == "" || request.AppID == "" || request.AppOwner == "" {
		utils.ValidatorMissingRequest(rw)
		return
	}

	dbObject := &application{
		AppName: request.AppName,
		AppID: request.AppID,
		AppOwner: request.AppOwner,
		Members: request.Members,
	}

	utils.DB.Create(dbObject)

	rsp, _ = json.Marshal(&utils.ResponseJSON{Code: 0, Data: dbObject })
	utils.HttpResponse("json", rw, rsp)
}

