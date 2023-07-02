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
// LightOn Switch Status
type LightOn struct {
  V OnV
  Err error // Output only
}
// LightMode Mode
type LightMode struct {
  V ModeV
  Err error // Output only
}
// LightBrightness Brightness
type LightBrightness struct {
  V BrightnessV
  Err error // Output only
}
// PTCBathHeaterMode Mode
type PTCBathHeaterMode struct {
  V ModeV
  Err error // Output only
}
// PTCBathHeaterTargetTemperature Target Temperature
type PTCBathHeaterTargetTemperature struct {
  V TargetTemperatureV
  Err error // Output only
}
// PTCBathHeaterTemperature Temperature
type PTCBathHeaterTemperature struct {
  V TemperatureV
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
// ModeV Mode
type ModeV uint8
const (
  ModeDay ModeV = 0
  ModeNight ModeV = 1
)
func (v ModeV) String() string {
  switch v {
  case ModeDay:
    return "Day"
  case ModeNight:
    return "Night"
  default:
    return fmt.Sprintf("unknown Mode: %v", uint8(v))
  }
}
// BrightnessV Brightness
type BrightnessV uint8
// TargetTemperatureV Target Temperature
type TargetTemperatureV uint8
// TemperatureV Temperature
type TemperatureV uint8