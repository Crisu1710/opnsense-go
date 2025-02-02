package unbound

import (
	"context"
	"github.com/browningluke/opnsense-go/pkg/api"
)

var HostAliasOpts = api.ReqOpts{
	AddEndpoint:         "/unbound/settings/addHostAlias",
	GetEndpoint:         "/unbound/settings/getHostAlias",
	UpdateEndpoint:      "/unbound/settings/setHostAlias",
	DeleteEndpoint:      "/unbound/settings/delHostAlias",
	ReconfigureEndpoint: unboundReconfigureEndpoint,
	Monad:               "alias",
}

// Data structs

type HostAlias struct {
	Enabled     string          `json:"enabled"`
	Host        api.SelectedMap `json:"host"`
	Hostname    string          `json:"hostname"`
	Domain      string          `json:"domain"`
	Description string          `json:"description"`
}

// CRUD operations

func (c *Controller) AddHostAlias(ctx context.Context, resource *HostAlias) (string, error) {
	return api.Add(c.Client(), ctx, HostAliasOpts, resource)
}

func (c *Controller) GetHostAlias(ctx context.Context, id string) (*HostAlias, error) {
	return api.Get(c.Client(), ctx, HostAliasOpts, &HostAlias{}, id)
}

func (c *Controller) UpdateHostAlias(ctx context.Context, id string, resource *HostAlias) error {
	return api.Update(c.Client(), ctx, HostAliasOpts, resource, id)
}

func (c *Controller) DeleteHostAlias(ctx context.Context, id string) error {
	return api.Delete(c.Client(), ctx, HostAliasOpts, id)
}
