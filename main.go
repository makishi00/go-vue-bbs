package main

import "github.com/makishi00/go-test/router"

func main() {
	r := router.GetRouter()
	r.Run(":8000")
}
