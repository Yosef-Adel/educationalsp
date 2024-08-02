package main

import (
	"bufio"
	"educationalsp/rpc"
	"fmt"
	"os"
)

func main() {
	fmt.Println("hi")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(rpc.Split)

	for scanner.Scan() {
		msg := scanner.Text()
		handelMessage(msg)
	}
}

func handelMessage(_ any) {}
