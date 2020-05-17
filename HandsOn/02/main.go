/* Create a data structure to pass to a template which contains information about
hotels including Name, Address, Town, Zip, Region. region can be:
Southern, Central, Northern, can hold an unlimited number of hotels" */

package main

import (
	"html/template"
	"log"
	"os"
)

type hotel struct {
	Name    string
	Address string
	Town    string
	ZIP     string
}

type regions struct {
	RegName string
	Hotels  []hotel
}

var ptrTemplate *template.Template

func init() {
	ptrTemplate = template.Must(template.ParseFiles("template.gohtml"))
}

func main() {

	hotelPerRegion := []regions{
		regions{
			RegName: "North",
			Hotels: []hotel{
				hotel{
					Name:    "North Hotel 1",
					Address: "North Hotel 1 Address",
					Town:    "North Hotel 1 Town",
					ZIP:     "North Hotel 1 ZIP",
				},
				hotel{
					Name:    "North Hotel 2",
					Address: "North Hotel 2 Address",
					Town:    "North Hotel 2 Town",
					ZIP:     "North Hotel 2 ZIP",
				},
				hotel{
					Name:    "North Hotel 3",
					Address: "North Hotel 3 Address",
					Town:    "North Hotel 3 Town",
					ZIP:     "North Hotel 3 ZIP",
				},
			},
		},
		regions{
			RegName: "South",
			Hotels: []hotel{
				hotel{
					Name:    "South Hotel 1",
					Address: "South Hotel 1 Address",
					Town:    "South Hotel 1 Town",
					ZIP:     "South Hotel 1 ZIP",
				},
				hotel{
					Name:    "South Hotel 2",
					Address: "South Hotel 2 Address",
					Town:    "South Hotel 2 Town",
					ZIP:     "South Hotel 2 ZIP",
				},
				hotel{
					Name:    "South Hotel 3",
					Address: "South Hotel 3 Address",
					Town:    "South Hotel 3 Town",
					ZIP:     "South Hotel 3 ZIP",
				},
			},
		},
	}

	err := ptrTemplate.ExecuteTemplate(os.Stdout, "template.gohtml", hotelPerRegion)
	if err != nil {
		log.Panicln(err)
	}

}
