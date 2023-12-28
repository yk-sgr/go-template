package api

import (
	"net/http"

	"github.com/yk-sgr/nexthook-api/pkg/domain"
)

// handleGetUser godoc
//	@Summary		Get user by ID
//	@Description	get a user by ID
//	@Tags			user
//	@Security		BearerAuth
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	domain.GetUserByIDResponse
//	@Failure		401	{object}	api.ErrorResponse
//	@Router			/users/{id} [get]
//	@Param			id	path	string	true	"User ID"
func (api *API) handleGetUserByID(w http.ResponseWriter, r *http.Request) {
	u := api.Authenticate(w, r)
	if u == nil {
		return
	}

	u, err := api.userService.GetUserByID(r.Context(), u.ID, u.ID)
	if err != nil {
		api.Error(w, err)
		return
	}

	api.SendJSON(w, http.StatusOK, &domain.GetUserByIDResponse{
		User: u,
	})
}
