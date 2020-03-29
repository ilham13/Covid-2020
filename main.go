package main

import "github.com/ilham13/Covid-2020/app"

func main() {
	a := app.Routes{}
	a.Initialize()
	a.Run(":8080")
}
