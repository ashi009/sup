package xiaomic16
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
  devices.RegisterProperty(reflect.TypeOf(AirConditionerOn{}), 2, 1)
  devices.RegisterProperty(reflect.TypeOf(AirConditionerMode{}), 2, 2)
  devices.RegisterProperty(reflect.TypeOf(AirConditionerFault{}), 2, 3)
  devices.RegisterProperty(reflect.TypeOf(AirConditionerTargetTemperature{}), 2, 4)
  devices.RegisterProperty(reflect.TypeOf(AirConditionerEco{}), 2, 7)
  devices.RegisterProperty(reflect.TypeOf(AirConditionerHeater{}), 2, 9)
  devices.RegisterProperty(reflect.TypeOf(AirConditionerDryer{}), 2, 10)
  devices.RegisterProperty(reflect.TypeOf(AirConditionerSleepMode{}), 2, 11)
  devices.RegisterProperty(reflect.TypeOf(AirConditionerTargetHumidity{}), 2, 14)
  devices.RegisterProperty(reflect.TypeOf(AirConditionerUnStraightBlowing{}), 2, 15)
  devices.RegisterProperty(reflect.TypeOf(FanControlFanLevel{}), 3, 2)
  devices.RegisterProperty(reflect.TypeOf(FanControlVerticalSwing{}), 3, 4)
  devices.RegisterProperty(reflect.TypeOf(EnvironmentTemperature{}), 4, 7)
  devices.RegisterProperty(reflect.TypeOf(EnvironmentRelativeHumidity{}), 4, 9)
  devices.RegisterProperty(reflect.TypeOf(AlarmAlarm{}), 5, 1)
  devices.RegisterProperty(reflect.TypeOf(IndicatorLightOn{}), 6, 1)
  devices.RegisterProperty(reflect.TypeOf(IndicatorLightBrightness{}), 6, 2)
  devices.RegisterProperty(reflect.TypeOf(Electricity{}), 8, 1)
  devices.RegisterProperty(reflect.TypeOf(ElecCount{}), 8, 3)
  devices.RegisterProperty(reflect.TypeOf(Clean{}), 9, 1)
  devices.RegisterProperty(reflect.TypeOf(Examine{}), 9, 2)
  devices.RegisterProperty(reflect.TypeOf(Error{}), 9, 3)
  devices.RegisterProperty(reflect.TypeOf(RunningDuration{}), 9, 5)
  devices.RegisterProperty(reflect.TypeOf(FanPercent{}), 10, 1)
  devices.RegisterProperty(reflect.TypeOf(Timer{}), 10, 3)
  devices.RegisterProperty(reflect.TypeOf(HumidityRange{}), 10, 6)
  devices.RegisterProperty(reflect.TypeOf(TemperatureHumidity{}), 10, 9)
  devices.RegisterProperty(reflect.TypeOf(IotLinkageIotTemp{}), 11, 1)
  devices.RegisterProperty(reflect.TypeOf(IotLinkageTempCtrl{}), 11, 5)
  devices.RegisterProperty(reflect.TypeOf(IotLinkageSmartSleepCtrl{}), 11, 6)
  devices.RegisterProperty(reflect.TypeOf(MachineStateIndoorPipeTemp{}), 12, 1)
  devices.RegisterProperty(reflect.TypeOf(MachineStateIndoorFanSpeed{}), 12, 3)
  devices.RegisterProperty(reflect.TypeOf(MachineStateRealHeaterSwitch{}), 12, 4)
  devices.RegisterProperty(reflect.TypeOf(MachineStateRealindoorFanLever{}), 12, 5)
  devices.RegisterProperty(reflect.TypeOf(MachineStateOutdoorTemp{}), 12, 7)
  devices.RegisterProperty(reflect.TypeOf(MachineStateOutdoorPipeTemp{}), 12, 8)
  devices.RegisterProperty(reflect.TypeOf(MachineStateOutdoorExhaustTemp{}), 12, 9)
  devices.RegisterProperty(reflect.TypeOf(MachineStateOutdoorFanLever{}), 12, 10)
  devices.RegisterProperty(reflect.TypeOf(MachineStateCompressorFrequency{}), 12, 11)
  devices.RegisterProperty(reflect.TypeOf(MachineStateFourwayValveSwitch{}), 12, 12)
  devices.RegisterProperty(reflect.TypeOf(MachineStateOutdoorCurrent{}), 12, 13)
  devices.RegisterProperty(reflect.TypeOf(MachineStateOutdoorMVoltage{}), 12, 14)
  devices.RegisterProperty(reflect.TypeOf(MachineStateExpansionValve{}), 12, 15)
  devices.RegisterProperty(reflect.TypeOf(MachineStateLongitudeLatitude{}), 12, 16)
  devices.RegisterProperty(reflect.TypeOf(MachineStateAmbientLight{}), 12, 19)
  devices.RegisterProperty(reflect.TypeOf(FaultValueFaultValue{}), 13, 1)
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
// DeviceInformationSerialNo Serial Number
type DeviceInformationSerialNo struct {
  V SerialNoV
  Err error // Output only
}
// AirConditionerOn Switch Status
type AirConditionerOn struct {
  V OnV
  Err error // Output only
}
// AirConditionerMode Mode
type AirConditionerMode struct {
  V ModeV
  Err error // Output only
}
// AirConditionerFault Device Fault
type AirConditionerFault struct {
  V FaultV
  Err error // Output only
}
// AirConditionerTargetTemperature Target Temperature
type AirConditionerTargetTemperature struct {
  V TargetTemperatureV
  Err error // Output only
}
// AirConditionerEco ECO
type AirConditionerEco struct {
  V EcoV
  Err error // Output only
}
// AirConditionerHeater Heater
type AirConditionerHeater struct {
  V HeaterV
  Err error // Output only
}
// AirConditionerDryer Dryer
type AirConditionerDryer struct {
  V DryerV
  Err error // Output only
}
// AirConditionerSleepMode Sleep Mode
type AirConditionerSleepMode struct {
  V SleepModeV
  Err error // Output only
}
// AirConditionerTargetHumidity Target Humidity
type AirConditionerTargetHumidity struct {
  V TargetHumidityV
  Err error // Output only
}
// AirConditionerUnStraightBlowing Unstraight Blowing
type AirConditionerUnStraightBlowing struct {
  V UnStraightBlowingV
  Err error // Output only
}
// FanControlFanLevel Fan Level
type FanControlFanLevel struct {
  V FanLevelV
  Err error // Output only
}
// FanControlVerticalSwing Vertical Swing
type FanControlVerticalSwing struct {
  V VerticalSwingV
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
// AlarmAlarm Alarm
type AlarmAlarm struct {
  V AlarmV
  Err error // Output only
}
// IndicatorLightOn Switch Status
type IndicatorLightOn struct {
  V OnV
  Err error // Output only
}
// IndicatorLightBrightness Brightness
type IndicatorLightBrightness struct {
  V BrightnessV
  Err error // Output only
}
// Electricity 耗电量
type Electricity struct {
  V ElectricityV
  Err error // Output only
}
// ElecCount 耗电量本次计数
type ElecCount struct {
  V ElecCountV
  Err error // Output only
}
// Clean 自清洁开关
type Clean struct {
  V CleanV
  Err error // Output only
}
// Examine 机组体检开关
type Examine struct {
  V ExamineV
  Err error // Output only
}
// Error 故障
type Error struct {
  V ErrorV
  Err error // Output only
}
// RunningDuration 累计运行时长
type RunningDuration struct {
  V RunningDurationV
  Err error // Output only
}
// FanPercent 风速百分比
type FanPercent struct {
  V FanPercentV
  Err error // Output only
}
// Timer timer
type Timer struct {
  V TimerV
  Err error // Output only
}
// HumidityRange humidity-range
type HumidityRange struct {
  V HumidityRangeV
  Err error // Output only
}
// TemperatureHumidity temperature-humidity
type TemperatureHumidity struct {
  V TemperatureHumidityV
  Err error // Output only
}
// IotLinkageIotTemp iot-temp
type IotLinkageIotTemp struct {
  V IotTempV
  Err error // Output only
}
// IotLinkageTempCtrl temp-ctrl
type IotLinkageTempCtrl struct {
  V TempCtrlV
  Err error // Output only
}
// IotLinkageSmartSleepCtrl smart-sleep-ctrl
type IotLinkageSmartSleepCtrl struct {
  V SmartSleepCtrlV
  Err error // Output only
}
// MachineStateIndoorPipeTemp indoor-pipe-temp
type MachineStateIndoorPipeTemp struct {
  V IndoorPipeTempV
  Err error // Output only
}
// MachineStateIndoorFanSpeed indoor-fan-speed
type MachineStateIndoorFanSpeed struct {
  V IndoorFanSpeedV
  Err error // Output only
}
// MachineStateRealHeaterSwitch real-heater-switch
type MachineStateRealHeaterSwitch struct {
  V RealHeaterSwitchV
  Err error // Output only
}
// MachineStateRealindoorFanLever realindoor-fan-lever
type MachineStateRealindoorFanLever struct {
  V RealindoorFanLeverV
  Err error // Output only
}
// MachineStateOutdoorTemp outdoor-temp
type MachineStateOutdoorTemp struct {
  V OutdoorTempV
  Err error // Output only
}
// MachineStateOutdoorPipeTemp outdoor-pipe-temp
type MachineStateOutdoorPipeTemp struct {
  V OutdoorPipeTempV
  Err error // Output only
}
// MachineStateOutdoorExhaustTemp outdoor-exhaust-temp
type MachineStateOutdoorExhaustTemp struct {
  V OutdoorExhaustTempV
  Err error // Output only
}
// MachineStateOutdoorFanLever outdoor-fan-lever
type MachineStateOutdoorFanLever struct {
  V OutdoorFanLeverV
  Err error // Output only
}
// MachineStateCompressorFrequency compressor-frequency
type MachineStateCompressorFrequency struct {
  V CompressorFrequencyV
  Err error // Output only
}
// MachineStateFourwayValveSwitch fourway-valve-switch
type MachineStateFourwayValveSwitch struct {
  V FourwayValveSwitchV
  Err error // Output only
}
// MachineStateOutdoorCurrent outdoor--current
type MachineStateOutdoorCurrent struct {
  V OutdoorCurrentV
  Err error // Output only
}
// MachineStateOutdoorMVoltage outdoor-m-voltage
type MachineStateOutdoorMVoltage struct {
  V OutdoorMVoltageV
  Err error // Output only
}
// MachineStateExpansionValve expansion-valve
type MachineStateExpansionValve struct {
  V ExpansionValveV
  Err error // Output only
}
// MachineStateLongitudeLatitude longitude-latitude
type MachineStateLongitudeLatitude struct {
  V LongitudeLatitudeV
  Err error // Output only
}
// MachineStateAmbientLight ambient-light
type MachineStateAmbientLight struct {
  V AmbientLightV
  Err error // Output only
}
// FaultValueFaultValue fault-value
type FaultValueFaultValue struct {
  V FaultValueV
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
// SerialNoV Serial Number
type SerialNoV string
// OnV Switch Status
type OnV bool
// ModeV Mode
type ModeV uint8
const (
  ModeCool ModeV = 2
  ModeDry ModeV = 3
  ModeFan ModeV = 4
  ModeHeat ModeV = 5
)
func (v ModeV) String() string {
  switch v {
  case ModeCool:
    return "Cool"
  case ModeDry:
    return "Dry"
  case ModeFan:
    return "Fan"
  case ModeHeat:
    return "Heat"
  default:
    return fmt.Sprintf("unknown Mode: %v", uint8(v))
  }
}
// FaultV Device Fault
type FaultV string
// TargetTemperatureV Target Temperature
type TargetTemperatureV float64
// EcoV ECO
type EcoV bool
// HeaterV Heater
type HeaterV bool
// DryerV Dryer
type DryerV bool
// SleepModeV Sleep Mode
type SleepModeV bool
// TargetHumidityV Target Humidity
type TargetHumidityV uint8
// UnStraightBlowingV Unstraight Blowing
type UnStraightBlowingV bool
// FanLevelV Fan Level
type FanLevelV uint8
const (
  FanLevelAuto FanLevelV = 0
  FanLevelLevel1 FanLevelV = 1
  FanLevelLevel2 FanLevelV = 2
  FanLevelLevel3 FanLevelV = 3
  FanLevelLevel4 FanLevelV = 4
  FanLevelLevel5 FanLevelV = 5
  FanLevelLevel6 FanLevelV = 6
  FanLevelLevel7 FanLevelV = 7
)
func (v FanLevelV) String() string {
  switch v {
  case FanLevelAuto:
    return "Auto"
  case FanLevelLevel1:
    return "Level1"
  case FanLevelLevel2:
    return "Level2"
  case FanLevelLevel3:
    return "Level3"
  case FanLevelLevel4:
    return "Level4"
  case FanLevelLevel5:
    return "Level5"
  case FanLevelLevel6:
    return "Level6"
  case FanLevelLevel7:
    return "Level7"
  default:
    return fmt.Sprintf("unknown FanLevel: %v", uint8(v))
  }
}
// VerticalSwingV Vertical Swing
type VerticalSwingV bool
// TemperatureV Temperature
type TemperatureV float64
// RelativeHumidityV Relative Humidity
type RelativeHumidityV uint8
// AlarmV Alarm
type AlarmV bool
// BrightnessV Brightness
type BrightnessV uint8
const (
  BrightnessAuto BrightnessV = 0
  BrightnessMedium BrightnessV = 1
  BrightnessHigh BrightnessV = 2
)
func (v BrightnessV) String() string {
  switch v {
  case BrightnessAuto:
    return "Auto"
  case BrightnessMedium:
    return "Medium"
  case BrightnessHigh:
    return "High"
  default:
    return fmt.Sprintf("unknown Brightness: %v", uint8(v))
  }
}
// ElectricityV 耗电量
type ElectricityV float64
// ElecCountV 耗电量本次计数
type ElecCountV uint16
// CleanV 自清洁开关
type CleanV string
// ExamineV 机组体检开关
type ExamineV string
// ErrorV 故障
type ErrorV string
// RunningDurationV 累计运行时长
type RunningDurationV float64
// FanPercentV 风速百分比
type FanPercentV uint8
// TimerV timer
type TimerV string
// HumidityRangeV humidity-range
type HumidityRangeV string
// TemperatureHumidityV temperature-humidity
type TemperatureHumidityV float64
// IotTempV iot-temp
type IotTempV float64
// TempCtrlV temp-ctrl
type TempCtrlV bool
// SmartSleepCtrlV smart-sleep-ctrl
type SmartSleepCtrlV bool
// IndoorPipeTempV indoor-pipe-temp
type IndoorPipeTempV float64
// IndoorFanSpeedV indoor-fan-speed
type IndoorFanSpeedV uint16
// RealHeaterSwitchV real-heater-switch
type RealHeaterSwitchV bool
// RealindoorFanLeverV realindoor-fan-lever
type RealindoorFanLeverV uint8
const (
  RealindoorFanLeverLevel0 RealindoorFanLeverV = 0
  RealindoorFanLeverLevel1 RealindoorFanLeverV = 1
  RealindoorFanLeverLevel2 RealindoorFanLeverV = 2
  RealindoorFanLeverLevel3 RealindoorFanLeverV = 3
  RealindoorFanLeverLevel4 RealindoorFanLeverV = 4
  RealindoorFanLeverLevel5 RealindoorFanLeverV = 5
  RealindoorFanLeverLevel6 RealindoorFanLeverV = 6
  RealindoorFanLeverLevel7 RealindoorFanLeverV = 7
  RealindoorFanLeverLevel8 RealindoorFanLeverV = 8
  RealindoorFanLeverLevel9 RealindoorFanLeverV = 9
)
func (v RealindoorFanLeverV) String() string {
  switch v {
  case RealindoorFanLeverLevel0:
    return "Level 0"
  case RealindoorFanLeverLevel1:
    return "Level 1"
  case RealindoorFanLeverLevel2:
    return "Level 2"
  case RealindoorFanLeverLevel3:
    return "Level 3"
  case RealindoorFanLeverLevel4:
    return "Level 4"
  case RealindoorFanLeverLevel5:
    return "Level 5"
  case RealindoorFanLeverLevel6:
    return "Level 6"
  case RealindoorFanLeverLevel7:
    return "Level 7"
  case RealindoorFanLeverLevel8:
    return "Level 8"
  case RealindoorFanLeverLevel9:
    return "Level 9"
  default:
    return fmt.Sprintf("unknown RealindoorFanLever: %v", uint8(v))
  }
}
// OutdoorTempV outdoor-temp
type OutdoorTempV float64
// OutdoorPipeTempV outdoor-pipe-temp
type OutdoorPipeTempV float64
// OutdoorExhaustTempV outdoor-exhaust-temp
type OutdoorExhaustTempV float64
// OutdoorFanLeverV outdoor-fan-lever
type OutdoorFanLeverV uint8
const (
  OutdoorFanLeverLevel0 OutdoorFanLeverV = 0
  OutdoorFanLeverLevel1 OutdoorFanLeverV = 1
  OutdoorFanLeverLevel2 OutdoorFanLeverV = 2
  OutdoorFanLeverLevel3 OutdoorFanLeverV = 3
  OutdoorFanLeverLevel4 OutdoorFanLeverV = 4
  OutdoorFanLeverLevel5 OutdoorFanLeverV = 5
  OutdoorFanLeverLevel6 OutdoorFanLeverV = 6
  OutdoorFanLeverLevel7 OutdoorFanLeverV = 7
  OutdoorFanLeverLevel8 OutdoorFanLeverV = 8
  OutdoorFanLeverLevel9 OutdoorFanLeverV = 9
)
func (v OutdoorFanLeverV) String() string {
  switch v {
  case OutdoorFanLeverLevel0:
    return "Level 0"
  case OutdoorFanLeverLevel1:
    return "Level 1"
  case OutdoorFanLeverLevel2:
    return "Level 2"
  case OutdoorFanLeverLevel3:
    return "Level 3"
  case OutdoorFanLeverLevel4:
    return "Level 4"
  case OutdoorFanLeverLevel5:
    return "Level 5"
  case OutdoorFanLeverLevel6:
    return "Level 6"
  case OutdoorFanLeverLevel7:
    return "Level 7"
  case OutdoorFanLeverLevel8:
    return "Level 8"
  case OutdoorFanLeverLevel9:
    return "Level 9"
  default:
    return fmt.Sprintf("unknown OutdoorFanLever: %v", uint8(v))
  }
}
// CompressorFrequencyV compressor-frequency
type CompressorFrequencyV float64
// FourwayValveSwitchV fourway-valve-switch
type FourwayValveSwitchV bool
// OutdoorCurrentV outdoor--current
type OutdoorCurrentV float64
// OutdoorMVoltageV outdoor-m-voltage
type OutdoorMVoltageV float64
// ExpansionValveV expansion-valve
type ExpansionValveV uint16
// LongitudeLatitudeV longitude-latitude
type LongitudeLatitudeV string
// AmbientLightV ambient-light
type AmbientLightV uint32
// FaultValueV fault-value
type FaultValueV int16
const (
  FaultValueNoFailure FaultValueV = 0
  FaultValueF24 FaultValueV = 4
  FaultValueF32 FaultValueV = 5
  FaultValueP1 FaultValueV = 6
  FaultValueP21 FaultValueV = 7
  FaultValueP22 FaultValueV = 8
  FaultValueP23 FaultValueV = 9
  FaultValueP24 FaultValueV = 10
  FaultValueP25 FaultValueV = 11
  FaultValueP4 FaultValueV = 12
  FaultValueP5 FaultValueV = 13
  FaultValueP6 FaultValueV = 14
  FaultValueP8 FaultValueV = 15
  FaultValuePA FaultValueV = 16
  FaultValuePC FaultValueV = 17
  FaultValueU2 FaultValueV = 18
  FaultValueU3 FaultValueV = 19
  FaultValueU4 FaultValueV = 20
  FaultValueU5 FaultValueV = 21
  FaultValueU61 FaultValueV = 22
  FaultValueU62 FaultValueV = 23
  FaultValueU8 FaultValueV = 24
  FaultValueU0 FaultValueV = 25
  FaultValueC1 FaultValueV = 26
  FaultValueC2 FaultValueV = 27
  FaultValueC3 FaultValueV = 28
  FaultValueC4 FaultValueV = 29
  FaultValueC5 FaultValueV = 30
  FaultValueF24CLEAR FaultValueV = -4
  FaultValueF32CLEAR FaultValueV = -5
  FaultValueP1CLEAR FaultValueV = -6
  FaultValueP21CLEAR FaultValueV = -7
  FaultValueP22CLEAR FaultValueV = -8
  FaultValueP23CLEAR FaultValueV = -9
  FaultValueP24CLEAR FaultValueV = -10
  FaultValueP25CLEAR FaultValueV = -11
  FaultValueP4CLEAR FaultValueV = -12
  FaultValueP5CLEAR FaultValueV = -13
  FaultValueP6CLEAR FaultValueV = -14
  FaultValueP8CLEAR FaultValueV = -15
  FaultValuePACLEAR FaultValueV = -16
  FaultValuePCCLEAR FaultValueV = -17
  FaultValueU2CLEAR FaultValueV = -18
  FaultValueU3CLEAR FaultValueV = -19
  FaultValueU4CLEAR FaultValueV = -20
  FaultValueU5CLEAR FaultValueV = -21
  FaultValueU61CLEAR FaultValueV = -22
  FaultValueU62CLEAR FaultValueV = -23
  FaultValueU8CLEAR FaultValueV = -24
  FaultValueU0CLEAR FaultValueV = -25
  FaultValueC1CLEAR FaultValueV = -26
  FaultValueC2CLEAR FaultValueV = -27
  FaultValueC3CLEAR FaultValueV = -28
  FaultValueC4CLEAR FaultValueV = -29
  FaultValueC5CLEAR FaultValueV = -30
)
func (v FaultValueV) String() string {
  switch v {
  case FaultValueNoFailure:
    return "No Failure"
  case FaultValueF24:
    return "F2.4"
  case FaultValueF32:
    return "F3.2"
  case FaultValueP1:
    return "P1"
  case FaultValueP21:
    return "P2.1"
  case FaultValueP22:
    return "P2.2"
  case FaultValueP23:
    return "P2.3"
  case FaultValueP24:
    return "P2.4"
  case FaultValueP25:
    return "P2.5"
  case FaultValueP4:
    return "P4"
  case FaultValueP5:
    return "P5"
  case FaultValueP6:
    return "P6"
  case FaultValueP8:
    return "P8"
  case FaultValuePA:
    return "PA"
  case FaultValuePC:
    return "PC"
  case FaultValueU2:
    return "U2"
  case FaultValueU3:
    return "U3"
  case FaultValueU4:
    return "U4"
  case FaultValueU5:
    return "U5"
  case FaultValueU61:
    return "U6.1"
  case FaultValueU62:
    return "U6.2"
  case FaultValueU8:
    return "U8"
  case FaultValueU0:
    return "U0"
  case FaultValueC1:
    return "C1"
  case FaultValueC2:
    return "C2"
  case FaultValueC3:
    return "C3"
  case FaultValueC4:
    return "C4"
  case FaultValueC5:
    return "C5"
  case FaultValueF24CLEAR:
    return "F2.4  CLEAR"
  case FaultValueF32CLEAR:
    return "F3.2 CLEAR"
  case FaultValueP1CLEAR:
    return "P1 CLEAR"
  case FaultValueP21CLEAR:
    return "P2.1 CLEAR"
  case FaultValueP22CLEAR:
    return "P2.2  CLEAR"
  case FaultValueP23CLEAR:
    return "P2.3 CLEAR"
  case FaultValueP24CLEAR:
    return "P2.4 CLEAR"
  case FaultValueP25CLEAR:
    return "P2.5 CLEAR"
  case FaultValueP4CLEAR:
    return "P4 CLEAR"
  case FaultValueP5CLEAR:
    return "P5 CLEAR"
  case FaultValueP6CLEAR:
    return "P6 CLEAR"
  case FaultValueP8CLEAR:
    return "P8 CLEAR"
  case FaultValuePACLEAR:
    return "PA CLEAR"
  case FaultValuePCCLEAR:
    return "PC CLEAR"
  case FaultValueU2CLEAR:
    return "U2 CLEAR"
  case FaultValueU3CLEAR:
    return "U3 CLEAR"
  case FaultValueU4CLEAR:
    return "U4  CLEAR"
  case FaultValueU5CLEAR:
    return "U5 CLEAR"
  case FaultValueU61CLEAR:
    return "U6.1 CLEAR"
  case FaultValueU62CLEAR:
    return "U6.2  CLEAR"
  case FaultValueU8CLEAR:
    return "U8 CLEAR"
  case FaultValueU0CLEAR:
    return "U0 CLEAR"
  case FaultValueC1CLEAR:
    return "C1 CLEAR"
  case FaultValueC2CLEAR:
    return "C2 CLEAR"
  case FaultValueC3CLEAR:
    return "C3 CLEAR"
  case FaultValueC4CLEAR:
    return "C4 CLEAR"
  case FaultValueC5CLEAR:
    return "C5 CLEAR"
  default:
    return fmt.Sprintf("unknown FaultValue: %v", int16(v))
  }
}