package parser

import (
	"fmt"
	"github.com/willianantunes/go-playground/internal/file-watcher/dtos"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// Just to make it easier to read it as it's too long
var rawRegexLogLine = strings.Replace(`^(?P<address>[\d\w:\-.]+) 
(?P<identificationProtocolUser>.+?) 
(?P<remoteUser>.+?) 
\[(?P<whenRequestWasMade>.+)\] 
\"(?P<requestLine>.+?)\" 
(?P<statusCode>[0-9]{3}) 
(?P<bytesSent>[0-9]+) 
\"?(?P<referer>.+?)\"? 
\"?((?P<userAgent>.+?)\"?)$`, "\n", "", -1)

// The regexLogLine itself
var regexLogLine = regexp.MustCompile(rawRegexLogLine)

// Indexes of each named group
var addressIndex = regexLogLine.SubexpIndex("address")
var identificationProtocolUserIndex = regexLogLine.SubexpIndex("identificationProtocolUser")
var remoteUserIndex = regexLogLine.SubexpIndex("remoteUser")
var whenRequestWasMadeIndex = regexLogLine.SubexpIndex("whenRequestWasMade")
var requestLineIndex = regexLogLine.SubexpIndex("requestLine")
var statusCodeIndex = regexLogLine.SubexpIndex("statusCode")
var bytesSentIndex = regexLogLine.SubexpIndex("bytesSent")
var refererIndex = regexLogLine.SubexpIndex("referer")
var userAgentIndex = regexLogLine.SubexpIndex("userAgent")

// Another regex, now for the request line part of a given log entry
var regexRequestLine = regexp.MustCompile(`^(?P<httpMethod>[A-Z]+) (?P<requestPath>.+?) (?P<version>.+)`)
var httpMethodIndex = regexRequestLine.SubexpIndex("httpMethod")
var requestPathIndex = regexRequestLine.SubexpIndex("requestPath")
var versionIndex = regexRequestLine.SubexpIndex("version")

// https://yourbasic.org/golang/format-parse-string-time-date-example/
var layoutToParseTime = "02/Jan/2006:15:04:05 -0700"

func ExtractLineDetails(line string) *dtos.LogEntry {
	matchesLogLine := regexLogLine.FindStringSubmatch(line)
	// Values from each named group
	addressValue := matchesLogLine[addressIndex]
	identificationProtocolUserValue := matchesLogLine[identificationProtocolUserIndex]
	remoteUserValue := matchesLogLine[remoteUserIndex]
	whenRequestWasMadeValue := matchesLogLine[whenRequestWasMadeIndex]
	requestLineIndexValue := matchesLogLine[requestLineIndex]
	statusCodeValue := matchesLogLine[statusCodeIndex]
	bytesSentValue := matchesLogLine[bytesSentIndex]
	refererValue := matchesLogLine[refererIndex]
	userAgentValue := matchesLogLine[userAgentIndex]
	// Arranging some data
	// TODO: I know it's too much the bitSize! Gonna fix it later 👀
	statusCodeValueAsInt, err := strconv.ParseInt(statusCodeValue, 10, 32)
	if err != nil {
		panic("Could not retrieve status code. Has the log file changed?! 🤔")
	}
	bytesSentValueAsInt, err := strconv.ParseInt(bytesSentValue, 10, 32)
	if err != nil {
		panic("Could not retrieve bytes sent. Has the log file changed?! 🤔")
	}
	matchesRequestLine := regexRequestLine.FindStringSubmatch(requestLineIndexValue)
	httpMethodValue := matchesRequestLine[httpMethodIndex]
	requestPathValue := matchesRequestLine[requestPathIndex]
	versionValue := matchesRequestLine[versionIndex]
	whenRequestWasMadeConverted, err := time.Parse(layoutToParseTime, whenRequestWasMadeValue)
	if err != nil {
		message := fmt.Sprintf("Could not convert %s to Time. Has the log file changed?! 🤔", whenRequestWasMadeValue)
		panic(message)
	}
	// Creating DTO
	return &dtos.LogEntry{
		Address:                     retrieveValueOtherwiseEmpty(addressValue),
		WhenRequestWasMade:          whenRequestWasMadeConverted,
		IdentificationProtocolUser:  retrieveValueOtherwiseEmpty(identificationProtocolUserValue),
		RemoteUserWhenAuthenticated: retrieveValueOtherwiseEmpty(remoteUserValue),
		RequestLine: dtos.RequestLine{
			HTTPMethod:  httpMethodValue,
			RequestPath: requestPathValue,
			Protocol:    versionValue,
		},
		StatusCode: int(statusCodeValueAsInt),
		BytesSent:  int(bytesSentValueAsInt),
		Referer:    retrieveValueOtherwiseEmpty(refererValue),
		UserAgent:  retrieveValueOtherwiseEmpty(userAgentValue),
	}
}

func retrieveValueOtherwiseEmpty(value string) string {
	if value != "-" {
		return value
	}
	return ""
}
