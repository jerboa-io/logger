Simple Go log extension

Usage

logger.SetLogLevel(int)

0 = no logging except fatal
1 = errors and fatal 
3 = info, error, and fatal
5 = debug, info, error, and fatal

set a writer
logger.SetOutput(io.Writer)
example: logger.SetOutput(os.Stdout)


writing logs:

logger functions take variadic interfaces

logger.Debug("this is a debug message", 42, "some more detail")

the various functions will only log if the correct logging level is set.
This is useful when you combine with environment variables. As an example your development environment may use level 5 whereas your production environment may use level 1 to reduce log sizes and noise.

Available functions:

logger.Debug()
logger.Info()
logger.Error()
logger.Fatal()
Fatal will utilise log.Fatal() which is equivalent to l.Print() followed by a call to os.Exit(1).




