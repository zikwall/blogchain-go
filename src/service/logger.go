package service

import (
	"fmt"
	"log"
	"os"
)

const (
	Reset  Color = "\033[0m"
	Red    Color = "\033[31m"
	Green  Color = "\033[32m"
	Yellow Color = "\033[33m"
	Cyan   Color = "\033[36m"
)

type (
	Color                   string
	BlogchainInternalLogger struct {
		handler  *os.File
		colorize BlogchainLoggerColorize
		debug    bool
	}
	BlogchainLoggerColorize struct{}
)

func NewBlogchainInternalLogger(isDebug bool) *BlogchainInternalLogger {
	l := new(BlogchainInternalLogger)
	l.handler = os.Stdout
	l.colorize = BlogchainLoggerColorize{}

	return l
}

func (b BlogchainInternalLogger) Info(message interface{}) {
	fmt.Printf("%s: %v \n", b.colorize.Colored("[BLOGCHAIN INFO]", Cyan), message)
}

func (b BlogchainInternalLogger) Warning(message interface{}) {
	fmt.Printf("%s: %v \n", b.colorize.Colored("[BLOGCHAIN WARNING]", Yellow), message)
}

func (b BlogchainInternalLogger) Error(error error) {
	if b.IsDebug() {
		log.Fatalf("%s: %v", b.colorize.Colored("[BLOGCHAIN ERROR]", Red), error)
	} else {
		b.Warning(error)
	}
}

func (b BlogchainInternalLogger) IsDebug() bool {
	return b.debug
}

func (c BlogchainLoggerColorize) Colored(content string, color Color) string {
	return string(color) + content + string(Reset)
}

func (b BlogchainInternalLogger) Close() error {
	return nil
}

func (b BlogchainInternalLogger) CloseMessage() string {
	return "Close internal blogchain logger"
}
