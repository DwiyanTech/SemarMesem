
package logger

import (
	"time"
	 "fmt"
	 "go.uber.org/zap"
)


// Init Variables
var zapLogger *zap.Logger


func SetupZapLogger(){
	config := zap.NewDevelopmentEncoderConfig()
    config.EncodeTime = zapcore.ISO8601TimeEncoder
    config.EncodeLevel = zapcore.CapitalColorLevelEncoder
    config.OutputPaths = []string{"stdout"}
    zapLogger, err := config.Build()

    if err != nil {
    	 panic("Error configure Logger",err)
   }
}

func FunctionCounter(functionName string){
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", functionName, time.Since(start))
	}
}


func Info(msg string ){
	zapLogger.Info(msg)
}


func Error(msg string ){

	zapLogger.Error(msg)

}


func Warn(msg string ){

	zapLogger.Warn(msg)


}