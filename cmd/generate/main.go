package main

import (
	"fmt"
	"os"
	"strings"

	"nsqd.cc/dich-markov/internal/generate"
)

func main() {
	seed := os.Args[1:]
	chain := generate.LoadChain("./model.json")
	fmt.Printf("Seed is: %s\n", strings.Join(seed, " "))

	data := generate.GenerateAdvanced(chain, strings.ToLower(strings.Join(seed, " ")))
	fmt.Printf("%s\n", data)
}
