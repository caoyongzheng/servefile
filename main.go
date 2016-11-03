package main

import (
	"flag"

	"github.com/go-martini/martini"
)

func main() {
	var port = flag.String("port", "8080", "port to listen")
	var content = flag.String("content", "", "file to server")
	flag.Parse()
	m := martini.Classic()
	m.Use(martini.Static(
		*content,
	))
	m.RunOnAddr(":" + *port)
}
