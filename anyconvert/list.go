package anyconvert

import (
	"fmt"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/any"
	pbvalue "github.com/tigerwill90/scenario-go-sdk/proto/value"
)

// AnyToFloatList unmarshal an arbitrary any.Any message to a []float32.
// Declaration & implementation generation are located in internal/proto/value package
// It expect an any.Any message which contain the url type for a
// value.ListValue. It returns an error if the target message in
// any.Any does not match or if an unmarshal error occurs.
func AnyToFloatList(value *any.Any) ([]float32, error) {
	listValue, err := AnyToListValue(value)
	if err != nil {
		return nil, fmt.Errorf("unable to convert to a []float32: %w", err)
	}
	vals := make([]float32, 0, len(listValue.Value))
	for _, v := range listValue.Value {
		val, err := AnyToFloat(v)
		if err != nil {
			return nil, fmt.Errorf("unable to convert to a []float32: %w", err)
		}
		vals = append(vals, val)
	}
	return vals, nil
}

func MustAnyToFloatList(value *any.Any) []float32 {
	v, err := AnyToFloatList(value)
	handlePanic(err)
	return v
}

// AnyToIntList unmarshal an arbitrary any.Any message to a []int32.
// Declaration & implementation generation are located in internal/proto/value package
// It expect an any.Any message which contain the url type for a
// value.ListValue. It returns an error if the target message in
// any.Any does not match or if an unmarshal error occurs.
func AnyToIntList(value *any.Any) ([]int32, error) {
	listValue, err := AnyToListValue(value)
	if err != nil {
		return nil, fmt.Errorf("unable to convert to a []int32: %w", err)
	}
	vals := make([]int32, 0, len(listValue.Value))
	for _, v := range listValue.Value {
		val, err := AnyToInt(v)
		if err != nil {
			return nil, fmt.Errorf("unable to convert to a []int32: %w", err)
		}
		vals = append(vals, val)
	}
	return vals, nil
}

func MustAnyToIntList(value *any.Any) []int32 {
	v, err := AnyToIntList(value)
	handlePanic(err)
	return v
}

// AnyToStringList unmarshal an arbitrary any.Any message to a []string.
// Declaration & implementation generation are located in internal/proto/value package
// It expect an any.Any message which contain the url type for a
// value.ListValue. It returns an error if the target message in
// any.Any does not match or if an unmarshal error occurs.
func AnyToStringList(value *any.Any) ([]string, error) {
	listValue, err := AnyToListValue(value)
	if err != nil {
		return nil, fmt.Errorf("unable to convert to a []string: %w", err)
	}
	vals := make([]string, 0, len(listValue.Value))
	for _, v := range listValue.Value {
		val, err := AnyToString(v)
		if err != nil {
			return nil, fmt.Errorf("unable to convert to a []string: %w", err)
		}
		vals = append(vals, val)
	}
	return vals, nil
}

func MustAnyToStringList(value *any.Any) []string {
	v, err := AnyToStringList(value)
	handlePanic(err)
	return v
}

// AnyToBoolList unmarshal an arbitrary any.Any message to a []bool.
// Declaration & implementation generation are located in internal/proto/value package
// It expect an any.Any message which contain the url type for a
// value.ListValue. It returns an error if the target message in
// any.Any does not match or if an unmarshal error occurs.
func AnyToBoolList(value *any.Any) ([]bool, error) {
	listValue, err := AnyToListValue(value)
	if err != nil {
		return nil, fmt.Errorf("unable to convert to a []bool: %w", err)
	}
	vals := make([]bool, 0, len(listValue.Value))
	for _, v := range listValue.Value {
		val, err := AnyToBool(v)
		if err != nil {
			return nil, fmt.Errorf("unable to convert to a []bool: %w", err)
		}
		vals = append(vals, val)
	}
	return vals, nil
}

func MustAnyToBoolList(value *any.Any) []bool {
	v, err := AnyToBoolList(value)
	handlePanic(err)
	return v
}

// FloatListToAny return an arbitrary any.Any message which contain the url type for a value.ListValue.
// Declaration & implementation generation are located in internal/proto/value package
// If an unmarshal error occurs, it return the error.
func FloatListToAny(values []float32) (*any.Any, error) {
	anyList := make([]*any.Any, 0, len(values))
	for _, v := range values {
		anyVal, err := FloatToAny(v)
		if err != nil {
			return nil, err
		}
		anyList = append(anyList, anyVal)
	}

	val, err := ListToAny(anyList)
	if err != nil {
		return nil, err
	}
	return val, nil
}

func MustFloatListToAny(values []float32) *any.Any {
	v, err := FloatListToAny(values)
	handlePanic(err)
	return v
}

// IntListToAny return an arbitrary any.Any message which contain the url type for a value.ListValue.
// Declaration & implementation generation are located in internal/proto/value package
// If an unmarshal error occurs, it return the error.
func IntListToAny(values []int32) (*any.Any, error) {
	anyList := make([]*any.Any, 0, len(values))
	for _, v := range values {
		anyVal, err := IntToAny(v)
		if err != nil {
			return nil, err
		}
		anyList = append(anyList, anyVal)
	}

	val, err := ListToAny(anyList)
	if err != nil {
		return nil, err
	}
	return val, nil
}

func MustIntListToAny(values []int32) *any.Any {
	v, err := IntListToAny(values)
	handlePanic(err)
	return v
}

// StringListToAny return an arbitrary any.Any message which contain the url type for a value.ListValue.
// Declaration & implementation generation are located in internal/proto/value package
// If an unmarshal error occurs, it return the error.
func StringListToAny(values []string) (*any.Any, error) {
	anyList := make([]*any.Any, 0, len(values))
	for _, v := range values {
		anyVal, err := StringToAny(v)
		if err != nil {
			return nil, err
		}
		anyList = append(anyList, anyVal)
	}

	val, err := ListToAny(anyList)
	if err != nil {
		return nil, err
	}
	return val, nil
}

func MustStringListToAny(values []string) *any.Any {
	v, err := StringListToAny(values)
	handlePanic(err)
	return v
}

// BoolListToAny return an arbitrary any.Any message which contain the url type for a value.ListValue.
// Declaration & implementation generation are located in internal/proto/value package
// If an unmarshal error occurs, it return the error.
func BoolListToAny(values []bool) (*any.Any, error) {
	anyList := make([]*any.Any, 0, len(values))
	for _, v := range values {
		anyVal, err := BoolToAny(v)
		if err != nil {
			return nil, err
		}
		anyList = append(anyList, anyVal)
	}

	val, err := ListToAny(anyList)
	if err != nil {
		return nil, err
	}
	return val, nil
}

func MustBoolListToAny(values []bool) *any.Any {
	v, err := BoolListToAny(values)
	handlePanic(err)
	return v
}

// ListToAny return an arbitrary any.Any message which contain the url type for a value.ListValue.
// Declaration & implementation generation are located in internal/proto/value package
// If an unmarshal error occurs, it return the error.
func ListToAny(values []*any.Any) (*any.Any, error) {
	listValue := &pbvalue.ListValue{Value: values}
	value, err := ptypes.MarshalAny(listValue)
	if err != nil {
		return nil, err
	}
	return value, err
}
