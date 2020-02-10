package main

import (
	"fmt"
	"os"
)

func generateSomeLogs() {
	fmt.Println(os.Getenv("LOGS"))
}
