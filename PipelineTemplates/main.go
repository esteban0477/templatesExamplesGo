package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
	"time"
)

var ptrTemplateGlob *template.Template

var funcMap = template.FuncMap{
	"ConvertTimeToUnix":    formatUnix,
	"ConvertTimeToKitchen": formatKitchen,
}

func init() {
	ptrTemplateGlob = template.Must(template.New("").Funcs(funcMap).ParseGlob("*.gohtml"))
}

func main() {

	ptrFile, err := os.Create("index.html")
	if err != nil {
		log.Panicln(err)
	}
	defer ptrFile.Close()

	fmt.Println(time.Now())
	err = ptrTemplateGlob.ExecuteTemplate(ptrFile, "template1.gohtml", time.Now())
	if err != nil {
		log.Panicln(err)
	}
}

func formatUnix(t time.Time) string {
	return t.Format(time.UnixDate)
}

func formatKitchen(t time.Time) string {
	return t.Format(time.Kitchen)
}
