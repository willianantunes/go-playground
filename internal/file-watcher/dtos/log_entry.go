package dtos

import (
	"fmt"
	"reflect"
	"time"
)

type RequestLine struct {
	HTTPMethod  string
	RequestPath string
	Protocol    string
}

func (requestLine RequestLine) String() string {
	return fmt.Sprintf("Protocol %v with method %v to request path %v", requestLine.Protocol, requestLine.HTTPMethod, requestLine.RequestPath)
}

type LogEntry struct {
	Address            string
	WhenRequestWasMade time.Time
	// https://en.wikipedia.org/wiki/Ident_protocol
	IdentificationProtocolUser  string
	RemoteUserWhenAuthenticated string
	RequestLine                 RequestLine
	StatusCode                  int
	BytesSent                   int
	Referer                     string
	UserAgent                   string
}

func (logEntry LogEntry) String() string {
	return fmt.Sprintf("(%s): Request from %v with status code %v through agent %v", logEntry.WhenRequestWasMade.Format(time.RFC1123Z), logEntry.Address, logEntry.StatusCode, logEntry.UserAgent)
}

func (logEntry LogEntry) Equal(logEntryArgument LogEntry) bool {
	bothTimesAreTheSame := logEntry.WhenRequestWasMade.Format(time.RFC1123Z) == logEntryArgument.WhenRequestWasMade.Format(time.RFC1123Z)
	// No problem at all since we're changing a copy
	logEntry.WhenRequestWasMade = logEntryArgument.WhenRequestWasMade
	return reflect.DeepEqual(logEntry, logEntryArgument) && bothTimesAreTheSame
}
