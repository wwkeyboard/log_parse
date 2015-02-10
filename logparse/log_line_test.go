package log_parse

import (
	"fmt"
	"testing"

	"github.com/wwkeyboard/log_parse/logparse"
)

func TestLogLineParse(t *testing.T) {
	// Given
	log_line := "162.243.191.67 - - [30/Jan/2015:04:15:37 +0000] \"GET / HTTP/1.1\" 304 0 \"-\" \"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/40.0.2214.93 Safari/537.36\""

	// When
	ll := log_parse.ParseLogLine(log_line)

	// Then
	if ll.IP != "162.243.191.67" {
		t.Error("IP didn't parse correctly")
	}

	if ll.At != "[30/Jan/2015:04:15:37 +0000]" {
		t.Error("Time not parsed correctle")
	}
}
