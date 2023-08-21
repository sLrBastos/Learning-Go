package stack

import "testing"

// TESTING STACK. SEEING IF PUSH, BY PUSHING 5, AND POP BY SEEING IF IT POPS THE 5 THAT WAS JUST PUSHED
func TestPushPop(t *testing.T) {
	c := new(Stack)
	c.Push(5)
	if c.Pop() != 5 {
		t.Log("Pop should have been 5")
		t.Fail()
	}
}
