package main

import "fmt"

func testDummy() (string, string, error) {
	return "dummy1", "dummy2", nil
}

func dummyFunc() {
	str, _, err := testDummy()
	fmt.Println(str, err)
}
