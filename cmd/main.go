package main

import (
	"fmt"
	"lawyerinyou-backend/pkg/database"
	"lawyerinyou-backend/pkg/logging"
	"lawyerinyou-backend/pkg/settings"
)

func init() {
	settings.Setup("./config/config.json")
	database.Setup()
	logging.Setup()
}

func main() {
	fmt.Println("Oke")
}
