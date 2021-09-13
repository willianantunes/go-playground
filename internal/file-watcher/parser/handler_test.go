package parser

import (
	"fmt"
	"github.com/willianantunes/go-playground/internal/file-watcher/dtos"
	"testing"
	"time"
)

func TestShouldExtractLineDetailsWhenLineHasAllFieldsSet(t *testing.T) {
	// Arrange
	referer := "https://www.willianantunes.com/blog/2021/08/query-compressed-logs-that-are-stored-in-s3-using-aws-athena/"
	requestPath := "/meujornal_api/req/3_0/items/id/3635320"
	sampleLinePlaceholder := `104.155.117.66 jafar iago [14/Jan/2020:14:17:44 +0000] "GET %s HTTP/1.1" 200 2060 "%s" "okhttp/3.11.0"`
	sampleLine := fmt.Sprintf(sampleLinePlaceholder, requestPath, referer)
	// Act
	logEntryFromSampleLine := ExtractLineDetails(sampleLine)
	// Assert
	expectedLogEntry := dtos.LogEntry{
		Address:                     "104.155.117.66",
		WhenRequestWasMade:          time.Date(2020, time.January, 14, 14, 17, 44, 0, time.UTC),
		IdentificationProtocolUser:  "jafar",
		Referer:                     referer,
		RemoteUserWhenAuthenticated: "iago",
		StatusCode:                  200,
		BytesSent:                   2060,
		UserAgent:                   "okhttp/3.11.0",
		RequestLine: dtos.RequestLine{
			HTTPMethod:  "GET",
			RequestPath: requestPath,
			Protocol:    "HTTP/1.1",
		},
	}
	if logEntryFromSampleLine.Equal(expectedLogEntry) == false {
		t.Errorf("ExtractLineDetails() = [%v], want [%v]", logEntryFromSampleLine, expectedLogEntry)
	}
}

func TestShouldExtractLineDetailsWhenLineHasSomeFieldsSet(t *testing.T) {
	// Arrange
	requestPath := "/api/v1/agrabah"
	sampleLinePlaceholder := `- - - [23/Jun/1989:08:25:59 -0300] "PUT %s HTTP/2.0" 500 206000 - -`
	sampleLine := fmt.Sprintf(sampleLinePlaceholder, requestPath)
	// Act
	logEntryFromSampleLine := ExtractLineDetails(sampleLine)
	// Assert
	expectedLogEntry := dtos.LogEntry{
		Address:                     "",
		WhenRequestWasMade:          time.Date(1989, time.June, 23, 8, 25, 59, 0, time.FixedZone("", -3*3600)),
		IdentificationProtocolUser:  "",
		Referer:                     "",
		RemoteUserWhenAuthenticated: "",
		StatusCode:                  500,
		BytesSent:                   206000,
		UserAgent:                   "",
		RequestLine: dtos.RequestLine{
			HTTPMethod:  "PUT",
			RequestPath: requestPath,
			Protocol:    "HTTP/2.0",
		},
	}
	if logEntryFromSampleLine.Equal(expectedLogEntry) == false {
		t.Errorf("ExtractLineDetails() = [%v], want [%v]", logEntryFromSampleLine, expectedLogEntry)
	}
}
