package errs

// Error err
type Error struct {
	Code    int
	Message string
}

// Error implement error method
func (e *Error) Error() string {
	return e.Message
}

// New create a custom error
func New(code int, message string) error {
	return &Error{Code: code, Message: message}
}

// Errors
var (
	ErrTokenInvalid    = New(1001, "Token is invalid")
	ErrTokenExpired    = New(1002, "Token is expired")
	ErrIssuerInvalid   = New(1003, "Issuer is invalid")
	ErrAudienceInvalid = New(1004, "Audience is invalid")
)
