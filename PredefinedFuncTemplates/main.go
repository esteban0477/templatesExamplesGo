package main

import (
	"log"
	"os"
	"text/template"
)

var ptrTemplateGlob *template.Template

func init() {
	ptrTemplateGlob = template.Must(template.New("").ParseGlob("*.gohtml"))
}

func main() {

	ptrFile, err := os.Create("index.html")
	if err != nil {
		log.Panicln(err)
	}
	defer ptrFile.Close()

	NameSlice := []string{"Michael", "Martha", "Jonas", "Ulrich"}
	IntSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	SomeStruct := struct {
		Names    []string
		Integers []int
	}{
		Names:    NameSlice,
		Integers: IntSlice,
	}

	err = ptrTemplateGlob.ExecuteTemplate(ptrFile, "template1.gohtml", SomeStruct)
	if err != nil {
		log.Panicln(err)
	}
}
