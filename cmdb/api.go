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
	"helianthus/database"
)

func GetApplication(rw http.ResponseWriter, req *http.Request) {
	var rst []modelApplication
  var res []zApplicationBase
	database.Conn.Find(&rst)
	
	for i := 0; i < len(rst); i++ {
		res = append(res, zApplicationBase{
			ApplicationName: rst[i].AppName,
			ApplicationID: rst[i].AppID,
			ApplicationOwner: rst[i].AppOwner,
			ApplicationMembers: utils.StringToSlice(rst[i].AppMembers, ","),
		})
	}

	rsp, _ := json.Marshal(&utils.ResponseJSON{Code: 0, Data: res })
	utils.HttpResponse("json", rw, rsp)
}

func GetApplicationByAppID(rw http.ResponseWriter, req *http.Request) {
	var rst modelApplication
	var res []zApplicationBase
	rst = modelApplication{AppID: utils.GetRouteName(req, "AppID")}
	database.Conn.Where(rst).Find(&rst)

	if rst.AppOwner == "" {
		rsp, _ := json.Marshal(&utils.ResponseJSON{Code: 1, Data : "No such application with AppID: " + rst.AppID})
		utils.HttpResponse("json", rw, rsp)
		return
	}

	res = append(res, zApplicationBase{
		ApplicationName: rst.AppName,
		ApplicationID: rst.AppID,
		ApplicationOwner: rst.AppOwner,
		ApplicationMembers: utils.StringToSlice(rst.AppMembers, ","),
	})
	rsp, _ := json.Marshal(&utils.ResponseJSON{Code: 0, Data: res})
	utils.HttpResponse("json", rw, rsp)
}

func CreateApplication(rw http.ResponseWriter, req *http.Request) {
	rsp := utils.ValidatePostJSON(req)
	if rsp != nil {
		utils.HttpResponse("json", rw, rsp)
		return
	}

	var request zApplicationBase

	json.NewDecoder(req.Body).Decode(&request)
	
	if request.ApplicationID == "" || request.ApplicationName == "" || request.ApplicationOwner == "" {
		utils.ValidatorMissingRequest(rw)
		return
	}

	dbObject := &modelApplication{
		AppName: request.ApplicationName,
		AppID: request.ApplicationID,
		AppOwner: request.ApplicationOwner,
		AppMembers: utils.SliceToString(request.ApplicationMembers, ","),
	}

	database.Conn.Create(dbObject)

	rsp, _ = json.Marshal(&utils.ResponseJSON{Code: 0, Data: dbObject })
	utils.HttpResponse("json", rw, rsp)
}

func UpdateApplicationByAppID(rw http.ResponseWriter, req *http.Request) {
	rsp := utils.ValidatePostJSON(req)
	if rsp != nil {
		utils.HttpResponse("json", rw, rsp)
		return
	}

	var request zApplicationBase
	json.NewDecoder(req.Body).Decode(&request)
	
	if utils.GetRouteName(req, "AppID") == "" {
		utils.ValidatorMissingRequest(rw)
		return
	}
	dbObject := &modelApplication{AppID: utils.GetRouteName(req, "AppID")}
	database.Conn.Where(dbObject).First(&dbObject)

	if request.ApplicationOwner != "" {
		dbObject.AppOwner = request.ApplicationOwner
	}
	if request.ApplicationMembers != nil {
		dbObject.AppMembers = utils.SliceToString(request.ApplicationMembers, ",")
	}
	if request.ApplicationName != "" {
		dbObject.AppName = request.ApplicationName
	}

	database.Conn.Save(&dbObject)
	rsp, _ = json.Marshal(&utils.ResponseJSON{Code: 0, Data: "Success Update" })
	utils.HttpResponse("json", rw, rsp)
}

