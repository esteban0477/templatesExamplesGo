package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
	"time"
)

type person struct {
	FirstName string
	LastName  string
	Age       int
}

var ptrTemplateGlob *template.Template

var funcMap = template.FuncMap{
	"ConvertToUpper": toupper,
	"ConvertToLower": tolower,
}

func init() {
	// ptrTemplateGlob = template.Must(template.ParseGlob("*.gohtml"))
	ptrTemplateGlob = template.Must(template.New("").Funcs(funcMap).ParseGlob("*.gohtml"))
}

func main() {

	ptrFile, err := os.Create("index.html")
	if err != nil {
		log.Panicln(err)
	}
	defer ptrFile.Close()

	people := []person{
		person{
			FirstName: "Michael",
			LastName:  "Kahnwald",
			Age:       28,
		},
		person{
			FirstName: "Jonas",
			LastName:  "Kahnwald",
			Age:       49,
		},
	}

	fmt.Println(time.Now())
	err = ptrTemplateGlob.ExecuteTemplate(ptrFile, "template1.gohtml", people)
	if err != nil {
		log.Panicln(err)
	}
}

func toupper(s string) string {
	sliceBytes := []byte(s)
	for index, value := range sliceBytes {
		if value >= 97 && value <= 122 {
			sliceBytes[index] = value - 32
		}
	}
	return string(sliceBytes)
}

func tolower(s string) string {
	sliceBytes := []byte(s)
	for index, value := range sliceBytes {
		if value >= 65 && value <= 90 {
			sliceBytes[index] = value + 32
		}
	}
	return string(sliceBytes)
}
