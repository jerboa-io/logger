package logger

import (
    "fmt"
    "io"
    "log"
    "strings"
    "time"
)
//default log level
var logLevel = 3
/*SetLogLevel updates the logging level
0 = no logging except fatal
1 = errors and fatal only
2 nci = errors and fatal only
3 = info error and fatal
4 nci = info error and fatal
5 = debug info error and fatal
*/
func SetLogLevel(level int) {
    logLevel = level
}
func SetOutput(w io.Writer) {
    log.SetOutput(w)
}
//Debug log debug information
func Debug(l ...interface{}) {
    if logLevel >= 5 {
        log.Printf(
            "%s\t%s\t%s",
            "DEBUG",
            time.Now().Format("2006-01-02 15:04"),
            fmt.Sprintf(strings.Repeat("%+v ", len(l)), l...),
        )
    }
}
//Info log information
func Info(l ...interface{}) {
    if logLevel >= 3 {
        log.Printf(
            "%s\t%s\t%s",
            "INFO",
            time.Now().Format("2006-01-02 15:04"),
            fmt.Sprintf(strings.Repeat("%+v ", len(l)), l...),
        )
    }
}
//Error log errors
func Error(l ...interface{}) {
    if logLevel >= 1 {
        log.Printf(
            "%s\t%s\t%s",
            "ERROR",
            time.Now().Format("2006-01-02 15:04"),
            fmt.Sprintf(strings.Repeat("%+v ", len(l)), l...),
        )
    }
}
//Fatal log errors and die
func Fatal(l ...interface{}) {
    log.Fatalf(
        "%s\t%s\t%s",
        "FATAL",
        time.Now().Format("2006-01-02 15:04"),
        fmt.Sprintf(strings.Repeat("%+v ", len(l)), l...),
    )
}