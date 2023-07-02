package xiaomiesh1
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
  devices.RegisterProperty(reflect.TypeOf(DeviceInformationSerialNo{}), 1, 5)
  devices.RegisterProperty(reflect.TypeOf(TelevisionInputControl{}), 2, 1)
  devices.RegisterProperty(reflect.TypeOf(TelevisionSleepMode{}), 2, 2)
  devices.RegisterProperty(reflect.TypeOf(SpeakerVolume{}), 4, 1)
  devices.RegisterProperty(reflect.TypeOf(SpeakerMute{}), 4, 2)
  devices.RegisterProperty(reflect.TypeOf(MessageRouterRequest{}), 5, 1)
  devices.RegisterProperty(reflect.TypeOf(MessageRouterResponse{}), 5, 2)
  devices.RegisterProperty(reflect.TypeOf(NoDisturbNoDisturb{}), 9, 1)
  devices.RegisterProperty(reflect.TypeOf(NoDisturbEnableTimePeriod{}), 9, 2)
  devices.RegisterProperty(reflect.TypeOf(RemoteControlStatePlaying{}), 7, 1)
  devices.RegisterProperty(reflect.TypeOf(SpeakerModeIsOn{}), 8, 1)
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
// DeviceInformationSerialNumber Device ID
type DeviceInformationSerialNumber struct {
  V SerialNumberV
  Err error // Output only
}
// DeviceInformationFirmwareRevision Current Firmware Version
type DeviceInformationFirmwareRevision struct {
  V FirmwareRevisionV
  Err error // Output only
}
// DeviceInformationSerialNo Serial Number
type DeviceInformationSerialNo struct {
  V SerialNoV
  Err error // Output only
}
// TelevisionInputControl TV Input Control
type TelevisionInputControl struct {
  V InputControlV
  Err error // Output only
}
// TelevisionSleepMode Sleep Mode
type TelevisionSleepMode struct {
  V SleepModeV
  Err error // Output only
}
// SpeakerVolume Volume
type SpeakerVolume struct {
  V VolumeV
  Err error // Output only
}
// SpeakerMute Mute
type SpeakerMute struct {
  V MuteV
  Err error // Output only
}
// MessageRouterRequest Request
type MessageRouterRequest struct {
  V RequestV
  Err error // Output only
}
// MessageRouterResponse Response
type MessageRouterResponse struct {
  V ResponseV
  Err error // Output only
}
// NoDisturbNoDisturb No Disturb
type NoDisturbNoDisturb struct {
  V NoDisturbV
  Err error // Output only
}
// NoDisturbEnableTimePeriod Enable Time Period
type NoDisturbEnableTimePeriod struct {
  V EnableTimePeriodV
  Err error // Output only
}
// RemoteControlStatePlaying state-playing
type RemoteControlStatePlaying struct {
  V StatePlayingV
  Err error // Output only
}
// SpeakerModeIsOn is-on
type SpeakerModeIsOn struct {
  V IsOnV
  Err error // Output only
}
// ManufacturerV Device Manufacturer
type ManufacturerV string
// ModelV Device Model
type ModelV string
// SerialNumberV Device ID
type SerialNumberV string
// FirmwareRevisionV Current Firmware Version
type FirmwareRevisionV string
// SerialNoV Serial Number
type SerialNoV string
// InputControlV TV Input Control
type InputControlV uint8
const (
  InputControlATV InputControlV = 1
  InputControlDTV InputControlV = 2
  InputControlAV InputControlV = 3
  InputControlHDMI1 InputControlV = 4
  InputControlHDMI2 InputControlV = 5
  InputControlHDMI3 InputControlV = 6
)
func (v InputControlV) String() string {
  switch v {
  case InputControlATV:
    return "ATV"
  case InputControlDTV:
    return "DTV"
  case InputControlAV:
    return "AV"
  case InputControlHDMI1:
    return "HDMI 1"
  case InputControlHDMI2:
    return "HDMI 2"
  case InputControlHDMI3:
    return "HDMI 3"
  default:
    return fmt.Sprintf("unknown InputControl: %v", uint8(v))
  }
}
// SleepModeV Sleep Mode
type SleepModeV bool
// VolumeV Volume
type VolumeV uint8
// MuteV Mute
type MuteV bool
// RequestV Request
type RequestV string
// ResponseV Response
type ResponseV string
// NoDisturbV No Disturb
type NoDisturbV bool
// EnableTimePeriodV Enable Time Period
type EnableTimePeriodV string
// StatePlayingV state-playing
type StatePlayingV bool
// IsOnV is-on
type IsOnV bool