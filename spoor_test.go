package spoor_test

import (
	"bytes"
	"github.com/abdulg/spoor"
	"log"
	"testing"
)

func TestInfo(t *testing.T) {
	var infoBuf bytes.Buffer
	var debugBuf bytes.Buffer
	info := log.New(&infoBuf, "info: ", 0)
	debug := log.New(&debugBuf, "debug: ", 0)
	testSpoor := spoor.New(info, debug)
	infoExpected := "info: Hello, world\n"
	debugExpected := ""

	testSpoor.Infof("Hello, %s", "world")
	if infoBuf.String() != infoExpected {
		t.Errorf("Expected '%s', received '%s'", infoExpected, infoBuf.String())
	}

	if debugBuf.String() != "" {
		t.Errorf("Expected '%s', received '%s'", debugExpected, infoBuf.String())
	}
}
