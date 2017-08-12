package api

import (
	"net/http"

	userService "github.com/tidepool-org/platform/user/service"
)

type Status struct {
	Version     string
	Environment string
	Server      interface{}
}

func (s *Standard) GetStatus(userServiceContext userService.Context) {
	status := &Status{
		Version: s.VersionReporter().Long(),
		Server:  s.StatusMiddleware().GetStatus(),
	}
	userServiceContext.RespondWithStatusAndData(http.StatusOK, status)
}
