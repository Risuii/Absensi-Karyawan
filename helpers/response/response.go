package response

import (
	"encoding/json"
	"net/http"
)

type Response interface {
	Err() (err error)
	JSON(w http.ResponseWriter) (err error)
}

type ResponseImpl struct {
	err    error
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

func Success(status string, data interface{}) (resp Response) {
	return &ResponseImpl{
		err:    nil,
		Status: status,
		Data:   data,
	}
}

func Error(status string, err error) (resp Response) {
	return &ResponseImpl{
		err:    err,
		Status: status,
		Data:   nil,
	}
}

func (r *ResponseImpl) getStatusCode(status string) (statusCode int) {
	switch status {
	case StatusOK:
		return http.StatusOK
	case StatusCreated:
		return http.StatusCreated
	case StatusBadRequest:
		return http.StatusBadRequest
	case StatusUnauthorized:
		return http.StatusUnauthorized
	case StatusForbiddend:
		return http.StatusForbidden
	case StatusNotFound:
		return http.StatusNotFound
	case StatusConflicted:
		return http.StatusConflict
	case StatusUnprocessableEntity:
		return http.StatusUnprocessableEntity
	case StatusInternalServerError:
		return http.StatusInternalServerError
	default:
		return http.StatusInternalServerError
	}
}

func (r *ResponseImpl) Err() (err error) {
	return r.err
}

func (r *ResponseImpl) JSON(w http.ResponseWriter) error {
	statusCode := r.getStatusCode(r.Status)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	return json.NewEncoder(w).Encode(r)
}
