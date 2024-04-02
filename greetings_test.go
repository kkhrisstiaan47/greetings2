package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "Camilo"

	want := regexp.MustCompile(`\b` + name + `\b`)

	msg, err := Hello("Camilo")

	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Camilo") = %q, %v, quiere coincidencia para %#q, nil`, msg, err, want)
	}
}

func TestNameEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, quiere "", error`, msg, err)
	}
}
