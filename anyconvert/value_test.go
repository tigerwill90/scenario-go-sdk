package anyconvert

import (
	"fmt"
	"google.golang.org/protobuf/types/known/anypb"
	"reflect"
	"testing"
)

func TestFromToAny(t *testing.T) {
	cases := []struct {
		name  string
		value interface{}
	}{
		{"string", "hello world"},
		{"float", 1.25},
		{"int", 5},
		{"bool", true},
		{"[]string", []string{"iron man", "thor", "hulk"}},
		{"[]float", []float32{1.2, 1.55, 100.565454}},
		{"[]int", []int{5, 6, 4, 10, 898, 87877}},
		{"[]bool", []bool{true, false, false}},
		{"map[string]string", map[string]string{"key1": "value1", "key2": "value2"}},
		{"map[string]float", map[string]float32{"key1": 1.2, "key2": 4568}},
		{"map[string]int", map[string]int32{"key1": 1, "key2": 4568}},
		{"map[string]bool", map[string]bool{"key1": true, "key2": false}},
		{"map[string][]string", map[string][]string{"key1": {"a", "e"}, "key2": {"i", "o"}}},
		{"map[string][]float", map[string][]float32{"key1": {3.41, 5.6}, "key2": {12, 21.456547}}},
		{"map[string][]int", map[string][]int32{"key1": {341, 56}, "key2": {12, 21456547}}},
		{"map[string][]bool", map[string][]bool{"key1": {true, true}, "key2": {false, true}}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			switch in := tc.value.(type) {
			case float32:
				any, err := ToAny(in)
				if err != nil {
					t.Fatal(err)
				}

				var out float32
				if err := FromAny(any, &out); err != nil {
					t.Error(err)
				}
				if !reflect.DeepEqual(in, out) {
					t.Errorf("want = %v, got = %v", in, out)
				}
			case int32:
				any, err := ToAny(in)
				if err != nil {
					t.Fatal(err)
				}

				var out int32
				if err := FromAny(any, &out); err != nil {
					t.Error(err)
				}
				if !reflect.DeepEqual(in, out) {
					t.Errorf("want = %v, got = %v", in, out)
				}
			case string:
				any, err := ToAny(in)
				if err != nil {
					t.Fatal(err)
				}

				var out string
				if err := FromAny(any, &out); err != nil {
					t.Error(err)
				}
				if !reflect.DeepEqual(in, out) {
					t.Errorf("want = %v, got = %v", in, out)
				}
			case bool:
				any, err := ToAny(in)
				if err != nil {
					t.Fatal(err)
				}

				var out bool
				if err := FromAny(any, &out); err != nil {
					t.Error(err)
				}
				if !reflect.DeepEqual(in, out) {
					t.Errorf("want = %v, got = %v", in, out)
				}
			case []float32:
				any, err := ToAny(in)
				if err != nil {
					t.Fatal(err)
				}

				var out []float32
				if err := FromAny(any, &out); err != nil {
					t.Error(err)
				}
				if !reflect.DeepEqual(in, out) {
					t.Errorf("want = %v, got = %v", in, out)
				}
			case []int32:
				any, err := ToAny(in)
				if err != nil {
					t.Fatal(err)
				}

				var out []int32
				if err := FromAny(any, &out); err != nil {
					t.Error(err)
				}
				if !reflect.DeepEqual(in, out) {
					t.Errorf("want = %v, got = %v", in, out)
				}
			case []string:
				any, err := ToAny(in)
				if err != nil {
					t.Fatal(err)
				}

				var out []string
				if err := FromAny(any, &out); err != nil {
					t.Error(err)
				}
				if !reflect.DeepEqual(in, out) {
					t.Errorf("want = %v, got = %v", in, out)
				}
			case []bool:
				any, err := ToAny(in)
				if err != nil {
					t.Fatal(err)
				}

				var out []bool
				if err := FromAny(any, &out); err != nil {
					t.Error(err)
				}
				if !reflect.DeepEqual(in, out) {
					t.Errorf("want = %v, got = %v", in, out)
				}
			case map[string]float32:
				any, err := ToAny(in)
				if err != nil {
					t.Fatal(err)
				}

				var out map[string]float32
				if err := FromAny(any, &out); err != nil {
					t.Error(err)
				}
				if !reflect.DeepEqual(in, out) {
					t.Errorf("want = %v, got = %v", in, out)
				}
			case map[string]int32:
				any, err := ToAny(in)
				if err != nil {
					t.Fatal(err)
				}

				var out map[string]int32
				if err := FromAny(any, &out); err != nil {
					t.Error(err)
				}
				if !reflect.DeepEqual(in, out) {
					t.Errorf("want = %v, got = %v", in, out)
				}
			case map[string]string:
				any, err := ToAny(in)
				if err != nil {
					t.Fatal(err)
				}

				var out map[string]string
				if err := FromAny(any, &out); err != nil {
					t.Error(err)
				}
				if !reflect.DeepEqual(in, out) {
					t.Errorf("want = %v, got = %v", in, out)
				}
			case map[string]bool:
				any, err := ToAny(in)
				if err != nil {
					t.Fatal(err)
				}

				var out map[string]bool
				if err := FromAny(any, &out); err != nil {
					t.Error(err)
				}
				if !reflect.DeepEqual(in, out) {
					t.Errorf("want = %v, got = %v", in, out)
				}
			case map[string][]float32:
				any, err := ToAny(in)
				if err != nil {
					t.Fatal(err)
				}

				var out map[string][]float32
				if err := FromAny(any, &out); err != nil {
					t.Error(err)
				}
				if !reflect.DeepEqual(in, out) {
					t.Errorf("want = %v, got = %v", in, out)
				}
			case map[string][]int32:
				any, err := ToAny(in)
				if err != nil {
					t.Fatal(err)
				}

				var out map[string][]int32
				if err := FromAny(any, &out); err != nil {
					t.Error(err)
				}
				if !reflect.DeepEqual(in, out) {
					t.Errorf("want = %v, got = %v", in, out)
				}
			case map[string][]string:
				any, err := ToAny(in)
				if err != nil {
					t.Fatal(err)
				}

				var out map[string][]string
				if err := FromAny(any, &out); err != nil {
					t.Error(err)
				}
				if !reflect.DeepEqual(in, out) {
					t.Errorf("want = %v, got = %v", in, out)
				}
			case map[string][]bool:
				any, err := ToAny(in)
				if err != nil {
					t.Fatal(err)
				}

				var out map[string][]bool
				if err := FromAny(any, &out); err != nil {
					t.Error(err)
				}
				if !reflect.DeepEqual(in, out) {
					t.Errorf("want = %v, got = %v", in, out)
				}
			}
		})
	}
}

func TestIsInvalidTypeErr(t *testing.T) {
	var b []byte
	_, err := ToAny(b)
	if !IsInvalidTypeErr(err) {
		t.Errorf("want = %s, got = %s", ErrInvalidType, err)
	}

	any := new(anypb.Any)
	err = FromAny(any, &b)
	if !IsInvalidTypeErr(err) {
		t.Errorf("want = %s, got = %s", ErrInvalidType, err)
	}
}

func TestIsRequiredPtrErr(t *testing.T) {
	var s string
	any := new(anypb.Any)
	err := FromAny(any, s)
	if !IsRequiredPtrErr(err) {
		t.Errorf("want = %s, got = %s", ErrPtrRequired, err)
	}
}

func Example() {
	any, err := ToAny("hello, world")
	if err != nil {
		panic(err)
	}

	var s string
	if err := FromAny(any, &s); err != nil {
		panic(err)
	}
	fmt.Println(s)
	// Output:
	// hello, world
}
