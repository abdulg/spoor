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
	}{
		{"Hello, info", "", "info: Hello, info\n", ""},
	}
	for _, test := range tests {
		var infoBuf bytes.Buffer
		var debugBuf bytes.Buffer
		info := log.New(&infoBuf, "info: ", 0)
		debug := log.New(&debugBuf, "debug: ", 0)
		testSpoor := spoor.New(info, debug)

		if test.infoLog != "" {
			testSpoor.Infof("%s", test.infoLog)
			if infoBuf.String() != test.infoExpected {
				t.Errorf("Expected '%s', received '%s'", test.infoExpected, infoBuf.String())
			}
		}

		if test.debugLog != "" {
			testSpoor.Debugf("%s", test.debugLog)
			if infoBuf.String() != test.debugExpected {
				t.Errorf("Expected '%s', received '%s'", test.debugExpected, debugBuf.String())
			}
		}
	}
}
