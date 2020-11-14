package anyconvert

import (
	"errors"
	"fmt"
	"github.com/golang/protobuf/ptypes/any"
	"reflect"
)

var (
	ErrInvalidType = errors.New("invalid type")
	ErrPtrRequired = errors.New("pointer required")
)

/*
	FromAny unserialize Any into a arbitrary value. Any provide
	great interoperability between programming language and is
	extensively used in the plugin protocol through gRPC.

	If the value given is not a pointer, FromAny will return an
	error of type ErrPtrRequired. It the value type is not supported,
	it will return an error of type ErrInvalidType. If the any type url
	is not supported, it will return the marshaling/unmarshaling error.

	Supported message type are
	- value.Float32Value
	- value.Int32Value
	- value.StringValue
	- value.BoolValue
	- value.ListValue
	- value.MapValue

	Declaration & implementation generation are located in internal/proto/value package

	Supported value type are
	- float32
	- int32
	- string
	- bool
	- []float32
	- []int32
	- []string
	- []bool
	- map[string]float32
	- map[string]int32
	- map[string]string
	- map[string]bool
	- map[string][]float32
	- map[string][]int32
	- map[string][]string
	- map[string][]bool
*/
func FromAny(any *any.Any, value interface{}) error {
	if !isPointer(value) {
		return ErrPtrRequired
	}

	switch val := value.(type) {
	case *float32:
		v, err := AnyToFloat(any)
		if err != nil {
			return err
		}
		*val = v
	case *int32:
		v, err := AnyToInt(any)
		if err != nil {
			return err
		}
		*val = v
	case *string:
		v, err := AnyToString(any)
		if err != nil {
			return err
		}
		*val = v
	case *bool:
		v, err := AnyToBool(any)
		if err != nil {
			return err
		}
		*val = v
	case *[]float32:
		v, err := AnyToFloatList(any)
		if err != nil {
			return err
		}
		*val = v
	case *[]int32:
		v, err := AnyToIntList(any)
		if err != nil {
			return err
		}
		*val = v
	case *[]string:
		v, err := AnyToStringList(any)
		if err != nil {
			return err
		}
		*val = v
	case *[]bool:
		v, err := AnyToBoolList(any)
		if err != nil {
			return err
		}
		*val = v
	case *map[string]float32:
		v, err := AnyToFloatMap(any)
		if err != nil {
			return err
		}
		*val = v
	case *map[string]int32:
		v, err := AnyToIntMap(any)
		if err != nil {
			return err
		}
		*val = v
	case *map[string]string:
		v, err := AnyToStringMap(any)
		if err != nil {
			return err
		}
		*val = v
	case *map[string]bool:
		v, err := AnyToBoolMap(any)
		if err != nil {
			return err
		}
		*val = v
	case *map[string][]float32:
		v, err := AnyToFloatListMap(any)
		if err != nil {
			return err
		}
		*val = v
	case *map[string][]int32:
		v, err := AnyToIntListMap(any)
		if err != nil {
			return err
		}
		*val = v
	case *map[string][]string:
		v, err := AnyToStringListMap(any)
		if err != nil {
			return err
		}
		*val = v
	case *map[string][]bool:
		v, err := AnyToBoolListMap(any)
		if err != nil {
			return err
		}
		*val = v
	default:
		return fmt.Errorf("%s type is not supported: %w", reflect.TypeOf(value).String(), ErrInvalidType)
	}
	return nil
}

/*
	ToAny serialize the given arbitrary value to a reference of type Any.
	Any provide great interoperability between programming language
	and is extensively used in the plugin protocol through gRPC. If the
	value type is not supported, ToAny return an error of type
	ErrInvalidType. If the any type url is not supported, it will
	return the marshaling/unmarshaling error.

	Supported message type are
	- value.Float32Value
	- value.Int32Value
	- value.StringValue
	- value.BoolValue
	- value.ListValue
	- value.MapValue

	Declaration & implementation generation are located in internal/proto/value package

	Supported type are
	- float32
	- int32
	- string
	- bool
	- []float32
	- []int32
	- []string
	- []bool
	- map[string]float32
	- map[string]int32
	- map[string]string
	- map[string]bool
	- map[string][]float32
	- map[string][]int32
	- map[string][]string
	- map[string][]bool
*/
func ToAny(value interface{}) (*any.Any, error) {
	switch val := value.(type) {
	case float32:
		any, err := FloatToAny(val)
		if err != nil {
			return nil, err
		}
		return any, nil
	case int32:
		any, err := IntToAny(val)
		if err != nil {
			return nil, err
		}
		return any, nil
	case string:
		any, err := StringToAny(val)
		if err != nil {
			return nil, err
		}
		return any, nil
	case bool:
		any, err := BoolToAny(val)
		if err != nil {
			return nil, err
		}
		return any, nil
	case []float32:
		any, err := FloatListToAny(val)
		if err != nil {
			return nil, err
		}
		return any, nil
	case []int32:
		any, err := IntListToAny(val)
		if err != nil {
			return nil, err
		}
		return any, nil
	case []string:
		any, err := StringListToAny(val)
		if err != nil {
			return nil, err
		}
		return any, nil
	case []bool:
		any, err := BoolListToAny(val)
		if err != nil {
			return nil, err
		}
		return any, nil
	case map[string]float32:
		any, err := FloatMapToAny(val)
		if err != nil {
			return nil, err
		}
		return any, nil
	case map[string]int32:
		any, err := IntMapToAny(val)
		if err != nil {
			return nil, err
		}
		return any, nil
	case map[string]string:
		any, err := StringMapToAny(val)
		if err != nil {
			return nil, err
		}
		return any, nil
	case map[string]bool:
		any, err := BoolMapToAny(val)
		if err != nil {
			return nil, err
		}
		return any, nil
	case map[string][]float32:
		any, err := FloatListMapToAny(val)
		if err != nil {
			return nil, err
		}
		return any, nil
	case map[string][]int32:
		any, err := IntListMapToAny(val)
		if err != nil {
			return nil, err
		}
		return any, nil
	case map[string][]string:
		any, err := StringListMapToAny(val)
		if err != nil {
			return nil, err
		}
		return any, nil
	case map[string][]bool:
		any, err := BoolListMapToAny(val)
		if err != nil {
			return nil, err
		}
		return any, nil
	}
	return nil, fmt.Errorf("%s type is not supported: %w", reflect.TypeOf(value).String(), ErrInvalidType)
}

func isPointer(val interface{}) bool {
	if reflect.ValueOf(val).Kind() != reflect.Ptr {
		return false
	}
	return true
}

// IsInvalidTypeErr return true if the error is an ErrInvalidType
func IsInvalidTypeErr(err error) bool {
	return errors.Is(err, ErrInvalidType)
}

// IsRequiredPtrErr return true if the error is an ErrPtrRequired
func IsRequiredPtrErr(err error) bool {
	return errors.Is(err, ErrPtrRequired)
}

func handlePanic(err error) {
	if err != nil {
		panic(err)
	}
}
