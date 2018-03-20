package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/rufusroflpunch/wicky/lists"
)

func main() {
	db, err := gorm.Open("sqlite3", "lists.db")
	if err != nil {
		panic("Can't open a database connection!")
	}
	defer db.Close()
	setupDb(db)

	// env := &lists.HandlerEnvironment{Db: db.Set("gorm:auto_preload", true)}
	env := &lists.HandlerEnvironment{Db: db}
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.CORS())
	e.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
		fmt.Printf("%s\n", reqBody)
	}))

	api := e.Group("/api")

	// Lists
	api.GET("/lists", lists.GetLists(env))
	api.PUT("/lists/:id", lists.UpdateList(env))
	api.POST("/lists", lists.CreateList(env))
	api.DELETE("/lists/:id", lists.DeleteList(env))

	// ListItems
	api.GET("/list_items", lists.GetListItems(env))
	api.PUT("/list_items/:id", lists.UpdateListItem(env))
	api.PUT("/list_items/:id/toggle_done", lists.ToggleListItemDone(env))
	api.POST("/list_items", lists.CreateListItem(env))
	api.DELETE("/list_items/:id", lists.DeleteListItem(env))

	// Categories
	api.GET("/categories", lists.GetCategories(env))
	api.PUT("/categories/:id", lists.UpdateCategory(env))
	api.POST("/categories", lists.CreateCategory(env))
	api.DELETE("/categories/:id", lists.DeleteCategory(env))

	e.Logger.Fatal(e.Start(":8080"))
}

func setupDb(db *gorm.DB) {
	db.LogMode(true)
	db.AutoMigrate(&lists.List{})
	db.AutoMigrate(&lists.ListItem{})
	db.AutoMigrate(&lists.Category{})
}
