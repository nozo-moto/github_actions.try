package main

import (
	"fmt"

	"github.com/google/uuid"
)

func getUUID() (string, error) {
	return uuid.New().String(), nil
}

func hoge() string {
	return "hoge"
}

func main() {
	fmt.Println("vim-go")
}
