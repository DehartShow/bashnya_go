package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	arg := os.Args[1:]
	if len(arg) != 1 {
		fmt.Println("Usage: go run hw2.go <command>")
		os.Exit(1)
	}
	var number, _ = strconv.Atoi(arg[0])
	if number > 12307 {
		fmt.Println("большое... возьми поменьше")
		os.Exit(1)
	}
	if (number%9 == 0) && (number%13 == 0) {
		fmt.Printf("service error\n")
		os.Exit(1)
	}

	for number < 12307 {
		if number < 0 {
			number = -number
		} else if number%7 == 0 {
			number = number * 39
		} else if number%9 == 0 {
			number = (number * 13) + 1
			continue
		} else {
			number = (number + 2) * 3
		}

		if (number%9 == 0) && (number%13 == 0) {
			fmt.Println("service error")
			os.Exit(1)
		} else {
			number += 1
		}
	}
	fmt.Printf("Полученное число: %d", number)

}
