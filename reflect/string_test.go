package reflect

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

type Person struct {
	Name   string
	Age    int
	Height float64
}

type User struct {
	ID       int
	Username string
	Email    string
	Active   bool
}

func structToJson(s interface{}) (string, error) {
	// reflect 패키지로 입력 구조체의 타입 정보를 가져옴
	t := reflect.TypeOf(s)

	// 입력 구조체의 값을 가져옴
	v := reflect.ValueOf(s)

	// JSON 데이터를 저장할 map
	jsonData := make(map[string]interface{})
	jsonDataType := make(map[string]interface{})
	// 구조체의 필드 수만큼 반복하며 필드명, 값, 타입을 map에 저장
	for i := 0; i < t.NumField(); i++ {
		// 필드명을 가져옴
		fieldName := t.Field(i).Name

		// 필드값을 가져옴
		fieldValue := v.Field(i)

		// 필드값이 포인터 타입이면, 포인터가 참조하는 값을 가져옴
		if fieldValue.Kind() == reflect.Ptr {

		}
		switch fieldValue.Kind() {
		case reflect.Ptr:
			if !fieldValue.IsNil() {
				fieldValue = fieldValue.Elem()
			} else {
				continue
			}
		case reflect.Array:

		}

		// 필드값의 타입을 가져옴
		fieldType := fieldValue.Type().String()

		// 필드명, 값, 타입을 map에 저장
		jsonData[fieldName] = fieldValue.Interface()
		jsonDataType[fieldName] = fieldType
	}

	// map을 JSON 문자열로 변환하여 반환
	jsonBytes, err := json.Marshal(jsonData)
	if err != nil {
		return "", err
	}
	return string(jsonBytes), nil
}
func getValue(v reflect.Value) string {
	switch v.Kind() {
	case reflect.Ptr:
		if !v.IsNil() {
			v = v.Elem()
		} else {
			//continue
		}
	case reflect.Array:

	}
	return ""
}

func main() {
	// Person 구조체로 테스트
	p := Person{"John", 30, 180.5}
	jsonStr, _ := structToJson(p)
	fmt.Println(jsonStr)

	// User 구조체로 테스트
	u := User{1, "johndoe", "johndoe@example.com", true}
	jsonStr, _ = structToJson(u)
	fmt.Println(jsonStr)

	// 포인터 타입을 가진 구조체로 테스트
	nu := &User{2, "janedoe", "janedoe@example.com", true}
	jsonStr, _ = structToJson(nu)
	fmt.Println(jsonStr)
}

type Person2 struct {
	Name    string
	Age     int
	Height  float64
	Account *User2
}

type User2 struct {
	ID       int
	Username string
	Email    string
	Active   bool
}

func TestReflectStructToJSONWithPointerField(t *testing.T) {
	u := User2{1, "j", "jjjj@example.com", true}
	p := Person2{
		Name:    "ryeom",
		Age:     25,
		Height:  165.2,
		Account: &u,
	}

	// User 구조체로 테스트

	jsonStr, err := structToJson(p)
	if err != nil {
		t.Errorf("ReflectStructToJSON failed: %v", err)
	}
	fmt.Println(jsonStr)
	//
	//expected := `{"Name":"ryeom","Age":25,"Height":165.2,"Account":{"ID":"1234","Name":"j"}}`
	//if jsonStr != expected {
	//	t.Errorf("ReflectStructToJSON failed. Expected:\n%s\nBut got:\n%s", expected, jsonStr)
	//}
}
