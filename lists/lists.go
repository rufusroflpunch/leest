package lists

import (
	"net/http"

	"github.com/labstack/echo"
)

// GetLists returns the lists.
func GetLists(c echo.Context) error {
	return c.JSON(http.StatusOK, []string{"Hello", "World"})
}
