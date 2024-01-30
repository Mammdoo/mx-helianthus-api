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

func (this *application) GetApplicationID() string {
	return this.Application
}

func (this *application) GetOwner() string {	
	return this.Owner
}

func (this *application) GetMembers() string {
	return this.Members
}


func GetApplicationID(rw http.ResponseWriter, req *http.Request) {
	this := &application{
		Application: "ufs-web",
		Owner: "sine",
		Members: "sine",
	}

	rsp, _ := json.Marshal(&utils.ResponseJSON{Code: 0, Data: this.GetApplicationID()})
	utils.HttpResponse("json", rw, rsp)
}