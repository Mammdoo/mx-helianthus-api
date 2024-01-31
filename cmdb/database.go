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
	ID 						uint  	`gorm:"primaryKey;autoIncrement"`
	AppID 				string  `gorm:"column:short_name"`
	AppName				string
	AppOwner      string
	Members 			string  `gorm:"column:monitor_member"`
}

func (application) TableName() string {
	return "application"
}

type user struct {
	Uid 		string  `gorm:"primaryKey"`
	Name 		string
	Email 	string
}

func migration() {
	utils.DB.AutoMigrate(&application{})
	utils.DB.AutoMigrate(&user{})
}

func init() {
	migration()
}