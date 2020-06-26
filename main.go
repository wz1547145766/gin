package main

import (
	"gin/sever"
)

func main() {

	r := sever.Routers()
	r.Run(":8000")
}
