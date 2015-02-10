package log_parse

import (
	"strings"
)

type LogLine struct {
	IP       string // address of the caller
	At       string
	Method   string // request method
	Path     string
	Response int
	Client   string
}

func ParseLogLine(s string) (ll *LogLine) {
	parts := strings.Split(s, " ")

	ll = new(LogLine)

	ll.IP = parts[0]
	ll.At = parts[3]
	return
}
