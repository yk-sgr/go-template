package api

import (
	"net/http"
	"strings"

	"github.com/yk-sgr/nexthook-api/pkg/domain"
)

func (api *API) Authenticate(w http.ResponseWriter, r *http.Request) *domain.User {
	tokenHeader := r.Header.Get("Authorization")
	splitted := strings.Split(tokenHeader, " ")

	// invalid token format
	if len(splitted) != 2 || splitted[0] != "Bearer" || len(splitted[1]) == 0 {
		api.SendUnauthorized(w)
		return nil
	}

	// validate token
	token := splitted[1]
	u, err := api.authService.ValidateToken(r.Context(), token)
	if err != nil {
		api.Error(w, err)
		return nil
	}
	return u
}
