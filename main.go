package main

import "fmt"

func main() {

	vals := make([]string, 3)
	vals[0] = "foo"
	vals[1] = "bar"
	vals[2] = "foobar"

	fmt.Printf("%v\n", vals)

}
