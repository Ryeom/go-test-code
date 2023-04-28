package test

import (
	"fmt"
	"regexp"
	"testing"
)

func TestRegexp(t *testing.T) {
	matched, _ := regexp.MatchString("[가-힣a-zA-Z]+\\([a-zA-Z0-9._]+\\)$", "asd강(sa1.wh_o23asdf)")
	fmt.Println(matched) // true

	matched, _ = regexp.MatchString("[a-zA-Z0-9._]+\\@$", "hello_test@")
	fmt.Println(matched) // true

	matched, _ = regexp.MatchString("[a-zA-Z0-9._]+\\@[a-zA-Z0-9]+\\.[a-zA-Z0-9]", "hi123@vivibibi.com")
	fmt.Println(matched) // true
	//
	//
	//regexp.MustCompile()
	//str := "김안녕(wnthfkdnr__23432501)"

}
