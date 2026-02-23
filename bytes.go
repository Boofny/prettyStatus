package main

import (
	"fmt"
	"time"
)

func main() {
	m := time.Now().Local().Local().Month()
	fmt.Println(byte(m))
	fmt.Println(uint8(m))
}
