package app

import (
	"fmt"
	log "github.com/V-Develop/funch-fullstack-assignment-backend/util"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
	Connection *gorm.DB
	Info       string
}

func InitDatabaseConfig(loger log.Loger) *Database {
	database = database.ReadDatabaseConfig(loger)
	database.ConnectDatabase(loger)
	return database
}

func (database *Database) ReadDatabaseConfig(loger log.Loger) *Database {
	var envs map[string]string
	envs, err := godotenv.Read(".env")
	if err != nil {
		loger.LogError("SERVER: failed to loading .env file: "+err.Error(), "0")
		panic(err.Error())
	}

	config := &Database{}

	config.Info = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		envs["DB_USER"],
		envs["DB_PASS"],
		envs["DB_HOST"],
		envs["DB_PORT"],
		envs["DB_NAME"])

	loger.LogInfo("SERVER: read database configuration successfully", "0")
	return config
}

func (database *Database) ConnectDatabase(loger log.Loger) {
	var err error

	db, err := gorm.Open(mysql.Open(database.Info), &gorm.Config{})
	if err != nil {
		loger.LogError("SERVER: failed to connect database mysql: "+err.Error(), "0")
		panic(err.Error())
	}

	loger.LogInfo("SERVER: connect database successfully", "0")
	database.Connection = db
}
