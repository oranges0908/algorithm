package main

import "fmt"

func testBackspaceCompare() {
	fmt.Println(getString("") == "")
	fmt.Println(getString("#") == "")
	fmt.Println(getString("a#") == "")
	fmt.Println(getString("a##") == "")
	fmt.Println(getString("ab##") == "")
	fmt.Println(getString("##") == "")
	fmt.Println(getString("##a") == "a")
	fmt.Println(getString("a##a") == "a")
	fmt.Println(getString("ab#a") == "aa")
	fmt.Println(backspaceCompare("ab#c", "ad#c") == true)
	fmt.Println(backspaceCompare("ab##", "c#d#") == true)
	fmt.Println(backspaceCompare("a#c", "b") == false)
}

func backspaceCompare(s string, t string) bool {
	rs := getString(s)
	rt := getString(t)
	return rs == rt
}

func getString(s string) string {
	rs := ""
	backspace := uint8('#')
	for j := 0; j < len(s); j++ {
		if s[j] == backspace {
			if len(rs)-1 >= 0 {
				rs = rs[:len(rs)-1]
			} else {
				rs = ""
			}
		} else {
			rs += string(s[j])
		}
	}
	return rs
}
