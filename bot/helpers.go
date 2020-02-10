package main

import (
	"fmt"
	"os"
)

func generateSomeLogs() {
	fmt.Printf(os.Getenv("LOGS"))
}
