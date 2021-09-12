package dtos

import (
	"fmt"
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
	return fmt.Sprintf("Request from %v with status code %v through agent %v", logEntry.Address, logEntry.StatusCode, logEntry.UserAgent)
}
