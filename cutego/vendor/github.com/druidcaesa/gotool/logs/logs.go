package logs

import (
	"log"
	"os"
)

var (
	debugLog *log.Logger
	infoLog  *log.Logger
	errorLog *log.Logger
)

type Logs struct {
}

func init() {
	log.Println("init ...")
	debugLog = log.New(os.Stdout, "[DEBUG] ", log.Ldate|log.Ltime|log.Lshortfile)
	infoLog = log.New(os.Stdout, "[INFO] ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLog = log.New(os.Stderr, "[ERROR] ", log.Ldate|log.Ltime|log.Lshortfile)
}

func (Logs) DebugLog() *log.Logger {
	return debugLog
}
func (Logs) ErrorLog() *log.Logger {
	return errorLog
}
func (Logs) InfoLog() *log.Logger {
	return infoLog
}
