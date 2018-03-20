package main

import "github.com/makishi00/go-vue-bbs/router"

func main() {
	r := router.GetRouter()
	r.Run(":8000")
}
