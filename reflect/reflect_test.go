package reflect

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"
	"testing"
)

type ss struct {
	Name     string
	Age      int
	Nickname string
	hello    string
	Va       *http.Client
}

func TestReflect(t *testing.T) {
	obj := ss{"van", 12, "vivi", "nol", nil}
	val := reflect.ValueOf(&obj).Elem()

	//for i := 0; i < val.NumField(); i++ {
	//	name := val.Type().Field(i).Name
	//	fmt.Println(fmt.Sprint("- ", name, ": ", val.Field(i).Interface(), "\n"))
	//}

	a, _ := val.Type().FieldByName("Nickname")
	fieldName := a.Name

	value := val.Field(a.Index[0]).Interface()
	fmt.Println(fieldName, value)
	fmt.Println(StructToString("a", &obj))

}

func StructToString(title string, p interface{}) string {
	str := "[" + title + "]\n"
	val := reflect.ValueOf(p).Elem()
	fmt.Println(reflect.TypeOf(p).String())
	for i := 0; i < val.NumField(); i++ {
		a := val.Type().Field(i)

		if !strings.Contains(a.Type.String(), ".") {
			name := a.Name
			str += fmt.Sprint("- ", name, ": ", val.Field(i).Interface(), "\n")
		}

	}
	return str
}
