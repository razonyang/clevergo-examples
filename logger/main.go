package main

import (
	"flag"

	"clevergo.tech/clevergo"
	"clevergo.tech/log"
	"github.com/sirupsen/logrus"
	"go.uber.org/zap"
)

var (
	logger    log.Logger
	loggerStr *string
)

func main() {
	loggerStr = flag.String("logger", "", "logrus, zap")
	flag.Parse()

	switch *loggerStr {
	case "logrus":
		logger = logrus.New()
	case "zap":
		zapLogger, _ := zap.NewDevelopment()
		logger = zapLogger.Sugar()
	}

	app := clevergo.New()
	app.Logger = logger
	app.Get("/log", handle)
	app.Run(":8080")
}

func handle(c *clevergo.Context) error {
	c.Logger().Debug("debug msg")
	c.Logger().Debugf("debugf msg")
	c.Logger().Info("debug")
	c.Logger().Infof("debugf")
	c.Logger().Warn("warn msg")
	c.Logger().Warnf("warnf msg")
	c.Logger().Error("error msg")
	c.Logger().Errorf("errorf msg")
	// c.Logger().Fatal("fatal msg")
	// c.Logger().Fatalf("fatalf msg")
	return nil
}
