package anyconvert

import (
	"fmt"
	"github.com/golang/protobuf/ptypes/any"
)

// AnyToFloatListMap unmarshal an arbitrary any.Any message to a map[string][]float32.
// Declaration & implementation generation are located in internal/proto/value package
// It expect an any.Any message which contain the url type for a
// value.MapValue. It returns an error if the target message in
// any.Any does not match or if an unmarshal error occurs.
func AnyToFloatListMap(value *any.Any) (map[string][]float32, error) {
	mapList := make(map[string][]float32)
	mapValue, err := AnyToMapValue(value)
	if err != nil {
		return nil, fmt.Errorf("unable to convert to a map[string][]float32: %w", err)
	}
	for k, value := range mapValue.Value {
		listValue, err := AnyToListValue(value)
		if err != nil {
			return nil, fmt.Errorf("unable to convert to a map[string][]float32: %w", err)
		}
		listType := make([]float32, 0, len(listValue.Value))
		for _, v := range listValue.Value {
			val, err := AnyToFloat(v)
			if err != nil {
				return nil, fmt.Errorf("unable to convert to a map[string][]float32: %w", err)
			}
			listType = append(listType, val)
		}
		mapList[k] = listType
	}
	return mapList, nil
}

func MustAnyToFloatListMap(value *any.Any) map[string][]float32 {
	v, err := AnyToFloatListMap(value)
	handlePanic(err)
	return v
}

// AnyToIntListMap unmarshal an arbitrary any.Any message to a map[string][]int32.
// Declaration & implementation generation are located in internal/proto/value package
// It expect an any.Any message which contain the url type for a
// value.MapValue. It returns an error if the target message in
// any.Any does not match or if an unmarshal error occurs.
func AnyToIntListMap(value *any.Any) (map[string][]int32, error) {
	mapList := make(map[string][]int32)
	mapValue, err := AnyToMapValue(value)
	if err != nil {
		return nil, fmt.Errorf("unable to convert to a map[string][]int32: %w", err)
	}
	for k, value := range mapValue.Value {
		listValue, err := AnyToListValue(value)
		if err != nil {
			return nil, fmt.Errorf("unable to convert to a map[string][]int32: %w", err)
		}
		listType := make([]int32, 0, len(listValue.Value))
		for _, v := range listValue.Value {
			val, err := AnyToInt(v)
			if err != nil {
				return nil, fmt.Errorf("unable to convert to a map[string][]int32: %w", err)
			}
			listType = append(listType, val)
		}
		mapList[k] = listType
	}
	return mapList, nil
}

func MustAnyToIntListMap(value *any.Any) map[string][]int32 {
	v, err := AnyToIntListMap(value)
	handlePanic(err)
	return v
}

// AnyToStringListMap unmarshal an arbitrary any.Any message to a map[string][]string.
// Declaration & implementation generation are located in internal/proto/value package
// It expect an any.Any message which contain the url type for a
// value.MapValue. It returns an error if the target message in
// any.Any does not match or if an unmarshal error occurs.
func AnyToStringListMap(value *any.Any) (map[string][]string, error) {
	mapList := make(map[string][]string)
	mapValue, err := AnyToMapValue(value)
	if err != nil {
		return nil, fmt.Errorf("unable to convert to a map[string][]string: %w", err)
	}
	for k, value := range mapValue.Value {
		listValue, err := AnyToListValue(value)
		if err != nil {
			return nil, fmt.Errorf("unable to convert to a map[string][]string: %w", err)
		}
		listType := make([]string, 0, len(listValue.Value))
		for _, v := range listValue.Value {
			val, err := AnyToString(v)
			if err != nil {
				return nil, fmt.Errorf("unable to convert to a map[string][]string: %w", err)
			}
			listType = append(listType, val)
		}
		mapList[k] = listType
	}
	return mapList, nil
}

func MustAnyToStringListMap(value *any.Any) map[string][]string {
	v, err := AnyToStringListMap(value)
	handlePanic(err)
	return v
}

// AnyToBoolMap unmarshal an arbitrary any.Any message to a map[string][]bool.
// Declaration & implementation generation are located in internal/proto/value package
// It expect an any.Any message which contain the url type for a
// value.MapValue. It returns an error if the target message in
// any.Any does not match or if an unmarshal error occurs.
func AnyToBoolListMap(value *any.Any) (map[string][]bool, error) {
	mapList := make(map[string][]bool)
	mapValue, err := AnyToMapValue(value)
	if err != nil {
		return nil, fmt.Errorf("unable to convert to a map[string][]bool: %w", err)
	}
	for k, value := range mapValue.Value {
		listValue, err := AnyToListValue(value)
		if err != nil {
			return nil, fmt.Errorf("unable to convert to a map[string][]bool: %w", err)
		}
		listType := make([]bool, 0, len(listValue.Value))
		for _, v := range listValue.Value {
			val, err := AnyToBool(v)
			if err != nil {
				return nil, fmt.Errorf("unable to convert to a map[string][]bool: %w", err)
			}
			listType = append(listType, val)
		}
		mapList[k] = listType
	}
	return mapList, nil
}

func MustAnyToBoolListMap(value *any.Any) map[string][]bool {
	v, err := AnyToBoolListMap(value)
	handlePanic(err)
	return v
}

// FloatListMapToAny return an arbitrary any.Any message which contain the url type for a value.MapValue.
// Declaration & implementation generation are located in internal/proto/value package
// If an unmarshal error occurs, it return the error.
func FloatListMapToAny(values map[string][]float32) (*any.Any, error) {
	mapListValue := make(map[string][]*any.Any)
	for k, value := range values {
		anyList := make([]*any.Any, 0, len(value))
		for _, v := range value {
			val, err := FloatToAny(v)
			if err != nil {
				return nil, err
			}
			anyList = append(anyList, val)
		}
		mapListValue[k] = anyList
	}
	return ListMapToAny(mapListValue)
}

func MustFloatListMapToAny(values map[string][]float32) *any.Any {
	v, err := FloatListMapToAny(values)
	handlePanic(err)
	return v
}

// IntListMapToAny return an arbitrary any.Any message which contain the url type for a value.MapValue.
// Declaration & implementation generation are located in internal/proto/value package
// If an unmarshal error occurs, it return the error.
func IntListMapToAny(values map[string][]int32) (*any.Any, error) {
	mapListValue := make(map[string][]*any.Any)
	for k, value := range values {
		anyList := make([]*any.Any, 0, len(value))
		for _, v := range value {
			val, err := IntToAny(v)
			if err != nil {
				return nil, err
			}
			anyList = append(anyList, val)
		}
		mapListValue[k] = anyList
	}
	return ListMapToAny(mapListValue)
}

func MustIntListMapToAny(values map[string][]int32) *any.Any {
	v, err := IntListMapToAny(values)
	handlePanic(err)
	return v
}

// StringListMapToAny return an arbitrary any.Any message which contain the url type for a value.MapValue.
// Declaration & implementation generation are located in internal/proto/value package
// If an unmarshal error occurs, it return the error.
func StringListMapToAny(values map[string][]string) (*any.Any, error) {
	mapListValue := make(map[string][]*any.Any)
	for k, value := range values {
		anyList := make([]*any.Any, 0, len(value))
		for _, v := range value {
			val, err := StringToAny(v)
			if err != nil {
				return nil, err
			}
			anyList = append(anyList, val)
		}
		mapListValue[k] = anyList
	}
	return ListMapToAny(mapListValue)
}

func MustStringListMapToAny(values map[string][]string) *any.Any {
	v, err := StringListMapToAny(values)
	handlePanic(err)
	return v
}

// BoolListMapToAny return an arbitrary any.Any message which contain the url type for a value.MapValue.
// Declaration & implementation generation are located in internal/proto/value package
// If an unmarshal error occurs, it return the error.
func BoolListMapToAny(values map[string][]bool) (*any.Any, error) {
	mapListValue := make(map[string][]*any.Any)
	for k, value := range values {
		anyList := make([]*any.Any, 0, len(value))
		for _, v := range value {
			val, err := BoolToAny(v)
			if err != nil {
				return nil, err
			}
			anyList = append(anyList, val)
		}
		mapListValue[k] = anyList
	}
	return ListMapToAny(mapListValue)
}

func MustBoolListMapToAny(values map[string][]bool) *any.Any {
	v, err := BoolListMapToAny(values)
	handlePanic(err)
	return v
}

// ListMapToAny return an arbitrary any.Any message which contain the url type for a value.MapValue.
// Declaration & implementation generation are located in internal/proto/value package
// If an unmarshal error occurs, it return the error.
func ListMapToAny(values map[string][]*any.Any) (*any.Any, error) {
	anyMap := make(map[string]*any.Any)
	for k, value := range values {
		listValue, err := ListToAny(value)
		if err != nil {
			return nil, err
		}
		anyMap[k] = listValue
	}
	mapValue, err := MapToAny(anyMap)
	if err != nil {
		return nil, err
	}
	return mapValue, nil
}
