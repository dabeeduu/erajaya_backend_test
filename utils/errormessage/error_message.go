package errormessage

const (
	ErrorInvalidBody            = "invalid request"
	ErrorInvalidQueryParam      = "invalid query param"
	ErrorInternalError          = "an unexpected error occurred"
	ErrorFailToQuery            = "failed to query database"
	ErrorFailToScanRows         = "failed to scan rows"
	ErrorFailToExecQuery        = "failed to exec query"
	ErrorFailToGetRowsAffected  = "failed to get rows affected"
	ErrorNoRowsAffected         = "no rows affected"
	ErrorFailToBumpRedisVersion = "failed to bump redis version"
	ErrorFailToSetRedisCache    = "failed to set redis cache"
)
