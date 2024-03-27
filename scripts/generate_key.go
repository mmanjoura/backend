package main

import (
	"fmt"

	"github.com/mmanjoura/niya-voyage/backend/pkg/auth"
)

func main() {
	fmt.Println(auth.GenerateRandomKey())
}
