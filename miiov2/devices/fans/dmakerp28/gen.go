package dmakerp28
import (
  "reflect"
  "fmt"

  "sup/miio/devices"
)
var _ fmt.Stringer
func init() {
  devices.RegisterProperty(reflect.TypeOf(DeviceInformationManufacturer{}), 1, 1)
  devices.RegisterProperty(reflect.TypeOf(DeviceInformationModel{}), 1, 2)
  devices.RegisterProperty(reflect.TypeOf(DeviceInformationSerialNumber{}), 1, 3)
  devices.RegisterProperty(reflect.TypeOf(DeviceInformationFirmwareRevision{}), 1, 4)
  devices.RegisterProperty(reflect.TypeOf(FanOn{}), 2, 1)
  devices.RegisterProperty(reflect.TypeOf(FanFanLevel{}), 2, 2)
  devices.RegisterProperty(reflect.TypeOf(FanMode{}), 2, 3)
  devices.RegisterProperty(reflect.TypeOf(FanHorizontalSwing{}), 2, 4)
  devices.RegisterProperty(reflect.TypeOf(FanHorizontalSwingIncludedAngle{}), 2, 5)
  devices.RegisterProperty(reflect.TypeOf(FanVerticalSwing{}), 2, 7)
  devices.RegisterProperty(reflect.TypeOf(FanVerticalSwingIncludedAngle{}), 2, 8)
  devices.RegisterProperty(reflect.TypeOf(IndicatorLightOn{}), 4, 1)
  devices.RegisterProperty(reflect.TypeOf(AlarmAlarm{}), 5, 1)
  devices.RegisterProperty(reflect.TypeOf(PhysicalControlLockedPhysicalControlsLocked{}), 7, 1)
  devices.RegisterProperty(reflect.TypeOf(EnvironmentTemperature{}), 9, 1)
  devices.RegisterProperty(reflect.TypeOf(EnvironmentRelativeHumidity{}), 9, 2)
  devices.RegisterProperty(reflect.TypeOf(OffDelayTimeOffDelayTime{}), 3, 1)
  devices.RegisterProperty(reflect.TypeOf(DmServiceSpeedLevel{}), 8, 1)
  devices.RegisterProperty(reflect.TypeOf(DmServiceSwingUpdownManual{}), 8, 2)
  devices.RegisterProperty(reflect.TypeOf(DmServiceSwingLrManual{}), 8, 3)
  devices.RegisterProperty(reflect.TypeOf(DmServiceStartLeft{}), 8, 4)
  devices.RegisterProperty(reflect.TypeOf(DmServiceStartRight{}), 8, 5)
  devices.RegisterProperty(reflect.TypeOf(DmServiceStartUp{}), 8, 6)
  devices.RegisterProperty(reflect.TypeOf(DmServiceStartDown{}), 8, 7)
  devices.RegisterProperty(reflect.TypeOf(DmServiceSwingAll{}), 8, 8)
  devices.RegisterProperty(reflect.TypeOf(DmServiceBackToCenter{}), 8, 9)
  devices.RegisterProperty(reflect.TypeOf(DmServiceOffToCenter{}), 8, 10)
}
// DeviceInformationManufacturer Device Manufacturer
type DeviceInformationManufacturer struct {
  V ManufacturerV
  Err error // Output only
}
// DeviceInformationModel Device Model
type DeviceInformationModel struct {
  V ModelV
  Err error // Output only
}
// DeviceInformationSerialNumber Device Serial Number
type DeviceInformationSerialNumber struct {
  V SerialNumberV
  Err error // Output only
}
// DeviceInformationFirmwareRevision Current Firmware Version
type DeviceInformationFirmwareRevision struct {
  V FirmwareRevisionV
  Err error // Output only
}
// FanOn Switch Status
type FanOn struct {
  V OnV
  Err error // Output only
}
// FanFanLevel Gear Fan Level 
type FanFanLevel struct {
  V FanLevelV
  Err error // Output only
}
// FanMode Mode
type FanMode struct {
  V ModeV
  Err error // Output only
}
// FanHorizontalSwing Horizontal Swing
type FanHorizontalSwing struct {
  V HorizontalSwingV
  Err error // Output only
}
// FanHorizontalSwingIncludedAngle Horizontal Swing Included Angle
type FanHorizontalSwingIncludedAngle struct {
  V HorizontalSwingIncludedAngleV
  Err error // Output only
}
// FanVerticalSwing Vertical Swing
type FanVerticalSwing struct {
  V VerticalSwingV
  Err error // Output only
}
// FanVerticalSwingIncludedAngle Vertical Swing Included Angle
type FanVerticalSwingIncludedAngle struct {
  V VerticalSwingIncludedAngleV
  Err error // Output only
}
// IndicatorLightOn Switch Status
type IndicatorLightOn struct {
  V OnV
  Err error // Output only
}
// AlarmAlarm Alarm
type AlarmAlarm struct {
  V AlarmV
  Err error // Output only
}
// PhysicalControlLockedPhysicalControlsLocked Physical Control Locked
type PhysicalControlLockedPhysicalControlsLocked struct {
  V PhysicalControlsLockedV
  Err error // Output only
}
// EnvironmentTemperature Temperature
type EnvironmentTemperature struct {
  V TemperatureV
  Err error // Output only
}
// EnvironmentRelativeHumidity Relative Humidity
type EnvironmentRelativeHumidity struct {
  V RelativeHumidityV
  Err error // Output only
}
// OffDelayTimeOffDelayTime off-delay-time
type OffDelayTimeOffDelayTime struct {
  V OffDelayTimeV
  Err error // Output only
}
// DmServiceSpeedLevel speed-level
type DmServiceSpeedLevel struct {
  V SpeedLevelV
  Err error // Output only
}
// DmServiceSwingUpdownManual swing-updown-manual
type DmServiceSwingUpdownManual struct {
  V SwingUpdownManualV
  Err error // Output only
}
// DmServiceSwingLrManual swing-lr-manual
type DmServiceSwingLrManual struct {
  V SwingLrManualV
  Err error // Output only
}
// DmServiceStartLeft start-left
type DmServiceStartLeft struct {
  V StartLeftV
  Err error // Output only
}
// DmServiceStartRight start-right
type DmServiceStartRight struct {
  V StartRightV
  Err error // Output only
}
// DmServiceStartUp start-up
type DmServiceStartUp struct {
  V StartUpV
  Err error // Output only
}
// DmServiceStartDown start-down
type DmServiceStartDown struct {
  V StartDownV
  Err error // Output only
}
// DmServiceSwingAll swing-all
type DmServiceSwingAll struct {
  V SwingAllV
  Err error // Output only
}
// DmServiceBackToCenter back-to-center
type DmServiceBackToCenter struct {
  V BackToCenterV
  Err error // Output only
}
// DmServiceOffToCenter off-to-center
type DmServiceOffToCenter struct {
  V OffToCenterV
  Err error // Output only
}
// ManufacturerV Device Manufacturer
type ManufacturerV string
// ModelV Device Model
type ModelV string
// SerialNumberV Device Serial Number
type SerialNumberV string
// FirmwareRevisionV Current Firmware Version
type FirmwareRevisionV string
// OnV Switch Status
type OnV bool
// FanLevelV Gear Fan Level 
type FanLevelV uint8
const (
  FanLevelLevel1 FanLevelV = 1
  FanLevelLevel2 FanLevelV = 2
  FanLevelLevel3 FanLevelV = 3
  FanLevelLevel4 FanLevelV = 4
)
func (v FanLevelV) String() string {
  switch v {
  case FanLevelLevel1:
    return "Level1"
  case FanLevelLevel2:
    return "Level2"
  case FanLevelLevel3:
    return "Level3"
  case FanLevelLevel4:
    return "Level4"
  default:
    return fmt.Sprintf("unknown FanLevel: %v", uint8(v))
  }
}
// ModeV Mode
type ModeV uint8
const (
  ModeStraightWind ModeV = 0
  ModeNaturalWind ModeV = 1
  ModeSmart ModeV = 2
  ModeSleep ModeV = 3
)
func (v ModeV) String() string {
  switch v {
  case ModeStraightWind:
    return "Straight Wind"
  case ModeNaturalWind:
    return "Natural Wind"
  case ModeSmart:
    return "Smart"
  case ModeSleep:
    return "Sleep"
  default:
    return fmt.Sprintf("unknown Mode: %v", uint8(v))
  }
}
// HorizontalSwingV Horizontal Swing
type HorizontalSwingV bool
// HorizontalSwingIncludedAngleV Horizontal Swing Included Angle
type HorizontalSwingIncludedAngleV uint16
const (
  HorizontalSwingIncludedAngle30 HorizontalSwingIncludedAngleV = 30
  HorizontalSwingIncludedAngle60 HorizontalSwingIncludedAngleV = 60
  HorizontalSwingIncludedAngle90 HorizontalSwingIncludedAngleV = 90
  HorizontalSwingIncludedAngle120 HorizontalSwingIncludedAngleV = 120
  HorizontalSwingIncludedAngle150 HorizontalSwingIncludedAngleV = 150
)
func (v HorizontalSwingIncludedAngleV) String() string {
  switch v {
  case HorizontalSwingIncludedAngle30:
    return "30"
  case HorizontalSwingIncludedAngle60:
    return "60"
  case HorizontalSwingIncludedAngle90:
    return "90"
  case HorizontalSwingIncludedAngle120:
    return "120"
  case HorizontalSwingIncludedAngle150:
    return "150"
  default:
    return fmt.Sprintf("unknown HorizontalSwingIncludedAngle: %v", uint16(v))
  }
}
// VerticalSwingV Vertical Swing
type VerticalSwingV bool
// VerticalSwingIncludedAngleV Vertical Swing Included Angle
type VerticalSwingIncludedAngleV uint8
const (
  VerticalSwingIncludedAngle30 VerticalSwingIncludedAngleV = 30
  VerticalSwingIncludedAngle60 VerticalSwingIncludedAngleV = 60
  VerticalSwingIncludedAngle90 VerticalSwingIncludedAngleV = 90
)
func (v VerticalSwingIncludedAngleV) String() string {
  switch v {
  case VerticalSwingIncludedAngle30:
    return "30"
  case VerticalSwingIncludedAngle60:
    return "60"
  case VerticalSwingIncludedAngle90:
    return "90"
  default:
    return fmt.Sprintf("unknown VerticalSwingIncludedAngle: %v", uint8(v))
  }
}
// AlarmV Alarm
type AlarmV bool
// PhysicalControlsLockedV Physical Control Locked
type PhysicalControlsLockedV bool
// TemperatureV Temperature
type TemperatureV float64
// RelativeHumidityV Relative Humidity
type RelativeHumidityV uint8
// OffDelayTimeV off-delay-time
type OffDelayTimeV uint16
// SpeedLevelV speed-level
type SpeedLevelV uint8
// SwingUpdownManualV swing-updown-manual
type SwingUpdownManualV uint8
const (
  SwingUpdownManualNONE SwingUpdownManualV = 0
  SwingUpdownManualUP SwingUpdownManualV = 1
  SwingUpdownManualDOWN SwingUpdownManualV = 2
)
func (v SwingUpdownManualV) String() string {
  switch v {
  case SwingUpdownManualNONE:
    return "NONE"
  case SwingUpdownManualUP:
    return "UP"
  case SwingUpdownManualDOWN:
    return "DOWN"
  default:
    return fmt.Sprintf("unknown SwingUpdownManual: %v", uint8(v))
  }
}
// SwingLrManualV swing-lr-manual
type SwingLrManualV uint8
const (
  SwingLrManualNONE SwingLrManualV = 0
  SwingLrManualLEFT SwingLrManualV = 1
  SwingLrManualRIGHT SwingLrManualV = 2
)
func (v SwingLrManualV) String() string {
  switch v {
  case SwingLrManualNONE:
    return "NONE"
  case SwingLrManualLEFT:
    return "LEFT"
  case SwingLrManualRIGHT:
    return "RIGHT"
  default:
    return fmt.Sprintf("unknown SwingLrManual: %v", uint8(v))
  }
}
// StartLeftV start-left
type StartLeftV bool
// StartRightV start-right
type StartRightV bool
// StartUpV start-up
type StartUpV bool
// StartDownV start-down
type StartDownV bool
// SwingAllV swing-all
type SwingAllV bool
// BackToCenterV back-to-center
type BackToCenterV bool
// OffToCenterV off-to-center
type OffToCenterV bool