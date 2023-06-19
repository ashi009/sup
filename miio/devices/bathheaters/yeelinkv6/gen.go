package yeelinkv6
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
  devices.RegisterProperty(reflect.TypeOf(LightOn{}), 2, 1)
  devices.RegisterProperty(reflect.TypeOf(LightMode{}), 2, 2)
  devices.RegisterProperty(reflect.TypeOf(LightBrightness{}), 2, 3)
  devices.RegisterProperty(reflect.TypeOf(PTCBathHeaterMode{}), 3, 1)
  devices.RegisterProperty(reflect.TypeOf(PTCBathHeaterTargetTemperature{}), 3, 2)
  devices.RegisterProperty(reflect.TypeOf(PTCBathHeaterTemperature{}), 3, 3)
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
// LightOn Switch Status
type LightOn struct {
  V MiotV2On
  Err error // Output only
}
// LightMode Mode
type LightMode struct {
  V MiotV2Mode
  Err error // Output only
}
// LightBrightness Brightness
type LightBrightness struct {
  V MiotV2Brightness
  Err error // Output only
}
// PTCBathHeaterMode Mode
type PTCBathHeaterMode struct {
  V MiotV2Mode
  Err error // Output only
}
// PTCBathHeaterTargetTemperature Target Temperature
type PTCBathHeaterTargetTemperature struct {
  V MiotV2TargetTemperature
  Err error // Output only
}
// PTCBathHeaterTemperature Temperature
type PTCBathHeaterTemperature struct {
  V MiotV2Temperature
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
// MiotV2Mode Mode
type MiotV2Mode uint8
const (
  MiotV2ModeDay MiotV2Mode = 0
  MiotV2ModeNight MiotV2Mode = 1
)
func (v MiotV2Mode) String() string {
  switch v {
  case MiotV2ModeDay:
    return "Day"
  case MiotV2ModeNight:
    return "Night"
  default:
    return fmt.Sprintf("unknown MiotV2Mode: %v", uint8(v))
  }
}
// MiotV2Brightness Brightness
type MiotV2Brightness uint8
// MiotV2TargetTemperature Target Temperature
type MiotV2TargetTemperature uint8
// MiotV2Temperature Temperature
type MiotV2Temperature uint8