package main

import (
	"fmt"

	"gitlab.com/alexnikita/gols/config"
)

func main() {
	value := config.Request("test", false)

	fmt.Printf("test value: %s\n", value)
}
