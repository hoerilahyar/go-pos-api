package errors

// General Errors
var (
	ErrInternalServer     = New("ERR0601", "Internal server error")
	ErrBadRequest         = New("ERR0602", "Bad request")
	ErrNotFound           = New("ERR0603", "Data not found")
	ErrValidation         = New("ERR0604", "Validation failed")
	ErrConflict           = New("ERR0605", "Conflict occurred")
	ErrTooManyRequests    = New("ERR0606", "Too many requests")
	ErrServiceUnavailable = New("ERR0607", "Service temporarily unavailable")
	ErrTimeout            = New("ERR0608", "Request timeout")
	ErrForbidden          = New("ERR0609", "Forbidden access")
	ErrNotImplemented     = New("ERR0610", "Feature not implemented yet")
)

// Auth & Authorization
var (
	ErrUnauthorized            = New("ERR0701", "Unauthorized access")
	ErrInvalidCredentials      = New("ERR0702", "Invalid username or password")
	ErrTokenExpired            = New("ERR0703", "Authentication token has expired")
	ErrTokenInvalid            = New("ERR0704", "Invalid authentication token")
	ErrAccessDenied            = New("ERR0705", "Access denied for this resource")
	ErrUserNotActive           = New("ERR0706", "User account is not active")
	ErrEmailExist              = New("ERR0707", "Email already exist")
	ErrUsernameExist           = New("ERR0708", "Username already exist")
	ErrAuthHeaderMissing       = New("ERR0709", "Authorization header is missing")
	ErrAuthHeaderInvalidFormat = New("ERR0710", "Invalid Authorization header format")
	ErrTokenNotYetValid        = New("ERR0711", "Token is not yet valid")
	ErrTokenExpiredSecondary   = New("ERR0712", "Token expired (secondary check)")
)

// Database / Storage
var (

	// Database Connection Errors
	ErrDBAccessDenied       = New("ERR0801", "Access denied for user")
	ErrDBNotFound           = New("ERR0802", "Database not found")
	ErrDBConnection         = New("ERR0803", "Failed to connect to database")
	ErrDBServerLost         = New("ERR0804", "Lost connection to MySQL server")
	ErrDBTooManyConnections = New("ERR0805", "Too many connections to the database")

	// Query / Record Related Errors
	ErrDuplicateEntry = New("ERR0806", "Duplicate data entry")
	ErrRecordNotFound = New("ERR0807", "Table or record not found")
	ErrFieldNotFound  = New("ERR0808", "Unknown column in field list")
	ErrSQLSyntax      = New("ERR0809", "SQL syntax error")
	ErrTruncatedValue = New("ERR0810", "Incorrect or truncated value")

	// Constraint & FK Errors
	ErrForeignKeyViolation    = New("ERR0811", "Foreign key constraint fails")
	ErrForeignKeyReferenced   = New("ERR0812", "Cannot delete or update: row is referenced")
	ErrMissingDefaultValue    = New("ERR0813", "Field does not have a default value")
	ErrNullValueInNonNullable = New("ERR0814", "Column cannot be null")

	// Other general DB error
	ErrDBTransaction = New("ERR0815", "Database transaction failed")
	ErrDataCorrupted = New("ERR0816", "Data is corrupted or invalid")

	ErrInvalidField         = New("ERR0817", "Invalid field specified in query")
	ErrQuerySyntax          = New("ERR0818", "Query syntax error")
	ErrInvalidCharacter     = New("ERR0819", "Invalid character or data format")
	ErrInvalidDataType      = New("ERR0820", "Invalid data type for field")
	ErrMissingRequiredField = New("ERR0821", "Missing required field value")
	ErrNullValue            = New("ERR0822", "Null value not allowed")
	ErrInternal             = New("ERR0823", "Internal database error")
)

// Business Logic
var (
	ErrInsufficientStock  = New("ERR0901", "Insufficient stock")
	ErrInvalidStatus      = New("ERR0902", "Invalid status for this operation")
	ErrBookingUnavailable = New("ERR0903", "Booking is unavailable")
	ErrPaymentFailed      = New("ERR0904", "Payment failed")
	ErrQuotaExceeded      = New("ERR0905", "Quota exceeded")
	ErrAlreadyProcessed   = New("ERR0906", "Data already processed")
)

// Configuration / System
var (
	ErrMissingEnv       = New("ERR1001", "Missing environment configuration")
	ErrServiceNotReady  = New("ERR1002", "Service is not ready")
	ErrDependencyFailed = New("ERR1003", "External dependency failure")
	ErrRateLimited      = New("ERR1004", "Too many requests from this client")
)

// Validation
var (
	ErrUsernameFormat = New("ERR1101", "username must not be an email")
	ErrEmailFormat    = New("ERR1101", "invalid email format")
)

// File
var (
	ErrFileNotFound     = New("ERR1201", "File not found")
	ErrFileUploadFailed = New("ERR1202", "Failed to upload file")
)

// Hashing / Convert
var (
	ErrHashPassword             = New("ERR1301", "Failed to hash password")
	ErrCheckHashPassword        = New("ERR1302", "Failed to hash password")
	ErrConvertStrToInt          = New("ERR1303", "Failed to convert string to int")
	ErrConvertIntToStr          = New("ERR1304", "Failed to convert int to string")
	ErrMarshalJson              = New("ERR1305", "failed to marshaling data")
	ErrTokenSignatureInvalid    = New("ERR1306", "Invalid token signature")
	ErrSignatureMethod          = New("ERR1307", "Unexpected signing method")
	ErrTokenParsingFailed       = New("ERR1308", "Error during token parsing")
	ErrTokenMalformed           = New("ERR1309", "Token is malformed")
	ErrTokenClaimsParsingFailed = New("ERR1310", "Could not parse token claims")
	ErrTokenInvalidAfterParsing = New("ERR1311", "Token is not valid after parsing")
	ErrGenerateToken            = New("ERR1312", "Failed to generate token")
)

// Crypt
var (
	ErrEncrypt = New("ERR1304", "Failed to encrypt data")
	ErrDecrypt = New("ERR1305", "Failed to decrypt data")
)

var (
	// User errors
	ErrUserList   = New("ERR1401", "Failed to list users")
	ErrUserDetail = New("ERR1402", "Failed to get user detail")
	ErrUserCreate = New("ERR1403", "Failed to create user")
	ErrUserUpdate = New("ERR1404", "Failed to update user")
	ErrUserDelete = New("ERR1405", "Failed to delete user")

	// Auth errors
	ErrAuthRegister = New("ERR1406", "Failed to register user")
	ErrAuthLogin    = New("ERR1407", "Failed to login")

	// Policy errors
	ErrPolicyList   = New("ERR1408", "Failed to list policies")
	ErrPolicyShow   = New("ERR1409", "Failed to get policy detail")
	ErrPolicyCreate = New("ERR1410", "Failed to create policy")
	ErrPolicyUpdate = New("ERR1411", "Failed to update policy")
	ErrPolicyDelete = New("ERR1412", "Failed to delete policy")
	ErrPolicyAssign = New("ERR1413", "Failed to assign policy")
	ErrPolicyRevoke = New("ERR1414", "Failed to revoke policy")

	// Category errors
	ErrCategoryList   = New("ERR1415", "Failed to list categories")
	ErrCategoryShow   = New("ERR1416", "Failed to get category detail")
	ErrCategoryCreate = New("ERR1417", "Failed to create category")
	ErrCategoryUpdate = New("ERR1418", "Failed to update category")
	ErrCategoryDelete = New("ERR1419", "Failed to delete category")

	// Product errors
	ErrProductList   = New("ERR1420", "Failed to list products")
	ErrProductShow   = New("ERR1421", "Failed to get product detail")
	ErrProductCreate = New("ERR1422", "Failed to create product")
	ErrProductUpdate = New("ERR1423", "Failed to update product")
	ErrProductDelete = New("ERR1424", "Failed to delete product")
)
