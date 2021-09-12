package dtos

import (
	"strings"
	"testing"
	"time"
)

func TestShouldCreateLogEntry(t *testing.T) {
	// Arrange
	address := "104.155.117.66"
	whenRequestWasMade := time.Date(2020, time.January, 14, 14, 17, 44, 0, time.UTC)
	requestLine := strings.Split("GET /meujornal_api/req/3_0/items/id/3635320 HTTP/1.1", " ")
	httpMethod, requestPath, protocol := requestLine[0], requestLine[1], requestLine[2]
	statusCode := 200
	bytesSent := 2060
	userAgent := "okhttp/3.11.0"
	// Act
	logEntry := LogEntry{
		Address:            address,
		WhenRequestWasMade: whenRequestWasMade,
		StatusCode:         statusCode,
		BytesSent:          bytesSent,
		UserAgent:          userAgent,
		RequestLine: RequestLine{
			HTTPMethod:  httpMethod,
			RequestPath: requestPath,
			Protocol:    protocol,
		},
	}
	// Just to explore them
	t.Log(logEntry)
	t.Log(logEntry.RequestLine)
}
