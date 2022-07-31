package LogManager

import (
	"github.com/sirupsen/logrus"
	"os"
)

//日志打印

var (
	ErrorLogger *logrus.Logger
)

func newLogger(loggerName string) (ret *logrus.Logger) {
	ret = logrus.New()
	flag := os.O_RDWR | os.O_CREATE | os.O_APPEND
	src, err := os.OpenFile("./logs./ " + loggerName + ".txt", flag, 0666)
	if err != nil {
		panic(any("fatal error:can not open logs file"))
	}
	ret.Out = src
	return ret
}

func init() {
	os.Mkdir("./logs", 0666)
	ErrorLogger = newLogger("error")
}