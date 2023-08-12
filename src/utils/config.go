package utils

import (
	"go-backend/logging"

	"github.com/pelletier/go-toml"
)


var Config toml.Tree

func LoadConfig() string {

	Config, err := toml.LoadFile("config.toml")
	if err != nil {
		panic(err)
	}

	if Config.Get("logger").(string) == "FileLogger" {
		logging.InitializeFileTransactionLog()
	} else if Config.Get("logger").(string) == "SQLogger" {
		logging.InitializeSQLTransactionLog()
	} else {
		panic("Unknown Logging format requested. Avaiable options are FileLogger and SQLogger")
	}
	

	return Config.Get("port").(string)
}
