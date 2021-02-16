package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"os"
)

func main() {
	scan := bufio.NewScanner(os.Stdin)
	fmt.Println("-- ASCII decoder --")
	scan.Scan()
	input := scan.Text()
	byteRes, err := hex.DecodeString(input)
	if err != nil {
		panic(err)
	}
	fmt.Println("-- Result --")
	fmt.Println(string(byteRes))
}
