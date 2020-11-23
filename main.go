package main

import (
	"telegram_robot/logic"
)

func main() {
	service := logic.NewService()
	service.Run()
}
