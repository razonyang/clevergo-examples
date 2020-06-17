package main

import (
	"flag"

	"clevergo.tech/log"
	"github.com/sirupsen/logrus"
	"go.uber.org/zap"
)

var (
	logger  log.Logger
	adapter string
)

func init() {
	flag.StringVar(&adapter, "adapter", "zap", "adapter: zap, logrus")
	flag.Parse()
	switch adapter {
	case "logrus":
		logger = logrus.New()
	default:
		logger = zap.NewExample().Sugar()
	}
}

func main() {
	logger.Debug("Debug")
	logger.Debugf("%s", "Debugf")
	logger.Info("Info")
	logger.Infof("%s", "Infof")
	logger.Warn("Warn")
	logger.Warnf("%s", "Warnf")
	logger.Error("Error")
	logger.Errorf("%s", "Errorf")
	//logger.Fatal("Fatal")
	//logger.Fatalf("%s", "Fatalf")
}
