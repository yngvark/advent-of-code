package main

import (
	"bufio"
	"fmt"
	"github.com/yngvark/advent-of-code/pkg/helloworld"
	"os"
	"strconv"
)

func main() {
	fmt.Println(helloworld.Hello())

	err := run()
	if err != nil {
		fmt.Println("Program error:")
		fmt.Println(err.Error())
	}
}

func run() error {
	file, err := os.Open("data.txt")
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	increaseCount := 0
	decreaseCount := 0
	lastDepth := -1
	for scanner.Scan() {
		line := scanner.Text()

		depth, err := toInt(line)
		if err != nil {
			return err
		}

		if lastDepth == -1 {
			lastDepth = depth
			continue
		}

		if depth > lastDepth {
			increaseCount++
		} else if depth < lastDepth {
			decreaseCount++
		}

		lastDepth = depth
	}

	fmt.Println("---")
	fmt.Printf("Increases: %d\n", increaseCount)
	fmt.Printf("Decreases: %d\n", decreaseCount)

	return nil
}

func toInt(str string) (int, error) {
	i, err := strconv.ParseInt(str, 10, 32)
	if err != nil {
		return -1, err
	}

	return int(i), nil
}
