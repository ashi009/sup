package main

import (
	"context"
	"log"
	"net"
	"time"

	"sup/miio"
	"sup/miio/devices/airconditioners/xiaomic16"

	"github.com/brutella/hap/accessory"
	"github.com/brutella/hap/characteristic"
	"github.com/brutella/hap/service"
)

type AirConditioner struct {
	*accessory.A
	Thermostat *service.Thermostat
	Fan        *service.FanV2
}

func NewAirConditioner() *AirConditioner {
	var ac AirConditioner
	ac.A = accessory.New(accessory.Info{
		Name:         "Air Conditioner",
		Manufacturer: "Xiaomi",
	}, accessory.TypeAirConditioner)
	ac.Thermostat = service.NewThermostat()
	ac.Thermostat.AddC(ac.Thermostat.CurrentRelativeHumidity.C)
	ac.Thermostat.AddC(ac.Thermostat.TargetRelativeHumidity.C)
	ac.Thermostat.TargetHeatingCoolingState.ValidVals = []int{
		characteristic.TargetHeatingCoolingStateOff,
		characteristic.TargetHeatingCoolingStateHeat,
		characteristic.TargetHeatingCoolingStateCool,
	}
	ac.AddS(ac.Thermostat.S)
	ac.Fan = service.NewFanV2()
	ac.Fan.Active.Permissions = []string{characteristic.PermissionRead, characteristic.PermissionEvents}
	ac.Fan.AddC(ac.Fan.RotationSpeed.C)
	ac.Fan.RotationSpeed.SetMinValue(1)
	ac.Fan.AddC(ac.Fan.SwingMode.C)
	ac.AddS(ac.Fan.S)
	return &ac
}

func ConnectAirConditioner() *AirConditioner {
	d, err := miio.ConenctDevice(
		&net.UDPAddr{
			IP:   net.IPv4(192, 168, 1, 168),
			Port: 54321,
		},
		"9e66782fd03c1f933f40193d181134c0",
	)
	if err != nil {
		panic(err)
	}
	ac := NewAirConditioner()
	go func() {
		for range time.NewTicker(10 * time.Second).C {
			var props struct {
				xiaomic16.AirConditionerOn
				xiaomic16.AirConditionerMode
				xiaomic16.EnvironmentTemperature
				xiaomic16.AirConditionerTargetTemperature
				xiaomic16.EnvironmentRelativeHumidity
				xiaomic16.AirConditionerTargetHumidity
				xiaomic16.FanPercent
				xiaomic16.FanControlVerticalSwing
			}
			if err := d.GetPropertiesV2(context.Background(), &props); err != nil {
				log.Println(err)
				continue
			}
			ac.Thermostat.CurrentTemperature.SetValue(float64(props.EnvironmentTemperature.V))
			ac.Thermostat.TargetTemperature.SetValue(float64(props.AirConditionerTargetTemperature.V))
			ac.Thermostat.CurrentRelativeHumidity.SetValue(float64(props.EnvironmentRelativeHumidity.V))
			ac.Thermostat.TargetRelativeHumidity.SetValue(float64(props.AirConditionerTargetHumidity.V))
			if props.AirConditionerOn.V {
				switch props.AirConditionerMode.V {
				case xiaomic16.ModeCool:
					ac.Thermostat.CurrentHeatingCoolingState.SetValue(characteristic.CurrentHeatingCoolingStateCool)
				case xiaomic16.ModeHeat:
					ac.Thermostat.CurrentHeatingCoolingState.SetValue(characteristic.CurrentHeatingCoolingStateHeat)
				default:
					ac.Thermostat.CurrentHeatingCoolingState.SetValue(3)
				}
				ac.Fan.Active.SetValue(characteristic.ActiveActive)
			} else {
				ac.Thermostat.CurrentHeatingCoolingState.SetValue(characteristic.CurrentHeatingCoolingStateOff)
				ac.Fan.Active.SetValue(characteristic.ActiveInactive)
			}
			ac.Fan.RotationSpeed.SetValue(float64(props.FanPercent.V))
			if props.FanControlVerticalSwing.V {
				ac.Fan.SwingMode.SetValue(characteristic.SwingModeSwingEnabled)
			} else {
				ac.Fan.SwingMode.SetValue(characteristic.SwingModeSwingDisabled)
			}
		}
	}()
	ac.Thermostat.TargetTemperature.OnValueRemoteUpdate(func(v float64) {
		nprop := struct {
			xiaomic16.AirConditionerTargetTemperature
		}{
			xiaomic16.AirConditionerTargetTemperature{V: xiaomic16.TargetTemperatureV(v)},
		}
		if err := d.SetPropertiesV2(context.Background(), &nprop); err != nil {
			log.Println(err)
		}
	})
	ac.Thermostat.TargetHeatingCoolingState.OnValueRemoteUpdate(func(v int) {
		var nprop struct {
			xiaomic16.AirConditionerOn
			xiaomic16.AirConditionerMode
		}
		switch v {
		case characteristic.TargetHeatingCoolingStateOff:
			nprop.AirConditionerOn.V = false
		case characteristic.TargetHeatingCoolingStateHeat:
			nprop.AirConditionerOn.V = true
			nprop.AirConditionerMode.V = xiaomic16.ModeHeat
		case characteristic.TargetHeatingCoolingStateCool:
			nprop.AirConditionerOn.V = true
			nprop.AirConditionerMode.V = xiaomic16.ModeCool
		}
		if err := d.SetPropertiesV2(context.Background(), &nprop); err != nil {
			log.Println(err)
		}
	})
	ac.Thermostat.TargetRelativeHumidity.OnValueRemoteUpdate(func(v float64) {
		nprop := struct {
			xiaomic16.AirConditionerTargetHumidity
		}{
			xiaomic16.AirConditionerTargetHumidity{V: xiaomic16.TargetHumidityV(v)},
		}
		if err := d.SetPropertiesV2(context.Background(), &nprop); err != nil {
			log.Println(err)
		}
	})
	ac.Fan.RotationSpeed.OnValueRemoteUpdate(func(v float64) {
		nprop := struct {
			xiaomic16.FanPercent
		}{
			xiaomic16.FanPercent{V: xiaomic16.FanPercentV(v)},
		}
		if err := d.SetPropertiesV2(context.Background(), &nprop); err != nil {
			log.Println(err)
		}
	})
	ac.Fan.SwingMode.OnValueRemoteUpdate(func(v int) {
		nprop := struct {
			xiaomic16.FanControlVerticalSwing
		}{
			xiaomic16.FanControlVerticalSwing{V: v == characteristic.SwingModeSwingEnabled},
		}
		if err := d.SetPropertiesV2(context.Background(), &nprop); err != nil {
			log.Println(err)
		}
	})
	return ac
}
