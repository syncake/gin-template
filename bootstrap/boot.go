package bootstrap

import (
	"fmt"
	"gin-template/main/app/constant"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

func Boot() {
	constant.Config = loadConfig()
	constant.Logger = setupLogger("service")
}

func loadConfig() *viper.Viper {
	config := viper.New()
	config.SetConfigName("default")
	config.SetConfigType("yaml")
	config.AddConfigPath("./config/")

	err := config.ReadInConfig()
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	return config
}

func setupLogger(logName string) *logrus.Logger {
	logger := logrus.New()
	logger.Level = logrus.DebugLevel

	//Logger.Formatter = new(logrus.JSONFormatter)
	logger.Formatter = new(logrus.TextFormatter) //default
	//Logger.Formatter.(*logrus.TextFormatter).DisableColors = true    // remove colors
	//Logger.Formatter.(*logrus.TextFormatter).DisableTimestamp = true // remove timestamp from test output

	file, err := os.OpenFile(
		fmt.Sprintf("logs/%s.log", logName),
		os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err == nil {
		logger.Out = file
	} else {
		logger.Out = os.Stdout
		logger.Info("Failed to log to file, using default stderr")
	}

	return logger
}
