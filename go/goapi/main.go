package main

import (
	"fmt"
	"goapi/controllers"
	"goapi/models"
)

func main() {
	fmt.Println(models.Db)
	controllers.StartServer()
}
