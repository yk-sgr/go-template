package api

import (
	"net/http"

	"github.com/yk-sgr/nexthook-api/pkg/domain"
)

// handleSignUp godoc
//	@Summary		User Sign Up
//	@Description	sign up a new user
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			body	body		domain.SignUpRequest	true	"Signup Object"
//	@Success		201		{object}	domain.SignUpResponse
//	@Failure		400		{object}	api.ErrorResponse
//	@Router			/auth/signup [post]
func (api *API) handleSignUp(w http.ResponseWriter, r *http.Request) {
	var dto *domain.SignUpRequest
	err := api.Receive(r, &dto)
	if err != nil {
		api.SendBadRequest(w)
		return
	}

	res, err := api.authService.SignUp(r.Context(), dto)
	if err != nil {
		api.Error(w, err)
		return
	}

	api.SendJSON(w, http.StatusCreated, res)
}

// handleSignIn godoc
//	@Summary		User Sign In
//	@Description	sign in a user
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			body	body		domain.SignInRequest	true	"Signin Object"
//	@Success		200		{object}	domain.SignInResponse
//	@Failure		400		{object}	api.ErrorResponse
//	@Router			/auth/signin [post]
func (api *API) handleSignIn(w http.ResponseWriter, r *http.Request) {
	var dto *domain.SignInRequest
	err := api.Receive(r, &dto)
	if err != nil {
		api.SendBadRequest(w)
		return
	}

	res, err := api.authService.SignIn(r.Context(), dto)
	if err != nil {
		api.Error(w, err)
		return
	}

	api.SendJSON(w, http.StatusOK, res)
}
