/*
 * @Copyright @ Shopee
 * @Author: Peidong Xu
 * @Email: peidong.xu@shopee.com
 * @Date: 2023/11/22 11:52 AM
 * @Version: kit
 */

package valuate

import (
	"math"
	"time"
)

type Value struct {
	Type ValueType
	// use integer store unsigned int value, uint ptr, int, float, bool value
	// use str store string value
	// combine integer and inf store time value
	// use inf store other value, e.g. object, array, duration

	integer int64
	str     string
	inf     interface{}
}

type ValueType int

const (
	UnknownType ValueType = iota
	BinaryType
	BoolType
	IntType
	UintType
	FloatType
	StringType
	DurationType
	ArrayType
	StructType
	TimeType
	TimeFullType
	PtrType

	trueValue = 1
)

var (
	_minTimeInt64 = time.Unix(0, math.MinInt64)
	_maxTimeInt64 = time.Unix(0, math.MaxInt64)
)

func (v Value) GetBytes() ([]byte, error) {
	if v.Type != BinaryType {
		return nil, nil
	}
	if ret, ok := v.inf.([]byte); ok {
		return ret, nil
	}
	return nil, GetBytesValueError
}

func (v Value) GetString() string {
	return v.str
}

func (v Value) GetUint() uint64 {
	return uint64(v.integer)
}

func (v Value) GetInt() int64 {
	return v.integer
}

func (v Value) GetFloat() float64 {
	return math.Float64frombits(uint64(v.integer))
}

func (v Value) GetBool() bool {
	return v.integer == trueValue
}

func (v Value) GetDuration() time.Duration {
	return time.Duration(v.integer)
}

func (v Value) GetTime() (time.Time, error) {
	if v.Type == TimeFullType {
		return v.inf.(time.Time), nil
	}

	if v.inf != nil {
		if lo, ok := v.inf.(*time.Location); ok {
			return time.Unix(0, v.integer).In(lo), nil
		}
		return time.Time{}, GetTimeValueError
	} else {
		// Fall back to UTC if location is nil.
		return time.Unix(0, v.integer), nil
	}
}

func (v Value) GetPtr() uintptr {
	return uintptr(v.integer)
}

func AnyValue(any interface{}) Value {
	switch val := any.(type) {
	case bool:
		return BoolValue(val)
	case *bool:
	case float64:
		return FloatValue(val)
	case *float64, *float32:
	case float32:
		return FloatValue(float64(val))
	case int:
		return IntValue(int64(val))
	case *int, *int64, *int32, *int16, *int8:
	case int64:
		return IntValue(val)
	case int32:
		return IntValue(int64(val))
	case int16:
		return IntValue(int64(val))
	case int8:
		return IntValue(int64(val))
	case string:
		return StringValue(val)
	case *string:
	case uint:
		return UintValue(uint64(val))
	case *uint, *uint64, *uint32, *uint16, *uint8:
	case uint64:
		return UintValue(val)
	case uint32:
		return UintValue(uint64(val))
	case uint16:
		return UintValue(uint64(val))
	case uint8:
		return UintValue(uint64(val))
	case uintptr:
		return UintPoint(val)
	case *uintptr:
	case time.Time:
		return TimeValue(val)
	case *time.Time:
	case time.Duration:
		return DurationValue(val)
	case *time.Duration:
	case []bool:
		return ArrayValue(Bools(val))
	case []float64:
		return ArrayValue(Float64s(val))
	case []float32:
		return ArrayValue(Float32s(val))
	case []int:
		return ArrayValue(Ints(val))
	case []int64:
		return ArrayValue(Int64s(val))
	case []int32:
		return ArrayValue(Int32s(val))
	case []int16:
		return ArrayValue(Int16s(val))
	case []int8:
		return ArrayValue(Int8s(val))
	case []string:
		return ArrayValue(Strings(val))
	case []uint:
		return ArrayValue(Uints(val))
	case []uint64:
		return ArrayValue(Uint64s(val))
	case []uint32:
		return ArrayValue(Uint32s(val))
	case []uint16:
		return ArrayValue(Uint16s(val))
	case []byte:
		return ByteString(val)
	case []uintptr:
		return ArrayValue(UintPoints(val))
	case []time.Duration:
		return ArrayValue(Durations(val))
	case []time.Time:
	default:
		// reflect struct or slice

	}
	return Value{
		Type: UnknownType,
		inf:  any,
	}

}

type Struct map[string]Value

func StructValue(obj Struct) Value {
	return Value{
		Type: StructType,
		inf:  obj,
	}
}

func ArrayValue(arr Array) Value {
	return Value{
		Type: ArrayType,
		inf:  arr,
	}
}

func UintPoint(v uintptr) Value {
	return Value{
		Type:    PtrType,
		integer: int64(v),
	}
}

func ByteString(v []byte) Value {
	return Value{
		Type: BinaryType,
		inf:  v,
	}
}

func UintValue(v uint64) Value {
	return Value{
		Type:    UintType,
		integer: int64(v),
	}
}

func IntValue(v int64) Value {
	return Value{
		Type:    IntType,
		integer: v,
	}
}

func FloatValue(v float64) Value {
	return Value{
		Type:    FloatType,
		integer: int64(math.Float64bits(v)),
	}
}

func StringValue(v string) Value {
	return Value{
		Type: StringType,
		str:  v,
	}
}

func BoolValue(v bool) Value {
	var val int64
	if v {
		val = trueValue
	}
	return Value{
		Type:    BoolType,
		integer: val,
	}
}

func TimeValue(val time.Time) Value {
	if val.Before(_minTimeInt64) || val.After(_maxTimeInt64) {
		return Value{Type: TimeFullType, inf: val}
	}
	return Value{Type: TimeType, integer: val.UnixNano(), inf: val.Location()}
}

func DurationValue(val time.Duration) Value {
	return Value{
		Type:    DurationType,
		integer: int64(val),
	}
}

type Comparable interface {
	Equal(other Comparable) bool
	Greater(other Comparable) bool
}
