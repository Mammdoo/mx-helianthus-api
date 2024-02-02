/*

LICENSE:  MIT
Author:   sine
Email:    sinerwr@gmail.com

*/

package database

import (
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"

	"helianthus/utils"
)

var (
	Conn *gorm.DB
)

type dbconfig struct {
	Username string
	Password string
	Host     string
	Port     string
	Database string
}

func (dsn *dbconfig) String() string {
	return dsn.Username + ":" + dsn.Password + "@tcp(" + dsn.Host + ":" + dsn.Port + ")/" + dsn.Database
}

func generateDSN() string {
	dsn := &dbconfig{
		Username: utils.GetEnvWithDefault("DB_USERNAME", "root"),
		Password: utils.GetEnvWithDefault("DB_PASSWORD", "this-is-secret-123"),
		Host:     utils.GetEnvWithDefault("DB_HOST", "localhost"),
		Port:     utils.GetEnvWithDefault("DB_PORT", "3306"),
		Database: utils.GetEnvWithDefault("DB_DATABASE", "helianthus"),
	}
	return dsn.String()
}

func init() {
	var err error
	if utils.GetEnvWithDefault("DB_DRIVER", "mysql") == "mysql" {
		Conn, err = gorm.Open(mysql.Open(generateDSN()), &gorm.Config{})
	}
	if utils.GetEnvWithDefault("DB_DRIVER", "mysql") == "postgres" {
		Conn, err = gorm.Open(postgres.Open(generateDSN()), &gorm.Config{})
	}
	if err != nil {
		panic(err)
	}
	cnp, err := Conn.DB()
	if err != nil {
		panic(err)
	}
	cnp.SetMaxIdleConns(20)
	cnp.SetMaxOpenConns(1000)
}