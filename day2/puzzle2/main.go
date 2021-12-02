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
	depth, distance, aim := 0, 0, 0
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
		switch v.instruction {
		case up:
			aim = aim - v.value
		case down:
			aim = aim + v.value
		case forward:
			distance = distance + v.value
			depth = depth + (aim * v.value)
		default:
			panic("ahhhh")
		}
	}
	fmt.Println("depth:", depth)
	fmt.Println("distance:", distance)
	fmt.Println("total:", depth*distance)
}
