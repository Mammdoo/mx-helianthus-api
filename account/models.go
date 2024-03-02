/*

LICENSE:  MIT
Author:   sine
Email:    sinerwr@gmail.com

*/

package account

import (
	"helianthus/database"
)

type modelAccount struct {
	UUID 					string  `gorm:"primaryKey"`
	Name 					string
	Email					string
	Mobile	      string
	Class			 		string
}

func (modelAccount) TableName() string {
	return "accounts"
}

func migration() {
	database.Conn.AutoMigrate(&modelAccount{})
}

func init() {
	migration()
}