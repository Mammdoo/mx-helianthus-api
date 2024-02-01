/*

LICENSE:  MIT
Author:   sine
Email:    sinerwr@gmail.com

*/

package cmdb

type zApplicationBase struct {
	ApplicationID 				string	`json:"app_id"`
	ApplicationName				string	`json:"app_name"`
	ApplicationOwner      string	`json:"app_owner"`
	ApplicationMembers 		[]string	`json:"app_members"`
}