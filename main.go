package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	b, error := reader.ReadByte()

	for nil == error {
		fmt.Printf(string(b))

		b, error = reader.ReadByte()
	}

	fmt.Println(error)
}
