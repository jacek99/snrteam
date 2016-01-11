package common
import "errors"

// configure via environment variables to be Docker friendly
const ENV_HTTP_PORT = "HTTP_PORT"
const ENV_DB_FOLDER = "ENV_DB_FOLDER"

// errors
var (
	RECORD_NOT_FOUND_ERROR = errors.New("record_not_found_error")
)