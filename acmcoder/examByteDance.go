package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	test()
}
func test() {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	s1 := input.Text()
	input.Scan()
	m := input.Text()
	input.Scan()
	i,_ := strconv.Atoi(input.Text())
	var s []string
	for j := i; j > 0; j--{
		input.Scan()
		s = append(s, input.Text())

	}
	s = append(s, s1,m)

	for _, v := range s {

		fmt.Print(v)
	}

}
