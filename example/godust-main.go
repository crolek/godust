package main

import (
	"fmt"

	"github.com/crolek/godust"
)

func main() {
	value := godust.RenderDustjs("../assets/test-template.js", "test", `{"name": "Chuck"}`)

	fmt.Println("value: ")
	fmt.Println(value)

}
