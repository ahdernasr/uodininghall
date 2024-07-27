package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
)

type Dish struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Allergies   string `json:"allergies"`
}

type Meal struct {
	Breakfast []Dish `json:"breakfast"`
	Lunch     []Dish `json:"lunch"`
	Dinner    []Dish `json:"dinner"`
	Other     []Dish `json:"other"`
}

// Define the Menu struct
type Menu struct {
	Grill         Meal `json:"grill"`
	MindBodySoul  Meal `json:"mindBodySoul"`
	PlantBase     Meal `json:"plantBase"`
	ServiceMinute Meal `json:"serviceMinute"`
	Trattoria     Meal `json:"trattoria"`
	WorldFlavours Meal `json:"worldFlavours"`
}

func main() {

	c := colly.NewCollector()

	menu := new(Menu)

	// Grill section - Breakfast, Lunch, Dinner

	c.OnHTML("div#cat-1", func(e *colly.HTMLElement) {

		mealStates := map[string]string{
			"Breakfast": "Breakfast",
			"Lunch":     "Lunch",
			"Dinner":    "Dinner",
			"Other":     "Other",
		}
		stateIndex := ""

		e.DOM.Children().Each(func(index int, item *goquery.Selection) {
			nameNode := item.Get(0)

			if nameNode != nil {
				if nameNode.Data == "h4" {
					switch getText(item) {
					case "Breakfast":
						stateIndex = "Breakfast"
					case "Lunch":
						stateIndex = "Lunch"
					case "Dinner":
						stateIndex = "Dinner"
					}

				} else if nameNode.Data == "p" {
					switch mealStates[stateIndex] {
					case "Breakfast":
						menu.Grill.Breakfast[len(menu.Grill.Breakfast)-1].Description = getText(item)
					case "Lunch":
						menu.Grill.Lunch[len(menu.Grill.Lunch)-1].Description = getText(item)
					case "Dinner":
						menu.Grill.Dinner[len(menu.Grill.Dinner)-1].Description = getText(item)
					default:
						menu.Grill.Other[len(menu.Grill.Other)-1].Description = getText(item)
					}
				} else {
					switch mealStates[stateIndex] {
					case "Breakfast":
						menu.Grill.Breakfast = append(menu.Grill.Breakfast, Dish{
							Name:        getText(item),
							Description: "",
							Allergies:   getTitle(item),
						})
					case "Lunch":
						menu.Grill.Lunch = append(menu.Grill.Lunch, Dish{
							Name:        getText(item),
							Description: "",
							Allergies:   getTitle(item),
						})
					case "Dinner":
						menu.Grill.Dinner = append(menu.Grill.Dinner, Dish{
							Name:        getText(item),
							Description: "",
							Allergies:   getTitle(item),
						})
					default:
						menu.Grill.Other = append(menu.Grill.Other, Dish{
							Name:        getText(item),
							Description: "",
							Allergies:   getTitle(item),
						})
					}
				}
			}

		})

	})

	// Mind Body Soul Section - Breakfast, Lunch, Dinner

	c.OnHTML("div#cat-2", func(e *colly.HTMLElement) {

		mealStates := map[string]string{
			"Breakfast": "Breakfast",
			"Lunch":     "Lunch",
			"Dinner":    "Dinner",
			"Other":     "Other",
		}
		stateIndex := ""

		e.DOM.Children().Each(func(index int, item *goquery.Selection) {
			nameNode := item.Get(0)

			if nameNode != nil {
				if nameNode.Data == "h4" {
					switch getText(item) {
					case "Breakfast":
						stateIndex = "Breakfast"
					case "Lunch":
						stateIndex = "Lunch"
					case "Dinner":
						stateIndex = "Dinner"
					}

				} else if nameNode.Data == "p" {
					switch mealStates[stateIndex] {
					case "Breakfast":
						menu.MindBodySoul.Breakfast[len(menu.MindBodySoul.Breakfast)-1].Description = getText(item)
					case "Lunch":
						menu.MindBodySoul.Lunch[len(menu.MindBodySoul.Lunch)-1].Description = getText(item)
					case "Dinner":
						menu.MindBodySoul.Dinner[len(menu.MindBodySoul.Dinner)-1].Description = getText(item)
					default:
						menu.MindBodySoul.Other[len(menu.MindBodySoul.Other)-1].Description = getText(item)
					}
				} else {
					switch mealStates[stateIndex] {
					case "Breakfast":
						menu.MindBodySoul.Breakfast = append(menu.MindBodySoul.Breakfast, Dish{
							Name:        getText(item),
							Description: "",
							Allergies:   getTitle(item),
						})
					case "Lunch":
						menu.MindBodySoul.Lunch = append(menu.MindBodySoul.Lunch, Dish{
							Name:        getText(item),
							Description: "",
							Allergies:   getTitle(item),
						})
					case "Dinner":
						menu.MindBodySoul.Dinner = append(menu.MindBodySoul.Dinner, Dish{
							Name:        getText(item),
							Description: "",
							Allergies:   getTitle(item),
						})
					default:
						menu.MindBodySoul.Other = append(menu.MindBodySoul.Other, Dish{
							Name:        getText(item),
							Description: "",
							Allergies:   getTitle(item),
						})
					}
				}
			}

		})

	})

	// Plant Base - List of items

	c.OnHTML("div#cat-3", func(e *colly.HTMLElement) {

		mealStates := map[string]string{
			"Breakfast": "Breakfast",
			"Lunch":     "Lunch",
			"Dinner":    "Dinner",
			"Other":     "Other",
		}
		stateIndex := ""

		e.DOM.Children().Each(func(index int, item *goquery.Selection) {
			nameNode := item.Get(0)

			if nameNode != nil {
				if nameNode.Data == "h4" {
					switch getText(item) {
					case "Breakfast":
						stateIndex = "Breakfast"
					case "Lunch":
						stateIndex = "Lunch"
					case "Dinner":
						stateIndex = "Dinner"
					}

				} else if nameNode.Data == "p" {
					switch mealStates[stateIndex] {
					case "Breakfast":
						menu.PlantBase.Breakfast[len(menu.PlantBase.Breakfast)-1].Description = getText(item)
					case "Lunch":
						menu.PlantBase.Lunch[len(menu.PlantBase.Lunch)-1].Description = getText(item)
					case "Dinner":
						menu.PlantBase.Dinner[len(menu.PlantBase.Dinner)-1].Description = getText(item)
					default:
						menu.PlantBase.Other[len(menu.PlantBase.Other)-1].Description = getText(item)
					}
				} else {
					switch mealStates[stateIndex] {
					case "Breakfast":
						menu.PlantBase.Breakfast = append(menu.PlantBase.Breakfast, Dish{
							Name:        getText(item),
							Description: "",
							Allergies:   getTitle(item),
						})
					case "Lunch":
						menu.PlantBase.Lunch = append(menu.PlantBase.Lunch, Dish{
							Name:        getText(item),
							Description: "",
							Allergies:   getTitle(item),
						})
					case "Dinner":
						menu.PlantBase.Dinner = append(menu.PlantBase.Dinner, Dish{
							Name:        getText(item),
							Description: "",
							Allergies:   getTitle(item),
						})
					default:
						menu.PlantBase.Other = append(menu.PlantBase.Other, Dish{
							Name:        getText(item),
							Description: "",
							Allergies:   getTitle(item),
						})
					}
				}
			}

		})

	})

	// Service Minute - Item, Breakfast, Lunch, Dinner

	c.OnHTML("div#cat-4", func(e *colly.HTMLElement) {

		mealStates := map[string]string{
			"Breakfast": "Breakfast",
			"Lunch":     "Lunch",
			"Dinner":    "Dinner",
			"Other":     "Other",
		}
		stateIndex := ""

		e.DOM.Children().Each(func(index int, item *goquery.Selection) {
			node := item.Get(0)
			if node != nil {
				if node.Data == "h4" {
					switch getText(item) {
					case "Breakfast":
						stateIndex = "Breakfast"
					case "Lunch":
						stateIndex = "Lunch"
					case "Dinner":
						stateIndex = "Dinner"
					}

				} else if node.Data == "p" {
					switch mealStates[stateIndex] {
					case "Breakfast":
						menu.ServiceMinute.Breakfast[len(menu.ServiceMinute.Breakfast)-1].Description = getText(item)
					case "Lunch":
						menu.ServiceMinute.Lunch[len(menu.ServiceMinute.Lunch)-1].Description = getText(item)
					case "Dinner":
						menu.ServiceMinute.Dinner[len(menu.ServiceMinute.Dinner)-1].Description = getText(item)
					default:
						menu.ServiceMinute.Other[len(menu.ServiceMinute.Other)-1].Description = getText(item)
					}
				} else {
					switch mealStates[stateIndex] {
					case "Breakfast":
						menu.ServiceMinute.Breakfast = append(menu.ServiceMinute.Breakfast, Dish{
							Name:        getText(item),
							Description: "",
							Allergies:   getTitle(item),
						})
					case "Lunch":
						menu.ServiceMinute.Lunch = append(menu.ServiceMinute.Lunch, Dish{
							Name:        getText(item),
							Description: "",
							Allergies:   getTitle(item),
						})
					case "Dinner":
						menu.ServiceMinute.Dinner = append(menu.ServiceMinute.Dinner, Dish{
							Name:        getText(item),
							Description: "",
							Allergies:   getTitle(item),
						})
					default:
						menu.ServiceMinute.Other = append(menu.ServiceMinute.Other, Dish{
							Name:        getText(item),
							Description: "",
							Allergies:   getTitle(item),
						})
					}
				}
			}

		})

	})

	// Trattoria - Breakfast, Lunch, Dinner

	c.OnHTML("div#cat-6", func(e *colly.HTMLElement) {

		mealStates := map[string]string{
			"Breakfast": "Breakfast",
			"Lunch":     "Lunch",
			"Dinner":    "Dinner",
			"Other":     "Other",
		}
		stateIndex := ""

		e.DOM.Children().Each(func(index int, item *goquery.Selection) {
			node := item.Get(0)
			if node != nil {
				if node.Data == "h4" {
					switch getText(item) {
					case "Breakfast":
						stateIndex = "Breakfast"
					case "Lunch":
						stateIndex = "Lunch"
					case "Dinner":
						stateIndex = "Dinner"
					}

				} else if node.Data == "p" {
					switch mealStates[stateIndex] {
					case "Breakfast":
						menu.Trattoria.Breakfast[len(menu.Trattoria.Breakfast)-1].Description = getText(item)
					case "Lunch":
						menu.Trattoria.Lunch[len(menu.Trattoria.Lunch)-1].Description = getText(item)
					case "Dinner":
						menu.Trattoria.Dinner[len(menu.Trattoria.Dinner)-1].Description = getText(item)
					default:
						menu.Trattoria.Other[len(menu.Trattoria.Other)-1].Description = getText(item)
					}
				} else {
					switch mealStates[stateIndex] {
					case "Breakfast":
						menu.Trattoria.Breakfast = append(menu.Trattoria.Breakfast, Dish{
							Name:        getText(item),
							Description: "",
							Allergies:   getTitle(item),
						})
					case "Lunch":
						menu.Trattoria.Lunch = append(menu.Trattoria.Lunch, Dish{
							Name:        getText(item),
							Description: "",
							Allergies:   getTitle(item),
						})
					case "Dinner":
						menu.Trattoria.Dinner = append(menu.Trattoria.Dinner, Dish{
							Name:        getText(item),
							Description: "",
							Allergies:   getTitle(item),
						})
					default:
						menu.Trattoria.Other = append(menu.Trattoria.Other, Dish{
							Name:        getText(item),
							Description: "",
							Allergies:   getTitle(item),
						})
					}
				}
			}

		})

	})

	// World Flavors - Breakfast, Lunch, Dinner

	c.OnHTML("div#cat-6", func(e *colly.HTMLElement) {

		mealStates := map[string]string{
			"Breakfast": "Breakfast",
			"Lunch":     "Lunch",
			"Dinner":    "Dinner",
			"Other":     "Other",
		}
		stateIndex := ""

		e.DOM.Children().Each(func(index int, item *goquery.Selection) {
			node := item.Get(0)
			if node != nil {
				if node.Data == "h4" {
					switch getText(item) {
					case "Breakfast":
						stateIndex = "Breakfast"
					case "Lunch":
						stateIndex = "Lunch"
					case "Dinner":
						stateIndex = "Dinner"
					}

				} else if node.Data == "p" {
					switch mealStates[stateIndex] {
					case "Breakfast":
						menu.WorldFlavours.Breakfast[len(menu.WorldFlavours.Breakfast)-1].Description = getText(item)
					case "Lunch":
						menu.WorldFlavours.Lunch[len(menu.WorldFlavours.Lunch)-1].Description = getText(item)
					case "Dinner":
						menu.WorldFlavours.Dinner[len(menu.WorldFlavours.Dinner)-1].Description = getText(item)
					default:
						menu.WorldFlavours.Other[len(menu.WorldFlavours.Other)-1].Description = getText(item)
					}
				} else {
					switch mealStates[stateIndex] {
					case "Breakfast":
						menu.WorldFlavours.Breakfast = append(menu.WorldFlavours.Breakfast, Dish{
							Name:        getText(item),
							Description: "",
							Allergies:   getTitle(item),
						})
					case "Lunch":
						menu.WorldFlavours.Lunch = append(menu.WorldFlavours.Lunch, Dish{
							Name:        getText(item),
							Description: "",
							Allergies:   getTitle(item),
						})
					case "Dinner":
						menu.WorldFlavours.Dinner = append(menu.WorldFlavours.Dinner, Dish{
							Name:        getText(item),
							Description: "",
							Allergies:   getTitle(item),
						})
					default:
						menu.WorldFlavours.Other = append(menu.WorldFlavours.Other, Dish{
							Name:        getText(item),
							Description: "",
							Allergies:   getTitle(item),
						})
					}
				}
			}

		})

	})

	// c.Visit("https://www.uottawa.ca/campus-life/eat-campus/where-eat/dining-hall/menu")
	c.Visit("https://web.archive.org/web/20230128143243/https://www.uottawa.ca/campus-life/eat-campus/where-eat/dining-hall/menu")
	c.Wait()

	jsonData, err := json.MarshalIndent(menu, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}
	fmt.Println(string(jsonData))
	// fmt.Println(menu)
}

func getText(s *goquery.Selection) string {
	cloned := s.Clone()
	cloned.Find("i").Remove()
	cloned.Find("span").Remove()
	text := strings.TrimSpace(cloned.Text())
	return text
}

func getTitle(s *goquery.Selection) string {
	img := s.Find("span img")
	title, exists := img.Attr("title")
	if exists {
		return title
	}
	return "None"
}
