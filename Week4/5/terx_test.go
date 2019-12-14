package main
import "testing"
func testAdder() {
	sum := add(2 , 2)
	expected := 4

	if sun != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}


}
