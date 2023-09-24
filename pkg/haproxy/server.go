package haproxy

import (
	"context"
	"github.com/browningluke/opnsense-go/pkg/api"
)

var ServerOpts = api.ReqOpts{
	AddEndpoint:         "/haproxy/settings/addServer",
	GetEndpoint:         "/haproxy/settings/getServer",
	UpdateEndpoint:      "/haproxy/settings/setServer",
	DeleteEndpoint:      "/haproxy/settings/delServer",
	ReconfigureEndpoint: haproxyReconfigureEndpoint,
	Monad:               "server",
}

// Data structs

type Server struct {
	Address              string              `json:"address"`
	Advanced             string              `json:"advanced"`
	CheckDownInterval    string              `json:"checkDownInterval"`
	CheckInterval        string              `json:"checkInterval"`
	Checkport            string              `json:"checkport"`
	Description          string              `json:"description"`
	Enabled              string              `json:"enabled"`
	LinkedResolver       api.SelectedMap     `json:"linkedResolver"`
	MaxConnections       string              `json:"maxConnections"`
	Mode                 api.SelectedMap     `json:"mode"`
	MultiplexerProtocol  api.SelectedMap     `json:"multiplexer_protocol"`
	Name                 string              `json:"name"`
	Number               string              `json:"number"`
	Port                 string              `json:"port"`
	ResolvePrefer        api.SelectedMap     `json:"resolvePrefer"`
	ResolverOpts         api.SelectedMap     `json:"resolverOpts"`
	ServiceName          string              `json:"serviceName"`
	Source               string              `json:"source"`
	Ssl                  string              `json:"ssl"`
	SslCA                api.SelectedMapList `json:"sslCA"`
	SslClientCertificate api.SelectedMap     `json:"sslClientCertificate"`
	SslCRL               api.SelectedMap     `json:"sslCRL"`
	SslSNI               string              `json:"sslSNI"`
	SslVerify            string              `json:"sslVerify"`
	Type                 api.SelectedMap     `json:"type"`
	UnixSocket           api.SelectedMap     `json:"unix_socket"`
	Weight               string              `json:"weight"`
}

// CRUD operations

func (c *Controller) AddServer(ctx context.Context, resource *Server) (string, error) {
	return api.Add(c.Client(), ctx, ServerOpts, resource)
}

func (c *Controller) GetServer(ctx context.Context, id string) (*Server, error) {
	return api.Get(c.Client(), ctx, ServerOpts, &Server{}, id)
}

func (c *Controller) UpdateServer(ctx context.Context, id string, resource *Server) error {
	return api.Update(c.Client(), ctx, ServerOpts, resource, id)
}

func (c *Controller) DeleteServer(ctx context.Context, id string) error {
	return api.Delete(c.Client(), ctx, ServerOpts, id)
}
