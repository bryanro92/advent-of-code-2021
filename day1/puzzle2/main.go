package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	t := 0
	a := []int{}

	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)

	for s.Scan() {
		v, _ := strconv.Atoi(s.Text()) // yeet
		a = append(a, v)
	}

	for i, _ := range a {
		if i == 0 { // skip first
			continue
		}
		if i+2 >= len(a) { // skip last three
			continue
		}
		s1 := a[i-1] + a[i] + a[i+1] // previous window
		s2 := a[i] + a[i+1] + a[i+2] // current window
		if s2 > s1 {
			t++
		}

	}
	fmt.Println(t)
}
