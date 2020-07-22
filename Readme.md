![Test](https://github.com/abdulg/spoor/workflows/Test/badge.svg?branch=trunk)
# Spoor

Spoor is a logging package with support for 2 log levels, _Info_, and _Debug_. The package includes a standard logger. It's accessible using `spoor.Infof` and `spoor.Debugf`. The arguments to either are the same as for `fmt.Printf`. The standard logger logs both Info and Debug level logs to `stderr` using a `log.Logger`. Remember that `log.logger` appends a newline if one is not present in the text to be logged. Use `spoor.DebugOn()` and `spoor.DebugOff()` to activate and deactivate the `debug` logger.

Alternatively you can create your own logger using `spoor.New(info *log.Logger, debug *log.Logger)`, where `info` and `debug` are the Info and Debug level loggers respectively.
