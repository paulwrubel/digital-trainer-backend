package main

import (
	"os"
	"os/signal"
	"strings"

	"github.com/sirupsen/logrus"
)

func main() {
	log := initLogger()
	log.Info("starting program")

	appData, err := initAppData(log)
	if err != nil {
		log.WithError(err).Fatalln("cannot initialize digital trainer data")
	}

	log.Infoln("starting API server")
	listenAndServe(log, appData)

	log.Infoln("blocking until signalled to shutdown")
	// make channel for interrupt signal
	c := make(chan os.Signal, 1)
	// tell os to send to chan when signal received
	signal.Notify(c, os.Interrupt)
	// wait for signal
	<-c

	log.Infoln("shutting down")
	os.Exit(0)
}

func initLogger() *logrus.Logger {
	log := logrus.New()
	log.SetOutput(os.Stdout)
	switch strings.ToUpper(os.Getenv("DTB_LOG_LEVEL")) {
	case "TRACE":
		log.SetLevel(logrus.TraceLevel)
	case "DEBUG":
		log.SetLevel(logrus.DebugLevel)
	case "INFO":
		log.SetLevel(logrus.InfoLevel)
	case "WARN":
		log.SetLevel(logrus.WarnLevel)
	case "ERROR":
		log.SetLevel(logrus.ErrorLevel)
	case "FATAL":
		log.SetLevel(logrus.FatalLevel)
	case "PANIC":
		log.SetLevel(logrus.PanicLevel)
	default:
		log.SetLevel(logrus.WarnLevel)
	}
	return log
}
