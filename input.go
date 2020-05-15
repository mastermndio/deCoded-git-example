package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	buf := bufio.NewReader(os.Stdin)
	fmt.Println("What is your age?: ")
	age, err := buf.ReadBytes('\n')

	if err != nil {
		fmt.Println(age)
	} else {
		fmt.Println(err)
	}
}
