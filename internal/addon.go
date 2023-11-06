package internal

import (
	"awesomeProject/pkg"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"plugin"
)

type RouteFunc func(c *echo.Echo)

type Addon struct {
	Routes []RouteFunc
}

var (
	Routes = []RouteFunc{}
)

type Export struct {
	Name string
	//Routes []RouteFunc
}

func (e *Export) Named() string {
	return "oi"
}

func ConfigureAddons() Addon {

	routes := Routes

	for _, addon := range pkg.GetAddons() {
		fmt.Println(addon)
		path := fmt.Sprintf("addon/%s/main.so", addon)
		plug, err := plugin.Open(path)

		if err != nil {
			fmt.Println(err)
			fmt.Println(fmt.Sprintf("Error when getting plugin %s", addon))
			break
		}
		symbol, err := plug.Lookup("Exporting")
		if err != nil {
			fmt.Println(err)
			log.Fatal(fmt.Sprintf("Error when looking up the plugin %s", addon))
		}
		var exported func() interface{}
		var ok bool

		exported, ok = symbol.(func() interface{})
		if !ok {
			log.Fatal(fmt.Sprintf("Error on getting exported values for %s", addon))
		}
		a := exported()
		stuf, err := json.Marshal(a)

		x := map[string]string{}

		err = json.Unmarshal([]byte(stuf), &x)

		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(string(stuf), x["Name"])

	}
	return Addon{
		Routes: routes,
	}

}
