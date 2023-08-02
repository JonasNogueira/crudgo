package main

import (
	"fmt"

	"github.com/JonasNogueira/crudgo/config"
	"github.com/JonasNogueira/crudgo/router"
)

func main() {

	err := config.Init()
	if err != nil {
		fmt.Println(err)
		return
	}

	router.Initialize()
}
