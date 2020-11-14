package anyconvert

import (
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/any"
	"github.com/tigerwill90/scenario-go-sdk/proto/value"
)

// AnyToListValue return a new message of the type value.ListValue.
// Declaration & implementation generation are located in internal/proto/value package
// It returns an error if the target message in any.Any does
// not match or if an unmarshal error occurs.
func AnyToListValue(any *any.Any) (*value.ListValue, error) {
	listValue := new(value.ListValue)
	if err := ptypes.UnmarshalAny(any, listValue); err != nil {
		return nil, err
	}
	return listValue, nil
}

// AnyToMapValue return a new message of the type value.MapValue.
// Declaration & implementation generation are located in internal/proto/value package
// It returns an error if the target message in any.Any does
// not match or if an unmarshal error occurs.
func AnyToMapValue(any *any.Any) (*value.MapValue, error) {
	mapValue := new(value.MapValue)
	if err := ptypes.UnmarshalAny(any, mapValue); err != nil {
		return nil, err
	}
	return mapValue, nil
}

// AnyToFloatValue return a new message of the type value.Float32Value.
// Declaration & implementation generation are located in internal/proto/value package
// It returns an error if the target message in any.Any does
// not match or if an unmarshal error occurs.
func AnyToFloatValue(any *any.Any) (*value.Float32Value, error) {
	val := new(value.Float32Value)
	if err := ptypes.UnmarshalAny(any, val); err != nil {
		return nil, err
	}
	return val, nil
}

// AnyToIntValue return a new message of the type value.Int32Value.
// Declaration & implementation generation are located in internal/proto/value package
// It returns an error if the target message in any.Any does
// not match or if an unmarshal error occurs.
func AnyToIntValue(any *any.Any) (*value.Int32Value, error) {
	val := new(value.Int32Value)
	if err := ptypes.UnmarshalAny(any, val); err != nil {
		return nil, err
	}
	return val, nil
}

// AnyToStringValue return a new message of the type value.StringValue.
// Declaration & implementation generation are located in internal/proto/value package
// It returns an error if the target message in any.Any does
// not match or if an unmarshal error occurs.
func AnyToStringValue(any *any.Any) (*value.StringValue, error) {
	val := new(value.StringValue)
	if err := ptypes.UnmarshalAny(any, val); err != nil {
		return nil, err
	}
	return val, nil
}

// AnyToBoolValue return a new message of the type value.BoolValue.
// Declaration & implementation generation are located in internal/proto/value package
// It returns an error if the target message in any.Any does
// not match or if an unmarshal error occurs.
func AnyToBoolValue(any *any.Any) (*value.BoolValue, error) {
	val := new(value.BoolValue)
	if err := ptypes.UnmarshalAny(any, val); err != nil {
		return nil, err
	}
	return val, nil
}
