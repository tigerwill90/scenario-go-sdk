package anyconvert

import (
	"fmt"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/any"
	pbvalue "github.com/tigerwill90/scenario-go-sdk/proto/value"
)

// AnyToFloat unmarshal an arbitrary any.Any message to a float32.
// Declaration & implementation generation are located in internal/proto/value package
// It expect an any.Any message which contain the url type for a
// value.Float32Value. It returns an error if the target message in
// any.Any does not match or if an unmarshal error occurs.
func AnyToFloat(value *any.Any) (float32, error) {
	val, err := AnyToFloatValue(value)
	if err != nil {
		return 0, fmt.Errorf("unable to convert to a float32: %w", err)
	}
	return val.Value, nil
}

func MustAnyToFloat(value *any.Any) float32 {
	v, err := AnyToFloat(value)
	handlePanic(err)
	return v
}

// AnyToInt unmarshal an arbitrary any.Any message to a int32.
// Declaration & implementation generation are located in internal/proto/value package
// It expect an any.Any message which contain the url type for a
// value.Int32Value. It returns an error if the target message in
// any.Any does not match or if an unmarshal error occurs.
func AnyToInt(value *any.Any) (int32, error) {
	val, err := AnyToIntValue(value)
	if err != nil {
		return 0, fmt.Errorf("unable to convert to an int32: %w", err)
	}
	return val.Value, nil
}

func MustAnyToInt(value *any.Any) int32 {
	v, err := AnyToInt(value)
	handlePanic(err)
	return v
}

// AnyToString unmarshal an arbitrary any.Any message to a string.
// Declaration & implementation generation are located in internal/proto/value package
// It expect an any.Any message which contain the url type for a
// value.StringValue. It returns an error if the target message in
// any.Any does not match or if an unmarshal error occurs.
func AnyToString(value *any.Any) (string, error) {
	val, err := AnyToStringValue(value)
	if err != nil {
		return "", fmt.Errorf("unable to convert to a string: %w", err)
	}
	return val.Value, nil
}

func MustAnyToString(value *any.Any) string {
	v, err := AnyToString(value)
	handlePanic(err)
	return v
}

// AnyToBool unmarshal an arbitrary any.Any message to a bool.
// Declaration & implementation generation are located in internal/proto/value package
// It expect an any.Any message which contain the url type for a
// value.BoolValue. It returns an error if the target message in
// any.Any does not match or if an unmarshal error occurs.
func AnyToBool(value *any.Any) (bool, error) {
	val, err := AnyToBoolValue(value)
	if err != nil {
		return false, fmt.Errorf("unable to convert to a bool: %w", err)
	}
	return val.Value, nil
}

func MustAnyToBool(value *any.Any) bool {
	v, err := AnyToBool(value)
	handlePanic(err)
	return v
}

// FloatToAny return an arbitrary any.Any message which contain the url type for a value.Float32Value.
// Declaration & implementation generation are located in internal/proto/value package
// Declaration & implementation generation are located in internal/proto/value package
// If an unmarshal error occurs, it return the error
func FloatToAny(value float32) (*any.Any, error) {
	val := &pbvalue.Float32Value{Value: value}
	v, err := ptypes.MarshalAny(val)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func MustFloatToAny(value float32) *any.Any {
	v, err := FloatToAny(value)
	handlePanic(err)
	return v
}

// IntToAny return an arbitrary any.Any message which contain the url type for a value.Int32Value.
// Declaration & implementation generation are located in internal/proto/value package
// If an unmarshal error occurs, it return the error
func IntToAny(value int32) (*any.Any, error) {
	val := &pbvalue.Int32Value{Value: value}
	v, err := ptypes.MarshalAny(val)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func MustIntToAny(value int32) *any.Any {
	v, err := IntToAny(value)
	handlePanic(err)
	return v
}

// StringToAny return an arbitrary any.Any message which contain the url type for a value.StringValue.
// Declaration & implementation generation are located in internal/proto/value package
// If an unmarshal error occurs, it return the error
func StringToAny(value string) (*any.Any, error) {
	val := &pbvalue.StringValue{Value: value}
	v, err := ptypes.MarshalAny(val)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func MustStringToAny(value string) *any.Any {
	v, err := StringToAny(value)
	handlePanic(err)
	return v
}

// BoolToAny return an arbitrary any.Any message which contain the url type for a value.BoolValue.
// Declaration & implementation generation are located in internal/proto/value package
// If an unmarshal error occurs, it return the error
func BoolToAny(value bool) (*any.Any, error) {
	val := &pbvalue.BoolValue{Value: value}
	v, err := ptypes.MarshalAny(val)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func MustBoolToAny(value bool) *any.Any {
	v, err := BoolToAny(value)
	handlePanic(err)
	return v
}
