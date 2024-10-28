package fail

type GerericError struct {
	Code    string
	Message string
}

func (ge *GerericError) Error() string {
	return ge.Message
}

type ValidationError struct {
	GerericError
}
type NotFoundError struct {
	GerericError
}
type AlreadyExistsError struct {
	GerericError
}
type InternalError struct {
	GerericError
}

func WithNotFoundError(code, message string) *NotFoundError {
	return &NotFoundError{GerericError{
		code, message,
	}}
}

func WithAlreadyExistsError(code, message string) *AlreadyExistsError {
	return &AlreadyExistsError{GerericError{
		code, message,
	}}
}

func WithValidationError(code, message string) *ValidationError {
	return &ValidationError{GerericError{
		code, message,
	}}
}

func WithInternalError(code, message string) *InternalError {
	return &InternalError{GerericError{
		code, message,
	}}
}
