Simple Go log extension

Usage

var log = logger.New()

log.SetLogLevel(int)

0 = no logging except fatal
1 = errors and fatal 
3 = info, error, and fatal
5 = debug, info, error, and fatal

set a writer
log.SetOutput(io.Writer)
example: log.SetOutput(os.Stdout)


writing logs:

logger functions take variadic interfaces

log.Debug("this is a debug message", 42, "some more detail")

the various functions will only log if the correct logging level is set.
This is useful when you combine with environment variables. As an example your development environment may use level 5 whereas your production environment may use level 1 to reduce log sizes and noise.

Available functions:

log.Debug()

log.Info()

log.Error()

log.Fatal()

Fatal will utilise log.Fatal() which is equivalent to l.Print() followed by a call to os.Exit(1).




