package errno

var (
	// Common errors
	OK                  = &Errno{Code: 0, Message: "OK"}
	InternalServerError = &Errno{Code: 10001, Message: "Internal server error"}
	ErrBind             = &Errno{Code: 10002, Message: "Error occurred while binding the request body to the struct."}

	ErrTokenGenerate = &Errno{Code: 20003, Message: "Error occurred while signing the JSON web token."}

	ErrTokenInvalid = &Errno{Code: 20101, Message: "The token was invalid."}
	ErrAuthFailed   = &Errno{Code: 20101, Message: "The sid or password was incorrect."}

	// upload errors
	ErrFileNotFound = &Errno{Code: 20301, Message: "File not found"}
	ErrUploadFailed = &Errno{Code: 20303, Message: "Fail to upload file"}
)