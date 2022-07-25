package logger
import (
    "fmt"
    "io"
    "log"
    "strconv"
    "strings"
    "time"
)
type Logger struct {
    logLevel int
}
//New returns a new logger with a default log level of 3
func New() *Logger {
    log.SetFlags(0)
    return &Logger{
        logLevel: 3, // info, warning, error, and fatal
    }
}
/*SetLogLevel updates the logging level
This accepts an int or string for convenience
0 = no logging except fatal
1 = errors and fatal only
2 = warning, errors, and fatal only
3 = info, warning, error, and fatal
4 nci = info error and fatal
5 = debug info error and fatal
*/
func (l *Logger) SetLogLevel(level interface{}) {
    switch v := level.(type) {
    case int:
        l.logLevel = v
    case int32:
        l.logLevel = int(v)
    case int64:
        l.logLevel = int(v)
    case string:
        newLevel,err := strconv.Atoi(v)
        if err != nil {
            l.Error(err)
            return
        }
        l.logLevel = newLevel
    default:
        l.Error("unexpected type when setting logging level")
    }
 
}
//SetOutput directs the log output to a single writer
func (l *Logger) SetOutput(w io.Writer) {
    log.SetOutput(w)
}

//SetOutputs directs the log output to a multi writer
func (l *Logger) SetOutputs(w []io.Writer) {
    log.SetOutput(io.MultiWriter(w...))
}

//Debug log debug information
func (l *Logger) Debug(v ...interface{}) {
    if l.logLevel >= 5 {
        log.Printf(
            "level=%s\tts=%s\tmsg=%s",
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
            "level=%s\tts=%s\tmsg=%s",
            "INFO",
            time.Now().Format("2006-01-02 15:04"),
            fmt.Sprintf(strings.Repeat("%+v ", len(v)), v...),
        )
    }
}
//Warn log information
func (l *Logger) Warn(v ...interface{}) {
    if l.logLevel >= 2 {
        log.Printf(
            "level=%s\tts=%s\tmsg=%s",
            "WARN",
            time.Now().Format("2006-01-02 15:04"),
            fmt.Sprintf(strings.Repeat("%+v ", len(v)), v...),
        )
    }
}
//Error log errors
func (l *Logger) Error(v ...interface{}) {
    if l.logLevel >= 1 {
        log.Printf(
            "level=%s\tts=%s\tmsg=%s",
            "ERROR",
            time.Now().Format("2006-01-02 15:04"),
            fmt.Sprintf(strings.Repeat("%+v ", len(v)), v...),
        )
    }
}
//Fatal log errors and die
func (l *Logger) Fatal(v ...interface{}) {
    log.Fatalf(
        "level=%s\tts=%s\tmsg=%s",
        "FATAL",
        time.Now().Format("2006-01-02 15:04"),
        fmt.Sprintf(strings.Repeat("%+v ", len(v)), v...),
    )
}