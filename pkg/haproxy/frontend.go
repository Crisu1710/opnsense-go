package haproxy

import (
	"context"
	"github.com/browningluke/opnsense-go/pkg/api"
)

var FrontendOpts = api.ReqOpts{
	AddEndpoint:         "/haproxy/settings/addFrontend",
	GetEndpoint:         "/haproxy/settings/getFrontend",
	UpdateEndpoint:      "/haproxy/settings/setFrontend",
	DeleteEndpoint:      "/haproxy/settings/delFrontend",
	ReconfigureEndpoint: haproxyReconfigureEndpoint,
	Monad:               "frontend",
}

// Data structs

type Frontend struct {
	AdvertisedProtocols          api.SelectedMap `json:"advertised_protocols"`
	BasicAuthEnabled             string          `json:"basicAuthEnabled"`
	BasicAuthGroups              string          `json:"basicAuthGroups"`
	BasicAuthUsers               string          `json:"basicAuthUsers"`
	Bind                         string          `json:"bind"`
	BindOptions                  string          `json:"bindOptions"`
	ConnectionBehaviour          api.SelectedMap `json:"connectionBehaviour"`
	CustomOptions                string          `json:"customOptions"`
	DefaultBackend               api.SelectedMap `json:"defaultBackend"`
	Description                  string          `json:"description"`
	Enabled                      string          `json:"enabled"`
	ForwardFor                   string          `json:"forwardFor"`
	Http2Enabled                 string          `json:"http2Enabled"`
	Http2EnabledNontls           string          `json:"http2Enabled_nontls"`
	LinkedActions                string          `json:"linkedActions"`
	LinkedCpuAffinityRules       string          `json:"linkedCpuAffinityRules"`
	LinkedErrorfiles             string          `json:"linkedErrorfiles"`
	LoggingDetailedLog           string          `json:"logging_detailedLog"`
	LoggingDontLogNormal         string          `json:"logging_dontLogNormal"`
	LoggingDontLogNull           string          `json:"logging_dontLogNull"`
	LoggingLogSeparateErrors     string          `json:"logging_logSeparateErrors"`
	LoggingSocketStats           string          `json:"logging_socketStats"`
	Mode                         api.SelectedMap `json:"mode"`
	Name                         string          `json:"name"`
	PrometheusEnabled            string          `json:"prometheus_enabled"`
	PrometheusPath               string          `json:"prometheus_path"`
	SslAdvancedEnabled           string          `json:"ssl_advancedEnabled"`
	SslBindOptions               api.SelectedMap `json:"ssl_bindOptions"`
	SslCertificates              api.SelectedMap `json:"ssl_certificates"`
	SslCipherList                string          `json:"ssl_cipherList"`
	SslCipherSuites              string          `json:"ssl_cipherSuites"`
	SslClientAuthCAs             api.SelectedMap `json:"ssl_clientAuthCAs"`
	SslClientAuthCRLs            api.SelectedMap `json:"ssl_clientAuthCRLs"`
	SslClientAuthEnabled         string          `json:"ssl_clientAuthEnabled"`
	SslClientAuthVerify          api.SelectedMap `json:"ssl_clientAuthVerify"`
	SslCustomOptions             string          `json:"ssl_customOptions"`
	SslDefaultCertificate        api.SelectedMap `json:"ssl_default_certificate"`
	SslEnabled                   string          `json:"ssl_enabled"`
	SslHstsEnabled               string          `json:"ssl_hstsEnabled"`
	SslHstsIncludeSubDomains     string          `json:"ssl_hstsIncludeSubDomains"`
	SslHstsMaxAge                string          `json:"ssl_hstsMaxAge"`
	SslHstsPreload               string          `json:"ssl_hstsPreload"`
	SslMaxVersion                api.SelectedMap `json:"ssl_maxVersion"`
	SslMinVersion                api.SelectedMap `json:"ssl_minVersion"`
	StickinessBytesInRatePeriod  string          `json:"stickiness_bytesInRatePeriod"`
	StickinessBytesOutRatePeriod string          `json:"stickiness_bytesOutRatePeriod"`
	StickinessConnRatePeriod     string          `json:"stickiness_connRatePeriod"`
	StickinessCounter            string          `json:"stickiness_counter"`
	StickinessCounterKey         string          `json:"stickiness_counter_key"`
	StickinessDataTypes          api.SelectedMap `json:"stickiness_dataTypes"`
	StickinessExpire             string          `json:"stickiness_expire"`
	StickinessHttpErrRatePeriod  string          `json:"stickiness_httpErrRatePeriod"`
	StickinessHttpReqRatePeriod  string          `json:"stickiness_httpReqRatePeriod"`
	StickinessLength             string          `json:"stickiness_length"`
	StickinessPattern            api.SelectedMap `json:"stickiness_pattern"`
	StickinessSessRatePeriod     string          `json:"stickiness_sessRatePeriod"`
	StickinessSize               string          `json:"stickiness_size"`
	TuningMaxConnections         string          `json:"tuning_maxConnections"`
	TuningShards                 string          `json:"tuning_shards"`
	TuningTimeoutClient          string          `json:"tuning_timeoutClient"`
	TuningTimeoutHttpKeepAlive   string          `json:"tuning_timeoutHttpKeepAlive"`
	TuningTimeoutHttpReq         string          `json:"tuning_timeoutHttpReq"`
}

// CRUD operations

func (c *Controller) AddFrontend(ctx context.Context, resource *Frontend) (string, error) {
	return api.Add(c.Client(), ctx, FrontendOpts, resource)
}

func (c *Controller) GetFrontend(ctx context.Context, id string) (*Frontend, error) {
	return api.Get(c.Client(), ctx, FrontendOpts, &Frontend{}, id)
}

func (c *Controller) UpdateFrontend(ctx context.Context, id string, resource *Frontend) error {
	return api.Update(c.Client(), ctx, FrontendOpts, resource, id)
}

func (c *Controller) DeleteFrontend(ctx context.Context, id string) error {
	return api.Delete(c.Client(), ctx, FrontendOpts, id)
}
