package constantvar

// env related
const (
	DB_DSN_KONVE = "DB_DSN_KONVE"
	DB_DSN_MIS   = "DB_DSN_MIS"
	PORT         = "PORT"
	AUTH_API_URL = "AUTH_API_URL"
)

// route related
const (
	ROUTE_SYNC_DATA        = "sync-data"
	ROUTE_SYNC_DUMMY       = "sync-dummy"
	ROUTE_CHECK_NASABAH_ID = "check-nasabah-id"
	ROUTE_API_AUTH         = "/marketline/validate"
)

// db related
const (
	SECONDARY_DB_MIS                = "dailysalesdb"
	DB_SQL_STRING_CONNECTION_FORMAT = "Server=%s;User ID=%s; Password=%s;Database=%s;Trusted_Connection=False;Encrypt=disable;"
)

// http related
const (
	AUTHORIZATION_SPECIAL_CASE = "Authorization"
	API_LOWER_CASE             = "api"
	ENCODED_DECODED_DB_API     = "EncodeDecodeDb"
)

// stage related
const (
	STAGE_DEVELOPMENT        = "DEVELOPMENT"
	STAGE_PRODUCTION         = "PRODUCTION"
	DB_LOG_LEVEL_DEVELOPMENT = 4
	DB_LOG_LEVEL_PRODUCTION  = 1
)

// http response related
const (
	HTTP_MESSAGE_INVALID_INPUT    = "%s tidak boleh null & kosong"
	HTTP_RESPONSE_FAILED_TO_FETCH = "failed to fetch"
	HTTP_RESPONSE_DATA_NOT_FOUND  = "data not found"
	HTTP_RESPONSE_SUCCESS         = "success"
)

// External URL
const (
	EXTERNAL_URL_MMS_FE = "10.61.5.142:8787"
)
