package convert

import (
	"reflect"
	"strings"
	"testing"
)

func TestToMap(t *testing.T) {

	resp := &UserStruct{
		CsUserId:          133243242325454,
		LastName:          "last name",
		MiddleName:        "middle name",
		FirstName:         "first name",
		Region:            "SG",
		OtherPhoneNumbers: []string{"123", "4561232", "425345"},
	}

	t.Log(resp.ToMap())
}

type UserStruct struct {
	CsUserId          int64    `json:"cs_user_id"`
	Region            string   `json:"region"`
	LastName          string   `json:"last_username"`
	MiddleName        string   `json:"middle_username"`
	FirstName         string   `json:"first_username"`
	Address           string   `json:"address"`
	OtherPhoneNumbers []string `json:"other_phone_numbers"`
}

func (user *UserStruct) ToMap() map[string]interface{} {
	return Struct2Map(user, func(value reflect.Value) interface{} {
		return ReflectValue2Str(value, func(v reflect.Value) string {
			switch value.Kind() {
			case reflect.Slice, reflect.Array:
				var (
					strVal []string
				)
				for i := 0; i < value.Len(); i++ {
					strInf := ReflectBaseType2Str(value.Index(i))
					strVal = append(strVal, strInf)
				}
				return strings.Join(strVal, CommaStr)
			default:
				return EmptyStr
			}
		})
	})
}
