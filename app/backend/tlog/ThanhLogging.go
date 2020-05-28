package tlog

import (
	"log"

	"go.uber.org/zap"
)

var logger *zap.Logger

var cfg = zap.NewDevelopmentConfig()

//Itf short for map[string]interface
type Itf map[string]interface{}

//Init ...
func Init() {
	var err error
	logger, err = cfg.Build()
	if err != nil {
		log.Println("Can not create log")
	}
	//logger.Info("Constructed logger success")
}

//GetLogger ...
func GetLogger() *zap.Logger {
	return logger
}

//Error log error msg
func Error(msg string, err error) {
	logger.Error(msg, zap.Error(err))
}

//Info ....
func Info(listinfo map[string]interface{}) {
	logger.Sugar().Info(listinfo)
}
