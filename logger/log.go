package logger

import (
	"fmt"
	"os"

	"github.com/spf13/viper"

	log "github.com/sirupsen/logrus"
)

// LogLevel : debug, info, trace, error, fatal
var LogLevel string

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	l := viper.GetString("log.level")
	if l == "debug" {
		log.SetLevel(log.DebugLevel)
	} else if l == "info" {
		log.SetLevel(log.InfoLevel)
	} else if l == "trace" {
		log.SetLevel(log.TraceLevel)
	} else if l == "error" {
		log.SetLevel(log.ErrorLevel)
	} else if l == "fatal" {
		log.SetLevel(log.FatalLevel)
	} else {
		log.SetLevel(log.TraceLevel)
		l = "trace"
	}
	LogLevel = l
	fmt.Println(fmt.Sprintf("Log level set to '%s'", l))

}

//Warn => Log Everything
func Warn(l interface{}) {
	log.WithFields(log.Fields{
		"SERVICE": "WINGO",
	}).Warnln(l)
}

//Info => Log Everything
func Info(l interface{}) {
	log.WithFields(log.Fields{
		"SERVICE": "WINGO",
	}).Infoln(l)
}

//Debug => Log Everything
func Debug(l ...interface{}) {
	log.WithFields(log.Fields{
		"SERVICE": "WINGO",
	}).Debugln(l...)
}

//Error => Log Everything
func Error(l ...interface{}) {
	log.WithFields(log.Fields{
		"SERVICE": "WINGO",
	}).Errorln(l...)
}

//Fatal => Log Everything
func Fatal(l interface{}) {
	log.WithFields(log.Fields{
		"SERVICE": "WINGO",
	}).Fatalln(l)
}

//CheckOrFatal => If err is not nil make it Fatal
func CheckOrFatal(e error) {
	if e != nil {
		Fatal(e)
	}
}

// DebugWriter log writer interface
type DebugWriter interface {
	Debug(v ...interface{})
}

// Logger default logger
type Logger struct {
	DebugWriter
}

// Print format & print log
func (logger Logger) Print(values ...interface{}) {
	Debug(values...)
}
