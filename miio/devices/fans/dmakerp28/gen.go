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
  V MiotV2Manufacturer
  Err error // Output only
}
// DeviceInformationModel Device Model
type DeviceInformationModel struct {
  V MiotV2Model
  Err error // Output only
}
// DeviceInformationSerialNumber Device Serial Number
type DeviceInformationSerialNumber struct {
  V MiotV2SerialNumber
  Err error // Output only
}
// DeviceInformationFirmwareRevision Current Firmware Version
type DeviceInformationFirmwareRevision struct {
  V MiotV2FirmwareRevision
  Err error // Output only
}
// FanOn Switch Status
type FanOn struct {
  V MiotV2On
  Err error // Output only
}
// FanFanLevel Gear Fan Level 
type FanFanLevel struct {
  V MiotV2FanLevel
  Err error // Output only
}
// FanMode Mode
type FanMode struct {
  V MiotV2Mode
  Err error // Output only
}
// FanHorizontalSwing Horizontal Swing
type FanHorizontalSwing struct {
  V MiotV2HorizontalSwing
  Err error // Output only
}
// FanHorizontalSwingIncludedAngle Horizontal Swing Included Angle
type FanHorizontalSwingIncludedAngle struct {
  V MiotV2HorizontalSwingIncludedAngle
  Err error // Output only
}
// FanVerticalSwing Vertical Swing
type FanVerticalSwing struct {
  V MiotV2VerticalSwing
  Err error // Output only
}
// FanVerticalSwingIncludedAngle Vertical Swing Included Angle
type FanVerticalSwingIncludedAngle struct {
  V MiotV2VerticalSwingIncludedAngle
  Err error // Output only
}
// IndicatorLightOn Switch Status
type IndicatorLightOn struct {
  V MiotV2On
  Err error // Output only
}
// AlarmAlarm Alarm
type AlarmAlarm struct {
  V MiotV2Alarm
  Err error // Output only
}
// PhysicalControlLockedPhysicalControlsLocked Physical Control Locked
type PhysicalControlLockedPhysicalControlsLocked struct {
  V MiotV2PhysicalControlsLocked
  Err error // Output only
}
// EnvironmentTemperature Temperature
type EnvironmentTemperature struct {
  V MiotV2Temperature
  Err error // Output only
}
// EnvironmentRelativeHumidity Relative Humidity
type EnvironmentRelativeHumidity struct {
  V MiotV2RelativeHumidity
  Err error // Output only
}
// OffDelayTimeOffDelayTime off-delay-time
type OffDelayTimeOffDelayTime struct {
  V DmakerOffDelayTime
  Err error // Output only
}
// DmServiceSpeedLevel speed-level
type DmServiceSpeedLevel struct {
  V DmakerSpeedLevel
  Err error // Output only
}
// DmServiceSwingUpdownManual swing-updown-manual
type DmServiceSwingUpdownManual struct {
  V DmakerSwingUpdownManual
  Err error // Output only
}
// DmServiceSwingLrManual swing-lr-manual
type DmServiceSwingLrManual struct {
  V DmakerSwingLrManual
  Err error // Output only
}
// DmServiceStartLeft start-left
type DmServiceStartLeft struct {
  V DmakerStartLeft
  Err error // Output only
}
// DmServiceStartRight start-right
type DmServiceStartRight struct {
  V DmakerStartRight
  Err error // Output only
}
// DmServiceStartUp start-up
type DmServiceStartUp struct {
  V DmakerStartUp
  Err error // Output only
}
// DmServiceStartDown start-down
type DmServiceStartDown struct {
  V DmakerStartDown
  Err error // Output only
}
// DmServiceSwingAll swing-all
type DmServiceSwingAll struct {
  V DmakerSwingAll
  Err error // Output only
}
// DmServiceBackToCenter back-to-center
type DmServiceBackToCenter struct {
  V DmakerBackToCenter
  Err error // Output only
}
// DmServiceOffToCenter off-to-center
type DmServiceOffToCenter struct {
  V DmakerOffToCenter
  Err error // Output only
}
// MiotV2Manufacturer Device Manufacturer
type MiotV2Manufacturer string
// MiotV2Model Device Model
type MiotV2Model string
// MiotV2SerialNumber Device Serial Number
type MiotV2SerialNumber string
// MiotV2FirmwareRevision Current Firmware Version
type MiotV2FirmwareRevision string
// MiotV2On Switch Status
type MiotV2On bool
// MiotV2FanLevel Gear Fan Level 
type MiotV2FanLevel uint8
const (
  MiotV2FanLevelLevel1 MiotV2FanLevel = 1
  MiotV2FanLevelLevel2 MiotV2FanLevel = 2
  MiotV2FanLevelLevel3 MiotV2FanLevel = 3
  MiotV2FanLevelLevel4 MiotV2FanLevel = 4
)
func (v MiotV2FanLevel) String() string {
  switch v {
  case MiotV2FanLevelLevel1:
    return "Level1"
  case MiotV2FanLevelLevel2:
    return "Level2"
  case MiotV2FanLevelLevel3:
    return "Level3"
  case MiotV2FanLevelLevel4:
    return "Level4"
  default:
    return fmt.Sprintf("unknown MiotV2FanLevel: %v", uint8(v))
  }
}
// MiotV2Mode Mode
type MiotV2Mode uint8
const (
  MiotV2ModeStraightWind MiotV2Mode = 0
  MiotV2ModeNaturalWind MiotV2Mode = 1
  MiotV2ModeSmart MiotV2Mode = 2
  MiotV2ModeSleep MiotV2Mode = 3
)
func (v MiotV2Mode) String() string {
  switch v {
  case MiotV2ModeStraightWind:
    return "Straight Wind"
  case MiotV2ModeNaturalWind:
    return "Natural Wind"
  case MiotV2ModeSmart:
    return "Smart"
  case MiotV2ModeSleep:
    return "Sleep"
  default:
    return fmt.Sprintf("unknown MiotV2Mode: %v", uint8(v))
  }
}
// MiotV2HorizontalSwing Horizontal Swing
type MiotV2HorizontalSwing bool
// MiotV2HorizontalSwingIncludedAngle Horizontal Swing Included Angle
type MiotV2HorizontalSwingIncludedAngle uint16
const (
  MiotV2HorizontalSwingIncludedAngle30 MiotV2HorizontalSwingIncludedAngle = 30
  MiotV2HorizontalSwingIncludedAngle60 MiotV2HorizontalSwingIncludedAngle = 60
  MiotV2HorizontalSwingIncludedAngle90 MiotV2HorizontalSwingIncludedAngle = 90
  MiotV2HorizontalSwingIncludedAngle120 MiotV2HorizontalSwingIncludedAngle = 120
  MiotV2HorizontalSwingIncludedAngle150 MiotV2HorizontalSwingIncludedAngle = 150
)
func (v MiotV2HorizontalSwingIncludedAngle) String() string {
  switch v {
  case MiotV2HorizontalSwingIncludedAngle30:
    return "30"
  case MiotV2HorizontalSwingIncludedAngle60:
    return "60"
  case MiotV2HorizontalSwingIncludedAngle90:
    return "90"
  case MiotV2HorizontalSwingIncludedAngle120:
    return "120"
  case MiotV2HorizontalSwingIncludedAngle150:
    return "150"
  default:
    return fmt.Sprintf("unknown MiotV2HorizontalSwingIncludedAngle: %v", uint16(v))
  }
}
// MiotV2VerticalSwing Vertical Swing
type MiotV2VerticalSwing bool
// MiotV2VerticalSwingIncludedAngle Vertical Swing Included Angle
type MiotV2VerticalSwingIncludedAngle uint8
const (
  MiotV2VerticalSwingIncludedAngle30 MiotV2VerticalSwingIncludedAngle = 30
  MiotV2VerticalSwingIncludedAngle60 MiotV2VerticalSwingIncludedAngle = 60
  MiotV2VerticalSwingIncludedAngle90 MiotV2VerticalSwingIncludedAngle = 90
)
func (v MiotV2VerticalSwingIncludedAngle) String() string {
  switch v {
  case MiotV2VerticalSwingIncludedAngle30:
    return "30"
  case MiotV2VerticalSwingIncludedAngle60:
    return "60"
  case MiotV2VerticalSwingIncludedAngle90:
    return "90"
  default:
    return fmt.Sprintf("unknown MiotV2VerticalSwingIncludedAngle: %v", uint8(v))
  }
}
// MiotV2Alarm Alarm
type MiotV2Alarm bool
// MiotV2PhysicalControlsLocked Physical Control Locked
type MiotV2PhysicalControlsLocked bool
// MiotV2Temperature Temperature
type MiotV2Temperature float64
// MiotV2RelativeHumidity Relative Humidity
type MiotV2RelativeHumidity uint8
// DmakerOffDelayTime off-delay-time
type DmakerOffDelayTime uint16
// DmakerSpeedLevel speed-level
type DmakerSpeedLevel uint8
// DmakerSwingUpdownManual swing-updown-manual
type DmakerSwingUpdownManual uint8
const (
  DmakerSwingUpdownManualNONE DmakerSwingUpdownManual = 0
  DmakerSwingUpdownManualUP DmakerSwingUpdownManual = 1
  DmakerSwingUpdownManualDOWN DmakerSwingUpdownManual = 2
)
func (v DmakerSwingUpdownManual) String() string {
  switch v {
  case DmakerSwingUpdownManualNONE:
    return "NONE"
  case DmakerSwingUpdownManualUP:
    return "UP"
  case DmakerSwingUpdownManualDOWN:
    return "DOWN"
  default:
    return fmt.Sprintf("unknown DmakerSwingUpdownManual: %v", uint8(v))
  }
}
// DmakerSwingLrManual swing-lr-manual
type DmakerSwingLrManual uint8
const (
  DmakerSwingLrManualNONE DmakerSwingLrManual = 0
  DmakerSwingLrManualLEFT DmakerSwingLrManual = 1
  DmakerSwingLrManualRIGHT DmakerSwingLrManual = 2
)
func (v DmakerSwingLrManual) String() string {
  switch v {
  case DmakerSwingLrManualNONE:
    return "NONE"
  case DmakerSwingLrManualLEFT:
    return "LEFT"
  case DmakerSwingLrManualRIGHT:
    return "RIGHT"
  default:
    return fmt.Sprintf("unknown DmakerSwingLrManual: %v", uint8(v))
  }
}
// DmakerStartLeft start-left
type DmakerStartLeft bool
// DmakerStartRight start-right
type DmakerStartRight bool
// DmakerStartUp start-up
type DmakerStartUp bool
// DmakerStartDown start-down
type DmakerStartDown bool
// DmakerSwingAll swing-all
type DmakerSwingAll bool
// DmakerBackToCenter back-to-center
type DmakerBackToCenter bool
// DmakerOffToCenter off-to-center
type DmakerOffToCenter bool