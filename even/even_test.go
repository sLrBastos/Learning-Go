// TESTING PACKAGES
package even

import (
	"fmt"
)

// TESTING EVEN FUNCTION . RESULT -> ok      even    0.001s. 2 IS EVEN

// func TestEven(t *testing.T) {
// 	if !Even(2) {
// 		t.Log("2 should be even!")
// 		t.Fail()
// 	}
// }

// OBVIOUSLY 2 SHOULD BE NOT ODD. IT IS UST SO WE CAN SEE THE TEST FUNCTION RETURN AN ERROR

// --- FAIL: TestEven (0.00s)
//    c:\Users\siste\Desktop\godeeznutz\godeeznutz-intro\even\even_test.go:14: 2 should be odd
//FAIL
//FAIL	deez/even	0.092s
//FAIL
// func TestEven(t *testing.T) {
// 	if Even(2) {
// 		t.Log("2 should be odd")
// 		t.Fail()
// 	}
// }

// The Go test suite also allows you to incorporate example functions which serve as documentation and as tests.
// These functions need to start with Example.

func ExampleEven() {
	if Even(2) {
		fmt.Printf("Is even\n")
	}
	// OUTPUT:
	//Is even
}

//Those last two comments lines 1 are part of the example, go test uses those to check the generated output with the text in the comments.
// If there is a mismatch the test fails.
