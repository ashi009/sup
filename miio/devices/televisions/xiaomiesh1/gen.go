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
  V MiotV2Manufacturer
  Err error // Output only
}
// DeviceInformationModel Device Model
type DeviceInformationModel struct {
  V MiotV2Model
  Err error // Output only
}
// DeviceInformationSerialNumber Device ID
type DeviceInformationSerialNumber struct {
  V MiotV2SerialNumber
  Err error // Output only
}
// DeviceInformationFirmwareRevision Current Firmware Version
type DeviceInformationFirmwareRevision struct {
  V MiotV2FirmwareRevision
  Err error // Output only
}
// DeviceInformationSerialNo Serial Number
type DeviceInformationSerialNo struct {
  V MiotV2SerialNo
  Err error // Output only
}
// TelevisionInputControl TV Input Control
type TelevisionInputControl struct {
  V MiotV2InputControl
  Err error // Output only
}
// TelevisionSleepMode Sleep Mode
type TelevisionSleepMode struct {
  V MiotV2SleepMode
  Err error // Output only
}
// SpeakerVolume Volume
type SpeakerVolume struct {
  V MiotV2Volume
  Err error // Output only
}
// SpeakerMute Mute
type SpeakerMute struct {
  V MiotV2Mute
  Err error // Output only
}
// MessageRouterRequest Request
type MessageRouterRequest struct {
  V MiotV2Request
  Err error // Output only
}
// MessageRouterResponse Response
type MessageRouterResponse struct {
  V MiotV2Response
  Err error // Output only
}
// NoDisturbNoDisturb No Disturb
type NoDisturbNoDisturb struct {
  V MiotV2NoDisturb
  Err error // Output only
}
// NoDisturbEnableTimePeriod Enable Time Period
type NoDisturbEnableTimePeriod struct {
  V MiotV2EnableTimePeriod
  Err error // Output only
}
// RemoteControlStatePlaying state-playing
type RemoteControlStatePlaying struct {
  V XiaomiStatePlaying
  Err error // Output only
}
// SpeakerModeIsOn is-on
type SpeakerModeIsOn struct {
  V XiaomiIsOn
  Err error // Output only
}
// MiotV2Manufacturer Device Manufacturer
type MiotV2Manufacturer string
// MiotV2Model Device Model
type MiotV2Model string
// MiotV2SerialNumber Device ID
type MiotV2SerialNumber string
// MiotV2FirmwareRevision Current Firmware Version
type MiotV2FirmwareRevision string
// MiotV2SerialNo Serial Number
type MiotV2SerialNo string
// MiotV2InputControl TV Input Control
type MiotV2InputControl uint8
const (
  MiotV2InputControlATV MiotV2InputControl = 1
  MiotV2InputControlDTV MiotV2InputControl = 2
  MiotV2InputControlAV MiotV2InputControl = 3
  MiotV2InputControlHDMI1 MiotV2InputControl = 4
  MiotV2InputControlHDMI2 MiotV2InputControl = 5
  MiotV2InputControlHDMI3 MiotV2InputControl = 6
)
func (v MiotV2InputControl) String() string {
  switch v {
  case MiotV2InputControlATV:
    return "ATV"
  case MiotV2InputControlDTV:
    return "DTV"
  case MiotV2InputControlAV:
    return "AV"
  case MiotV2InputControlHDMI1:
    return "HDMI 1"
  case MiotV2InputControlHDMI2:
    return "HDMI 2"
  case MiotV2InputControlHDMI3:
    return "HDMI 3"
  default:
    return fmt.Sprintf("unknown MiotV2InputControl: %v", uint8(v))
  }
}
// MiotV2SleepMode Sleep Mode
type MiotV2SleepMode bool
// MiotV2Volume Volume
type MiotV2Volume uint8
// MiotV2Mute Mute
type MiotV2Mute bool
// MiotV2Request Request
type MiotV2Request string
// MiotV2Response Response
type MiotV2Response string
// MiotV2NoDisturb No Disturb
type MiotV2NoDisturb bool
// MiotV2EnableTimePeriod Enable Time Period
type MiotV2EnableTimePeriod string
// XiaomiStatePlaying state-playing
type XiaomiStatePlaying bool
// XiaomiIsOn is-on
type XiaomiIsOn bool