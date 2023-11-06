package main

import "fmt"

//import "github.com/labstack/echo/v4"

func main() {
	fmt.Println(Exporting())
}

type Export struct {
	Name string
	//Routes []RouteFunc
}

func Exporting() interface{} {
	return &Export{}
}
