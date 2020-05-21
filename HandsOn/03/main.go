/* Create a data structure to pass to a template which contains information about
Restaurant Menu, including Breakfast, Lunch and Dinner items. */

package main

import (
	"html/template"
	"log"
	"os"
)

type optionMenu struct {
	TypeFood   string
	Beverage   string
	MainCourse string
	Starter    string
	Dessert    string
}

type restaurant struct {
	RestaurantName string
	FoodMenu       []optionMenu
}

var ptrTemplate *template.Template

func init() {
	ptrTemplate = template.Must(template.ParseFiles("template.gohtml"))
}

func main() {

	Restaurants := []restaurant{
		restaurant{
			RestaurantName: "Restaurant1",
			FoodMenu: []optionMenu{
				optionMenu{
					TypeFood:   "Breakfast",
					Beverage:   "Chocolate",
					MainCourse: "Fried eggs with pancackes",
					Starter:    "N/A",
					Dessert:    "N/A",
				},
				optionMenu{
					TypeFood:   "Lunch",
					Beverage:   "Coke",
					MainCourse: "Sirloin with rice",
					Starter:    "Tomato Soup",
					Dessert:    "Macadamia Cheesecake",
				},
				optionMenu{
					TypeFood:   "Dinner",
					Beverage:   "Coke",
					MainCourse: "Spaguetthi on tomato sauce",
					Starter:    `Cesars Salad`,
					Dessert:    "Vanilla Icecream",
				},
			},
		},
		restaurant{
			RestaurantName: "Restaurant1",
			FoodMenu: []optionMenu{
				optionMenu{
					TypeFood:   "Breakfast",
					Beverage:   "Chocolate",
					MainCourse: "Fried eggs with pancackes",
					Starter:    "N/A",
					Dessert:    "N/A",
				},
				optionMenu{
					TypeFood:   "Lunch",
					Beverage:   "Coke",
					MainCourse: "Sirloin with rice",
					Starter:    "Tomato Soup",
					Dessert:    "Macadamia Cheesecake",
				},
				optionMenu{
					TypeFood:   "Dinner",
					Beverage:   "Coke",
					MainCourse: "Spaguetthi on tomato sauce",
					Starter:    "Cesars Salad",
					Dessert:    "Vanilla Icecream",
				},
			},
		},
	}

	err := ptrTemplate.ExecuteTemplate(os.Stdout, "template.gohtml", Restaurants)
	if err != nil {
		log.Panicln(err)
	}

}
