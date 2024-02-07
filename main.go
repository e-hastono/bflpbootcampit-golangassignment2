package main

import (
	"github.com/e-hastono/bflpbootcamptit-golangassignment2/routers"
)

func main() {

	routers.StartServer().Run(":8080")
}
