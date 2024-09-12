package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "William"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("William")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Juan") = %q, %v, quiere coincidencia para %#q, nil`, msg, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	//fmt.Println(msg)
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, quiere "", error`, msg, err)
	}
}
