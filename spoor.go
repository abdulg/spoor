// The spoor package provides a logger with 2 levels, Info and Debug.
package spoor

import (
	"log"
	"os"
)

// SpoorInterface is an interface that define 2 loggers, Info and Debug, and a set of methods to turn the debug logger on/off.
type SpoorInterface interface {
	Infof(format string, v ...any)
	Info(v ...any)
	Debugf(format string, v ...any)
	Debug(v ...any)
	DebugOn()
	DebugOff()
}

// A Spoor represents an object that writes text to one of two loggers
type Spoor struct {
	info      *log.Logger
	debug     *log.Logger
	debugFlag bool
}

// New creates a new Spoor. The info variable provides a log.Logger for info level log output. The debug
// variable provides a log.Logger for debug level log output.
func New(info *log.Logger, debug *log.Logger) *Spoor {
	debugFlag := os.Getenv("SPOOR_DEBUG") != ""
	return &Spoor{info, debug, debugFlag}
}

// Infof prints to the info logger.
// Arguments are handled in the manner of log.Printf
func (l *Spoor) Infof(format string, v ...any) {
	l.info.Printf(format, v...)
}

// Info prints to the info logger.
// Arguments are handled in the manner of log.Printf
func (l *Spoor) Info(v ...any) {
	l.info.Print(v...)
}

// DebugOn activates the debug logger
func (l *Spoor) DebugOn() {
	l.debugFlag = true
}

// DebugOff deactivates the debug logger
func (l *Spoor) DebugOff() {
	l.debugFlag = false
}

// Debugf prints to the debug logger if the debug flag is set.
// Arguments are handled in the manner of log.Printf
func (l *Spoor) Debugf(format string, v ...any) {
	if l.debugFlag {
		l.debug.Printf(format, v...)
	}
}

// Debug prints to the debug logger if the debug flag is set.
// Arguments are handled in the manner of log.Printf
func (l *Spoor) Debug(v ...any) {
	if l.debugFlag {
		l.debug.Print(v...)
	}
}

var info = log.New(os.Stderr, "info: ", log.LstdFlags)
var debug = log.New(os.Stderr, "debug: ", log.LstdFlags)
var std = New(info, debug)

func Default() *Spoor {
	return std
}

// These function write to the std Spoor

// Infof prints to the info logger on the std Spoor
func Infof(format string, v ...any) {
	std.Infof(format, v...)
}

// Info prints to the info logger on the std Spoor
func Info(v ...any) {
	std.Info(v...)
}

// Debug prints to the info logger on the std Spoor
func Debug(v ...any) {
	std.Debug(v...)
}

// Debugf prints to the info logger on the std Spoor
func Debugf(format string, v ...any) {
	std.Debugf(format, v...)
}

// DebugOn activates the debug logger
func DebugOn() {
	std.DebugOn()
}

// DebugOff deactivates the debug logger
func DebugOff() {
	std.DebugOff()
}
