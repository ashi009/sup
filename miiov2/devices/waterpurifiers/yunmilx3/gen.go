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
// WaterPurifierStatus Status
type WaterPurifierStatus struct {
  V StatusV
  Err error // Output only
}
// TotalDissolvedSolidsSensorTdsIn Total Dissolved Solids for Input Water
type TotalDissolvedSolidsSensorTdsIn struct {
  V TdsInV
  Err error // Output only
}
// TotalDissolvedSolidsSensorTdsOut Total Dissolved Solids for Output Water
type TotalDissolvedSolidsSensorTdsOut struct {
  V TdsOutV
  Err error // Output only
}
// PPFilterFilterUsedTime Filter Used Time
type PPFilterFilterUsedTime struct {
  V FilterUsedTimeV
  Err error // Output only
}
// PPFilterFilterUsedFlow Filter Used Flow
type PPFilterFilterUsedFlow struct {
  V FilterUsedFlowV
  Err error // Output only
}
// C1FilterFilterUsedTime Filter Used Time
type C1FilterFilterUsedTime struct {
  V FilterUsedTimeV
  Err error // Output only
}
// C1FilterFilterUsedFlow Filter Used Flow
type C1FilterFilterUsedFlow struct {
  V FilterUsedFlowV
  Err error // Output only
}
// C2FilterFilterUsedTime Filter Used Time
type C2FilterFilterUsedTime struct {
  V FilterUsedTimeV
  Err error // Output only
}
// C2FilterFilterUsedFlow Filter Used Flow
type C2FilterFilterUsedFlow struct {
  V FilterUsedFlowV
  Err error // Output only
}
// ROFilterFilterUsedTime Filter Used Time
type ROFilterFilterUsedTime struct {
  V FilterUsedTimeV
  Err error // Output only
}
// ROFilterFilterUsedFlow Filter Used Flow
type ROFilterFilterUsedFlow struct {
  V FilterUsedFlowV
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
// StatusV Status
type StatusV uint8
const (
  StatusIdle StatusV = 1
  StatusPurifying StatusV = 2
)
func (v StatusV) String() string {
  switch v {
  case StatusIdle:
    return "Idle"
  case StatusPurifying:
    return "Purifying"
  default:
    return fmt.Sprintf("unknown Status: %v", uint8(v))
  }
}
// TdsInV Total Dissolved Solids for Input Water
type TdsInV uint16
// TdsOutV Total Dissolved Solids for Output Water
type TdsOutV uint16
// FilterUsedTimeV Filter Used Time
type FilterUsedTimeV uint16
// FilterUsedFlowV Filter Used Flow
type FilterUsedFlowV uint16