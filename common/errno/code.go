package errno

var (
	// Common errors
	OK                  = &Errno{Code: 0, Message: "OK"}
	InternalServerError = &Errno{Code: 10001, Message: "Internal server error"}
	ErrBind             = &Errno{Code: 10002, Message: "Error occurred while binding the request body to the struct."}

	ErrTokenGenerate = &Errno{Code: 20003, Message: "Error occurred while signing the JSON web token."}

	ErrTokenInvalid          = &Errno{Code: 20101, Message: "The token was invalid."}
	ErrAuthFailed            = &Errno{Code: 20102, Message: "The sid or password was incorrect."}
	ErrRequiredParamsMissing = &Errno{Code: 20103, Message: "Missing required query params"}

	// upload errors
	ErrFileNotFound = &Errno{Code: 20301, Message: "File not found"}
	ErrUploadFailed = &Errno{Code: 20302, Message: "Fail to upload file"}
	ErrParsing      = &Errno{Code: 20303, Message: "Error happened in parsing file"}
	ErrUploadScore  = &Errno{Code: 20304, Message: "Error happened in upload score"}

	// login errors
	ErrGetTeacherRole = &Errno{Code: 20401, Message: "Fail to get teacher's role"}

	// attchment errors
	ErrDeleteAttachment = &Errno{Code: 20501, Message: "fail to delete attachment"}
)
