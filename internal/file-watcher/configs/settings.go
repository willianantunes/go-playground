package configs

import (
	"github.com/joho/godotenv"
	"path"
)

// Once I have a remote interpreter, this won't be needed anymore

var projectRootPath = retrieveFullPathFromRelative("../../../")
var whereDotFileIs = path.Join(projectRootPath, ".env.development")
var _ = godotenv.Load(whereDotFileIs)

// App settings

var AccessLogLocation = getEnvValueOrPanic("PARSER_ACCESS_LOG_LOCATION")
