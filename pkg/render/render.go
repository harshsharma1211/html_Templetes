package render

import (
	"fmt"
	"html/template"
	"net/http"
)

func Rendertemp(w http.ResponseWriter, html string) {

	parsedtemplates, _ := template.ParseFiles("./templates/" + html)
	err := parsedtemplates.Execute(w, nil)
	if err != nil {

		fmt.Println("error passing template ", err)
		return
	}

}
