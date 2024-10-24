package fail

type gerericError struct {
	Code    string
	Message string
}

func (ge *gerericError) Error() string {
	return ge.Message
}

type ValidationError struct {
	gerericError
}
type NotFoundError struct {
	gerericError
}
type AlreadyExistsError struct {
	gerericError
}
type InternalError struct {
	gerericError
}

func WithNotFoundError(code, message string) *NotFoundError {
	return &NotFoundError{gerericError{
		code, message,
	}}
}

func WithAlreadyExistsError(code, message string) *AlreadyExistsError {
	return &AlreadyExistsError{gerericError{
		code, message,
	}}
}

func WithValidationError(code, message string) *ValidationError {
	return &ValidationError{gerericError{
		code, message,
	}}
}

func WithInternalError(code, message string) *InternalError {
	return &InternalError{gerericError{
		code, message,
	}}
}
