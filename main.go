package main

import (
	"fmt"
	"github.com/abhub23/go-utility/rim"
)

func main() {
	resp, err := rim.Flames("abdullah", "rimsha")
	if err != nil {
		fmt.Printf("error occured %v\n", err)
	}
	fmt.Println(resp)
}
