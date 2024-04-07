/*
 * @Copyright @ Shopee
 * @Author: Peidong Xu
 * @Email: peidong.xu@shopee.com
 * @Date: 2023/11/28 4:38 PM
 * @Version: kit
 */

package valuate

import "time"

type Array interface {
	Values() []Value
}

type Booleans []bool

func (b Booleans) Values() (ret []Value) {

	for i := range b {

		ret = append(ret, BoolValue(b[i]))
	}
	return
}

type Strings []string

func (s Strings) Values() (ret []Value) {
	for i := range s {

		ret = append(ret, StringValue(s[i]))
	}

	return
}

type Int64s []int64

func (s Int64s) Values() (ret []Value) {
	for i := range s {

		ret = append(ret, IntValue(s[i]))
	}

	return
}

type Int32s []int32

func (s Int32s) Values() (ret []Value) {
	for i := range s {

		ret = append(ret, IntValue(int64(s[i])))
	}

	return
}

type Integers []int

func (s Integers) Values() (ret []Value) {
	for i := range s {

		ret = append(ret, IntValue(int64(s[i])))
	}

	return
}

type Int16s []int16

func (s Int16s) Values() (ret []Value) {
	for i := range s {

		ret = append(ret, IntValue(int64(s[i])))
	}

	return
}

type Int8s []int8

func (s Int8s) Values() (ret []Value) {
	for i := range s {

		ret = append(ret, IntValue(int64(s[i])))
	}

	return
}

type Uints []uint

func (s Uints) Values() (ret []Value) {
	for i := range s {

		ret = append(ret, UintValue(uint64(s[i])))
	}

	return
}

type Uint64s []uint64

func (s Uint64s) Values() (ret []Value) {
	for i := range s {

		ret = append(ret, UintValue(s[i]))
	}

	return
}

type Uint32s []uint32

func (s Uint32s) Values() (ret []Value) {
	for i := range s {

		ret = append(ret, UintValue(uint64(s[i])))
	}

	return
}

type Uint16s []uint16

func (s Uint16s) Values() (ret []Value) {
	for i := range s {

		ret = append(ret, UintValue(uint64(s[i])))
	}

	return
}

type Float64s []float64

func (s Float64s) Values() (ret []Value) {
	for i := range s {
		ret = append(ret, FloatValue(s[i]))
	}

	return
}

type Float32s []float32

func (s Float32s) Values() (ret []Value) {
	for i := range s {
		ret = append(ret, FloatValue(float64(s[i])))
	}

	return
}

type UintPoints []uintptr

func (s UintPoints) Values() (ret []Value) {
	for i := range s {
		ret = append(ret, UintPoint(s[i]))
	}

	return
}

type Durations []time.Duration

func (s Durations) Values() (ret []Value) {
	for i := range s {
		ret = append(ret, DurationValue(s[i]))
	}
	return
}

type Values []Value

func (s Values) Values() []Value {
	return s
}

func (s Values) Get() Any {
	ret := make([]interface{}, 0)

	for _, v := range s {
		ret = append(ret, v.Get())
	}

	return ret
}
