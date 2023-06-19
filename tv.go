package main

import (
	"github.com/brutella/hap/accessory"
	"github.com/brutella/hap/characteristic"
	"github.com/brutella/hap/log"
	"github.com/brutella/hap/service"
)

func ConnectTV() *accessory.Television {
	// d, err := miio.ConenctDevice(
	// 	&net.UDPAddr{
	// 		IP:   net.IPv4(192, 168, 1, 168),
	// 		Port: 54321,
	// 	},
	// 	"9e66782fd03c1f933f40193d181134c0",
	// )
	// if err != nil {
	// 	log.Fatal(err)
	// }
	tv := accessory.NewTelevision(accessory.Info{
		Name:         "TV",
		Manufacturer: "Xiaomi",
	})

	tv.Television.SleepDiscoveryMode.SetValue(characteristic.SleepDiscoveryModeNotDiscoverable)

	for _, cfg := range []struct {
		InputID uint64
		Name    string
		DevType int
		SrcType int
	}{
		{1, "ATV", characteristic.InputDeviceTypeTv, characteristic.InputSourceTypeTuner},
		{2, "DTV", characteristic.InputDeviceTypeTv, characteristic.InputSourceTypeTuner},
		{3, "AV", characteristic.InputDeviceTypeTv, characteristic.InputSourceTypeCompositeVideo},
		{4, "HDMI 1", characteristic.InputDeviceTypeOther, characteristic.InputSourceTypeHdmi},
		{5, "HDMI 2", characteristic.InputDeviceTypeOther, characteristic.InputSourceTypeHdmi},
		{6, "HDMI 3", characteristic.InputDeviceTypeOther, characteristic.InputSourceTypeHdmi},
	} {
		src := service.NewInputSource()
		src.Id = cfg.InputID
		src.ConfiguredName.Permissions = []string{characteristic.PermissionRead}
		src.IsConfigured.SetValue(characteristic.IsConfiguredConfigured)
		src.InputDeviceType.SetValue(cfg.DevType)
		src.AddC(src.InputDeviceType.C)
		src.InputSourceType.SetValue(cfg.SrcType)
		src.ConfiguredName.SetValue(cfg.Name)
		// src.CurrentVisibilityState.SetValue(characteristic.CurrentVisibilityStateHidden)
		// src.TargetVisibilityState.SetValue(characteristic.TargetVisibilityStateShown)
		// src.AddC(src.TargetVisibilityState.C)
		tv.AddS(src.S)
	}

	tv.Television.AddC(tv.Television.PowerModeSelection.C)
	tv.Television.ActiveIdentifier.ValidVals = []int{1, 2, 3, 4, 5, 6}
	tv.Television.RemoteKey.ValidVals = []int{
		characteristic.RemoteKeyArrowUp,
		characteristic.RemoteKeyArrowDown,
		characteristic.RemoteKeyArrowLeft,
		characteristic.RemoteKeyArrowRight,
		characteristic.RemoteKeySelect,
		characteristic.RemoteKeyBack,
		characteristic.RemoteKeyExit,
	}
	tv.Television.RemoteKey.OnSetRemoteValue(func(v int) error {
		log.Debug.Printf("RemoteKey: %d", v)
		return nil
	})

	return tv
}
