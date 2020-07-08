package facebook

// AppError is wrap error apps
type AppError struct {
	Code int
	Err  error
}

const (
	// AppErrorInvalidMethod is invalid http method
	AppErrorInvalidMethod = iota

	// AppErrorInvalidVersion invalid  version
	AppErrorInvalidVersion

	// AppErrorInvalidToken invalid  token
	AppErrorInvalidToken
)

const (
	//MessageAppErrorInvalid is error general app
	messageAppErrorInvalid = "Unknown Error APP"

	// MessageAppErrorInvalidMethod is message error for invalid http method
	messageAppErrorInvalidMethod = "Invalid HTTP Method"

	// messageAppErrorInvalidVersion error version api facebook
	messageAppErrorInvalidVersion = "Invalid Version Facebook API"

	// messageAppErrorInvalidToken error version api facebook
	messageAppErrorInvalidToken = "Invalid Facebook Access Token"
)

// Error is implementattion from error
func (ae AppError) Error() string {
	switch ae.Code {
	case AppErrorInvalidMethod:
		return messageAppErrorInvalidMethod

	case AppErrorInvalidVersion:
		return messageAppErrorInvalidVersion

	case AppErrorInvalidToken:
		return messageAppErrorInvalidToken

	default:
		return messageAppErrorInvalid
	}
}
