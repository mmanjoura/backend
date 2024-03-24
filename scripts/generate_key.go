package main

import (
	"fmt"

	"github.com/mmanjoura/niya-voyage-v2/backend-v2/pkg/auth"
)

func main() {
	fmt.Println(auth.GenerateRandomKey())
}
