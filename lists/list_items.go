package lists

import (
	"net/http"

	"github.com/labstack/echo"
)

// ListItem represents a single item on a list.
type ListItem struct {
	Model
	Description string `json:"description"`
	Done        bool   `json:"done"`
	ListID      uint   `json:"list_id"`
}

// GetListItems returns the lists.
func GetListItems(env *HandlerEnvironment) echo.HandlerFunc {
	var listItems []ListItem
	return func(c echo.Context) error {
		env.Db.Find(&listItems)
		return c.JSON(http.StatusOK, listItems)
	}
}

// UpdateListItem updates a list item.
func UpdateListItem(env *HandlerEnvironment) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		var listItem, updatedListItem List
		env.Db.First(&listItem, id)
		if err := c.Bind(&updatedListItem); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Request is invalid."})
		}
		env.Db.Model(&listItem).Update(updatedListItem)
		return c.NoContent(http.StatusOK)
	}
}

// CreateListItem creates a new list item.
func CreateListItem(env *HandlerEnvironment) echo.HandlerFunc {
	return func(c echo.Context) error {
		var listItem ListItem
		if err := c.Bind(&listItem); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Request is invalid."})
		}
		env.Db.Create(&listItem)
		return c.NoContent(http.StatusCreated)
	}
}

// DeleteListItem deletes a list item completely.
func DeleteListItem(env *HandlerEnvironment) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		env.Db.Where("ID = ?", id).Delete(&ListItem{})
		return c.NoContent(http.StatusOK)
	}
}
