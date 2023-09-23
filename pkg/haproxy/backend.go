package haproxy

import (
	"context"
	"github.com/browningluke/opnsense-go/pkg/api"
)

var BackendOpts = api.ReqOpts{
	AddEndpoint:         "/haproxy/settings/addBackend",
	GetEndpoint:         "/haproxy/settings/getBackend",
	UpdateEndpoint:      "/haproxy/settings/setBackend",
	DeleteEndpoint:      "/haproxy/settings/delBackend",
	ReconfigureEndpoint: haproxyReconfigureEndpoint,
	Monad:               "backend",
}

// Data structs

type Backend struct {
	Algorithm                    api.SelectedMap     `json:"algorithm"`
	BaAdvertisedProtocols        api.SelectedMap     `json:"ba_advertised_protocols"`
	BasicAuthEnabled             string              `json:"basicAuthEnabled"`
	BasicAuthGroups              api.SelectedMapList `json:"basicAuthGroups"`
	BasicAuthUsers               api.SelectedMapList `json:"basicAuthUsers"`
	CheckDownInterval            string              `json:"checkDownInterval"`
	CheckInterval                string              `json:"checkInterval"`
	CustomOptions                string              `json:"customOptions"`
	Description                  string              `json:"description"`
	Enabled                      string              `json:"enabled"`
	HealthCheck                  api.SelectedMap     `json:"healthCheck"`
	HealthCheckEnabled           string              `json:"healthCheckEnabled"`
	HealthCheckFall              string              `json:"healthCheckFall"`
	HealthCheckLogStatus         string              `json:"healthCheckLogStatus"`
	HealthCheckRise              string              `json:"healthCheckRise"`
	Http2Enabled                 string              `json:"http2Enabled"`
	Http2EnabledNontls           string              `json:"http2Enabled_nontls"`
	LinkedActions                api.SelectedMapList `json:"linkedActions"`
	LinkedErrorfiles             api.SelectedMapList `json:"linkedErrorfiles"`
	LinkedFcgi                   api.SelectedMap     `json:"linkedFcgi"`
	LinkedMailer                 api.SelectedMap     `json:"linkedMailer"`
	LinkedResolver               api.SelectedMap     `json:"linkedResolver"`
	LinkedServers                api.SelectedMapList `json:"linkedServers"`
	Mode                         api.SelectedMap     `json:"mode"`
	Name                         string              `json:"name"`
	Persistence                  api.SelectedMap     `json:"persistence"`
	PersistenceCookiemode        api.SelectedMap     `json:"persistence_cookiemode"`
	PersistenceCookiename        string              `json:"persistence_cookiename"`
	PersistenceStripquotes       string              `json:"persistence_stripquotes"`
	ProxyProtocol                api.SelectedMap     `json:"proxyProtocol"`
	RandomDraws                  string              `json:"random_draws"`
	ResolvePrefer                api.SelectedMap     `json:"resolvePrefer"`
	ResolverOpts                 api.SelectedMapList `json:"resolverOpts"`
	Source                       string              `json:"source"`
	StickinessBytesInRatePeriod  string              `json:"stickiness_bytesInRatePeriod"`
	StickinessBytesOutRatePeriod string              `json:"stickiness_bytesOutRatePeriod"`
	StickinessConnRatePeriod     string              `json:"stickiness_connRatePeriod"`
	StickinessCookielength       string              `json:"stickiness_cookielength"`
	StickinessCookiename         string              `json:"stickiness_cookiename"`
	StickinessDataTypes          api.SelectedMapList `json:"stickiness_dataTypes"`
	StickinessExpire             string              `json:"stickiness_expire"`
	StickinessHttpErrRatePeriod  string              `json:"stickiness_httpErrRatePeriod"`
	StickinessHttpReqRatePeriod  string              `json:"stickiness_httpReqRatePeriod"`
	StickinessPattern            api.SelectedMap     `json:"stickiness_pattern"`
	StickinessSessRatePeriod     string              `json:"stickiness_sessRatePeriod"`
	StickinessSize               string              `json:"stickiness_size"`
	TuningCaching                string              `json:"tuning_caching"`
	TuningDefaultserver          string              `json:"tuning_defaultserver"`
	TuningHttpreuse              api.SelectedMap     `json:"tuning_httpreuse"`
	TuningNoport                 string              `json:"tuning_noport"`
	TuningRetries                string              `json:"tuning_retries"`
	TuningTimeoutCheck           string              `json:"tuning_timeoutCheck"`
	TuningTimeoutConnect         string              `json:"tuning_timeoutConnect"`
	TuningTimeoutServer          string              `json:"tuning_timeoutServer"`
}

// CRUD operations

func (c *Controller) AddBackend(ctx context.Context, resource *Backend) (string, error) {
	return api.Add(c.Client(), ctx, BackendOpts, resource)
}

func (c *Controller) GetBackend(ctx context.Context, id string) (*Backend, error) {
	return api.Get(c.Client(), ctx, BackendOpts, &Backend{}, id)
}

func (c *Controller) UpdateBackend(ctx context.Context, id string, resource *Backend) error {
	return api.Update(c.Client(), ctx, BackendOpts, resource, id)
}

func (c *Controller) DeleteBackend(ctx context.Context, id string) error {
	return api.Delete(c.Client(), ctx, BackendOpts, id)
}
