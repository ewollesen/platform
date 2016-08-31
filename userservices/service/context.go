package service

/* CHECKLIST
 * [x] Uses interfaces as appropriate
 * [x] Private package variables use underscore prefix
 * [x] All parameters validated
 * [x] All errors handled
 * [x] Reviewed for concurrency safety
 * [x] Code complete
 * [x] Full test coverage
 */

import (
	"github.com/tidepool-org/platform/service"
	"github.com/tidepool-org/platform/userservices/client"
)

type Context interface {
	service.Context

	UserServicesClient() client.Client

	AuthenticationDetails() client.AuthenticationDetails
	SetAuthenticationDetails(authenticationDetails client.AuthenticationDetails)
}

type HandlerFunc func(context Context)