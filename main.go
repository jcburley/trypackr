package main

import (
	"fmt"
	"log"

	"github.com/gobuffalo/packr/v2"
)

func main() {
	// set up a new box by giving it a name and an optional (relative) path to a folder on disk:
	box := packr.New("My Box", "./templates")

	s, err := box.FindString("admin/index.html")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(s)
}
