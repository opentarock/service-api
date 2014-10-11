package log

import (
	"os"

	"gopkg.in/inconshreveable/log15.v2"
)

func New(ctx ...interface{}) log15.Logger {
	logger := log15.New(ctx...)
	logger.SetHandler(log15.StreamHandler(os.Stdout, log15.LogfmtFormat()))
	return logger
}
