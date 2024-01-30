/*

LICENSE:  MIT
Author:   sine
Email:    sinerwr@gmail.com

*/

package utils

import (
	"net/http"
)

type ResponseJSON struct {
	Code int64       `json:"code"`
	Data interface{} `json:"data"`
}

type ResponseXML struct {
	Code int64       `xml:"code"`
	Data interface{} `xml:"data"`
}

func HttpResponse(contentType string, rw http.ResponseWriter, rsp []byte) {
	switch contentType {
	case "xml":
		rw.Header().Add("Content-Type", "application/xml")
	case "json":
		rw.Header().Add("Content-Type", "application/json")
	default:
		rw.Header().Add("Content-Type", "text/plain")
	}
	rw.Write(rsp)
}
