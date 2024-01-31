/*

LICENSE:  MIT
Author:   sine
Email:    sinerwr@gmail.com

*/

package cmdb

type requestApplication struct {
	AppID 				string	`json:"app_id"`
	AppName				string	`json:"app_name"`
	AppOwner      string	`json:"app_owner"`
	Members 			string	`json:"app_member"`
}