package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/rufusroflpunch/leest/lists"
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

	db.Create(&lists.Category{Name: "Personal"})
	db.Create(&lists.Category{Name: "Work"})
	db.Create(&lists.Category{Name: "Shopping"})
	db.Create(&lists.Category{Name: "Bucket Lists"})
}
