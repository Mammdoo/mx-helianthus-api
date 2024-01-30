/*

LICENSE:  MIT
Author:   sine
Email:    sinerwr@gmail.com

*/

package utils

import (
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
)

var (
	DB *gorm.DB
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
		Username: GetEnvWithDefault("DB_USERNAME", "root"),
		Password: GetEnvWithDefault("DB_PASSWORD", "this-is-secret-123"),
		Host:     GetEnvWithDefault("DB_HOST", "localhost"),
		Port:     GetEnvWithDefault("DB_PORT", "3306"),
		Database: GetEnvWithDefault("DB_DATABASE", "helianthus"),
	}
	return dsn.String()
}

func init() {
	var err error
	if GetEnvWithDefault("DB_DRIVER", "mysql") == "mysql" {
		DB, err = gorm.Open(mysql.Open(generateDSN()), &gorm.Config{})
	}
	if GetEnvWithDefault("DB_DRIVER", "mysql") == "postgres" {
		DB, err = gorm.Open(postgres.Open(generateDSN()), &gorm.Config{})
	}
	if err != nil {
		panic(err)
	}
	cnp, err := DB.DB()
	if err != nil {
		panic(err)
	}
	cnp.SetMaxIdleConns(20)
	cnp.SetMaxOpenConns(1000)
}