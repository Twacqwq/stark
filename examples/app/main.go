package main

import (
	"fmt"
	"log"

	"github.com/FarmerChillax/stark"
	"github.com/FarmerChillax/stark/app"
)

func main() {
	appBuilder, err := app.NewBuilder(&stark.Application{
		Name: "app-demo",
		Host: "0.0.0.0",
		Port: 6000,
	})
	if err != nil {
		log.Fatalln("app.New err: ", err)
	}

	fmt.Printf("%+v\n", appBuilder)
}
