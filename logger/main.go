package logger

import (
	"io"
	"log"
	"os"
)

var (
	Warning *log.Logger
	Info    *log.Logger
	Error   *log.Logger
)

func MakeLoggers() error {
	file, err := os.OpenFile("./log/logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		return err
	}

	mw := io.MultiWriter(os.Stdout, file)

	Warning = log.New(mw, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	Info = log.New(mw, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(mw, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	return nil
}
