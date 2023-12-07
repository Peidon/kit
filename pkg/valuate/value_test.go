/*
 * @Copyright @ Shopee
 * @Author: Peidong Xu
 * @Email: peidong.xu@shopee.com
 * @Date: 2023/11/28 5:44 PM
 * @Version: kit
 */

package valuate

import (
	"testing"
	"time"
)

func TestTimeValue(t *testing.T) {
	// testing time value
	n := time.Now()
	v := AnyValue(n)
	if (v.Type != TimeFullType && v.Type != TimeType) || n.Sub(v.GetTime()) != 0 {
		t.Error("time value error")
		return
	}
}

func TestArrayValue(t *testing.T) {
	// testing array value
	b := []bool{true, false}
	v := AnyValue(b)
	a := v.GetArray()

	for i := range a {
		if a[i].Type != BoolType || a[i].GetBool() != b[i] {
			t.Error("array value error")
			return
		}
	}

	s := []string{"a", "b", "c"}
	v = AnyValue(s)
	a = v.GetArray()

	for i := range a {
		if a[i].Type != StringType || a[i].GetString() != s[i] {
			t.Error("array value error")
			return
		}
	}
}

func TestValue_GetPtr(t *testing.T) {
	abc := &ABC{}
	v := AnyValue(abc)
	if v.Type != StructType {
		t.Error("struct ptr error")
		return
	}
}

func TestStructValue(t *testing.T) {

	abc := ABC{
		a: 12,
		B: "abc",
		c: true,

		D: []int64{1, 2, 3},

		E: &ABC{
			a: 3,
			B: "efg",
		},

		f: 1234.2345,
	}

	// testing struct value
	v := AnyValue(abc)
	s := v.GetStruct()

	v = s.Field("a")
	if v.Type != IntType || v.GetInt() != int64(abc.a) {
		t.Error("struct value error")
		return
	}

	// testing void
	if v := s.Field("gg"); v.Type != VoidType {
		t.Error("struct void field error")
	}

	v = s.Field("b")
	if v.Type != StringType || v.GetString() != abc.B {
		t.Error("struct value error")
		return
	}

	v = s.Field("c")
	if v.Type != BoolType || v.GetBool() != abc.c {
		t.Error("struct bool value error")
		return
	}

	// testing struct ptr value
	v = AnyValue(&abc)
	s = v.GetStruct()

	v = s.Field("a")
	if v.Type != IntType || v.GetInt() != int64(abc.a) {
		t.Error("struct int value error")
		return
	}

	// testing nested struct value
	v = s.Field("abc")
	if v.Type != StructType {
		t.Error("struct nested error")
		return
	}
	ns := v.GetStruct()
	v = ns.Field("a")
	if v.Type != IntType || v.GetInt() != int64(abc.E.a) {
		t.Error("struct nested error")
		return
	}

	v = s.Field("f")
	if v.Type != FloatType || v.GetFloat() != abc.f {
		t.Error("struct float value error")
		return
	}

	v = s.Field("d")
	if v.Type != ArrayType {
		t.Error("struct field array value error")
		return
	}
	a := v.GetArray()
	for i := range a {
		ar := abc.D.([]int64)
		if a[i].Type != IntType || a[i].GetInt() != ar[i] {
			t.Error("struct field array value error")
			return
		}
	}

}

type ABC struct {
	a int
	B string `json:"b"`
	c bool

	D interface{} `json:"d"`
	E *ABC        `json:"abc"`

	f float64

	G GG

	GgList []*GG
}

type GG struct {
	I []byte
	h uint
	J time.Time
}
