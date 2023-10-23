package greetings

import (
	"fmt"
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "Jose"
	want := regexp.MustCompile(`\b` + name + `\b`)
	fmt.Println(want)
	msg, err := Hello("Jose")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello ("Jose") = %q, %v, quiere coincidencia para %#q, nil`, msg, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello ("") = %q, %v, quiere "", error`, msg, err)
	}
}
