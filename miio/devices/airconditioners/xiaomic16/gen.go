package xiaomic16

import (
	"fmt"
	"reflect"

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
	V   MiotV2Manufacturer
	Err error // Output only
}

// DeviceInformationModel Device Model
type DeviceInformationModel struct {
	V   MiotV2Model
	Err error // Output only
}

// DeviceInformationSerialNumber Device Serial Number
type DeviceInformationSerialNumber struct {
	V   MiotV2SerialNumber
	Err error // Output only
}

// DeviceInformationFirmwareRevision Current Firmware Version
type DeviceInformationFirmwareRevision struct {
	V   MiotV2FirmwareRevision
	Err error // Output only
}

// DeviceInformationSerialNo Serial Number
type DeviceInformationSerialNo struct {
	V   MiotV2SerialNo
	Err error // Output only
}

// AirConditionerOn Switch Status
type AirConditionerOn struct {
	V   MiotV2On
	Err error // Output only
}

// AirConditionerMode Mode
type AirConditionerMode struct {
	V   MiotV2Mode
	Err error // Output only
}

// AirConditionerFault Device Fault
type AirConditionerFault struct {
	V   MiotV2Fault
	Err error // Output only
}

// AirConditionerTargetTemperature Target Temperature
type AirConditionerTargetTemperature struct {
	V   MiotV2TargetTemperature
	Err error // Output only
}

// AirConditionerEco ECO
type AirConditionerEco struct {
	V   MiotV2Eco
	Err error // Output only
}

// AirConditionerHeater Heater
type AirConditionerHeater struct {
	V   MiotV2Heater
	Err error // Output only
}

// AirConditionerDryer Dryer
type AirConditionerDryer struct {
	V   MiotV2Dryer
	Err error // Output only
}

// AirConditionerSleepMode Sleep Mode
type AirConditionerSleepMode struct {
	V   MiotV2SleepMode
	Err error // Output only
}

// AirConditionerTargetHumidity Target Humidity
type AirConditionerTargetHumidity struct {
	V   MiotV2TargetHumidity
	Err error // Output only
}

// AirConditionerUnStraightBlowing Unstraight Blowing
type AirConditionerUnStraightBlowing struct {
	V   MiotV2UnStraightBlowing
	Err error // Output only
}

// FanControlFanLevel Fan Level
type FanControlFanLevel struct {
	V   MiotV2FanLevel
	Err error // Output only
}

// FanControlVerticalSwing Vertical Swing
type FanControlVerticalSwing struct {
	V   MiotV2VerticalSwing
	Err error // Output only
}

// EnvironmentTemperature Temperature
type EnvironmentTemperature struct {
	V   MiotV2Temperature
	Err error // Output only
}

// EnvironmentRelativeHumidity Relative Humidity
type EnvironmentRelativeHumidity struct {
	V   MiotV2RelativeHumidity
	Err error // Output only
}

// AlarmAlarm Alarm
type AlarmAlarm struct {
	V   MiotV2Alarm
	Err error // Output only
}

// IndicatorLightOn Switch Status
type IndicatorLightOn struct {
	V   MiotV2On
	Err error // Output only
}

// IndicatorLightBrightness Brightness
type IndicatorLightBrightness struct {
	V   MiotV2Brightness
	Err error // Output only
}

// Electricity 耗电量
type Electricity struct {
	V   XiaomiElectricity
	Err error // Output only
}

// ElecCount 耗电量本次计数
type ElecCount struct {
	V   XiaomiElecCount
	Err error // Output only
}

// Clean 自清洁开关
type Clean struct {
	V   XiaomiClean
	Err error // Output only
}

// Examine 机组体检开关
type Examine struct {
	V   XiaomiExamine
	Err error // Output only
}

// Error 故障
type Error struct {
	V   XiaomiError
	Err error // Output only
}

// RunningDuration 累计运行时长
type RunningDuration struct {
	V   XiaomiRunningDuration
	Err error // Output only
}

// FanPercent 风速百分比
type FanPercent struct {
	V   XiaomiFanPercent
	Err error // Output only
}

// Timer timer
type Timer struct {
	V   XiaomiTimer
	Err error // Output only
}

// HumidityRange humidity-range
type HumidityRange struct {
	V   XiaomiHumidityRange
	Err error // Output only
}

// TemperatureHumidity temperature-humidity
type TemperatureHumidity struct {
	V   XiaomiTemperatureHumidity
	Err error // Output only
}

// IotLinkageIotTemp iot-temp
type IotLinkageIotTemp struct {
	V   XiaomiIotTemp
	Err error // Output only
}

// IotLinkageTempCtrl temp-ctrl
type IotLinkageTempCtrl struct {
	V   XiaomiTempCtrl
	Err error // Output only
}

// IotLinkageSmartSleepCtrl smart-sleep-ctrl
type IotLinkageSmartSleepCtrl struct {
	V   XiaomiSmartSleepCtrl
	Err error // Output only
}

// MachineStateIndoorPipeTemp indoor-pipe-temp
type MachineStateIndoorPipeTemp struct {
	V   XiaomiIndoorPipeTemp
	Err error // Output only
}

// MachineStateIndoorFanSpeed indoor-fan-speed
type MachineStateIndoorFanSpeed struct {
	V   XiaomiIndoorFanSpeed
	Err error // Output only
}

// MachineStateRealHeaterSwitch real-heater-switch
type MachineStateRealHeaterSwitch struct {
	V   XiaomiRealHeaterSwitch
	Err error // Output only
}

// MachineStateRealindoorFanLever realindoor-fan-lever
type MachineStateRealindoorFanLever struct {
	V   XiaomiRealindoorFanLever
	Err error // Output only
}

// MachineStateOutdoorTemp outdoor-temp
type MachineStateOutdoorTemp struct {
	V   XiaomiOutdoorTemp
	Err error // Output only
}

// MachineStateOutdoorPipeTemp outdoor-pipe-temp
type MachineStateOutdoorPipeTemp struct {
	V   XiaomiOutdoorPipeTemp
	Err error // Output only
}

// MachineStateOutdoorExhaustTemp outdoor-exhaust-temp
type MachineStateOutdoorExhaustTemp struct {
	V   XiaomiOutdoorExhaustTemp
	Err error // Output only
}

// MachineStateOutdoorFanLever outdoor-fan-lever
type MachineStateOutdoorFanLever struct {
	V   XiaomiOutdoorFanLever
	Err error // Output only
}

// MachineStateCompressorFrequency compressor-frequency
type MachineStateCompressorFrequency struct {
	V   XiaomiCompressorFrequency
	Err error // Output only
}

// MachineStateFourwayValveSwitch fourway-valve-switch
type MachineStateFourwayValveSwitch struct {
	V   XiaomiFourwayValveSwitch
	Err error // Output only
}

// MachineStateOutdoorCurrent outdoor--current
type MachineStateOutdoorCurrent struct {
	V   XiaomiOutdoorCurrent
	Err error // Output only
}

// MachineStateOutdoorMVoltage outdoor-m-voltage
type MachineStateOutdoorMVoltage struct {
	V   XiaomiOutdoorMVoltage
	Err error // Output only
}

// MachineStateExpansionValve expansion-valve
type MachineStateExpansionValve struct {
	V   XiaomiExpansionValve
	Err error // Output only
}

// MachineStateLongitudeLatitude longitude-latitude
type MachineStateLongitudeLatitude struct {
	V   XiaomiLongitudeLatitude
	Err error // Output only
}

// MachineStateAmbientLight ambient-light
type MachineStateAmbientLight struct {
	V   XiaomiAmbientLight
	Err error // Output only
}

// FaultValueFaultValue fault-value
type FaultValueFaultValue struct {
	V   XiaomiFaultValue
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

// MiotV2SerialNo Serial Number
type MiotV2SerialNo string

// MiotV2On Switch Status
type MiotV2On bool

// MiotV2Mode Mode
type MiotV2Mode uint8

const (
	MiotV2ModeCool MiotV2Mode = 2
	MiotV2ModeDry  MiotV2Mode = 3
	MiotV2ModeFan  MiotV2Mode = 4
	MiotV2ModeHeat MiotV2Mode = 5
)

func (v MiotV2Mode) String() string {
	switch v {
	case MiotV2ModeCool:
		return "Cool"
	case MiotV2ModeDry:
		return "Dry"
	case MiotV2ModeFan:
		return "Fan"
	case MiotV2ModeHeat:
		return "Heat"
	default:
		return fmt.Sprintf("unknown MiotV2Mode: %v", uint8(v))
	}
}

// MiotV2Fault Device Fault
type MiotV2Fault string

// MiotV2TargetTemperature Target Temperature
type MiotV2TargetTemperature float64

// MiotV2Eco ECO
type MiotV2Eco bool

// MiotV2Heater Heater
type MiotV2Heater bool

// MiotV2Dryer Dryer
type MiotV2Dryer bool

// MiotV2SleepMode Sleep Mode
type MiotV2SleepMode bool

// MiotV2TargetHumidity Target Humidity
type MiotV2TargetHumidity uint8

// MiotV2UnStraightBlowing Unstraight Blowing
type MiotV2UnStraightBlowing bool

// MiotV2FanLevel Fan Level
type MiotV2FanLevel uint8

const (
	MiotV2FanLevelAuto   MiotV2FanLevel = 0
	MiotV2FanLevelLevel1 MiotV2FanLevel = 1
	MiotV2FanLevelLevel2 MiotV2FanLevel = 2
	MiotV2FanLevelLevel3 MiotV2FanLevel = 3
	MiotV2FanLevelLevel4 MiotV2FanLevel = 4
	MiotV2FanLevelLevel5 MiotV2FanLevel = 5
	MiotV2FanLevelLevel6 MiotV2FanLevel = 6
	MiotV2FanLevelLevel7 MiotV2FanLevel = 7
)

func (v MiotV2FanLevel) String() string {
	switch v {
	case MiotV2FanLevelAuto:
		return "Auto"
	case MiotV2FanLevelLevel1:
		return "Level1"
	case MiotV2FanLevelLevel2:
		return "Level2"
	case MiotV2FanLevelLevel3:
		return "Level3"
	case MiotV2FanLevelLevel4:
		return "Level4"
	case MiotV2FanLevelLevel5:
		return "Level5"
	case MiotV2FanLevelLevel6:
		return "Level6"
	case MiotV2FanLevelLevel7:
		return "Level7"
	default:
		return fmt.Sprintf("unknown MiotV2FanLevel: %v", uint8(v))
	}
}

// MiotV2VerticalSwing Vertical Swing
type MiotV2VerticalSwing bool

// MiotV2Temperature Temperature
type MiotV2Temperature float64

// MiotV2RelativeHumidity Relative Humidity
type MiotV2RelativeHumidity uint8

// MiotV2Alarm Alarm
type MiotV2Alarm bool

// MiotV2Brightness Brightness
type MiotV2Brightness uint8

const (
	MiotV2BrightnessAuto   MiotV2Brightness = 0
	MiotV2BrightnessMedium MiotV2Brightness = 1
	MiotV2BrightnessHigh   MiotV2Brightness = 2
)

func (v MiotV2Brightness) String() string {
	switch v {
	case MiotV2BrightnessAuto:
		return "Auto"
	case MiotV2BrightnessMedium:
		return "Medium"
	case MiotV2BrightnessHigh:
		return "High"
	default:
		return fmt.Sprintf("unknown MiotV2Brightness: %v", uint8(v))
	}
}

// XiaomiElectricity 耗电量
type XiaomiElectricity float64

// XiaomiElecCount 耗电量本次计数
type XiaomiElecCount uint16

// XiaomiClean 自清洁开关
type XiaomiClean string

// XiaomiExamine 机组体检开关
type XiaomiExamine string

// XiaomiError 故障
type XiaomiError string

// XiaomiRunningDuration 累计运行时长
type XiaomiRunningDuration float64

// XiaomiFanPercent 风速百分比
type XiaomiFanPercent uint8

// XiaomiTimer timer
type XiaomiTimer string

// XiaomiHumidityRange humidity-range
type XiaomiHumidityRange string

// XiaomiTemperatureHumidity temperature-humidity
type XiaomiTemperatureHumidity float64

// XiaomiIotTemp iot-temp
type XiaomiIotTemp float64

// XiaomiTempCtrl temp-ctrl
type XiaomiTempCtrl bool

// XiaomiSmartSleepCtrl smart-sleep-ctrl
type XiaomiSmartSleepCtrl bool

// XiaomiIndoorPipeTemp indoor-pipe-temp
type XiaomiIndoorPipeTemp float64

// XiaomiIndoorFanSpeed indoor-fan-speed
type XiaomiIndoorFanSpeed uint16

// XiaomiRealHeaterSwitch real-heater-switch
type XiaomiRealHeaterSwitch bool

// XiaomiRealindoorFanLever realindoor-fan-lever
type XiaomiRealindoorFanLever uint8

const (
	XiaomiRealindoorFanLeverLevel0 XiaomiRealindoorFanLever = 0
	XiaomiRealindoorFanLeverLevel1 XiaomiRealindoorFanLever = 1
	XiaomiRealindoorFanLeverLevel2 XiaomiRealindoorFanLever = 2
	XiaomiRealindoorFanLeverLevel3 XiaomiRealindoorFanLever = 3
	XiaomiRealindoorFanLeverLevel4 XiaomiRealindoorFanLever = 4
	XiaomiRealindoorFanLeverLevel5 XiaomiRealindoorFanLever = 5
	XiaomiRealindoorFanLeverLevel6 XiaomiRealindoorFanLever = 6
	XiaomiRealindoorFanLeverLevel7 XiaomiRealindoorFanLever = 7
	XiaomiRealindoorFanLeverLevel8 XiaomiRealindoorFanLever = 8
	XiaomiRealindoorFanLeverLevel9 XiaomiRealindoorFanLever = 9
)

func (v XiaomiRealindoorFanLever) String() string {
	switch v {
	case XiaomiRealindoorFanLeverLevel0:
		return "Level 0"
	case XiaomiRealindoorFanLeverLevel1:
		return "Level 1"
	case XiaomiRealindoorFanLeverLevel2:
		return "Level 2"
	case XiaomiRealindoorFanLeverLevel3:
		return "Level 3"
	case XiaomiRealindoorFanLeverLevel4:
		return "Level 4"
	case XiaomiRealindoorFanLeverLevel5:
		return "Level 5"
	case XiaomiRealindoorFanLeverLevel6:
		return "Level 6"
	case XiaomiRealindoorFanLeverLevel7:
		return "Level 7"
	case XiaomiRealindoorFanLeverLevel8:
		return "Level 8"
	case XiaomiRealindoorFanLeverLevel9:
		return "Level 9"
	default:
		return fmt.Sprintf("unknown XiaomiRealindoorFanLever: %v", uint8(v))
	}
}

// XiaomiOutdoorTemp outdoor-temp
type XiaomiOutdoorTemp float64

// XiaomiOutdoorPipeTemp outdoor-pipe-temp
type XiaomiOutdoorPipeTemp float64

// XiaomiOutdoorExhaustTemp outdoor-exhaust-temp
type XiaomiOutdoorExhaustTemp float64

// XiaomiOutdoorFanLever outdoor-fan-lever
type XiaomiOutdoorFanLever uint8

const (
	XiaomiOutdoorFanLeverLevel0 XiaomiOutdoorFanLever = 0
	XiaomiOutdoorFanLeverLevel1 XiaomiOutdoorFanLever = 1
	XiaomiOutdoorFanLeverLevel2 XiaomiOutdoorFanLever = 2
	XiaomiOutdoorFanLeverLevel3 XiaomiOutdoorFanLever = 3
	XiaomiOutdoorFanLeverLevel4 XiaomiOutdoorFanLever = 4
	XiaomiOutdoorFanLeverLevel5 XiaomiOutdoorFanLever = 5
	XiaomiOutdoorFanLeverLevel6 XiaomiOutdoorFanLever = 6
	XiaomiOutdoorFanLeverLevel7 XiaomiOutdoorFanLever = 7
	XiaomiOutdoorFanLeverLevel8 XiaomiOutdoorFanLever = 8
	XiaomiOutdoorFanLeverLevel9 XiaomiOutdoorFanLever = 9
)

func (v XiaomiOutdoorFanLever) String() string {
	switch v {
	case XiaomiOutdoorFanLeverLevel0:
		return "Level 0"
	case XiaomiOutdoorFanLeverLevel1:
		return "Level 1"
	case XiaomiOutdoorFanLeverLevel2:
		return "Level 2"
	case XiaomiOutdoorFanLeverLevel3:
		return "Level 3"
	case XiaomiOutdoorFanLeverLevel4:
		return "Level 4"
	case XiaomiOutdoorFanLeverLevel5:
		return "Level 5"
	case XiaomiOutdoorFanLeverLevel6:
		return "Level 6"
	case XiaomiOutdoorFanLeverLevel7:
		return "Level 7"
	case XiaomiOutdoorFanLeverLevel8:
		return "Level 8"
	case XiaomiOutdoorFanLeverLevel9:
		return "Level 9"
	default:
		return fmt.Sprintf("unknown XiaomiOutdoorFanLever: %v", uint8(v))
	}
}

// XiaomiCompressorFrequency compressor-frequency
type XiaomiCompressorFrequency float64

// XiaomiFourwayValveSwitch fourway-valve-switch
type XiaomiFourwayValveSwitch bool

// XiaomiOutdoorCurrent outdoor--current
type XiaomiOutdoorCurrent float64

// XiaomiOutdoorMVoltage outdoor-m-voltage
type XiaomiOutdoorMVoltage float64

// XiaomiExpansionValve expansion-valve
type XiaomiExpansionValve uint16

// XiaomiLongitudeLatitude longitude-latitude
type XiaomiLongitudeLatitude string

// XiaomiAmbientLight ambient-light
type XiaomiAmbientLight uint32

// XiaomiFaultValue fault-value
type XiaomiFaultValue int16

const (
	XiaomiFaultValueNoFailure XiaomiFaultValue = 0
	XiaomiFaultValueF24       XiaomiFaultValue = 4
	XiaomiFaultValueF32       XiaomiFaultValue = 5
	XiaomiFaultValueP1        XiaomiFaultValue = 6
	XiaomiFaultValueP21       XiaomiFaultValue = 7
	XiaomiFaultValueP22       XiaomiFaultValue = 8
	XiaomiFaultValueP23       XiaomiFaultValue = 9
	XiaomiFaultValueP24       XiaomiFaultValue = 10
	XiaomiFaultValueP25       XiaomiFaultValue = 11
	XiaomiFaultValueP4        XiaomiFaultValue = 12
	XiaomiFaultValueP5        XiaomiFaultValue = 13
	XiaomiFaultValueP6        XiaomiFaultValue = 14
	XiaomiFaultValueP8        XiaomiFaultValue = 15
	XiaomiFaultValuePA        XiaomiFaultValue = 16
	XiaomiFaultValuePC        XiaomiFaultValue = 17
	XiaomiFaultValueU2        XiaomiFaultValue = 18
	XiaomiFaultValueU3        XiaomiFaultValue = 19
	XiaomiFaultValueU4        XiaomiFaultValue = 20
	XiaomiFaultValueU5        XiaomiFaultValue = 21
	XiaomiFaultValueU61       XiaomiFaultValue = 22
	XiaomiFaultValueU62       XiaomiFaultValue = 23
	XiaomiFaultValueU8        XiaomiFaultValue = 24
	XiaomiFaultValueU0        XiaomiFaultValue = 25
	XiaomiFaultValueC1        XiaomiFaultValue = 26
	XiaomiFaultValueC2        XiaomiFaultValue = 27
	XiaomiFaultValueC3        XiaomiFaultValue = 28
	XiaomiFaultValueC4        XiaomiFaultValue = 29
	XiaomiFaultValueC5        XiaomiFaultValue = 30
	XiaomiFaultValueF24CLEAR  XiaomiFaultValue = -4
	XiaomiFaultValueF32CLEAR  XiaomiFaultValue = -5
	XiaomiFaultValueP1CLEAR   XiaomiFaultValue = -6
	XiaomiFaultValueP21CLEAR  XiaomiFaultValue = -7
	XiaomiFaultValueP22CLEAR  XiaomiFaultValue = -8
	XiaomiFaultValueP23CLEAR  XiaomiFaultValue = -9
	XiaomiFaultValueP24CLEAR  XiaomiFaultValue = -10
	XiaomiFaultValueP25CLEAR  XiaomiFaultValue = -11
	XiaomiFaultValueP4CLEAR   XiaomiFaultValue = -12
	XiaomiFaultValueP5CLEAR   XiaomiFaultValue = -13
	XiaomiFaultValueP6CLEAR   XiaomiFaultValue = -14
	XiaomiFaultValueP8CLEAR   XiaomiFaultValue = -15
	XiaomiFaultValuePACLEAR   XiaomiFaultValue = -16
	XiaomiFaultValuePCCLEAR   XiaomiFaultValue = -17
	XiaomiFaultValueU2CLEAR   XiaomiFaultValue = -18
	XiaomiFaultValueU3CLEAR   XiaomiFaultValue = -19
	XiaomiFaultValueU4CLEAR   XiaomiFaultValue = -20
	XiaomiFaultValueU5CLEAR   XiaomiFaultValue = -21
	XiaomiFaultValueU61CLEAR  XiaomiFaultValue = -22
	XiaomiFaultValueU62CLEAR  XiaomiFaultValue = -23
	XiaomiFaultValueU8CLEAR   XiaomiFaultValue = -24
	XiaomiFaultValueU0CLEAR   XiaomiFaultValue = -25
	XiaomiFaultValueC1CLEAR   XiaomiFaultValue = -26
	XiaomiFaultValueC2CLEAR   XiaomiFaultValue = -27
	XiaomiFaultValueC3CLEAR   XiaomiFaultValue = -28
	XiaomiFaultValueC4CLEAR   XiaomiFaultValue = -29
	XiaomiFaultValueC5CLEAR   XiaomiFaultValue = -30
)

func (v XiaomiFaultValue) String() string {
	switch v {
	case XiaomiFaultValueNoFailure:
		return "No Failure"
	case XiaomiFaultValueF24:
		return "F2.4"
	case XiaomiFaultValueF32:
		return "F3.2"
	case XiaomiFaultValueP1:
		return "P1"
	case XiaomiFaultValueP21:
		return "P2.1"
	case XiaomiFaultValueP22:
		return "P2.2"
	case XiaomiFaultValueP23:
		return "P2.3"
	case XiaomiFaultValueP24:
		return "P2.4"
	case XiaomiFaultValueP25:
		return "P2.5"
	case XiaomiFaultValueP4:
		return "P4"
	case XiaomiFaultValueP5:
		return "P5"
	case XiaomiFaultValueP6:
		return "P6"
	case XiaomiFaultValueP8:
		return "P8"
	case XiaomiFaultValuePA:
		return "PA"
	case XiaomiFaultValuePC:
		return "PC"
	case XiaomiFaultValueU2:
		return "U2"
	case XiaomiFaultValueU3:
		return "U3"
	case XiaomiFaultValueU4:
		return "U4"
	case XiaomiFaultValueU5:
		return "U5"
	case XiaomiFaultValueU61:
		return "U6.1"
	case XiaomiFaultValueU62:
		return "U6.2"
	case XiaomiFaultValueU8:
		return "U8"
	case XiaomiFaultValueU0:
		return "U0"
	case XiaomiFaultValueC1:
		return "C1"
	case XiaomiFaultValueC2:
		return "C2"
	case XiaomiFaultValueC3:
		return "C3"
	case XiaomiFaultValueC4:
		return "C4"
	case XiaomiFaultValueC5:
		return "C5"
	case XiaomiFaultValueF24CLEAR:
		return "F2.4  CLEAR"
	case XiaomiFaultValueF32CLEAR:
		return "F3.2 CLEAR"
	case XiaomiFaultValueP1CLEAR:
		return "P1 CLEAR"
	case XiaomiFaultValueP21CLEAR:
		return "P2.1 CLEAR"
	case XiaomiFaultValueP22CLEAR:
		return "P2.2  CLEAR"
	case XiaomiFaultValueP23CLEAR:
		return "P2.3 CLEAR"
	case XiaomiFaultValueP24CLEAR:
		return "P2.4 CLEAR"
	case XiaomiFaultValueP25CLEAR:
		return "P2.5 CLEAR"
	case XiaomiFaultValueP4CLEAR:
		return "P4 CLEAR"
	case XiaomiFaultValueP5CLEAR:
		return "P5 CLEAR"
	case XiaomiFaultValueP6CLEAR:
		return "P6 CLEAR"
	case XiaomiFaultValueP8CLEAR:
		return "P8 CLEAR"
	case XiaomiFaultValuePACLEAR:
		return "PA CLEAR"
	case XiaomiFaultValuePCCLEAR:
		return "PC CLEAR"
	case XiaomiFaultValueU2CLEAR:
		return "U2 CLEAR"
	case XiaomiFaultValueU3CLEAR:
		return "U3 CLEAR"
	case XiaomiFaultValueU4CLEAR:
		return "U4  CLEAR"
	case XiaomiFaultValueU5CLEAR:
		return "U5 CLEAR"
	case XiaomiFaultValueU61CLEAR:
		return "U6.1 CLEAR"
	case XiaomiFaultValueU62CLEAR:
		return "U6.2  CLEAR"
	case XiaomiFaultValueU8CLEAR:
		return "U8 CLEAR"
	case XiaomiFaultValueU0CLEAR:
		return "U0 CLEAR"
	case XiaomiFaultValueC1CLEAR:
		return "C1 CLEAR"
	case XiaomiFaultValueC2CLEAR:
		return "C2 CLEAR"
	case XiaomiFaultValueC3CLEAR:
		return "C3 CLEAR"
	case XiaomiFaultValueC4CLEAR:
		return "C4 CLEAR"
	case XiaomiFaultValueC5CLEAR:
		return "C5 CLEAR"
	default:
		return fmt.Sprintf("unknown XiaomiFaultValue: %v", int16(v))
	}
}
