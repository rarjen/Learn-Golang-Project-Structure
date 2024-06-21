package constantvar

// HEADER
const (
	HEADER_PLATFORM_TYPE        = "X-Platform"
	HEADER_PLATFORM_TYPE_MOBILE = "mobile"
	HEADER_PLATFORM_TYPE_WEB    = "web"
)

// CONTEXT
const (
	CONTEXT_PLATFORM_TYPE = "PLATFORM_TYPE"
	CONTEXT_AUTHORIZATION = "Authorization"
)

// HTTP
const (
	HTTP_MESSAGE_INVALID_INPUT = "%s tidak boleh kosong"
	ERROR_EOF                  = "EOF"
	HTTP_INVALID_BODY_REQUEST  = "invalid body request"
	HTTP_UNKOWN_REQUEST        = "unknown error"
)

// CONSTANT
const (
	CONFIG_STAGE_DEV  = "DEVELOPMENT"
	CONFIG_STAGE_PROD = "PRODUCTION"
)
