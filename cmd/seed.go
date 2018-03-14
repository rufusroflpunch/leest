package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/rufusroflpunch/wicky/lists"
)

func main() {
	db, _ := gorm.Open("sqlite3", "lists.db")
	defer db.Close()

	db.Exec("DELETE FROM lists;")
	db.Exec("DELETE FROM list_items;")
	db.Exec("DELETE FROM categories;")

	db.AutoMigrate(&lists.List{})
	db.AutoMigrate(&lists.ListItem{})
	db.AutoMigrate(&lists.Category{})

	db.Create(&lists.Category{
		Name: "General",
		Lists: []lists.List{
			lists.List{
				Name:     "A List",
				Archived: false,
				ListItems: []lists.ListItem{
					lists.ListItem{Description: "Item #1", Done: false},
				},
			},
			lists.List{
				Name:     "Another List",
				Archived: false,
				ListItems: []lists.ListItem{
					lists.ListItem{Description: "Item #1", Done: false},
					lists.ListItem{Description: "Item #2", Done: false},
					lists.ListItem{Description: "Item #3", Done: false},
				},
			},
		},
	})
}
