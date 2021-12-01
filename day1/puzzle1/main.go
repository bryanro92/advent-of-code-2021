package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	t, i, p := 0, 0, 0
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	s := bufio.NewScanner(f)
	for s.Scan() {
		if i == 0 {
			p, _ = strconv.Atoi(s.Text())
			i++
			continue
		}
		v, _ := strconv.Atoi(s.Text())
		if v > p {
			t++
		}
		p, _ = strconv.Atoi(s.Text())
		i++
	}
	fmt.Println(t)
}
