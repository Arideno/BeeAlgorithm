package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	algorithm := NewBeeAlgorithm(27, 3)
	algorithm.Start()
	fmt.Println(time.Since(now))
}
