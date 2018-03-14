package lists

import (
	"net/http"

	"github.com/labstack/echo"
)

// List represents a list.
type List struct {
	Model
	Name       string     `json:"name"`
	Archived   bool       `json:"archived"`
	ListItems  []ListItem `json:"list_items"`
	CategoryID uint       `json:"category_id"`
}

// GetLists returns the lists.
func GetLists(env *HandlerEnvironment) echo.HandlerFunc {
	var lists []List
	return func(c echo.Context) error {
		env.Db.Find(&lists)
		return c.JSON(http.StatusOK, lists)
	}
}

// UpdateList updates a list.
func UpdateList(env *HandlerEnvironment) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		var list, updatedList List
		env.Db.First(&list, id)
		if err := c.Bind(&updatedList); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Request is invalid."})
		}
		env.Db.Model(&list).Update(updatedList)
		return c.NoContent(http.StatusOK)
	}
}

// CreateList creates a new list.
func CreateList(env *HandlerEnvironment) echo.HandlerFunc {
	return func(c echo.Context) error {
		var list List
		if err := c.Bind(&list); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Request is invalid."})
		}
		env.Db.Create(&list)
		return c.NoContent(http.StatusCreated)
	}
}

// DeleteList deletes a list completely.
func DeleteList(env *HandlerEnvironment) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		env.Db.Where("ID = ?", id).Delete(&List{})
		return c.NoContent(http.StatusOK)
	}
}
