package miio

import (
	"context"
	"net"
	"testing"
	"time"

	"sup/miio/devices/airconditioners/xiaomic16"
	"sup/miio/devices/bathheaters/yeelinkv6"
	"sup/miio/devices/fans/dmakerp28"
)

func TestTV(t *testing.T) {
	d, err := ConenctDevice(
		&net.UDPAddr{
			IP:   net.IPv4(192, 168, 1, 203),
			Port: 54321,
		},
		"6531383872346b79584e364434536370",
	)
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.Background()

	info, err := d.Info(ctx)
	if err != nil {
		t.Fatal(err)
	}
	t.Error(info)

	// props, err := d.GetProperties(Properties{
	// 	{DeviceID: "518128374", PropertyID: PropertyID{2, 1}},
	// 	{DeviceID: "518128374", PropertyID: PropertyID{2, 2}},
	// 	{DeviceID: "518128374", PropertyID: PropertyID{4, 1}},
	// 	{DeviceID: "518128374", PropertyID: PropertyID{4, 2}},
	// 	{DeviceID: "518128374", PropertyID: PropertyID{7, 1}},
	// 	// {PropertyID: PropertyID{2, 4}},
	// 	// {PropertyID: PropertyID{3, 2}},
	// 	// {PropertyID: PropertyID{3, 4}},
	// })
	// if err != nil {
	// 	t.Fatal(err)
	// }
	// for _, prop := range props {
	// 	t.Errorf("%#v", prop)
	// }

	props, err := d.GetProperties(ctx, Properties{
		{DeviceID: d.ID(), PropertyID: PropertyID{2, 1}},
	})
	if err != nil {
		t.Fatal(err)
	}
	for _, prop := range props {
		t.Errorf("%#v", prop)
	}

	props, err = d.SetProperties(ctx, Properties{
		{DeviceID: d.ID(), PropertyID: PropertyID{2, 1}, Value: false},
	})
	if err != nil {
		t.Fatal(err)
	}
	for _, prop := range props {
		t.Errorf("%#v", prop)
	}

	// var res interface{}
	// if err := d.Call("action", struct {
	// 	DeviceID string `json:"did"`
	// 	ActionID
	// 	In []interface{} `json:"in"`
	// }{
	// 	DeviceID: "518128374",
	// 	ActionID: ActionID{7, 2},
	// }, &res); err != nil {
	// 	t.Error(err)
	// }
	// t.Error(res)

	// //set_properties "[{'did':'<device_id>','value':True,'siid':2,'piid':1}]"
	// res = nil
	// if err := d.Call("set_properties", []interface{}{
	// 	map[string]interface{}{
	// 		"did":   "nothing",
	// 		"value": false,
	// 		"siid":  2,
	// 		"piid":  1,
	// 	},
	// }, &res); err != nil {
	// 	t.Fatal(err)
	// }
	// t.Error(res)
}

func TestAC(t *testing.T) {
	d, err := ConenctDevice(
		&net.UDPAddr{
			IP:   net.IPv4(192, 168, 1, 168),
			Port: 54321,
		},
		"9e66782fd03c1f933f40193d181134c0",
	)

	ctx := context.Background()
	ctx, _ = context.WithTimeout(ctx, 5*time.Second)

	info, err := d.Info(ctx)
	if err != nil {
		t.Fatal(err)
	}
	t.Error(info)

	var props struct {
		xiaomic16.AirConditionerOn
		xiaomic16.AirConditionerMode
		xiaomic16.AirConditionerTargetTemperature
		xiaomic16.AirConditionerTargetHumidity
		xiaomic16.EnvironmentTemperature
		xiaomic16.EnvironmentRelativeHumidity
	}
	if err := d.GetPropertiesV2(ctx, &props); err != nil {
		t.Error(err)
	}
	// t.Error(pretty.Sprint(props))
	t.Errorf("%#v", props)

	// var nprops = struct {
	// 	xiaomic16.AirConditionerOn
	// 	xiaomic16.AirConditionerMode
	// 	xiaomic16.AirConditionerTargetTemperature
	// }{
	// 	xiaomic16.AirConditionerOn{V: true},
	// 	xiaomic16.AirConditionerMode{V: xiaomic16.ModeCool},
	// 	xiaomic16.AirConditionerTargetTemperature{V: 27.5},
	// }

	var nprops = struct {
		xiaomic16.AirConditionerOn
		xiaomic16.FanPercent
		xiaomic16.FanControlVerticalSwing
		xiaomic16.AirConditionerUnStraightBlowing
	}{
		xiaomic16.AirConditionerOn{V: false},
		xiaomic16.FanPercent{V: 1},
		xiaomic16.FanControlVerticalSwing{V: true},
		xiaomic16.AirConditionerUnStraightBlowing{V: true},
	}
	if err := d.SetPropertiesV2(ctx, &nprops); err != nil {
		t.Error(err)
	}
	t.Errorf("%#v", nprops)
}

func TestAC2(t *testing.T) {
	d, err := ConenctDevice(
		&net.UDPAddr{
			IP:   net.IPv4(192, 168, 1, 111),
			Port: 54321,
		},
		"473f98a1ece8d8b8ae8252d5785b8bdd",
	)

	ctx := context.Background()
	ctx, _ = context.WithTimeout(ctx, 5*time.Second)

	info, err := d.Info(ctx)
	if err != nil {
		t.Fatal(err)
	}
	t.Error(info)

	var props struct {
		xiaomic16.AirConditionerOn
		xiaomic16.AirConditionerMode
		xiaomic16.AirConditionerTargetTemperature
		xiaomic16.AirConditionerTargetHumidity
		xiaomic16.EnvironmentTemperature
		xiaomic16.EnvironmentRelativeHumidity
	}
	if err := d.GetPropertiesV2(ctx, &props); err != nil {
		t.Error(err)
	}
	// t.Error(pretty.Sprint(props))
	t.Errorf("%#v", props)

	var nprops = struct {
		// xiaomic16.AirConditionerOn
		xiaomic16.FanPercent
		xiaomic16.FanControlVerticalSwing
		xiaomic16.AirConditionerUnStraightBlowing
	}{
		// xiaomic16.AirConditionerOn{V: true},
		xiaomic16.FanPercent{V: 1},
		xiaomic16.FanControlVerticalSwing{V: true},
		xiaomic16.AirConditionerUnStraightBlowing{V: true},
	}
	if err := d.SetPropertiesV2(ctx, &nprops); err != nil {
		t.Error(err)
	}
	t.Errorf("%#v", nprops)
}

func TestVent(t *testing.T) {
	d, err := ConenctDevice(
		&net.UDPAddr{
			IP:   net.IPv4(192, 168, 1, 209),
			Port: 54321,
		},
		"b84fe5ab47826e1575f9c1c3d22a1393",
	)
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.Background()
	ctx, _ = context.WithTimeout(ctx, 5*time.Second)

	// info, err := d.Info(ctx)
	// if err != nil {
	// 	t.Fatal(err)
	// }
	// t.Error(info)

	var props struct {
		yeelinkv6.PTCBathHeaterTemperature
		yeelinkv6.PTCBathHeaterMode
		yeelinkv6.PTCBathHeaterTargetTemperature
	}
	if err := d.GetPropertiesV2(ctx, &props); err != nil {
		t.Error(err)
	}
	// t.Error(pretty.Sprint(props))
	t.Errorf("%#v", props)
}

func TestFan(t *testing.T) {
	d, err := ConenctDevice(
		&net.UDPAddr{
			IP:   net.IPv4(192, 168, 1, 148),
			Port: 54321,
		},
		"6bcf039baf2374bd44d8a455d95b7738",
	)

	ctx := context.Background()
	ctx, _ = context.WithTimeout(ctx, 5*time.Second)

	info, err := d.Info(ctx)
	if err != nil {
		t.Fatal(err)
	}
	t.Error(info)

	nprops := struct {
		dmakerp28.FanOn
	}{
		dmakerp28.FanOn{V: false},
	}
	if err := d.SetPropertiesV2(ctx, &nprops); err != nil {
		t.Error(err)
	}
	// t.Error(pretty.Sprint(props))
	t.Errorf("%v", nprops)
}
