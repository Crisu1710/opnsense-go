package haproxy

import (
	"github.com/browningluke/opnsense-go/pkg/api"
)

const haproxyReconfigureEndpoint = "/haproxy/service/reconfigure"

// Controller for routes
type Controller struct {
	Api *api.Client
}

func (c *Controller) Client() *api.Client {
	return c.Api
}
