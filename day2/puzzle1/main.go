package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type instruction struct {
	instruction string
	value       int
}

const (
	forward string = "forward"
	up      string = "up"
	down    string = "down"
)

func main() {
	depth, distance := 0, 0
	instructions := []instruction{}

	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		line := string(s.Text())
		i := strings.Split(line, " ")
		v, _ := strconv.Atoi(i[1]) // yeet
		instruction := instruction{
			instruction: i[0],
			value:       v,
		}
		instructions = append(instructions, instruction)
	}

	for _, v := range instructions {
		if strings.EqualFold(v.instruction, up) {
			depth = depth - v.value
		}
		if strings.EqualFold(v.instruction, down) {
			depth = depth + v.value
		}
		if strings.EqualFold(v.instruction, forward) {
			distance = distance + v.value
		}
	}
	fmt.Println("depth:", depth)
	fmt.Println("distance:", distance)
	fmt.Println("total:", depth*distance)

}
