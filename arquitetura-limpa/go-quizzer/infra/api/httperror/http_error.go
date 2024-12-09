package httperror

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Marlliton/go-quizzer/domain/fail"
)

type httpError struct {
	Code    string
	Message string
}

func WriteError(err error, w http.ResponseWriter) {
	var internalError *fail.InternalError
	if errors.As(err, &internalError) {
		writeTypedError(w, http.StatusInternalServerError, &internalError.GerericError)
	}
	var notFoundError *fail.NotFoundError
	if errors.As(err, &notFoundError) {
		writeTypedError(w, http.StatusNotFound, &notFoundError.GerericError)
	}
	var validationError *fail.ValidationError
	if errors.As(err, &validationError) {
		writeTypedError(w, http.StatusBadRequest, &validationError.GerericError)
	}
	var alreadyExistsError *fail.AlreadyExistsError
	if errors.As(err, &alreadyExistsError) {
		writeTypedError(w, http.StatusConflict, &alreadyExistsError.GerericError)
	}

	writeTypedError(w, http.StatusInternalServerError, &internalError.GerericError)
}

func writeTypedError(w http.ResponseWriter, code int, err *fail.GerericError) {
	w.Header().Set("Content-Type", "application/json")

	errPayload := &httpError{
		Code:    err.Code,
		Message: err.Message,
	}

	body, _ := json.Marshal(errPayload)
	w.WriteHeader(code)
	w.Write(body)
}
