package main

import logger "github.com/isayme/go-logger"

func main() {
	logger.SetLevel("trace")

	logger.Trace("Hi")
	logger.Infof("Hi, %s", "世界")
	logger.Warnw("Hi")
	logger.Errorw("Hi", "str", "世界", "int", 5)

	logger.Warn("Hi")
	logger.Warnf("Hi, %s", "世界")
	logger.Warnw("Hi")
	logger.Warnw("Hi", "str", "世界", "int", 5)
	logger.Warnw("Hi", "5", "世界")

	logger.SetFormat(logger.FORMAT_JSON)
	logger.Trace("Hi")
	logger.Infof("Hi, %s", "世界")
	logger.SetFormat(logger.FORMAT_CONSOLE)
	logger.Warnw("Hi")
	logger.Errorw("Hi", "str", "世界", "int", 5)

	logger.SetLevel("trace")
	logger.Trace("should shown")
	logger.SetLevel("warn")
	logger.Trace("should not shown")
}
