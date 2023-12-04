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
	"reflect"
	"time"
)

type Value struct {
	Type ValueType
	// use integer store unsigned int value, uint ptr, int, float, bool value, duration
	// use str store string value
	// combine integer and inf store time value
	// use inf store other value, e.g. struct, array

	integer int64
	str     string
	inf     interface{}
}

type ValueType int8

const (
	UnknownType ValueType = iota
	VoidType
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

	noopTime = time.Time{}
	timeType = reflect.TypeOf(noopTime)
)

func (v Value) Get() Any {
	switch v.Type {
	case BinaryType:
		return v.GetBytes()
	case BoolType:
		return v.GetBool()
	case IntType:
		return v.GetInt()
	case UintType:
		return v.GetUint()
	case FloatType:
		return v.GetFloat()
	case StringType:
		return v.GetString()
	case DurationType:
		return v.GetDuration()
	case ArrayType:
		return v.GetArray()
	case StructType:
		return v.GetStruct()
	case TimeType, TimeFullType:
		return v.GetTime()
	case PtrType:
		return v.GetPtr()
	}
	return nil
}

func (v Value) GetBytes() []byte {
	if v.Type != BinaryType {
		return nil
	}
	if ret, ok := v.inf.([]byte); ok {
		return ret
	}
	return nil
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

func (v Value) GetTime() time.Time {
	if v.Type == TimeFullType {
		return v.inf.(time.Time)
	}

	if v.inf != nil {
		if lo, ok := v.inf.(*time.Location); ok {
			return time.Unix(0, v.integer).In(lo)
		}
		return time.Time{}
	} else {
		// Fall back to UTC if location is nil.
		return time.Unix(0, v.integer)
	}
}

func (v Value) GetPtr() uintptr {
	return uintptr(v.integer)
}

func (v Value) GetStruct() Struct {
	if v.Type != StructType {
		return nil
	}
	if val, ok := v.inf.(Struct); ok {
		return val
	}
	return nil
}

func (v Value) GetArray() []Value {
	if v.Type != ArrayType {
		return nil
	}
	if val, ok := v.inf.([]Value); ok {
		return val
	}
	return nil
}

func AnyValue(any interface{}) Value {
	switch val := any.(type) {
	case Value:
		return val
	case Array:
		return ArrayValue(val)
	case Struct:
		return StructValue(val, "valuate.Struct")
	case bool:
		return BoolValue(val)
	case float64:
		return FloatValue(val)
	case float32:
		return FloatValue(float64(val))
	case int:
		return IntValue(int64(val))
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
	case uint:
		return UintValue(uint64(val))
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
	case time.Time:
		return TimeValue(val)
	case time.Duration:
		return DurationValue(val)
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
	case reflect.Value:
		// reflect
		return reflectValue(val)
	default:
		// ptr, struct or slice
		return reflectValue(reflect.ValueOf(val))
	}
}

func reflectValue(val reflect.Value) Value {

	switch val.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return IntValue(val.Int())

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return UintValue(val.Uint())

	case reflect.String:
		return StringValue(val.String())

	case reflect.Bool:
		return BoolValue(val.Bool())

	case reflect.Float32, reflect.Float64:
		return FloatValue(val.Float())

	case reflect.Pointer:
		return AnyValue(val.Elem())

	case reflect.Interface:
		if val.CanInterface() {
			return AnyValue(val.Interface())
		}

	case reflect.Map:
		s := Struct{}
		tp := val.Type()

		keys := val.MapKeys()

		for i := range keys {
			k := keys[i]
			v := val.MapIndex(k)

			s[k.String()] = AnyValue(v)
		}

		return StructValue(s, tp.String())

	case reflect.Struct:
		// special
		if val.Type() == timeType {
			if !val.CanInterface() {
				return TimeValue(noopTime)
			}
			inf := val.Interface()
			if tim, ok := inf.(time.Time); ok {
				return TimeValue(tim)
			}
		}

		// normal
		s := Struct{}
		tp := val.Type()

		for i := 0; i < tp.NumField(); i++ {
			field := tp.Field(i)
			tagStr, ok := field.Tag.Lookup(TagJson)
			if !ok {
				// if no json tag, use field name
				tagStr = field.Name
			}
			key, ignore, omitempty := getKey(tagStr)
			if ignore || key == EmptyStr {
				continue
			}

			value := reflect.Indirect(val).Field(i)
			if ok && value.IsZero() && omitempty {
				continue
			}

			s[key] = AnyValue(value)
		}

		return StructValue(s, tp.String())

	case reflect.Slice, reflect.Array:
		arr := make([]Value, 0, val.Len())

		for i := 0; i < val.Len(); i++ {
			v := val.Index(i)
			arr = append(arr, AnyValue(v))
		}

		return Value{
			Type: ArrayType,
			inf:  arr,
		}

	}
	return Value{
		Type: UnknownType,
		str:  val.Kind().String(),
		inf:  nil,
	}
}

// Struct 优先使用json tag 中的名字作为 key
// 如果没有 json tag 则使用Field Name 作为 key
type Struct map[string]Value

func (s Struct) Field(k string) Value {
	if v, ok := s[k]; ok {
		return v
	}
	return Value{
		Type: VoidType,
	}
}

func StructValue(obj Struct, ty string) Value {
	return Value{
		Type: StructType,
		str:  ty,
		inf:  obj,
	}
}

func ArrayValue(arr Array) Value {
	return Value{
		Type: ArrayType,
		inf:  arr.Values(),
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

func (v Value) Equal(other Comparable) bool {
	if ot, ok := other.(Value); ok {
		return ot.Type == v.Type &&
			ot.str == v.str &&
			ot.integer == v.integer &&
			reflect.DeepEqual(ot.inf, v.inf)
	}

	return false
}

func (v Value) Greater(other Comparable) bool {
	ot, ok := other.(Value)
	if !ok {
		return false
	}

	a := v.Get()
	b := ot.Get()

	if isNumber(a) && isNumber(b) {
		return toFloat64(a) > toFloat64(b)
	}

	if (v.Type == TimeFullType || v.Type == TimeType) &&
		(ot.Type == TimeType || ot.Type == TimeFullType) {
		return v.GetTime().After(ot.GetTime())
	}

	return false
}
