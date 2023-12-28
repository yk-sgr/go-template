package api

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/yk-sgr/nexthook-api/pkg/domain"
)

// Receive decodes the request body into dst.
func (api *API) Receive(r *http.Request, dst interface{}) error {
	return json.NewDecoder(r.Body).Decode(dst)
}

// SendJSON encodes data to JSON and sends it to the client.
func (api *API) SendJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func (api *API) Error(w http.ResponseWriter, error interface{}) {
	switch err := error.(type) {
	case *domain.NotFoundError:
		api.SendNotFound(w)
	case *domain.AlreadyExistsError:
		api.SendError(w, http.StatusConflict, err)
	case *domain.ValidationError:
		api.SendBadRequest(w)
	case *domain.UnauthorizedError:
		api.SendUnauthorized(w)
	case *domain.ForbiddenError:
		api.SendUnauthorized(w)
	default:
		api.SendInternalServerError(w)
	}
}

// SendError sends an error message to the client.
func (api *API) SendError(w http.ResponseWriter, status int, err error) {
	api.SendJSON(w, status, &ErrorResponse{Error: err.Error()})
}

// SendNotFound sends a not found error to the client.
func (api *API) SendNotFound(w http.ResponseWriter) {
	api.SendError(w, http.StatusNotFound, errors.New("not found"))
}

// SendBadRequest sends a bad request error to the client.
func (api *API) SendBadRequest(w http.ResponseWriter) {
	api.SendError(w, http.StatusBadRequest, errors.New("bad request"))
}

// SendUnauthorized sends an unauthorized error to the client.
func (api *API) SendUnauthorized(w http.ResponseWriter) {
	api.SendError(w, http.StatusUnauthorized, errors.New("unauthorized"))
}

// SendInternalServerError sends an internal server error to the client.
func (api *API) SendInternalServerError(w http.ResponseWriter) {
	api.SendError(w, http.StatusInternalServerError, errors.New("internal server error"))
}

// SendForbidden sends a forbidden error to the client.
func (api *API) SendForbidden(w http.ResponseWriter) {
	api.SendError(w, http.StatusForbidden, errors.New("forbidden"))
}
