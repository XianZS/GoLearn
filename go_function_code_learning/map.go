package main

import (
	"fmt"
)

/*
	map:=make(map[key_type]value_type,(length))
	del:
		delete(map,key_name)
*/

func main() {
	fmt.Println("Go map using:")
	var c int
	scan, err := fmt.Scan(&c)
	if err != nil {
		fmt.Println(scan)
		return
	}
	ms := make(map[int]int)
	fmt.Printf("Please inter %v cople key and value:\n", c)
	for i := 1; i <= c; i++ {
		var k, v int
		_, err := fmt.Scan(&k, &v)
		if err != nil {
			return
		}
		ms[k] = v
	}
	fmt.Println(ms)
	//delete(ms, 1)
}
