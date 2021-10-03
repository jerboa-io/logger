package logger
import (
    "fmt"
    "io"
    "log"
    "strings"
    "time"
)
type Logger struct {
    logLevel int
}
//New returns a new logger with a default log level of 3
func New() *Logger {
    return &Logger{
        logLevel: 3,
    }
}
/*SetLogLevel updates the logging level
0 = no logging except fatal
1 = errors and fatal only
2 nci = errors and fatal only
3 = info error and fatal
4 nci = info error and fatal
5 = debug info error and fatal
*/
func (l *Logger) SetLogLevel(level int) {
    l.logLevel = level
}
//SetOutput directs the log output to a writer
func (l *Logger) SetOutput(w io.Writer) {
    log.SetOutput(w)
}
//Debug log debug information
func (l *Logger) Debug(v ...interface{}) {
    if l.logLevel >= 5 {
        log.Printf(
            "%s\t%s\t%s",
            "DEBUG",
            time.Now().Format("2006-01-02 15:04"),
            fmt.Sprintf(strings.Repeat("%+v ", len(v)), v...),
        )
    }
}
//Info log information
func (l *Logger) Info(v ...interface{}) {
    if l.logLevel >= 3 {
        log.Printf(
            "%s\t%s\t%s",
            "INFO",
            time.Now().Format("2006-01-02 15:04"),
            fmt.Sprintf(strings.Repeat("%+v ", len(v)), v...),
        )
    }
}
//Error log errors
func (l *Logger) Error(v ...interface{}) {
    if l.logLevel >= 1 {
        log.Printf(
            "%s\t%s\t%s",
            "ERROR",
            time.Now().Format("2006-01-02 15:04"),
            fmt.Sprintf(strings.Repeat("%+v ", len(v)), v...),
        )
    }
}
//Fatal log errors and die
func (l *Logger) Fatal(v ...interface{}) {
    log.Fatalf(
        "%s\t%s\t%s",
        "FATAL",
        time.Now().Format("2006-01-02 15:04"),
        fmt.Sprintf(strings.Repeat("%+v ", len(v)), v...),
    )
}