// SHAUN PAUL SINGH BHARTH
// ANONYMISER FUNCTIONS FOR GOLANG SAMPLE JUNIOR INTERVIEW
// 4th March 2022

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("HELLO")
	var a []int = make([]int, 3, 5)
	a = append(a, 55)

	var s string = "HELLO"
	s = s + "s"

	fmt.Println(s)
	fmt.Println(cleanStringData("HELLO"))
	fmt.Println(cleanIntData(2000))
	fmt.Println(cleanEmailAddress("aaaaa@bbbbbb.com"))
}

func cleanEmailAddress(add string) string {
	var parta = strings.Split(add, "@")
	var a string = cleanStringData(parta[0])
	var b string = cleanStringData(parta[1])

	return a + "@" + b
}

func cleanStringData(in string) string {

	var temp string = ""
	for i := 0; i < len(in); i++ {

		temp = temp + strconv.Itoa(i)
	}

	return temp
}

func cleanIntData(in int) int {

	return in / in % in
}
