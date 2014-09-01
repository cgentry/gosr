package http

const (
	INVALID_USER_PWD =  "Invalid user or password"
	INVALID_SUBSCRIBER = "Invalid subscriber or password"

	TOKEN_INCOMPLETE =  "Authorization token is incomplete."
	TOKEN_MISSING_PARM = "Authorization token is missing userid or hmac value"
	TOKEN_MISSING = "Authorization token is missing"
	MD5_MISMATCH = "Checksum mismatch for Content-MD5"
	MD5_MISSING = "No Content-MD5 in header"
	TIMESTAMP_MISSING = "No date/time specified for key check"
	TIMESTAMP_RANGE = "Time is outside of time window"

	SECRET_INVALID = "The shared secret cannot be empty"

	SIGNATURE_INVALID="Signature on request is invalid"

)
