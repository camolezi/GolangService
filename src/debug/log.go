package debug

import (
	"log"
	"os"
)

//This is for logging errors and debug information during devlopment

//LogLevel defines the level of logging of the aplication
type LogLevel int

const (
	//ErrorLevel enable only errors logs
	ErrorLevel LogLevel = 0
	//WarningLevel enable only errors and warnings
	WarningLevel LogLevel = 1
	//DebugLevel enable debug,warning and errors logs
	DebugLevel LogLevel = 2
)

//Logger is the interface used for logging
type Logger interface {
	Warning() *log.Logger
	Error() *log.Logger
	Debug() *log.Logger
}

type logger struct {
	debugLog   *log.Logger
	warningLog *log.Logger
	errorLog   *log.Logger
	logLevel   LogLevel
}

func (l *logger) Warning() *log.Logger {
	return l.warningLog
}

func (l *logger) Error() *log.Logger {
	return l.errorLog
}

func (l *logger) Debug() *log.Logger {
	return l.debugLog
}

//dontWrite is used when we dont want debug or warning messages
type dontWrite struct {
}

func (*dontWrite) Write(p []byte) (n int, err error) {
	return 0, nil
}

//NewLogger return a new usable logger
func NewLogger(level LogLevel) Logger {

	errorLog := log.New(os.Stderr, "[ERROR]  ", log.Ldate|log.Lshortfile|log.Ltime)

	warningLog := log.New(&dontWrite{}, "", 0)
	debugLog := log.New(&dontWrite{}, "", 0)

	//Turn warnings on
	if level >= 1 {
		warningLog = log.New(os.Stderr, "[WARNING]  ", log.Ldate|log.Lshortfile|log.Ltime)
	}

	//Turn debug logs on
	if level >= 2 {
		debugLog = log.New(os.Stderr, "[DEBUG]  ", log.Lshortfile)
	}

	return &logger{errorLog: errorLog, warningLog: warningLog, debugLog: debugLog, logLevel: level}
}
