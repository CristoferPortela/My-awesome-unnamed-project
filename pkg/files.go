package pkg

import (
	"fmt"
	"io/ioutil"
	"regexp"
)

func GetHtmlFrom(from string) []string {
	files, _ := ioutil.ReadDir(from)

	var htmls []string

	for _, file := range files {
		r, _ := regexp.Compile("html$")
		if r.MatchString(file.Name()) {
			htmls = append(htmls, fmt.Sprintf("%s/%s", from, file.Name()))
		}
	}
	return htmls
}

func GetAddons() []string {
	files, _ := ioutil.ReadDir("addon")
	var plugins []string

	for _, file := range files {
		if file.IsDir() {
			plugins = append(plugins, file.Name())
		}
	}
	return plugins
}
