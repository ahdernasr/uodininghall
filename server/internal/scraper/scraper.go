package main

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
)

type Meal struct {
	Breakfast []string
	Lunch     []string
	Dinner    []string
}

type Menu struct {
	Grill         Meal
	MindBodySoul  Meal
	PlantBase     Meal
	ServiceMinute Meal
	Trattoria     Meal
	WorldFlavours Meal
}

func main() {

	c := colly.NewCollector()

	menu := new(Menu)

	c.OnHTML("div#cat-1", func(e *colly.HTMLElement) {

		var mealStates [3]string = [3]string{"Breakfast", "Lunch", "Dinner"}
		stateIndex := -2

		e.DOM.Children().Each(func(index int, item *goquery.Selection) {
			node := item.Get(0)
			if node != nil {
				if node.Data == "h4" {
					stateIndex++
				} else {
					if stateIndex > -1 {
						switch mealStates[stateIndex] {
						case "Breakfast":
							menu.Grill.Breakfast = append(menu.Grill.Breakfast, getText(item))
						case "Lunch":
							menu.Grill.Lunch = append(menu.Grill.Lunch, getText(item))
						case "Dinner":
							menu.Grill.Dinner = append(menu.Grill.Dinner, getText(item))
						}
					}
				}
			}

		})

	})

	c.OnHTML("div#cat-2", func(e *colly.HTMLElement) {

		var mealStates [3]string = [3]string{"Breakfast", "Lunch", "Dinner"}
		stateIndex := -2

		e.DOM.Children().Each(func(index int, item *goquery.Selection) {
			node := item.Get(0)
			if node != nil {
				if node.Data == "h4" {
					stateIndex++
				} else {
					if stateIndex > -1 {
						switch mealStates[stateIndex] {
						case "Breakfast":
							menu.MindBodySoul.Breakfast = append(menu.MindBodySoul.Breakfast, getText(item))
						case "Lunch":
							menu.MindBodySoul.Lunch = append(menu.MindBodySoul.Lunch, getText(item))
						case "Dinner":
							menu.MindBodySoul.Dinner = append(menu.MindBodySoul.Dinner, getText(item))
						}
					}
				}
			}

		})

	})

	c.OnHTML("div#cat-3", func(e *colly.HTMLElement) {

		var mealStates [3]string = [3]string{"Breakfast", "Lunch", "Dinner"}
		stateIndex := -2

		e.DOM.Children().Each(func(index int, item *goquery.Selection) {
			node := item.Get(0)
			if node != nil {
				if node.Data == "h4" {
					stateIndex++
				} else {
					if stateIndex > -1 {
						switch mealStates[stateIndex] {
						case "Breakfast":
							menu.PlantBase.Breakfast = append(menu.PlantBase.Breakfast, getText(item))
						case "Lunch":
							menu.PlantBase.Lunch = append(menu.PlantBase.Lunch, getText(item))
						case "Dinner":
							menu.PlantBase.Dinner = append(menu.PlantBase.Dinner, getText(item))
						}
					}
				}
			}

		})

	})

	c.OnHTML("div#cat-4", func(e *colly.HTMLElement) {

		var mealStates [3]string = [3]string{"Breakfast", "Lunch", "Dinner"}
		stateIndex := -2

		e.DOM.Children().Each(func(index int, item *goquery.Selection) {
			node := item.Get(0)
			if node != nil {
				if node.Data == "h4" {
					stateIndex++
				} else {
					if stateIndex > -1 {
						switch mealStates[stateIndex] {
						case "Breakfast":
							menu.ServiceMinute.Breakfast = append(menu.ServiceMinute.Breakfast, getText(item))
						case "Lunch":
							menu.ServiceMinute.Lunch = append(menu.ServiceMinute.Lunch, getText(item))
						case "Dinner":
							menu.ServiceMinute.Dinner = append(menu.ServiceMinute.Dinner, getText(item))
						}
					}
				}
			}

		})

	})

	c.OnHTML("div#cat-5", func(e *colly.HTMLElement) {

		var mealStates [3]string = [3]string{"Breakfast", "Lunch", "Dinner"}
		stateIndex := -2

		e.DOM.Children().Each(func(index int, item *goquery.Selection) {
			node := item.Get(0)
			if node != nil {
				if node.Data == "h4" {
					stateIndex++
				} else {
					if stateIndex > -1 {
						switch mealStates[stateIndex] {
						case "Breakfast":
							menu.Trattoria.Breakfast = append(menu.Trattoria.Breakfast, getText(item))
						case "Lunch":
							menu.Trattoria.Lunch = append(menu.Trattoria.Lunch, getText(item))
						case "Dinner":
							menu.Trattoria.Dinner = append(menu.Trattoria.Dinner, getText(item))
						}
					}
				}
			}

		})

	})

	c.OnHTML("div#cat-6", func(e *colly.HTMLElement) {

		var mealStates [3]string = [3]string{"Breakfast", "Lunch", "Dinner"}
		stateIndex := -2

		e.DOM.Children().Each(func(index int, item *goquery.Selection) {
			node := item.Get(0)
			if node != nil {
				if node.Data == "h4" {
					stateIndex++
				} else {
					if stateIndex > -1 {
						switch mealStates[stateIndex] {
						case "Breakfast":
							menu.WorldFlavours.Breakfast = append(menu.WorldFlavours.Breakfast, getText(item))
						case "Lunch":
							menu.WorldFlavours.Lunch = append(menu.WorldFlavours.Lunch, getText(item))
						case "Dinner":
							menu.WorldFlavours.Dinner = append(menu.WorldFlavours.Dinner, getText(item))
						}
					}
				}
			}

		})

	})

	c.Visit("https://www.uottawa.ca/campus-life/eat-campus/where-eat/dining-hall/menu")
	c.Wait()
	fmt.Println(menu)
}

func getText(s *goquery.Selection) string {
	cloned := s.Clone()
	cloned.Find("i").Remove()
	cloned.Find("span").Remove()
	text := strings.TrimSpace(cloned.Text())
	return text
}
