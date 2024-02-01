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
)

func GetApplication(rw http.ResponseWriter, req *http.Request) {
	var rst []modelApplication
  var res []zApplicationBase
	utils.DB.Find(&rst)
	
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
	utils.DB.Limit(1).Find(&rst)

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

	utils.DB.Create(dbObject)

	rsp, _ = json.Marshal(&utils.ResponseJSON{Code: 0, Data: dbObject })
	utils.HttpResponse("json", rw, rsp)
}

