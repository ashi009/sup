package yunmilx3
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
  devices.RegisterProperty(reflect.TypeOf(WaterPurifierStatus{}), 2, 1)
  devices.RegisterProperty(reflect.TypeOf(TotalDissolvedSolidsSensorTdsIn{}), 3, 1)
  devices.RegisterProperty(reflect.TypeOf(TotalDissolvedSolidsSensorTdsOut{}), 3, 2)
  devices.RegisterProperty(reflect.TypeOf(PPFilterFilterUsedTime{}), 4, 1)
  devices.RegisterProperty(reflect.TypeOf(PPFilterFilterUsedFlow{}), 4, 2)
  devices.RegisterProperty(reflect.TypeOf(C1FilterFilterUsedTime{}), 5, 1)
  devices.RegisterProperty(reflect.TypeOf(C1FilterFilterUsedFlow{}), 5, 2)
  devices.RegisterProperty(reflect.TypeOf(C2FilterFilterUsedTime{}), 6, 1)
  devices.RegisterProperty(reflect.TypeOf(C2FilterFilterUsedFlow{}), 6, 2)
  devices.RegisterProperty(reflect.TypeOf(ROFilterFilterUsedTime{}), 7, 1)
  devices.RegisterProperty(reflect.TypeOf(ROFilterFilterUsedFlow{}), 7, 2)
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
// WaterPurifierStatus Status
type WaterPurifierStatus struct {
  V MiotV2Status
  Err error // Output only
}
// TotalDissolvedSolidsSensorTdsIn Total Dissolved Solids for Input Water
type TotalDissolvedSolidsSensorTdsIn struct {
  V MiotV2TdsIn
  Err error // Output only
}
// TotalDissolvedSolidsSensorTdsOut Total Dissolved Solids for Output Water
type TotalDissolvedSolidsSensorTdsOut struct {
  V MiotV2TdsOut
  Err error // Output only
}
// PPFilterFilterUsedTime Filter Used Time
type PPFilterFilterUsedTime struct {
  V MiotV2FilterUsedTime
  Err error // Output only
}
// PPFilterFilterUsedFlow Filter Used Flow
type PPFilterFilterUsedFlow struct {
  V MiotV2FilterUsedFlow
  Err error // Output only
}
// C1FilterFilterUsedTime Filter Used Time
type C1FilterFilterUsedTime struct {
  V MiotV2FilterUsedTime
  Err error // Output only
}
// C1FilterFilterUsedFlow Filter Used Flow
type C1FilterFilterUsedFlow struct {
  V MiotV2FilterUsedFlow
  Err error // Output only
}
// C2FilterFilterUsedTime Filter Used Time
type C2FilterFilterUsedTime struct {
  V MiotV2FilterUsedTime
  Err error // Output only
}
// C2FilterFilterUsedFlow Filter Used Flow
type C2FilterFilterUsedFlow struct {
  V MiotV2FilterUsedFlow
  Err error // Output only
}
// ROFilterFilterUsedTime Filter Used Time
type ROFilterFilterUsedTime struct {
  V MiotV2FilterUsedTime
  Err error // Output only
}
// ROFilterFilterUsedFlow Filter Used Flow
type ROFilterFilterUsedFlow struct {
  V MiotV2FilterUsedFlow
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
// MiotV2Status Status
type MiotV2Status uint8
const (
  MiotV2StatusIdle MiotV2Status = 1
  MiotV2StatusPurifying MiotV2Status = 2
)
func (v MiotV2Status) String() string {
  switch v {
  case MiotV2StatusIdle:
    return "Idle"
  case MiotV2StatusPurifying:
    return "Purifying"
  default:
    return fmt.Sprintf("unknown MiotV2Status: %v", uint8(v))
  }
}
// MiotV2TdsIn Total Dissolved Solids for Input Water
type MiotV2TdsIn uint16
// MiotV2TdsOut Total Dissolved Solids for Output Water
type MiotV2TdsOut uint16
// MiotV2FilterUsedTime Filter Used Time
type MiotV2FilterUsedTime uint16
// MiotV2FilterUsedFlow Filter Used Flow
type MiotV2FilterUsedFlow uint16