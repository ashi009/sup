package devices

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
)

type PropertyStatusCode int

const (
	PropertyStatusCodeOK            PropertyStatusCode = 0
	PropertyStatusCodePending       PropertyStatusCode = 1
	PropertyStatusCodeUnreadable    PropertyStatusCode = -4001
	PropertyStatusCodeUnwritable    PropertyStatusCode = -4002
	PropertyStatusCodeNotFound      PropertyStatusCode = -4003
	PropertyStatusCodeInternal      PropertyStatusCode = -4004
	PropertyStatusCodeInvalidValue  PropertyStatusCode = -4005
	PropertyStatusCodeInvalidParams PropertyStatusCode = -4006
	PropertyStatusCodeInvalidDID    PropertyStatusCode = -4007
)

func (e PropertyStatusCode) String() string {
	switch e {
	case PropertyStatusCodeOK:
		return "成功"
	case PropertyStatusCodePending:
		return "接收到请求，但操作还没有完成"
	case PropertyStatusCodeUnreadable:
		return "属性不可读"
	case PropertyStatusCodeUnwritable:
		return "属性不可写"
	case PropertyStatusCodeNotFound:
		return "属性、方法、事件不存在"
	case PropertyStatusCodeInternal:
		return "其他内部错误"
	case PropertyStatusCodeInvalidValue:
		return "属性value错误"
	case PropertyStatusCodeInvalidParams:
		return "方法in参数错误"
	case PropertyStatusCodeInvalidDID:
		return "did错误"
	default:
		return fmt.Sprintf("未知错误: %d", e)
	}
}

type PropertyError struct {
	StatusCode PropertyStatusCode
}

func (e PropertyError) Error() string {
	return e.StatusCode.String()
}

type PropertyBinding struct {
	ServiceID  int `json:"siid"`
	PropertyID int `json:"piid"`
}

var propertyRegistry = make(map[reflect.Type]PropertyBinding)

// RegisterProperty registers a property type.
// Internal use only.
func RegisterProperty(t reflect.Type, siid, piid int) {
	propertyRegistry[t] = PropertyBinding{siid, piid}
}

func MarshalProperties(v interface{}) ([]byte, error) {
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr {
		panic("v must be a pointer")
	}
	rv = rv.Elem()
	if rv.Kind() != reflect.Struct {
		panic("v must be a pointer to a struct")
	}
	type Property struct {
		DeviceID string `json:"did"`
		PropertyBinding
		Value interface{} `json:"value"`
	}
	props := make([]Property, 0, rv.NumField())
	for i := 0; i < rv.NumField(); i++ {
		f := rv.Field(i)
		binding, ok := propertyRegistry[f.Type()]
		if !ok {
			continue
		}
		props = append(props, Property{
			DeviceID:        "nothing",
			PropertyBinding: binding,
			Value:           f.Field(0).Interface(),
		})
	}
	return json.Marshal(props)
}

func UnmarshalProperties(b []byte, v interface{}) error {
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr {
		panic("v must be a pointer")
	}
	rv = rv.Elem()
	if rv.Kind() != reflect.Struct {
		panic("v must be a pointer to a struct")
	}
	props := make(map[PropertyBinding]reflect.Value)
	for i := 0; i < rv.NumField(); i++ {
		f := rv.Field(i)
		binding, ok := propertyRegistry[f.Type()]
		if !ok {
			continue
		}
		props[binding] = f
	}
	type Property struct {
		DeviceID string `json:"did"`
		PropertyBinding
		Value json.RawMessage    `json:"value"`
		Code  PropertyStatusCode `json:"code"`
	}
	var res []Property
	if err := json.Unmarshal(b, &res); err != nil {
		return err
	}
	var errs []error
	for _, rp := range res {
		p, ok := props[rp.PropertyBinding]
		if !ok {
			errs = append(errs, fmt.Errorf("unknown property binding: %v", rp.PropertyBinding))
		}
		delete(props, rp.PropertyBinding)
		if rp.Code != PropertyStatusCodeOK {
			err := PropertyError{rp.Code}
			p.Field(1).Set(reflect.ValueOf(err))
			errs = append(errs, err)
			continue
		}
		if len(rp.Value) > 0 {
			if err := json.Unmarshal(rp.Value, p.Field(0).Addr().Interface()); err != nil {
				p.Field(1).Set(reflect.ValueOf(err))
				errs = append(errs, err)
				continue
			}
		}
	}
	for _, p := range props {
		errs = append(errs, fmt.Errorf("missing property: %v", p.Type()))
	}
	return errors.Join(errs...)
}
