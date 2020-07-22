package spoor_test

import (
	"bytes"
	"github.com/abdulg/spoor"
	"log"
	"testing"
)

func TestInfo(t *testing.T) {
	var tests = []struct {
		infoLog       string
		debugLog      string
		infoExpected  string
		debugExpected string
		debugOn       bool
	}{
		{"Hello, info", "", "info: Hello, info\n", "debug: \n", true},
		{"Hello, info", "", "info: Hello, info\n", "", false},
		{"", "Hello, debug", "info: \n", "debug: Hello, debug\n", true},
		{"", "Hello, debug", "info: \n", "", false},
		{"Hello, info", "Hello, debug", "info: Hello, info\n", "debug: Hello, debug\n", true},
	}
	for _, test := range tests {
		var infoBuf bytes.Buffer
		var debugBuf bytes.Buffer
		info := log.New(&infoBuf, "info: ", 0)
		debug := log.New(&debugBuf, "debug: ", 0)
		testSpoor := spoor.New(info, debug)

		testSpoor.Infof("%s", test.infoLog)
		if infoBuf.String() != test.infoExpected {
			t.Errorf("Expected '%s', received '%s'", test.infoExpected, infoBuf.String())
		}

		if test.debugOn == true {
			testSpoor.DebugOn()
		}
		testSpoor.Debugf("%s", test.debugLog)
		if debugBuf.String() != test.debugExpected {
			t.Errorf("Expected '%s', received '%s'", test.debugExpected, debugBuf.String())
		}
	}
}
