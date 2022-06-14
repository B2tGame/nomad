package noop

import (
	"context"
	"time"
  "fmt"
  "github.com/hashicorp/nomad/plugins/device"
	"github.com/hashicorp/nomad/plugins/shared/structs"
)

func (d *NoopDevice) doFingerprint(ctx context.Context, devices chan *device.FingerprintResponse) {
	defer close(devices)
	ticker := time.NewTimer(0)
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			ticker.Reset(time.Minute)
		}
		d.writeFingerprintToChannel(devices)
	}
}

func (d *NoopDevice) writeFingerprintToChannel(devicesChan chan<- *device.FingerprintResponse) {
    devices := make([]*device.Device, 0, d.gpusCount)
		for idx := 0; idx < d.gpusCount; idx++ {
			devices = append(devices, &device.Device{
				ID:      "noop-" + fmt.Sprintf("%04d", idx),
				Healthy: true,
				HwLocality: &device.DeviceLocality{},
			})
		}

		deviceGroup := &device.DeviceGroup{
			Vendor:  vendor,
			Type:    deviceType,
			Name:    "Noop GPU",
			Devices: devices,
			Attributes: map[string]*structs.Attribute{},
		}
		devicesChan <- device.NewFingerprint(deviceGroup)
}
