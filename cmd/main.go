package main

import (
	"fmt"
	"lawyerinyou-backend/pkg/settings"
)

func init() {
	settings.Setup("./config/config.json")
}

func main() {
	fmt.Println("Oke")
}
