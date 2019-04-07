package controllers

import (
	"context"

	"github.com/MontFerret/ferret-server/server/http"

	"github.com/MontFerret/ferret-server/pkg/auth"
	"github.com/MontFerret/ferret-server/pkg/common"
	"github.com/MontFerret/ferret-server/server/http/api/restapi/operations"
	"github.com/MontFerret/ferret-server/server/logging"
	"github.com/go-openapi/runtime/middleware"
	"github.com/pkg/errors"
)

type Auth struct {
	auth *auth.Service
}

func NewAuth(service *auth.Service) (*Auth, error) {
	if service == nil {
		return nil, errors.Wrap(common.ErrMissedArgument, "auth service")
	}

	return &Auth{service}, nil
}

func (ctl *Auth) TokenByCredentials(params operations.TokenByCredentialsParams) middleware.Responder {
	logger := logging.FromRequest(params.HTTPRequest)
	ctx := context.Background()

	user, err := ctl.auth.VerifyCredentials(ctx, auth.Credentials{
		Name:     *params.Body.Username,
		Password: []byte(*params.Body.Password),
	})
	if err != nil {
		logger.Error().
			Timestamp().
			Err(err).
			Str("username", *params.Body.Username).
			Msg("verify credentials")

		return http.Unauthorized()
	}

	if user == nil {
		return http.Unauthorized()
	}

	token, err := ctl.auth.BuildToken(user)
	if err != nil {
		logger.Error().
			Timestamp().
			Err(err).
			Str("username", user.Name).
			Msg("build token")

		return http.Unauthorized()
	}

	return operations.NewTokenByCredentialsOK().
		WithPayload(&operations.TokenByCredentialsOKBody{
			Token: &token,
		})
}
