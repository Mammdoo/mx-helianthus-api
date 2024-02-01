/*

LICENSE:  MIT
Author:   sine
Email:    sinerwr@gmail.com

*/

package cmdb

import (
	"helianthus/utils"
)

type modelApplication struct {
	ID 						uint  	`gorm:"primaryKey;autoIncrement"`
	AppID 				string  `gorm:"column:short_name"`
	AppName				string
	AppOwner      string
	AppMembers 		string  `gorm:"column:monitor_member"`
}

func (modelApplication) TableName() string {
	return "application"
}

type modelUser struct {
	Uid 		string  `gorm:"primaryKey"`
	Name 		string
	Email 	string
}

func migration() {
	utils.DB.AutoMigrate(&modelApplication{})
	// utils.DB.AutoMigrate(&modelUser{})
}

func init() {
	migration()
}