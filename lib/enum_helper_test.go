package lib

import "testing"

func TestMakeDisplayString(t *testing.T) {
	deviceState := "in_use"
	userDisplay := MakeDisplayString(deviceState)
	expected := "In use"
	if userDisplay != expected {
		t.Errorf("wanted: %s got: %s\n", expected, userDisplay)
	}
}


// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
// func TestHelloName(t *testing.T) {
//     name := "Gladys"
//     want := regexp.MustCompile(`\b`+name+`\b`)
//     msg, err := Hello("Gladys")
//     if !want.MatchString(msg) || err != nil {
//         t.Fatalf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
//     }
// }
//
