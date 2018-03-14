package lists

import (
	"net/http"

	"github.com/labstack/echo"
)

// Category is a category is a lists.
type Category struct {
	Model
	Name  string `json:"name"`
	Lists []List `json:"lists"`
}

// GetCategories returns the categories.
func GetCategories(env *HandlerEnvironment) echo.HandlerFunc {
	var categories []Category
	return func(c echo.Context) error {
		env.Db.Find(&categories)
		return c.JSON(http.StatusOK, categories)
	}
}

// UpdateCategory updates a category.
func UpdateCategory(env *HandlerEnvironment) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		var category, updatedCategory Category
		env.Db.First(&category, id)
		if err := c.Bind(&updatedCategory); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Request is invalid."})
		}
		env.Db.Model(&category).Update(updatedCategory)
		return c.NoContent(http.StatusOK)
	}
}

// CreateCategory creates a new category.
func CreateCategory(env *HandlerEnvironment) echo.HandlerFunc {
	return func(c echo.Context) error {
		var category Category
		if err := c.Bind(&category); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Request is invalid."})
		}
		env.Db.Create(&category)
		return c.NoContent(http.StatusCreated)
	}
}

// DeleteCategory deletes a category completely.
func DeleteCategory(env *HandlerEnvironment) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		env.Db.Where("ID = ?", id).Delete(&Category{})
		return c.NoContent(http.StatusOK)
	}
}
