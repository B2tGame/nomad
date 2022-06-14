package noop

import (
	"context"
	"time"
	log "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/nomad/helper/pluginutils/loader"
	"github.com/hashicorp/nomad/plugins/base"
	"github.com/hashicorp/nomad/plugins/device"
	"github.com/hashicorp/nomad/plugins/shared/hclspec"
	"github.com/kr/pretty"
)

const (
	pluginName = "noop"
	pluginVersion = "v0.1.0"
	vendor = "noop"
	deviceType = device.DeviceTypeGPU
)


var (

	PluginID = loader.PluginID{
		Name:       pluginName,
		PluginType: base.PluginTypeDevice,
	}

	PluginConfig = &loader.InternalPluginConfig{
		Factory: func(ctx context.Context, logger log.Logger) interface{} { return NewNoopDevice(logger) },
	}


	pluginInfo = &base.PluginInfoResponse{
		Type:              base.PluginTypeDevice,
		PluginApiVersions: []string{device.ApiVersion010},
		PluginVersion:     pluginVersion,
		Name:              pluginName,
	}

	configSpec = hclspec.NewObject(map[string]*hclspec.Spec{
		"gpus": hclspec.NewDefault(
			hclspec.NewAttr("gpus", "number", false),
			hclspec.NewLiteral("0"),
		),
	})
)


type Config struct {
	GpusCount int `codec:"gpus"`
}
type NoopDevice struct {
	logger log.Logger
	gpusCount int
}

func NewNoopDevice(log log.Logger) *NoopDevice {
	return &NoopDevice{
		logger:  log.Named(pluginName),
	}
}

func (d *NoopDevice) PluginInfo() (*base.PluginInfoResponse, error) {
	return pluginInfo, nil
}

func (d *NoopDevice) ConfigSchema() (*hclspec.Spec, error) {
	return configSpec, nil
}

func (d *NoopDevice) SetConfig(c *base.Config) error {
	var config Config
	if err := base.MsgPackDecode(c.PluginConfig, &config); err != nil {
		return err
	}

	d.logger.Info("config set", "config", log.Fmt("% #v", pretty.Formatter(config)))
	d.gpusCount = config.GpusCount
	return nil
}

func (d *NoopDevice) Fingerprint(ctx context.Context) (<-chan *device.FingerprintResponse, error) {
	outCh := make(chan *device.FingerprintResponse)
	go d.doFingerprint(ctx, outCh)
	return outCh, nil
}



func (d *NoopDevice) Stats(ctx context.Context, interval time.Duration) (<-chan *device.StatsResponse, error) {
	return make(chan *device.StatsResponse), nil
}

func (d *NoopDevice) Reserve(deviceIDs []string) (*device.ContainerReservation, error) {
	return &device.ContainerReservation{}, nil
}
