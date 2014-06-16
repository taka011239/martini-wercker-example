package main

import "github.com/go-martini/martini"

func main() {
	m := martini.Classic()
	m.Get("/", home)
	m.Run()
}

func home(params martini.Params) (int, string) {
	return 200, "hello"
}
