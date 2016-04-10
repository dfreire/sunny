package main

import (
	"fmt"

	log "github.com/Sirupsen/logrus"
	"github.com/rifflock/lfshook"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(fmt.Errorf("Couldn't read the config file"))
	}

	log.AddHook(lfshook.NewHook(lfshook.PathMap{
		log.DebugLevel: "./log/log.txt",
		// log.InfoLevel:  "./log/log.txt",
		// log.WarnLevel:  "./log/log.txt",
		// log.ErrorLevel: "./log/log.txt",
		// log.FatalLevel: "./log/log.txt",
		// log.PanicLevel: "./log/log.txt",
	}))

	log.SetFormatter(&log.JSONFormatter{})
	// log.SetFormatter(&log.TextFormatter{FullTimestamp: true})
	// log.SetFormatter(&log.TextFormatter{})

	// hook, err := logrus_syslog.NewSyslogHook("", "", syslog.LOG_INFO, "")
	// if err == nil {
	// 	log.AddHook(hook)
	// }

	log.WithFields(log.Fields{"SUNNY_ENV": viper.Get("SUNNY_ENV")}).Info("initialized")
}

func main() {
}
