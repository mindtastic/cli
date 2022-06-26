package main

import (
	"fmt"

	"github.com/mindtastic/cli/openapi"
)

func main() {
	config := openapi.NewConfiguration()
	fmt.Println("Config: ", config)
}
