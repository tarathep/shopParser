package main

import (
	"encoding/csv"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

type ItemShop struct {
	ID      string `xml:"ID"`
	Points  int    `xml:"Points"`
	Comment string `xml:"Comment"`
}
type Category struct {
	ID    int        `xml:"ID"`
	Name  string     `xml:"Name"`
	Items []ItemShop `xml:"Item"`
}

type ECOShop struct {
	Category []Category `xml:"Category"`
}

//----

type Item struct {
	ID       string
	PictID   string
	IconID   string
	Name     string
	Category string
}

var itemCSV = "item.csv"

func main() {
	var categorys []Category
	getCategorys := GetCategory()
	items := GetItemFromCSV(itemCSV)

	for ID, cate := range getCategorys {
		ID = ID + 1
		var itemShops []ItemShop
		for _, item := range items {
			if item.Category == cate {
				itemShop := ItemShop{
					ID:      item.ID,
					Points:  0,
					Comment: "Free",
				}
				itemShops = append(itemShops, itemShop)

			}
		}

		categorys = append(categorys,
			Category{
				ID:    ID,
				Name:  cate,
				Items: itemShops,
			})
		fmt.Println(ID, cate)
	}

	ecoshop := ECOShop{}
	ecoshop.Category = categorys

	file, _ := xml.MarshalIndent(ecoshop, "", " ")
	_ = ioutil.WriteFile("ECOShop.xml", file, 0644)
}

func GetCategory() []string {
	var types []string
	var result []string
	types = append(types, "")
	items := GetItemFromCSV(itemCSV)
	for _, item := range items {
		if !isMatch(types, item.Category) {
			types = append(types, item.Category)
			result = append(result, item.Category)
		}
	}
	return result
}

func isMatch(types []string, search string) bool {
	for _, t := range types {
		if t == search {
			return true
		}
	}
	return false
}

func GetItemFromCSV(filePath string) []Item {
	var items []Item

	//Load a csv file.
	f, _ := os.Open(filePath)

	//Create a new reader.
	r := csv.NewReader(f)
	for {
		record, err := r.Read()
		// Stop at EOF.
		if err == io.EOF {
			break
		}

		if err != nil {
			panic(err)

		}

		items = append(items,
			Item{
				ID:       record[0],
				PictID:   record[1],
				IconID:   record[2],
				Name:     record[3],
				Category: record[4],
			})
	}
	return items
}
