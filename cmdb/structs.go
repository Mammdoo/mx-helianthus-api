/*

LICENSE:  MIT
Author:   sine
Email:    sinerwr@gmail.com

*/

package cmdb

import (
	"helianthus/utils"
)

type application struct {
	Application 	string
	Owner        	string
	Members 			string
}

func (application) TableName() string {
	return "application"
}


func init() {
	utils.DB.AutoMigrate(&application{})
}