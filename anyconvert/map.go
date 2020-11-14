package anyconvert

import (
	"fmt"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/any"
	pbvalue "github.com/tigerwill90/scenario-go-sdk/proto/value"
)

// AnyToFloatMap unmarshal an arbitrary any.Any message to a map[string]float32.
// Declaration & implementation generation are located in internal/proto/value package
// It expect an any.Any message which contain the url type for a
// value.MapValue. It returns an error if the target message in
// any.Any does not match or if an unmarshal error occurs.
func AnyToFloatMap(value *any.Any) (map[string]float32, error) {
	mapValue, err := AnyToMapValue(value)
	if err != nil {
		return nil, fmt.Errorf("unable to convert to a map[string]float32: %w", err)
	}
	vals := make(map[string]float32)
	for k, v := range mapValue.Value {
		val, err := AnyToFloat(v)
		if err != nil {
			return nil, fmt.Errorf("unable to convert to a map[string]float32: %w", err)
		}
		vals[k] = val
	}
	return vals, nil
}

func MustAnyToFloatMap(value *any.Any) map[string]float32 {
	v, err := AnyToFloatMap(value)
	handlePanic(err)
	return v
}

// AnyToIntMap unmarshal an arbitrary any.Any message to a map[string]int32.
// Declaration & implementation generation are located in internal/proto/value package
// It expect an any.Any message which contain the url type for a
// value.MapValue. It returns an error if the target message in
// any.Any does not match or if an unmarshal error occurs.
func AnyToIntMap(value *any.Any) (map[string]int32, error) {
	mapValue, err := AnyToMapValue(value)
	if err != nil {
		return nil, fmt.Errorf("unable to convert to a map[string]int32: %w", err)
	}
	vals := make(map[string]int32)
	for k, v := range mapValue.Value {
		val, err := AnyToInt(v)
		if err != nil {
			return nil, fmt.Errorf("unable to convert to a map[string]int32: %w", err)
		}
		vals[k] = val
	}
	return vals, nil
}

func MustAnyToIntMap(value *any.Any) map[string]int32 {
	v, err := AnyToIntMap(value)
	handlePanic(err)
	return v
}

// AnyToStringMap unmarshal an arbitrary any.Any message to a map[string]string.
// Declaration & implementation generation are located in internal/proto/value package
// It expect an any.Any message which contain the url type for a
// value.MapValue. It returns an error if the target message in
// any.Any does not match or if an unmarshal error occurs.
func AnyToStringMap(value *any.Any) (map[string]string, error) {
	mapValue, err := AnyToMapValue(value)
	if err != nil {
		return nil, fmt.Errorf("unable to convert to a map[string]string: %w", err)
	}
	vals := make(map[string]string)
	for k, v := range mapValue.Value {
		val, err := AnyToString(v)
		if err != nil {
			return nil, fmt.Errorf("unable to convert to a map[string]string: %w", err)
		}
		vals[k] = val
	}
	return vals, nil
}

func MustAnyToStringMap(value *any.Any) map[string]string {
	v, err := AnyToStringMap(value)
	handlePanic(err)
	return v
}

// AnyToBoolMap unmarshal an arbitrary any.Any message to a map[string]bool.
// Declaration & implementation generation are located in internal/proto/value package
// It expect an any.Any message which contain the url type for a
// value.MapValue. It returns an error if the target message in
// any.Any does not match or if an unmarshal error occurs.
func AnyToBoolMap(value *any.Any) (map[string]bool, error) {
	mapValue, err := AnyToMapValue(value)
	if err != nil {
		return nil, fmt.Errorf("unable to convert to a map[string]bool: %w", err)
	}
	vals := make(map[string]bool)
	for k, v := range mapValue.Value {
		val, err := AnyToBool(v)
		if err != nil {
			return nil, fmt.Errorf("unable to convert to a map[string]bool: %w", err)
		}
		vals[k] = val
	}
	return vals, nil
}

func MustAnyToBoolMap(value *any.Any) map[string]bool {
	v, err := AnyToBoolMap(value)
	handlePanic(err)
	return v
}

// FloatMapToAny return an arbitrary any.Any message which contain the url type for a value.MapValue.
// Declaration & implementation generation are located in internal/proto/value package
// If an unmarshal error occurs, it return the error.
func FloatMapToAny(values map[string]float32) (*any.Any, error) {
	anyMap := make(map[string]*any.Any)
	for k, v := range values {
		val, err := FloatToAny(v)
		if err != nil {
			return nil, err
		}
		anyMap[k] = val
	}
	return MapToAny(anyMap)
}

func MustFloatMapToAny(values map[string]float32) *any.Any {
	v, err := FloatMapToAny(values)
	handlePanic(err)
	return v
}

// IntMapToAny return an arbitrary any.Any message which contain the url type for a value.MapValue.
// Declaration & implementation generation are located in internal/proto/value package
// If an unmarshal error occurs, it return the error.
func IntMapToAny(values map[string]int32) (*any.Any, error) {
	anyMap := make(map[string]*any.Any)
	for k, v := range values {
		val, err := IntToAny(v)
		if err != nil {
			return nil, err
		}
		anyMap[k] = val
	}
	return MapToAny(anyMap)
}

func MustIntMapToAny(values map[string]int32) *any.Any {
	v, err := IntMapToAny(values)
	handlePanic(err)
	return v
}

// StringMapToAny return an arbitrary any.Any message which contain the url type for a value.MapValue.
// Declaration & implementation generation are located in internal/proto/value package
// If an unmarshal error occurs, it return the error.
func StringMapToAny(values map[string]string) (*any.Any, error) {
	anyMap := make(map[string]*any.Any)
	for k, v := range values {
		val, err := StringToAny(v)
		if err != nil {
			return nil, err
		}
		anyMap[k] = val
	}
	return MapToAny(anyMap)
}

func MustStringMapToAny(values map[string]string) *any.Any {
	v, err := StringMapToAny(values)
	handlePanic(err)
	return v
}

// BoolMapToAny return an arbitrary any.Any message which contain the url type for a value.MapValue.
// Declaration & implementation generation are located in internal/proto/value package
// If an unmarshal error occurs, it return the error.
func BoolMapToAny(values map[string]bool) (*any.Any, error) {
	anyMap := make(map[string]*any.Any)
	for k, v := range values {
		val, err := BoolToAny(v)
		if err != nil {
			return nil, err
		}
		anyMap[k] = val
	}
	return MapToAny(anyMap)
}

func MustBoolMapToAny(values map[string]bool) *any.Any {
	v, err := BoolMapToAny(values)
	handlePanic(err)
	return v
}

// MapToAny return an arbitrary any.Any message which contain the url type for a value.MapValue.
// Declaration & implementation generation are located in internal/proto/value package
// If an unmarshal error occurs, it return the error.
func MapToAny(values map[string]*any.Any) (*any.Any, error) {
	mapValue := &pbvalue.MapValue{Value: values}
	value, err := ptypes.MarshalAny(mapValue)
	if err != nil {
		return nil, err
	}
	return value, nil
}
